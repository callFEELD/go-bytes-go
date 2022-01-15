package src

import (
	"encoding/binary"
	"math"
	"reflect"
)

type IEncoder interface {
	countByteLen(data interface{}) (error, uint64)

	SetByteOrder(order ByteOrder)
	Encode(data interface{}) (error, []byte)

	encodeUint8(stream []byte, pos uint64, value uint8) ([]byte, uint64)
	encodeUint16(stream []byte, pos uint64, value uint16) ([]byte, uint64)
	encodeUint32(stream []byte, pos uint64, value uint32) ([]byte, uint64)
	encodeUint64(stream []byte, pos uint64, value uint64) ([]byte, uint64)
	encodeFloat32(stream []byte, pos uint64, value float32) ([]byte, uint64)
	encodeFloat64(stream []byte, pos uint64, value float64) ([]byte, uint64)
}

type Encoder struct {
	IEncoder
	order binary.ByteOrder
}

func NewEncoder() *Encoder {
	e := new(Encoder)
	e.order = binary.ByteOrder(binary.LittleEndian)
	return e
}

func (e *Encoder) SetByteOrder(order ByteOrder) {
	if order == BigEndian {
		e.order = binary.ByteOrder(binary.BigEndian)
	} else {
		e.order = binary.ByteOrder(binary.LittleEndian)
	}
}

func (e Encoder) Encode(data interface{}) (error, []byte) {
	// only allocate the amount of bytes which are necessary
	value := reflect.ValueOf(data)
	err, length := countByteLen(value)
	if err != nil {
		return err, []byte{}
	}

	stream := make([]byte, length)
	pos := uint64(0)

	stream, pos = e.encode(value, stream, pos)

	return nil, stream
}

func (e Encoder) encode(value reflect.Value, stream []byte, pos uint64) ([]byte, uint64) {
	switch value.Kind() {
	case reflect.Struct:
		return e.encodeStruct(value, stream, pos)
	case reflect.Array:
		return e.encodeArray(value, stream, pos)
	case reflect.Slice:
		return e.encodeArray(value, stream, pos)
	case reflect.String:
		return e.encodeArray(value, stream, pos)
	default:
		return e.encodeSingle(value, stream, pos)
	}
}

func (e Encoder) encodeStruct(structure reflect.Value, stream []byte, pos uint64) ([]byte, uint64) {
	for i := 0; i < structure.NumField(); i++ {
		field := structure.Field(i)
		stream, pos = e.encode(field, stream, pos)
	}

	return stream, pos
}

func (e Encoder) encodeArray(array reflect.Value, stream []byte, pos uint64) ([]byte, uint64) {
	stream, pos = e.encodeUint32(stream, pos, uint32(array.Len()))

	for i := 0; i < array.Len(); i++ {
		elem := array.Index(i)
		stream, pos = e.encode(elem, stream, pos)
	}

	return stream, pos
}

func (e Encoder) encodeSingle(value reflect.Value, stream []byte, pos uint64) ([]byte, uint64) {
	switch value.Kind() {
	case reflect.Uint8:
		return e.encodeUint8(stream, pos, uint8(value.Uint()))
	case reflect.Int8:
		return e.encodeUint8(stream, pos, uint8(value.Int()))
	case reflect.Uint16:
		return e.encodeUint16(stream, pos, uint16(value.Uint()))
	case reflect.Int16:
		return e.encodeUint16(stream, pos, uint16(value.Int()))
	case reflect.Uint32:
		return e.encodeUint32(stream, pos, uint32(value.Uint()))
	case reflect.Uint:
		return e.encodeUint32(stream, pos, uint32(value.Uint()))
	case reflect.Int32:
		return e.encodeUint32(stream, pos, uint32(value.Int()))
	case reflect.Int:
		return e.encodeUint32(stream, pos, uint32(value.Int()))
	case reflect.Uint64:
		return e.encodeUint64(stream, pos, value.Uint())
	case reflect.Int64:
		return e.encodeUint64(stream, pos, uint64(value.Int()))
	case reflect.Float32:
		return e.encodeFloat32(stream, pos, float32(value.Float()))
	case reflect.Float64:
		return e.encodeFloat64(stream, pos, value.Float())
	default:
		return stream, pos
	}
}

func (e Encoder) encodeUint8(stream []byte, pos uint64, value uint8) ([]byte, uint64) {
	stream[pos] = value
	pos += 1

	return stream, pos
}

func (e Encoder) encodeUint16(stream []byte, pos uint64, value uint16) ([]byte, uint64) {
	byteLen := uint64(2)
	e.order.PutUint16(stream[pos:pos+byteLen], value)
	pos += byteLen

	return stream, pos
}

func (e Encoder) encodeUint32(stream []byte, pos uint64, value uint32) ([]byte, uint64) {
	byteLen := uint64(4)
	e.order.PutUint32(stream[pos:pos+byteLen], value)
	pos += byteLen

	return stream, pos
}

func (e Encoder) encodeUint64(stream []byte, pos uint64, value uint64) ([]byte, uint64) {
	byteLen := uint64(8)
	e.order.PutUint64(stream[pos:pos+byteLen], value)
	pos += byteLen

	return stream, pos
}

func (e Encoder) encodeFloat32(stream []byte, pos uint64, value float32) ([]byte, uint64) {
	byteLen := uint64(4)
	e.order.PutUint32(stream[pos:pos+byteLen], math.Float32bits(value))
	pos += byteLen

	return stream, pos
}
func (e Encoder) encodeFloat64(stream []byte, pos uint64, value float64) ([]byte, uint64) {
	byteLen := uint64(8)
	e.order.PutUint64(stream[pos:pos+byteLen], math.Float64bits(value))
	pos += byteLen

	return stream, pos
}
