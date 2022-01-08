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
	err := recursivelyRunOver(data, func(data interface{}, dataType reflect.Kind) error {
		byteLen, exists := BYTELEN_MAP[dataType]
		if !exists {
			return ShouldNotHappenError
		}
		length += uint64(byteLen)

		return nil
	})

	return err, length
}

func recursivelyRunOver(data interface{}, handle func(data interface{}, dataType reflect.Kind) error) error {
	var err error = nil

	dataType := reflect.TypeOf(data).Kind()

	if !isSupported(data) {
		return UnsupportedDataTypeError
	}

	if reflect.TypeOf(data).Kind() == reflect.Struct {
		err = recursivelyRunOverStructure(data, handle)
	} else if reflect.TypeOf(data).Kind() == reflect.Array ||
		reflect.TypeOf(data).Kind() == reflect.Slice ||
		reflect.TypeOf(data).Kind() == reflect.String {
		err = recursivelyRunOverArray(data, handle)
	} else { // normal data type to handle
		err = handle(data, dataType)
	}

	return err
}

func recursivelyRunOverStructure(data interface{}, handle func(data interface{}, dataType reflect.Kind) error) error {
	structure := reflect.ValueOf(data)

	for i := 0; i < structure.NumField(); i++ {
		field := structure.Field(i).Interface()
		err := recursivelyRunOver(field, handle)
		if err != nil {
			return err
		}
	}

	return nil
}

func recursivelyRunOverArray(data interface{}, handle func(data interface{}, dataType reflect.Kind) error) error {
	array := reflect.ValueOf(data)

	for i := 0; i < array.Len(); i++ {
		elem := array.Index(i).Interface()
		err := recursivelyRunOver(elem, handle)
		if err != nil {
			return err
		}
	}

	return nil
}
