package core

type Codec interface {
	Encode(message Message) (typeName string, raw []byte, err error)
	Decode(typeName string, bytes []byte) (message Message, err error)
}
