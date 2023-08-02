package misc

import "fmt"

/*
1. Deep Copy:

Copying is the data itself, creating a new object. The newly created object does not share memory with the original object. The newly created object creates a new memory address in memory, and the value of the original object will not be affected when the new object value is modified. Since the memory address is different, when you release the memory address, you can release it separately.

Data of value type is all deep copy by default, Array, Int, String, Struct, Float, Bool.

2. Shallow Copy:

Copy the data address, only copy the pointer to the object. At this time, the memory address of the new object and the old object is the same, and the old object will change when the new object value is modified. When the memory address is released, the memory address is also released.

All data of reference type are light copy, Slice and Map by default.
*/

func CopyDeepShallow() {
	Int_1 := 0
	Int_2 := Int_1
	Array_1 := [3]int{1, 2, 3}
	Array_2 := Array_1
	Struct_1 := struct {
		name string
		age  int
	}{"Harry", 41}
	Struct_2 := Struct_1
	fmt.Println("Observation: All the data types will have different memory address, however some may have values pointing to the same data type.")
	fmt.Println("Deep Copy : Data Types")
	fmt.Printf("\nInt: %p - %p", &Int_1, &Int_2)
	Int_2 = Int_2 + 1
	fmt.Printf("\nInt: %v - %v", Int_1, Int_2)
	fmt.Printf("\nArray: %p - %p", &Array_1, &Array_2)
	Array_2[0] = 99
	fmt.Printf("\nArray: %v - %v", Array_1, Array_2)
	fmt.Printf("\nStruct: %p - %p", &Struct_1, &Struct_2)
	Struct_2.name = "Hardy"
	fmt.Printf("\nStruct: %v - %v", Struct_1, Struct_2)

	Slice_1 := []int{1, 2, 3}
	Slice_2 := Slice_1
	Map_1 := map[int]string{1: "a", 2: "b", 3: "c"}
	Map_2 := Map_1
	fmt.Println()
	fmt.Println("Shallow Copy : Data Types")
	fmt.Printf("\nSlice: %p - %p", &Slice_1, &Slice_2)
	Slice_2[0] = 99
	fmt.Printf("\nSlice: %v - %v", Slice_1, Slice_2)
	fmt.Printf("\nMap: %p - %p", &Map_1, &Map_2)
	Map_2[1] = "xxx"
	fmt.Printf("\nMap: %v - %v", Map_1, Map_2)

}
