package tm

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func UnmarshalMessage(m *TypedMessage) (proto.Message, error) {
	messageType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(m.MessageTypeName))
	if err != nil {
		return nil, err
	}

	protoMessage := messageType.New().Interface()
	return protoMessage, proto.Unmarshal(m.MessageData, protoMessage)
}

func MarshalMessage(message proto.Message) (*TypedMessage, error) {
	typeName := string(proto.MessageName(message))
	bytes, err := proto.Marshal(message)
	return &TypedMessage{
		MessageTypeName: typeName,
		MessageData:     bytes,
	}, err
}

type AskResponder[T proto.Message] func(message T)

func (f AskResponder[T]) Reply(message T) {
	f(message)
}
