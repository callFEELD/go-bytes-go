package tests

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	gobytesgo "go-bytes-go/src"
	"testing"
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

func BenchmarkEncode(b *testing.B) {
	e := gobytesgo.NewEncoder()
	hotel := Hotel{}

	alex := Person{Name: "Alex", Age: 32, Gender: Female}
	thomas := Person{Name: "Thomas", Age: 40, Gender: Male}
	dan := Person{Name: "Dan", Age: 18, Gender: Male}

	hotel.Manager = alex
	hotel.Guests = append(hotel.Guests, thomas)
	hotel.Guests = append(hotel.Guests, dan)

	for i := 0; i < b.N; i++ {
		e.Encode(hotel)
	}
}

func BenchmarkEncodeJSON(b *testing.B) {
	hotel := Hotel{}

	alex := Person{Name: "Alex", Age: 32, Gender: Female}
	thomas := Person{Name: "Thomas", Age: 40, Gender: Male}
	dan := Person{Name: "Dan", Age: 18, Gender: Male}

	hotel.Manager = alex
	hotel.Guests = append(hotel.Guests, thomas)
	hotel.Guests = append(hotel.Guests, dan)

	for i := 0; i < b.N; i++ {
		json.Marshal(hotel)
	}
}

func BenchmarkEncodeGOB(b *testing.B) {
	var buff bytes.Buffer
	e := gob.NewEncoder(&buff)
	hotel := Hotel{}

	alex := Person{Name: "Alex", Age: 32, Gender: Female}
	thomas := Person{Name: "Thomas", Age: 40, Gender: Male}
	dan := Person{Name: "Dan", Age: 18, Gender: Male}

	hotel.Manager = alex
	hotel.Guests = append(hotel.Guests, thomas)
	hotel.Guests = append(hotel.Guests, dan)

	for i := 0; i < b.N; i++ {
		e.Encode(hotel)
	}
}
