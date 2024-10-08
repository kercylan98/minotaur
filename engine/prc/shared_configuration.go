package prc

import (
	"github.com/kercylan98/minotaur/engine/prc/codec"
	"github.com/kercylan98/minotaur/toolkit"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"time"
)

const (
	SharedPolicyDecisionStop SharedPolicyDecision = iota
	SharedPolicyDecisionRestart
)

type SharedPolicyDecision uint8
type ErrorPolicyDecisionHandler = toolkit.ErrorPolicyDecisionHandler[SharedPolicyDecision]
type FunctionalErrorPolicyDecisionHandler = toolkit.FunctionalErrorPolicyDecisionHandler[SharedPolicyDecision]

// newSharedConfiguration 创建一个资源控制器的共享配置
func newSharedConfiguration() *SharedConfiguration {
	return &SharedConfiguration{
		codec:                   codec.NewProtobuf(),
		consecutiveRestartLimit: 10,
	}
}

// SharedConfiguration 共享配置
type SharedConfiguration struct {
	runtimeErrorHandler     ErrorPolicyDecisionHandler       // 运行时错误处理器，当处理器不存在时将会引发 panic
	codec                   codec.Codec                      // 编解码器
	sharedStartHook         SharedStartHook                  // 当开启共享时的钩子
	consecutiveRestartLimit int                              // 连续重启限制
	restartInterval         func(count int) time.Duration    // 重启间隔
	unknownReceiverRedirect func(message Message) *ProcessId // 未知接收者重定向
	grpcServerHooks         []GRPCLaunchBeforeHook           // GRPC 服务器钩子
	shareOpenedHooks        []ShareOpenedHook                // 共享连接打开时钩子
	shareClosedHooks        []SharedClosedHook               // 共享连接关闭时钩子
}

// WithShareClosedHooks 设置共享连接关闭时钩子
func (c *SharedConfiguration) WithShareClosedHooks(hooks ...SharedClosedHook) *SharedConfiguration {
	c.shareClosedHooks = append(c.shareClosedHooks, hooks...)
	return c
}

// WithShareOpenedHooks 设置共享连接打开时钩子
func (c *SharedConfiguration) WithShareOpenedHooks(hooks ...ShareOpenedHook) *SharedConfiguration {
	c.shareOpenedHooks = append(c.shareOpenedHooks, hooks...)
	return c
}

// WithGRPCLaunchBeforeHooks 设置 GRPC 服务器钩子，该方法将在创建 GRPC 服务器后调用
func (c *SharedConfiguration) WithGRPCLaunchBeforeHooks(hooks ...GRPCLaunchBeforeHook) *SharedConfiguration {
	c.grpcServerHooks = append(c.grpcServerHooks, hooks...)
	return c
}

// WithUnknownReceiverRedirect 设置未知接收者重定向
func (c *SharedConfiguration) WithUnknownReceiverRedirect(redirect func(message Message) *ProcessId) *SharedConfiguration {
	c.unknownReceiverRedirect = redirect
	return c
}

// WithFixedRestartInterval 使用固定间隔设置重启间隔。
//   - 该配置将会覆盖 WithRestartInterval 方法的设置。
func (c *SharedConfiguration) WithFixedRestartInterval(interval time.Duration) *SharedConfiguration {
	c.restartInterval = func(count int) time.Duration {
		return interval
	}
	return c
}

// WithRestartInterval 使用退避指数设置重启间隔，maxRetries 为最大重试次数，baseDelay 为基础延迟，maxDelay 为最大延迟。
//   - 该配置将会覆盖 WithFixedRestartInterval 方法的设置。
func (c *SharedConfiguration) WithRestartInterval(baseDelay, maxDelay time.Duration) *SharedConfiguration {
	c.restartInterval = func(count int) time.Duration {
		return chrono.StandardExponentialBackoff(count, c.consecutiveRestartLimit, baseDelay, maxDelay)
	}
	return c
}

// WithConsecutiveRestartLimit 设置连续重启限制，当 limit > 0 且连续重启失败到达 limit 时，将进行停止，而非继续重启。
//   - 如果需要控制重启间隔可使用 WithRestartInterval 或 WithFixedRestartInterval 方法。
func (c *SharedConfiguration) WithConsecutiveRestartLimit(limit int) *SharedConfiguration {
	c.consecutiveRestartLimit = limit
	return c
}

// WithSharedHook 设置共享钩子
func (c *SharedConfiguration) WithSharedHook(hook SharedStartHook) *SharedConfiguration {
	c.sharedStartHook = hook
	return c
}

// WithCodec 设置编解码器
func (c *SharedConfiguration) WithCodec(codec codec.Codec) *SharedConfiguration {
	c.codec = codec
	return c
}

// WithRuntimeErrorHandler 设置运行时错误处理器
func (c *SharedConfiguration) WithRuntimeErrorHandler(handler ErrorPolicyDecisionHandler) *SharedConfiguration {
	c.runtimeErrorHandler = handler
	return c
}
