package prc

import (
	"context"
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/engine/prc/codec"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/puzpuzpuz/xsync/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"net"
	"sync/atomic"
	"time"
)

const (
	sharedStateClosed uint32 = iota
	sharedStateSharing
	sharedStateShared
	sharedStateClosing
	sharedStateDead
)

// NewShared 创建一个资源控制器的共享
func NewShared(rc *ResourceController, configurator ...SharedConfigurator) *Shared {
	s := &Shared{
		config:  newSharedConfiguration(),
		rc:      rc,
		streams: xsync.NewMapOf[PhysicalAddress, sharedStream](),
	}

	for _, c := range configurator {
		c.Configure(s.config)
	}

	s.streamServer = newSharedServer(s)
	return s
}

// Shared 是用于对资源控制器进行网络共享的数据结构，它任需要主动的向特定已知的资源管理器发起交互。
type Shared struct {
	config       *SharedConfiguration
	streamServer *sharedServer
	rc           *ResourceController // 共享的资源控制器
	grpc         *grpc.Server
	streams      *xsync.MapOf[PhysicalAddress, sharedStream]
	state        atomic.Uint32
	restartCount int
	restartTimer atomic.Pointer[time.Timer]
}

// GetResourceController 获取资源控制器
func (s *Shared) GetResourceController() *ResourceController {
	return s.rc
}

// Dead 设置共享彻底关闭，将无法再继续重启
func (s *Shared) Dead() {
	if s == nil {
		return
	}
	s.Close()
	timer := s.restartTimer.Load()
	if timer != nil {
		timer.Stop()
	}
	s.state.Store(sharedStateDead)
}

// GetCodec 获取网络共享数据的编解码器
func (s *Shared) GetCodec() codec.Codec {
	return s.config.codec
}

// Share 共享资源控制器，开始监听网络活动
func (s *Shared) Share() error {
	if !s.state.CompareAndSwap(sharedStateClosed, sharedStateSharing) {
		return errors.New("shared is already sharing")
	}
	defer s.state.CompareAndSwap(sharedStateSharing, sharedStateClosed) // 如果共享失败，将状态重置为关闭

	listener, err := net.Listen("tcp", s.rc.GetPhysicalAddress())
	if err != nil {
		return fmt.Errorf("shared listen error: %w", err)
	}

	// 可能存在端口为 0 的情况，使用新物理地址替代
	s.rc.config.physicalAddress = listener.Addr().String()

	s.grpc = grpc.NewServer()
	s.grpc.RegisterService(&Shared_ServiceDesc, s.streamServer)
	for _, hook := range s.config.grpcServerHooks {
		hook(s.grpc)
	}

	go func() {
		s.runtimeError(s.grpc.Serve(listener))
	}()

	s.rc.RegisterResolver(FunctionalPhysicalAddressResolver(func(id *ProcessId) Process {
		pa := id.GetPhysicalAddress()
		process, exist := s.streams.Load(pa)
		if exist {
			return process.(sharedStream)
		}

		var err error
		process, err = s.open(pa)
		if err != nil {
			panic(err)
		}
		return process
	}))

	s.state.Store(sharedStateShared)
	if s.config.sharedStartHook != nil {
		s.config.sharedStartHook.OnSharedStart()
	}
	s.rc.logger().Debug("ResourceController", log.String("feature", "shared"), log.String("listen", s.rc.GetPhysicalAddress()))
	return nil
}

// Close 关闭共享，可指定错误进行关闭，如果指定错误并且存在对运行时的错误处理器，那么将执行
func (s *Shared) Close(err ...error) {
	if s.state.Load() == sharedStateShared {
		if len(err) > 0 {
			for _, e := range err {
				if e != nil {
					s.runtimeError(e)
					return
				}
			}
		}
	}
	if !s.state.CompareAndSwap(sharedStateShared, sharedStateClosing) {
		return
	}

	s.streams.Range(func(key PhysicalAddress, value sharedStream) bool {
		s.detachStream(key)
		return true
	})
	s.grpc.GracefulStop()
	s.streams.Clear()

	s.state.Store(sharedStateClosed)
	s.rc.logger().Debug("ResourceController", log.String("feature", "shared"), log.String("info", "closed"))
}

// open 打开一个资源控制器
func (s *Shared) open(address PhysicalAddress) (sharedStream, error) {
	cc, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := NewSharedClient(cc)
	server, err := client.StreamHandler(context.Background())
	if err != nil {
		return nil, err
	}

	// 与服务器发起握手
	if err = server.Send(&SharedMessage{
		MessageType: &SharedMessage_Handshake{Handshake: &Handshake{Address: s.rc.GetPhysicalAddress()}},
	}); err != nil {
		return nil, err
	}

	stream := newServerStream(address, s, server, cc)
	go func() {
		if err = s.streaming(address, stream); err != nil {
			// 连接断开，无需重连？下次使用会重新建连
			log.Debug("ResourceController", log.String("feature", "shared"), log.String("info", "streaming error"), log.Err(err))
		}
	}()

	return stream, nil
}

func (s *Shared) attachStream(address PhysicalAddress, stream sharedStream) {
	s.streams.Store(address, stream)
}

func (s *Shared) detachStream(address PhysicalAddress) {
	stream, loaded := s.streams.LoadAndDelete(address)
	if loaded {
		_ = stream.Send(&SharedMessage{
			MessageType: &SharedMessage_Farewell{&Farewell{Address: s.rc.GetPhysicalAddress()}},
		})
		stream.Close()
	}
}

//goland:noinspection t
func (s *Shared) runtimeError(err error) {
	if err == nil {
		return
	}
	if s.config.runtimeErrorHandler != nil {
		switch s.config.runtimeErrorHandler.Handle(err) {
		case SharedPolicyDecisionRestart:
			s.Close()
			if err = s.Share(); err != nil {
				s.restartCount++
				if s.restartCount <= s.config.consecutiveRestartLimit || s.config.consecutiveRestartLimit <= 0 {
					s.rc.logger().Debug("ResourceController", log.String("feature", "shared"), log.String("info", "runtime error restart"), log.Err(err))
					var next time.Duration
					if s.config.restartInterval != nil {
						next = s.config.restartInterval(s.restartCount)
					}
					oldTimer := s.restartTimer.Swap(time.AfterFunc(next, func() {
						s.runtimeError(err)
					}))
					if oldTimer != nil {
						oldTimer.Stop()
					}
					return
				}
				s.rc.logger().Error("ResourceController", log.String("feature", "shared"), log.String("info", "runtime error restart failed, stop!"), log.Err(err))
				return
			}
			s.restartCount = 0
			return
		case SharedPolicyDecisionStop:
			s.rc.logger().Debug("ResourceController", log.String("feature", "shared"), log.String("info", "runtime error stop"), log.Err(err))
			s.Close()
			return
		}
	}
	panic(err)
}

func (s *Shared) streaming(address PhysicalAddress, stream sharedStream) (err error) {
	s.attachStream(address, stream)
	defer func() {
		s.detachStream(address)
	}()

	var message *SharedMessage
	for {
		message, err = stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			if _, exist := s.streams.Load(address); !exist {
				return nil // 本地关闭
			}
			return err
		}

		switch m := message.MessageType.(type) {
		case *SharedMessage_DeliveryMessage:
			s.onDeliveryMessage(stream, address, m.DeliveryMessage)
		case *SharedMessage_BatchDeliveryMessage:
			s.onBatchDeliveryMessage(stream, address, m.BatchDeliveryMessage)
		case *SharedMessage_Farewell:
			return nil
		}
	}
}

func (s *Shared) onDeliveryMessage(stream sharedStream, address PhysicalAddress, m *DeliveryMessage) {
	message, err := s.config.codec.Decode(m.MessageType, m.MessageData)
	if err != nil {
		panic(err)
	}
	switch v := message.(type) {
	case *SharedErrorMessage:
		message = errors.New(v.Message)
	}

	sender, receiver := m.Sender, m.Receiver
	receiverProcess := s.rc.GetProcess(receiver)
	if receiverProcess == nil && s.config.unknownReceiverRedirect != nil {
		receiver = s.config.unknownReceiverRedirect(message)
		if receiver != nil {
			receiverProcess = s.rc.GetProcess(receiver)
		}
	}

	if m.System {
		receiverProcess.DeliverySystemMessage(receiver, sender, nil, message)
	} else {
		receiverProcess.DeliveryUserMessage(receiver, sender, nil, message)
	}
}

func (s *Shared) onBatchDeliveryMessage(stream sharedStream, address PhysicalAddress, message *BatchDeliveryMessage) {
	for _, deliveryMessage := range message.Messages {
		s.onDeliveryMessage(stream, address, deliveryMessage)
	}
}
