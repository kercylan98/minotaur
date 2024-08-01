package supervision

import (
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func NewAccidentRecord(primeCulprit, victim *prc.ProcessId, supervisor Supervisor, message, reason prc.Message, strategy Strategy, state *AccidentState, stack []byte) *AccidentRecord {
	return &AccidentRecord{
		PrimeCulprit: primeCulprit,
		Victim:       victim,
		Supervisor:   supervisor,
		Message:      message,
		Reason:       reason,
		Strategy:     strategy,
		State:        state,
		Stack:        stack,
	}
}

// AccidentRecord 事故记录
type AccidentRecord struct {
	PrimeCulprit *prc.ProcessId // 事故元凶（通常是消息发送人，可能不存在或逃逸）
	Victim       *prc.ProcessId // 事故受害者
	Supervisor   Supervisor     // 事故监管者，将由其进行决策
	Message      prc.Message    // 造成事故发生的消息
	Reason       prc.Message    // 事故原因
	Strategy     Strategy       // 受害人携带的监督策略，应由责任人执行
	State        *AccidentState // 事故状态
	Stack        []byte         // 事件堆栈
}

// CrossAccidentRecord 转换为支持跨网络传输的事故记录
func (ar *AccidentRecord) CrossAccidentRecord(strategyName StrategyName, shared *prc.Shared) (*CrossAccidentRecord, error) {
	typeName, data, err := shared.GetCodec().Encode(ar.Message)
	if err != nil {
		return nil, err
	}
	car := &CrossAccidentRecord{
		Strategy:     strategyName,
		PrimeCulprit: ar.PrimeCulprit,
		Victim:       ar.Victim,
		Message: &prc.DeliveryMessage{
			MessageType: typeName,
			MessageData: data,
		},
		State: &CrossAccidentState{
			AccidentTimes: make([]*timestamppb.Timestamp, len(ar.State.accidentTimes)),
		},
		Stack: ar.Stack,
	}

	for _, t := range ar.State.accidentTimes {
		car.State.AccidentTimes = append(car.State.AccidentTimes, timestamppb.New(t))
	}

	if ar.Supervisor != nil {
		car.Supervisor = ar.Supervisor.Ref()
	}

	switch v := ar.Reason.(type) {
	case error:
		car.Reason = v.Error()
	default:
		car.Reason = fmt.Sprintf("%v", v)
	}

	return car, nil
}

func (car *CrossAccidentRecord) AccidentRecord(supervisor Supervisor, shared *prc.Shared, strategy Strategy) (*AccidentRecord, error) {
	msg, err := shared.GetCodec().Decode(car.Message.MessageType, car.Message.MessageData)
	if err != nil {
		return nil, err
	}

	ar := &AccidentRecord{
		PrimeCulprit: car.PrimeCulprit,
		Victim:       car.Victim,
		Supervisor:   supervisor,
		Message:      msg,
		Strategy:     strategy,
		State: &AccidentState{
			accidentTimes: make([]time.Time, len(car.State.AccidentTimes)),
		},
		Stack: car.Stack,
	}
	for _, t := range car.State.AccidentTimes {
		ar.State.accidentTimes = append(ar.State.accidentTimes, t.AsTime())
	}

	if car.Reason != charproc.None {
		ar.Reason = errors.New(car.Reason)
	}

	return ar, nil
}
