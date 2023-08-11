package types

import (
	"fmt"
)

/*
Reference Link : https://golangbyexample.com/slice-in-golang/

A slice points to an underlying array and is internally represented by a slice header.
Unlike array, the size of a slice is flexible and can be changed.

Internally a slice is represented by three things.
1. Pointer to the underlying array
2. Current length of the underlying array
3. Total Capacity which is the maximum capacity to which the underlying array can expand.

	type SliceHeader struct {
	        Pointer uintptr
	        Len  int
	        Cap  int
	}

/golangbyexample.com/wp-content/uploads/2020/05/slice.jpg

############# Creating a Slice ###############

There are four ways of creating a slice
1. Using the []<type>{} format
2. Creating a slice from another slice or array
3. Using make function
4. Using new function

2. Create a slice from Array OR Slice
The format for creating a new slice by re-slicing an existing array is
--> [n]sample[start_index:end_index-1]
The above operation will return a new slice from the array starting from index start to index end-1.
So the element at index end is not included in the newly created slice.
While re-slicing , both start and end index is optional.
- The default value of the start index is zero
- The default value of the end index is the length of the array
- length of newly created slice = (end–start)
- capacity of newly created slice = (length_of_array–start)
- The newly created slices still refer the original array.

3. Using make function
func make([]{type}, length, capacity int) []{type}

Capacity is an optional parameter while creating slice using the make function. When capacity is omitted,
the capacity of the slice is equal length specified for the slice. When using make function, behind the scenes
go allocates an array equal to the capacity. All the elements of the allocated array are initialized with
default zero value of the type.

4. Using the new function
new is a builtin function provided by go that can also be used to create a slice.
It is not a very popular way of creating a slice as make is much more flexible in terms of functionalities .
It is not generally used and also using new function returns a pointer to nil slice. Let’s see an example.
In below example we are using the dereferencing operator ‘*’ as new function returns a pointer to the nil slice.

############# Length vs Capacity ##############

Before moving further, let’s emphasis on understanding the caveats of length and capacity.
Let’s create a simple slice with capacity greater than length.

numbers := make([]int, 3, 5)

Accessing the slice behind its length will result in a run time error “Index out of range”.
It doesn’t matter if the accessed index is within the capacity. So the below line will cause the run time error.

numbers[4] = 5

The length of the slice can be increased up to its capacity by re-slicing. So below re-slice will increase the length from 3 to 5.
numbers = numbers[0:5]

The length of the slice can also be decreased using re-slicing. So below re-slice will decrease the length from 3 to 2
numbers = numbers[0:2]

The advantage of having capacity is that array of size capacity can be pre-allocated during the initialization.
This is a performance boost as if more elements are needed to include in this array then space is already allocated for them.

############## Append to a Slice ###############

func append(slice []Type, elems ...Type) []Type
s := []int{1,2,3}
s = append(s, 4, 5) // append multiple elements
s2 = append(s2, s...) // append another slice

############# Copy a slice   ################
go builtin package provides copy function that can be used to copy a slice. Below is the signature of this function.
It takes in two slices dst and src, and copies data from src to dst. It returns the number of elements copied.

func copy(dst, src []Type) int
There are two cases to be considered while using the copy function:

If the length of src is greater than the length of dst, then the number of elements copied is the length of dst
If the length of dst is greater than the length of src, then the number of elements copied is the length of src
Basically the number of elements copied is minimum of length of (src, dst).

############## Multidimensional Slices #################
As the multi-dimensional array is an array of arrays, similarly multi-dimensional slice is a slice of slices.
To understand this, let's first look at the definition of a slice.

Data field in the slice header is a pointer to the underlying array. For a one dimensional slice, we have below declaration

oneDSlice := make([]int, 2)
To declare a two dimensional slice the declaration would be

twoDSlice = make([][]int, 2)
Above declaration means that we want to create a slice of 2 slices. Carefully understand this point.
But wait a second here, we haven't specified the second dimension here, meaning what is the length of each of the inner 2 slices.
In case of slice, each of the inner slice has to be explicitly intialized like below

	for i := range twoDSlice {
	    twoDSlice[i] = make([]int, 3)
	}

So using range on the original slice, we specify the length each of 2 slices using make.
Below is one other way of doing the same but with slice elements specified

var twoDSlice = make([][]int, 2)
twoDSlice[0] = []int{1, 2, 3}
twoDSlice[1] = []int{4, 5, 6}
Basically, with the above declaration, we create a slice of 2*3 dimensions which is a two-dimensional slice.
The same idea can be extended to two-dimension, three-dimension, and so on.
*/

func SliceCreation() {
	fmt.Println("Slice Creation!!")

	fmt.Println()
	fmt.Println("1. Direct Initialisation")
	fmt.Println("s := []int{1, 2, 3, 4, 5}")
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v len: %v  capacity: %v", s, len(s), cap(s))
	fmt.Println()

	fmt.Println()
	fmt.Println("2. Using make() function")
	fmt.Println("s1 := make([]int, 3)")
	s1 := make([]int, 3)
	fmt.Printf("%v len: %v  capacity: %v", s1, len(s1), cap(s1))
	fmt.Println()

	fmt.Println()
	fmt.Println("3. Using new() function")
	fmt.Println("s1 := new([]int)")
	s2 := new([]int)
	fmt.Printf("%v len: %v  capacity: %v", *s2, len(*s2), cap(*s2)) // dereference to use the slice
	fmt.Println()

	fmt.Println()
	fmt.Println("4. Using re-slicing (Array or Slice)")
	fmt.Println("s3 := s[:]")
	s3 := s[:]
	fmt.Printf("%v len: %v  capacity: %v", s3, len(s3), cap(s3))
	fmt.Println()

}

func SliceAppend() {
	fmt.Println()
	fmt.Println("Slice Append!!")

	fmt.Println()
	fmt.Println("Original Slice : s := []int{1, 2, 3}")
	s := []int{1, 2, 3}
	fmt.Printf("%v len: %v  capacity: %v", s, len(s), cap(s))
	fmt.Println()

	fmt.Println("1. Append single element: s = append(s, 4) ")
	s = append(s, 4) /// append single element
	fmt.Printf("%v len: %v  capacity: %v", s, len(s), cap(s))
	fmt.Println()

	fmt.Println("2. Append multiple element: s = append(s, 5, 6, 7) ")
	s = append(s, 5, 6, 7) // append multiple elements
	fmt.Printf("%v len: %v  capacity: %v", s, len(s), cap(s))
	fmt.Println()

	fmt.Println()
	fmt.Println("3. Append one slice to another : s = append(s, s1...) ")
	s1 := []int{8, 9}
	s2 := []int{10}
	s1 = append(s1, s2...) // append one slice to another
	s = append(s, s1...)   // append one slice to another
	fmt.Printf("%v len: %v  capacity: %v", s, len(s), cap(s))
	fmt.Println()

}

func SliceIterate() {
	fmt.Println()
	fmt.Println("Slice Iterate!!")

	fmt.Println()
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("1. Using for loop")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i])
	}
	fmt.Println()

	fmt.Println("2. Using for-range loop")
	for i, v := range s {
		fmt.Printf("s[%v]=%v ", i, v)
	}
	fmt.Println()
}
