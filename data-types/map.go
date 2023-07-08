package types

import (
	"fmt"
)

/*
Reference Link : https://golangbyexample.com/maps-in-golang/

Creation of Map : key value pairs, pointer type

Maps are golang builtin datatype similar to the hash table which maps a key to a value.
Map is an unordered collection where each key is unique while values can be the same for two or more different keys.
The advantages of using a map are that it provides fast retrieval, search, insert, and delete operations.

Maps are referenced data types. When you assign one map to another both refer to the same underlying map.
Below is the format for a map
map[key_type]value_type{}

############## Allowed Key types in a Map ##############

The map key can be any type that is comparable. Some of the comparable types as defined by go specification are
- boolean
- numeric
- string,
- pointer
- channel
- interface types
- structs – if all it’s field type is comparable
- array – if the type of value of array element is comparable

Some of the types which are not comparable as per go specification and which cannot be used as a key in a map are.
- Slice
- Map
- Function

########## Allowed Value types in a Map #############
Value can be of any type in a map.


############# Creation of Map ###############
1. Without key/value pairs
   m := map[int]int{}
2. With key/value pairs
   m1 := map[int]int{1:1, 2:2}
3. Nil Map
	var m2 map[int]string
	A map can also be declared with var keyword, but it creates a nil map as default zero value of map is nil.
	Adding any key value pair to that map will cause a panic.
	Use Case:
	One use case of having a map declared with var keyword is when an already existing map needs to be assigned to it
	or when we want to assign the result of a function.
*/

func MapCreation() {
	fmt.Println("Map Creation !!")

	fmt.Println("1. Using Map literal")
	fmt.Println("m := map[int]string{1: \"a\", 2: \"b\"}")
	m := map[int]string{1: "a", 2: "b"}
	fmt.Println(m)

	fmt.Println()

	fmt.Println("2. Using make() function")
	fmt.Println("m1 := make(map[string]int)")
	m1 := make(map[string]int)
	m1["a"] = 1
	m1["b"] = 2
	fmt.Println(m)
	fmt.Println()
}

func MapIterate() {
	fmt.Println("Map Iteration !!")

	m := map[int]string{
		1: "a", 2: "b", 3: "c",
	}

	for i, v := range m {
		fmt.Printf("%v:%v ,", i, v)

	}

	fmt.Println()
}
