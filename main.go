package main

import "fmt"

func main() {
	ints := map[string]int64{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	}
	floats := map[string]float64{
		"foo": 1.5,
		"bar": 2.5,
		"baz": 3.5,
	}

	fmt.Printf("sumInts: %v\n", sumInts(ints))
	fmt.Printf("sumFloats: %v\n", sumFloats(floats))


	// Call generic function with explicit "subtype" params
	fmt.Printf("SumIntsOrFloats explicit: %v\n", SumIntsOrFloats[string, int64](ints))
	fmt.Printf("SumIntsOrFloats explicit: %v\n", SumIntsOrFloats[string, float64](floats))
	// You can omit type arguments in calling code when the Go compiler
	// can infer the types you want to use. The compiler infers type
	// arguments from the types of function arguments. Note that this
	// isn’t always possible. For example, if you needed to call a
	// generic function that had no arguments, you would need to include
	// the type arguments in the function call.
	fmt.Printf("SumIntsOrFloats inferred: %v\n", SumIntsOrFloats(ints))
	fmt.Printf("SumIntsOrFloats inferred: %v\n", SumIntsOrFloats(floats))
	
	// Call generic function with explicit "subtype" params
	fmt.Printf("SumNumbers explicit: %v\n", SumNumbers[string, int64](ints))
	fmt.Printf("SumNumbers explicit: %v\n", SumNumbers[string, float64](floats))
	// You can omit type arguments in calling code when the Go compiler
	// can infer the types you want to use. The compiler infers type
	// arguments from the types of function arguments. Note that this
	// isn’t always possible. For example, if you needed to call a
	// generic function that had no arguments, you would need to include
	// the type arguments in the function call.
	fmt.Printf("SumNumbers inferred: %v\n", SumNumbers(ints))
	fmt.Printf("SumNumbers inferred: %v\n", SumNumbers(floats))
}

func sumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func sumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats is a generic function that sums the values of map `m`.
// It supports both int64 and float64 as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Declaring your own constraints
type Number interface {
	int64 | float64
}

// Now you could rewrite SumIntsOrFloats with the new constraint.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
