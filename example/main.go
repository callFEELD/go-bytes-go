package main

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
