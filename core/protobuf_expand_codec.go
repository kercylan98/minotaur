package core

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"io"
	"time"
)

const (
	builtTypeString builtTypeName = iota
	builtTypeUint8
	builtTypeUint16
	builtTypeUint32
	builtTypeUint64
	builtTypeInt8
	builtTypeInt16
	builtTypeInt32
	builtTypeInt64
	builtTypeFloat32
	builtTypeFloat64
	builtTypeBool
	builtTypeTime
	builtTypeTimeDuration
	builtTypeStringPtr
	builtTypeUint8Ptr
	builtTypeUint16Ptr
	builtTypeUint32Ptr
	builtTypeUint64Ptr
	builtTypeInt8Ptr
	builtTypeInt16Ptr
	builtTypeInt32Ptr
	builtTypeInt64Ptr
	builtTypeFloat32Ptr
	builtTypeFloat64Ptr
	builtTypeBoolPtr
	builtTypeStringSlice
	builtTypeUint8Slice
	builtTypeUint16Slice
	builtTypeUint32Slice
	builtTypeUint64Slice
	builtTypeInt8Slice
	builtTypeInt16Slice
	builtTypeInt32Slice
	builtTypeInt64Slice
	builtTypeFloat32Slice
	builtTypeFloat64Slice
	builtTypeBoolSlice
	builtTypeJsonRAW
)

var _ Codec = (*ProtobufExpandCodec)(nil)

type builtTypeName uint8

func (n builtTypeName) String() string {
	return convert.Uint8ToString(uint8(n))
}

type ProtobufExpandCodec struct{}

func NewProtobufExpandCodec() *ProtobufExpandCodec {
	return &ProtobufExpandCodec{}
}

func (p *ProtobufExpandCodec) builtInEncode(message Message) (typeName builtTypeName, raw []byte, err error) {
	err = fmt.Errorf("%w, but got %T", ErrorMessageMustIsProtoMessage, message)
	switch v := message.(type) {
	case string:
		return builtTypeString, []byte(v), nil
	case uint8:
		return p.writeLittleEndian(builtTypeUint8, v)
	case uint16:
		return p.writeLittleEndian(builtTypeUint16, v)
	case uint32:
		return p.writeLittleEndian(builtTypeUint32, v)
	case uint64:
		return p.writeLittleEndian(builtTypeUint64, v)
	case int8:
		return p.writeLittleEndian(builtTypeInt8, v)
	case int16:
		return p.writeLittleEndian(builtTypeInt16, v)
	case int32:
		return p.writeLittleEndian(builtTypeInt32, v)
	case int64:
		return p.writeLittleEndian(builtTypeInt64, v)
	case float32:
		return p.writeLittleEndian(builtTypeFloat32, v)
	case float64:
		return p.writeLittleEndian(builtTypeFloat64, v)
	case bool:
		return p.writeLittleEndian(builtTypeBool, v)
	case *string:
		return builtTypeStringPtr, []byte(*v), nil
	case *uint8:
		return p.writeLittleEndian(builtTypeUint8Ptr, v)
	case *uint16:
		return p.writeLittleEndian(builtTypeUint16Ptr, v)
	case *uint32:
		return p.writeLittleEndian(builtTypeUint32Ptr, v)
	case *uint64:
		return p.writeLittleEndian(builtTypeUint64Ptr, v)
	case *int8:
		return p.writeLittleEndian(builtTypeInt8Ptr, v)
	case *int16:
		return p.writeLittleEndian(builtTypeInt16Ptr, v)
	case *int32:
		return p.writeLittleEndian(builtTypeInt32Ptr, v)
	case *int64:
		return p.writeLittleEndian(builtTypeInt64Ptr, v)
	case *float32:
		return p.writeLittleEndian(builtTypeFloat32Ptr, v)
	case *float64:
		return p.writeLittleEndian(builtTypeFloat64Ptr, v)
	case *bool:
		return p.writeLittleEndian(builtTypeBoolPtr, v)
	case []string:
		var buf bytes.Buffer
		for _, s := range v {
			length := len(s)
			if err = binary.Write(&buf, binary.LittleEndian, uint32(length)); err != nil {
				return
			}

			if _, err = buf.Write([]byte(s)); err != nil {
				return
			}
		}
		return builtTypeStringSlice, buf.Bytes(), nil
	case []uint8:
		return p.writeLittleEndian(builtTypeUint8Slice, v)
	case []uint16:
		return p.writeLittleEndian(builtTypeUint16Slice, v)
	case []uint32:
		return p.writeLittleEndian(builtTypeUint32Slice, v)
	case []uint64:
		return p.writeLittleEndian(builtTypeUint64Slice, v)
	case []int8:
		return p.writeLittleEndian(builtTypeInt8Slice, v)
	case []int16:
		return p.writeLittleEndian(builtTypeInt16Slice, v)
	case []int32:
		return p.writeLittleEndian(builtTypeInt32Slice, v)
	case []int64:
		return p.writeLittleEndian(builtTypeInt64Slice, v)
	case []float32:
		return p.writeLittleEndian(builtTypeFloat32Slice, v)
	case []float64:
		return p.writeLittleEndian(builtTypeFloat64Slice, v)
	case []bool:
		return p.writeLittleEndian(builtTypeBoolSlice, v)
	case time.Time:
		return builtTypeTime, []byte(v.Format(time.RFC3339)), nil
	case time.Duration:
		return p.writeLittleEndian(builtTypeTimeDuration, int64(v))
	case json.RawMessage:
		return builtTypeJsonRAW, v, nil
	}
	return
}

func (p *ProtobufExpandCodec) builtInDecode(typeName builtTypeName, raw []byte) (message Message, err error) {
	err = fmt.Errorf("%w, but got %s", ErrorMessageMustIsProtoMessage, typeName)
	switch typeName {
	case builtTypeString:
		return string(raw), nil
	case builtTypeUint8:
		var v uint8
		return v, p.readLittleEndian(raw, &v)
	case builtTypeUint16:
		var v uint16
		return v, p.readLittleEndian(raw, &v)
	case builtTypeUint32:
		var v uint32
		return v, p.readLittleEndian(raw, &v)
	case builtTypeUint64:
		var v uint64
		return v, p.readLittleEndian(raw, &v)
	case builtTypeInt8:
		var v int8
		return v, p.readLittleEndian(raw, &v)
	case builtTypeInt16:
		var v int16
		return v, p.readLittleEndian(raw, &v)
	case builtTypeInt32:
		var v int32
		return v, p.readLittleEndian(raw, &v)
	case builtTypeInt64:
		var v int64
		return v, p.readLittleEndian(raw, &v)
	case builtTypeFloat32:
		var v float32
		return v, p.readLittleEndian(raw, &v)
	case builtTypeFloat64:
		var v float64
		return v, p.readLittleEndian(raw, &v)
	case builtTypeBool:
		var v bool
		return v, p.readLittleEndian(raw, &v)
	case builtTypeStringPtr:
		var v = string(raw)
		return &v, nil
	case builtTypeUint8Ptr:
		var v uint8
		err = p.readLittleEndian(raw, &v)
		message = &v
	case builtTypeUint16Ptr:
		var v uint16
		err = p.readLittleEndian(raw, &v)
		message = &v
	case builtTypeUint32Ptr:
		var v uint32
		err = p.readLittleEndian(raw, &v)
		message = &v
	case builtTypeUint64Ptr:
		var v uint64
		err = p.readLittleEndian(raw, &v)
		message = &v
	case builtTypeInt8Ptr:
		var v int8
		err = p.readLittleEndian(raw, &v)
		message = &v
	case builtTypeInt16Ptr:
		var v int16
		err = p.readLittleEndian(raw, &v)
		message = &v
	case builtTypeInt32Ptr:
		var v int32
		err = p.readLittleEndian(raw, &v)
		message = &v
	case builtTypeInt64Ptr:
		var v int64
		err = p.readLittleEndian(raw, &v)
		message = &v
	case builtTypeFloat32Ptr:
		var v float32
		err = p.readLittleEndian(raw, &v)
		message = &v
	case builtTypeFloat64Ptr:
		var v float64
		err = p.readLittleEndian(raw, &v)
		message = &v
	case builtTypeBoolPtr:
		var v bool
		err = p.readLittleEndian(raw, &v)
		message = &v
	case builtTypeStringSlice:
		buf := bytes.NewReader(raw)
		var result []string

		for {
			var length uint32
			err = binary.Read(buf, binary.LittleEndian, &length)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return nil, err
			}

			strBytes := make([]byte, length)
			_, err = buf.Read(strBytes)
			if err != nil {
				return nil, err
			}

			result = append(result, string(strBytes))
		}

		return result, nil
	case builtTypeUint8Slice:
		var v = make([]uint8, p.readLittleEndianLength(raw))
		return v, p.readLittleEndian(raw[4:], &v)
	case builtTypeUint16Slice:
		var v = make([]uint16, p.readLittleEndianLength(raw))
		return v, p.readLittleEndian(raw[4:], &v)
	case builtTypeUint32Slice:
		var v = make([]uint32, p.readLittleEndianLength(raw))
		return v, p.readLittleEndian(raw[4:], &v)
	case builtTypeUint64Slice:
		var v = make([]uint64, p.readLittleEndianLength(raw))
		return v, p.readLittleEndian(raw[4:], &v)
	case builtTypeInt8Slice:
		var v = make([]int8, p.readLittleEndianLength(raw))
		return v, p.readLittleEndian(raw[4:], &v)
	case builtTypeInt16Slice:
		var v = make([]int16, p.readLittleEndianLength(raw))
		return v, p.readLittleEndian(raw[4:], &v)
	case builtTypeInt32Slice:
		var v = make([]int32, p.readLittleEndianLength(raw))
		return v, p.readLittleEndian(raw[4:], &v)
	case builtTypeInt64Slice:
		var v = make([]int64, p.readLittleEndianLength(raw))
		return v, p.readLittleEndian(raw[4:], &v)
	case builtTypeFloat32Slice:
		var v = make([]float32, p.readLittleEndianLength(raw))
		return v, p.readLittleEndian(raw[4:], &v)
	case builtTypeFloat64Slice:
		var v = make([]float64, p.readLittleEndianLength(raw))
		return v, p.readLittleEndian(raw[4:], &v)
	case builtTypeBoolSlice:
		var v = make([]bool, p.readLittleEndianLength(raw))
		return v, p.readLittleEndian(raw[4:], &v)
	case builtTypeTime:
		t, err := time.Parse(time.RFC3339, string(raw))
		return t, err
	case builtTypeTimeDuration:
		var v int64
		return time.Duration(v), p.readLittleEndian(raw, &v)
	case builtTypeJsonRAW:
		return json.RawMessage(raw), nil
	}
	return
}

func (p *ProtobufExpandCodec) Encode(message Message) (typeName string, bytes []byte, err error) {
	pm, ok := message.(proto.Message)
	if !ok {
		n, b, e := p.builtInEncode(message)
		if e != nil {
			return "", b, e
		}
		return n.String(), b, e
	}

	typeName = string(proto.MessageName(pm))
	bytes, err = proto.Marshal(pm)
	return
}

func (p *ProtobufExpandCodec) Decode(typeName string, bytes []byte) (message Message, err error) {
	messageType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(typeName))
	if err != nil {
		return p.builtInDecode(builtTypeName(convert.StringToUint8(typeName)), bytes)
	}

	protoMessage := messageType.New().Interface()
	err = proto.Unmarshal(bytes, protoMessage)
	return protoMessage, err
}

func (p *ProtobufExpandCodec) writeLittleEndian(typeName builtTypeName, v any) (name builtTypeName, raw []byte, err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.LittleEndian, v); err == nil {
		raw = buf.Bytes()
	}
	name = typeName

	length := len(raw)
	buf = new(bytes.Buffer)
	if err = binary.Write(buf, binary.LittleEndian, uint32(length)); err != nil {
		return
	}
	raw = append(buf.Bytes(), raw...)
	return
}
func (p *ProtobufExpandCodec) readLittleEndianLength(raw []byte) uint32 {
	var length uint32
	buf := bytes.NewReader(raw)
	_ = binary.Read(buf, binary.LittleEndian, &length)
	return length
}

func (p *ProtobufExpandCodec) readLittleEndian(raw []byte, receive any) (err error) {
	buf := bytes.NewReader(raw)
	err = binary.Read(buf, binary.LittleEndian, receive)
	return
}
