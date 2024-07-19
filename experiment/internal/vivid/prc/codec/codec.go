package codec

type Codec interface {
	Encode(message any) (typeName string, raw []byte, err error)
	Decode(typeName string, bytes []byte) (message any, err error)
}
