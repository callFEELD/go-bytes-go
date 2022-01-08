# Go Bytes Go

A simple encoding lib to print/parse structured and unstructured data over a non self-describing bytestream.

This lib will use the included reflection in the Go programming language to automatically encode/decode the given data.

## Example
```go
import (
	"fmt"
	gobytesgo "go-bytes-go/src"
)

type Gender uint8

const (
	Female Gender = iota
	Male
)

type Person struct {
	Name   string
	Age    uint8
	Gender Gender
}

type Hotel struct {
	Manager Person
	Guests  []Person
}

func main() {
	// Create the Encoder and select Big Endian
	e := gobytesgo.NewEncoder()
	e.SetByteOrder(gobytesgo.BigEndian)

	// Prepare the structure which needs to be encoded
	hotel := Hotel{}

	alex := Person{Name: "Alex", Age: 32, Gender: Female}
	thomas := Person{Name: "Thomas", Age: 40, Gender: Male}
	dan := Person{Name: "Dan", Age: 18, Gender: Male}

	hotel.Manager = alex
	hotel.Guests = append(hotel.Guests, thomas)
	hotel.Guests = append(hotel.Guests, dan)

	err, stream := e.Encode(hotel)
	if err != nil {
		fmt.Println(err)
	}

	for _, b := range stream {
		fmt.Printf("%8b ", b)
	}
	fmt.Println()

	fmt.Printf("Stream of bytes length: %d\n", len(stream))
}
```
[Code](./example/main.go)

## Data Type mapping
This table is also reflecting the supported datatypes.

| Go Data Type | Amount of Bytes    |
|--------------|--------------------|
 | bool         | 1                  |
 | u/int8       | 1                  |
 | u/int16      | 2                  |
 | u/int32      | 4                  |
 | u/int64      | 6                  |
 | u/int        | 4 (assuming int32) |
 | float32      | 4                  |
 | float64      | 8                  |
 | array        | 4 + len * dataType |
 | slice        | 4 + len * dataType |
 | string       | 4 + len * dataType |           
| struct       | 0 (*)              |

(*) Using this encoding lib, it is assumed that the Encoder/Decoder knows the underlying structure.

## Limitations
- Maximum length of arrays and strings are limited by the MAX(uint32) number

## ToDo's
- [X] Encoder
- [ ] Decoder
- [ ] Documentation
- [ ] Examples 