package src

import "reflect"

type ByteOrder int

const (
	LittleEndian ByteOrder = iota
	BigEndian
)

var BYTELEN_MAP = map[reflect.Kind]uint{
	reflect.Uint8:   1,
	reflect.Int8:    1,
	reflect.Uint16:  2,
	reflect.Int16:   2,
	reflect.Uint32:  4,
	reflect.Int32:   4,
	reflect.Int:     4,
	reflect.Uint:    4,
	reflect.Uint64:  8,
	reflect.Int64:   8,
	reflect.Float32: 4,
	reflect.Float64: 8,
	reflect.Struct:  0,
	reflect.Array:   4,
	reflect.String:  4,
	reflect.Slice:   4,
}

var SUPPORTED_TYPES = []reflect.Kind{
	reflect.String,
	reflect.Struct,
	reflect.Array,
	reflect.Slice,
	reflect.Bool,
	reflect.Int,
	reflect.Uint,
	reflect.Float32,
	reflect.Float64,
	reflect.Int8,
	reflect.Uint8,
	reflect.Int16,
	reflect.Uint16,
	reflect.Int32,
	reflect.Uint32,
	reflect.Int64,
	reflect.Uint64,
}
