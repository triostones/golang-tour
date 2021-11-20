package main

/*
bool
string
int  int8  int16  int32  int64
unit8  uint16  uint32  uint64 uintptr
byte // alias for uint8
rune // alias for int32
	// represents a Unicode code point

float32 float64
complex64 complex128
*/

import "fmt"
import "math/cmplx"

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T value: %v\n", z, z)
}
