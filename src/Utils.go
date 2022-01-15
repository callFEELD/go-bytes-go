package src

import (
	"reflect"
)

func countByteLen(data reflect.Value) (error, uint64) {
	switch data.Kind() {
	case reflect.Array:
		return countByteLenArray(data)
	case reflect.Slice:
		return countByteLenArray(data)
	case reflect.String:
		return nil, Int32ByteLen + uint64(data.Len())*Int8ByteLen
	case reflect.Struct:
		return countByteLenStruct(data)
	case reflect.Uint8:
		return nil, Int8ByteLen
	case reflect.Int8:
		return nil, Int8ByteLen
	case reflect.Uint16:
		return nil, Int16ByteLen
	case reflect.Int16:
		return nil, Int16ByteLen
	case reflect.Uint32:
		return nil, Int32ByteLen
	case reflect.Int32:
		return nil, Int32ByteLen
	case reflect.Int:
		return nil, Int32ByteLen
	case reflect.Uint:
		return nil, Int32ByteLen
	case reflect.Uint64:
		return nil, Int64ByteLen
	case reflect.Int64:
		return nil, Int64ByteLen
	case reflect.Float32:
		return nil, Real32ByteLen
	case reflect.Float64:
		return nil, Real64ByteLen
	default:
		return UnsupportedDataTypeError, 0
	}
}

func countByteLenArray(data reflect.Value) (error, uint64) {
	var err error = nil
	var elemByteLen uint64 = 0

	arrayLen := uint64(data.Len())
	if arrayLen > 0 {
		err, elemByteLen = countByteLen(data.Index(0))
		if err != nil {
			return nil, 0
		}
	}

	return nil, Int32ByteLen + arrayLen*elemByteLen
}

func countByteLenStruct(data reflect.Value) (error, uint64) {
	var err error = nil
	var byteLen uint64 = 0
	var elemByteLen uint64 = 0

	for i := 0; i < data.NumField(); i++ {
		err, elemByteLen = countByteLen(data.Field(i))
		if err != nil {
			return nil, 0
		}
		byteLen += elemByteLen
	}

	return nil, byteLen
}
