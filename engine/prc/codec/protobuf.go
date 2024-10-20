package codec

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func NewProtobuf() *Protobuf {
	return new(Protobuf)
}

type Protobuf struct{}

func (p *Protobuf) Encode(message any) (typeName string, bytes []byte, err error) {
	pm, ok := message.(proto.Message)
	if !ok {
		return "", nil, fmt.Errorf("message is not a proto.Message, got %T", message)
	}

	typeName = string(proto.MessageName(pm))
	bytes, err = proto.Marshal(pm)
	return
}

func (p *Protobuf) Decode(typeName string, bytes []byte) (message any, err error) {
	messageType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(typeName))
	if err != nil {
		return nil, fmt.Errorf("message is not a proto.Message, got %T", message)
	}

	protoMessage := messageType.New().Interface()
	err = proto.Unmarshal(bytes, protoMessage)
	return protoMessage, err
}
