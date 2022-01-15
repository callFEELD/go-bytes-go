package src

type ByteOrder int

const (
	LittleEndian ByteOrder = iota
	BigEndian
)

const (
	Int8ByteLen   uint64 = 1
	Int16ByteLen  uint64 = 2
	Int32ByteLen  uint64 = 4
	Int64ByteLen  uint64 = 8
	Real32ByteLen uint64 = 4
	Real64ByteLen uint64 = 8
)
