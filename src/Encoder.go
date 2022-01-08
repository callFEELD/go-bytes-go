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
	d := new(Encoder)
	d.order = binary.ByteOrder(binary.LittleEndian)
	return d
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
	err, length := countByteLen(data)
	if err != nil {
		return err, []byte{}
	}

	stream := make([]byte, length)
	pos := uint64(0)

	// recursively go through all data types
	err = recursivelyRunOver(data, func(data interface{}, dataType reflect.Kind) error {
		switch dataType {
		case reflect.Struct: // no encoding special encoding to do on start of structures
			return nil
		case reflect.Array: // on strings and arrays we need to encode the length as an uint32 value
			stream, pos = e.encodeUint32(stream, pos, uint32(reflect.ValueOf(data).Len()))
		case reflect.String:
			stream, pos = e.encodeUint32(stream, pos, uint32(reflect.ValueOf(data).Len()))
		case reflect.Uint8:
			stream, pos = e.encodeUint8(stream, pos, uint8(reflect.ValueOf(data).Uint()))
		case reflect.Int8:
			stream, pos = e.encodeUint8(stream, pos, uint8(reflect.ValueOf(data).Int()))
		case reflect.Uint16:
			stream, pos = e.encodeUint16(stream, pos, uint16(reflect.ValueOf(data).Uint()))
		case reflect.Int16:
			stream, pos = e.encodeUint16(stream, pos, uint16(reflect.ValueOf(data).Int()))
		case reflect.Uint32:
			stream, pos = e.encodeUint32(stream, pos, uint32(reflect.ValueOf(data).Uint()))
		case reflect.Uint:
			stream, pos = e.encodeUint32(stream, pos, uint32(reflect.ValueOf(data).Uint()))
		case reflect.Int32:
			stream, pos = e.encodeUint32(stream, pos, uint32(reflect.ValueOf(data).Int()))
		case reflect.Int:
			stream, pos = e.encodeUint32(stream, pos, uint32(reflect.ValueOf(data).Int()))
		case reflect.Uint64:
			stream, pos = e.encodeUint64(stream, pos, reflect.ValueOf(data).Uint())
		case reflect.Int64:
			stream, pos = e.encodeUint64(stream, pos, uint64(reflect.ValueOf(data).Int()))
		case reflect.Float32:
			stream, pos = e.encodeFloat32(stream, pos, float32(reflect.ValueOf(data).Float()))
		case reflect.Float64:
			stream, pos = e.encodeFloat64(stream, pos, reflect.ValueOf(data).Float())
		}

		return nil
	})

	return err, stream
}

func (e Encoder) encodeUint8(stream []byte, pos uint64, value uint8) ([]byte, uint64) {
	byteLen := uint64(BYTELEN_MAP[reflect.Uint8])
	stream[pos] = value
	pos += byteLen

	return stream, pos
}

func (e Encoder) encodeUint16(stream []byte, pos uint64, value uint16) ([]byte, uint64) {
	byteLen := uint64(BYTELEN_MAP[reflect.Uint16])
	e.order.PutUint16(stream[pos:pos+byteLen], value)
	pos += byteLen

	return stream, pos
}

func (e Encoder) encodeUint32(stream []byte, pos uint64, value uint32) ([]byte, uint64) {
	byteLen := uint64(BYTELEN_MAP[reflect.Uint32])
	e.order.PutUint32(stream[pos:pos+byteLen], value)
	pos += byteLen

	return stream, pos
}

func (e Encoder) encodeUint64(stream []byte, pos uint64, value uint64) ([]byte, uint64) {
	byteLen := uint64(BYTELEN_MAP[reflect.Uint64])
	e.order.PutUint64(stream[pos:pos+byteLen], value)
	pos += byteLen

	return stream, pos
}

func (e Encoder) encodeFloat32(stream []byte, pos uint64, value float32) ([]byte, uint64) {
	byteLen := uint64(BYTELEN_MAP[reflect.Float32])
	e.order.PutUint32(stream[pos:pos+byteLen], math.Float32bits(value))
	pos += byteLen

	return stream, pos
}
func (e Encoder) encodeFloat64(stream []byte, pos uint64, value float64) ([]byte, uint64) {
	byteLen := uint64(BYTELEN_MAP[reflect.Float64])
	e.order.PutUint64(stream[pos:pos+byteLen], math.Float64bits(value))
	pos += byteLen

	return stream, pos
}
