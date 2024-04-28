package codec

type Encoder interface {
	Encode(src []byte) ([]byte, error)
}

type Decoder interface {
	Decode(src []byte) ([]byte, error)
}

type Codec interface {
	Encoder
	Decoder
}
