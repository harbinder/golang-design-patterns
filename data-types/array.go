package types

import (
	"fmt"
)

/*
Array in golang is a bit different than other languages
It is a data type with :
- store multiple elements of same type
- contiguous memory allocation
- fixed size

An array is a contiguous collection of elements of the same type. It is an ordered sequence of elements
stored contiguously in memory

Array are value type in go.
So an array variable name is not a pointer to the first element in fact it denotes the entire array. A copy of the array will be created when
- An array variable is assigned to another array variable.
- An array variable is passed as an argument to a function.
*/
func ArrayDeclaration() {
	fmt.Println("Array Declations!!")
	var a [3]int
	a = [3]int{1, 2} // As length is 3, third element will get default value
	fmt.Println()
	fmt.Printf("a: %v, length: %v", a, len(a))

	a1 := []bool{true, false, false} // length calculated by compiler
	fmt.Println()
	fmt.Printf("a: %v, length: %v", a1, len(a1))

	a2 := [3]int{} // initialised with default value of data type
	fmt.Println()
	fmt.Printf("a: %v, length: %v", a2, len(a2))

	a3 := [...]int{} // length calculated by compiler
	fmt.Println()
	fmt.Printf("a: %v, length: %v", a3, len(a3))
	fmt.Println()
}

/*
An array can be iterated using:

- Using for loop
- Using for-range loop
*/
func ArrayIterate() {
	fmt.Println()
	fmt.Println("Array Iterate!!")
	fmt.Println()

	a := [...]string{"a", "b", "c", "d"}
	len := len(a)
	fmt.Println("1. Using for loop")
	for i := 0; i < len; i++ {
		fmt.Printf("%v ", a[i])
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("2. Using for-range loop")
	for i, v := range a {
		fmt.Printf("a[%v]=%v ", i, v)
	}
	fmt.Println()
}

func ArrayMultiDimention() {
	fmt.Println()
	fmt.Println("Array Multi Dimension!!")
	fmt.Println()
	a := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i, row := range a {
		fmt.Println("Row:", i+1)
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Println()
	}
}
