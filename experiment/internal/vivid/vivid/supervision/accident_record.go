package supervision

import (
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/experiment/internal/vivid/prc"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func NewAccidentRecord(primeCulprit, victim *prc.ProcessRef, supervisor Supervisor, message, reason prc.Message, strategy Strategy, state *AccidentState) *AccidentRecord {
	return &AccidentRecord{
		PrimeCulprit: primeCulprit,
		Victim:       victim,
		Supervisor:   supervisor,
		Message:      message,
		Reason:       reason,
		Strategy:     strategy,
		State:        state,
	}
}

// AccidentRecord 事故记录
type AccidentRecord struct {
	PrimeCulprit *prc.ProcessRef // 事故元凶（通常是消息发送人，可能不存在或逃逸）
	Victim       *prc.ProcessRef // 事故受害者
	Supervisor   Supervisor      // 事故监管者，将由其进行决策
	Message      prc.Message     // 造成事故发生的消息
	Reason       prc.Message     // 事故原因
	Strategy     Strategy        // 受害人携带的监督策略，应由责任人执行
	State        *AccidentState  // 事故状态
}

// CrossAccidentRecord 转换为支持跨网络传输的事故记录
func (ar *AccidentRecord) CrossAccidentRecord(strategyName StrategyName, shared *prc.Shared) (*CrossAccidentRecord, error) {
	typeName, data, err := shared.GetCodec().Encode(ar.Message)
	if err != nil {
		return nil, err
	}
	car := &CrossAccidentRecord{
		Strategy:     strategyName,
		PrimeCulprit: ar.PrimeCulprit.GetId(),
		Victim:       ar.Victim.GetId(),
		Message: &prc.DeliveryMessage{
			MessageType: typeName,
			MessageData: data,
		},
		State: &CrossAccidentState{
			AccidentTimes: make([]*timestamppb.Timestamp, len(ar.State.accidentTimes)),
		},
	}

	for _, t := range ar.State.accidentTimes {
		car.State.AccidentTimes = append(car.State.AccidentTimes, timestamppb.New(t))
	}

	if ar.Supervisor != nil {
		car.Supervisor = ar.Supervisor.Ref().GetId()
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
		PrimeCulprit: prc.NewProcessRef(car.PrimeCulprit),
		Victim:       prc.NewProcessRef(car.Victim),
		Supervisor:   supervisor,
		Message:      msg,
		Strategy:     strategy,
		State: &AccidentState{
			accidentTimes: make([]time.Time, len(car.State.AccidentTimes)),
		},
	}
	for _, t := range car.State.AccidentTimes {
		ar.State.accidentTimes = append(ar.State.accidentTimes, t.AsTime())
	}

	if car.Reason != charproc.None {
		ar.Reason = errors.New(car.Reason)
	}

	return ar, nil
}
