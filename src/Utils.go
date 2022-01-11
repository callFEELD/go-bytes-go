package src

import (
	"reflect"
)

func isSupported(data interface{}) bool {
	for _, t := range SUPPORTED_TYPES {
		if reflect.TypeOf(data).Kind() == t {
			return true
		}
	}

	return false
}

func countByteLen(data interface{}) (error, uint64) {
	var length uint64 = 0

	// recursively go through all data types
	err := recursivelyRunOver(reflect.ValueOf(data), func(data reflect.Value, dataType reflect.Kind) error {
		byteLen, exists := BYTELEN_MAP[dataType]
		if !exists {
			return ShouldNotHappenError
		}
		length += uint64(byteLen)

		return nil
	})

	return err, length
}

func recursivelyRunOver(value reflect.Value, handle func(value reflect.Value, dataType reflect.Kind) error) error {
	var err error = nil

	dataType := value.Kind()

	if !isSupported(dataType) {
		return UnsupportedDataTypeError
	}

	if dataType == reflect.Struct {
		err = recursivelyRunOverStructure(value, handle)
	} else if dataType == reflect.Array ||
		dataType == reflect.Slice ||
		dataType == reflect.String {
		err = recursivelyRunOverArray(value, handle)
	} else { // normal data type to handle
		err = handle(value, dataType)
	}

	return err
}

func recursivelyRunOverStructure(structure reflect.Value, handle func(value reflect.Value, dataType reflect.Kind) error) error {
	for i := 0; i < structure.NumField(); i++ {
		field := structure.Field(i)
		err := recursivelyRunOver(field, handle)
		if err != nil {
			return err
		}
	}

	return nil
}

func recursivelyRunOverArray(array reflect.Value, handle func(value reflect.Value, dataType reflect.Kind) error) error {
	for i := 0; i < array.Len(); i++ {
		elem := array.Index(i)
		err := recursivelyRunOver(elem, handle)
		if err != nil {
			return err
		}
	}

	return nil
}
