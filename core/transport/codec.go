package transport

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/core"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"io"
	"time"
)

type Codec interface {
	Encode(message core.Message) (typeName string, raw []byte, err error)
	Decode(typeName string, bytes []byte) (message core.Message, err error)
}

const (
	builtTypeString       builtTypeName = "$1"
	builtTypeUint8        builtTypeName = "$2"
	builtTypeUint16       builtTypeName = "$3"
	builtTypeUint32       builtTypeName = "$4"
	builtTypeUint64       builtTypeName = "$5"
	builtTypeInt8         builtTypeName = "$6"
	builtTypeInt16        builtTypeName = "$7"
	builtTypeInt32        builtTypeName = "$8"
	builtTypeInt64        builtTypeName = "$9"
	builtTypeFloat32      builtTypeName = "$10"
	builtTypeFloat64      builtTypeName = "$11"
	builtTypeBool         builtTypeName = "$12"
	builtTypeTime         builtTypeName = "$13"
	builtTypeTimeDuration builtTypeName = "$14"
	builtTypeStringPtr    builtTypeName = "*$1"
	builtTypeUint8Ptr     builtTypeName = "*$2"
	builtTypeUint16Ptr    builtTypeName = "*$3"
	builtTypeUint32Ptr    builtTypeName = "*$4"
	builtTypeUint64Ptr    builtTypeName = "*$5"
	builtTypeInt8Ptr      builtTypeName = "*$6"
	builtTypeInt16Ptr     builtTypeName = "*$7"
	builtTypeInt32Ptr     builtTypeName = "*$8"
	builtTypeInt64Ptr     builtTypeName = "*$9"
	builtTypeFloat32Ptr   builtTypeName = "*$10"
	builtTypeFloat64Ptr   builtTypeName = "*$11"
	builtTypeBoolPtr      builtTypeName = "*$12"
	builtTypeStringSlice  builtTypeName = "[]1"
	builtTypeUint8Slice   builtTypeName = "[]2"
	builtTypeUint16Slice  builtTypeName = "[]3"
	builtTypeUint32Slice  builtTypeName = "[]4"
	builtTypeUint64Slice  builtTypeName = "[]5"
	builtTypeInt8Slice    builtTypeName = "[]6"
	builtTypeInt16Slice   builtTypeName = "[]7"
	builtTypeInt32Slice   builtTypeName = "[]8"
	builtTypeInt64Slice   builtTypeName = "[]9"
	builtTypeFloat32Slice builtTypeName = "[]10"
	builtTypeFloat64Slice builtTypeName = "[]11"
	builtTypeBoolSlice    builtTypeName = "[]12"
)

type builtTypeName = string

type protobufCodec struct{}

func (p *protobufCodec) builtInEncode(message core.Message) (typeName string, raw []byte, err error) {
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
	}
	return
}

func (p *protobufCodec) builtInDecode(typeName string, raw []byte) (message core.Message, err error) {
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
	}
	return
}

func (p *protobufCodec) Encode(message core.Message) (typeName string, bytes []byte, err error) {
	pm, ok := message.(proto.Message)
	if !ok {
		return p.builtInEncode(message)
	}

	typeName = string(proto.MessageName(pm))
	bytes, err = proto.Marshal(pm)
	return
}

func (p *protobufCodec) Decode(typeName string, bytes []byte) (message core.Message, err error) {
	messageType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(typeName))
	if err != nil {
		return p.builtInDecode(typeName, bytes)
	}

	protoMessage := messageType.New().Interface()
	err = proto.Unmarshal(bytes, protoMessage)
	return protoMessage, err
}

func (p *protobufCodec) writeLittleEndian(typeName builtTypeName, v any) (name builtTypeName, raw []byte, err error) {
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
func (p *protobufCodec) readLittleEndianLength(raw []byte) uint32 {
	var length uint32
	buf := bytes.NewReader(raw)
	_ = binary.Read(buf, binary.LittleEndian, &length)
	return length
}

func (p *protobufCodec) readLittleEndian(raw []byte, receive any) (err error) {
	buf := bytes.NewReader(raw)
	err = binary.Read(buf, binary.LittleEndian, receive)
	return
}
