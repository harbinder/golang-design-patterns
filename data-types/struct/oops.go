package structs

import "fmt"

/*
#### Inheritance ####
	https://golangbyexample.com/oop-inheritance-golang-complete/
	struct EMBEDDING - inherit properties and methods of parent
	interface implement - inherit sub-type ie using child struct object as argument for parent call. By using a Parent interface
	struct & interface - besides above 2 features, we can use type heirarchy with combination of struct embedding & interface implementing

Point 1: The very basic use case of inheritance is child type should be able to access common data and methods of a parent type.
This is done in GO via EMBEDDING. The base struct is embedded in child struct and base’s data and methods can directly be
accessed by child struct.
Point 2: One of the limitations  above  is that you CANNOT pass the child type to a function that expects the base type as
GO does not allow TYPE inheritance.
Point 3: Basically  sub-typing is not possible in GO by just using EMBEDDING.
Let’s try to fix this error. This is where GO interfaces come into the picture.
Point 4: So iinstead of creating memthod direectly for BASE struct, it should be implemented via Interface.
So when CHILD  struct embeds the BASE struct, it also implements  the Interface.
Now we can pass either BASE or CHILD struct object as agruments to same function
Point 5:
But there is one more limitation. Let’s say CHILD and BASE structs both have one more function M2(), which is being called by extended BASE struct
method M1(). Ideally, if M1() of BASE struct is iinvoked by object of CHILD struct, then M2() of CHILD struct should be invoked.
BUT this is not the case iin GO.
Point 6: One way to fix the above problem is to make M2() as a property which is of type function in the BASE struct.
This is possible in GO as functions are first-class variables in GO.

#### Polymorphism ####
	1. Compile time - Not supported in Go
 			- Method overloadiing
 			- Operator overloading
			But we can  use Variadic functions to. achieve similar things"
	2. RunTime - Supported in Go

#### Encapsulation ####
	Exported & Non-Exported at Package level only
	Nothing like Public/Private/Protected

#### Abstract Class ####
	https://golangbyexample.com/go-abstract-class/
*/

/*
######## EXAMPLE: TYPE HIERARCHY  ########

We have an interface and struct for ANY generic ANIMAL
(iAnimal) and (animal)

On similar lines, we have interface & struct each for Aquatic and Non-Aquatic Animals

iAquatic interface : embedding iAnimal interface
aquatic struct: embedding animal struct

iNonAquatic interface : embedding iAnimal
nonAquatic struct: embedding animal

Finally, we create 2 Concrete Types
shark struct - embedding aquatic struct
lion struct - embedding  nonAquatic struct

	In snippet below, see how we are able to create a hierarchy (see below).
	This is the idiomatic way of go to create type hierarchy and we are able to achieve this by using embedding both
	on struct level and on the interface level. The point to be noted here is that if you want distinction in your
	type hierarchy where lets say a “shark”  should not be both “iAquatic” and “iNonAquatic”  , then there should be
	at least one method in the method sets of “iAquatic” and “iNonAquatic”  which is not present in the other.
	In our example “swim”  and “walk”  are those methods.
*/
type iAnimal interface {
	breathe()
}
type animal struct {
}

// animal struct implements iAnimal interface
func (a *animal) breathe() {
	fmt.Println("Animal breate")
}

// iAquatic interface embedds iAnimal interface
type iAquatic interface {
	iAnimal
	swim()
}

// aquatic struct embedds animal struct
type aquatic struct {
	animal
}

// aquatic sturct implements iAquatic interface
func (a *aquatic) swim() {
	fmt.Println("Aquatic swim")
}

// iNonAquatic interface embedds iAnimal interface
type iNonAquatic interface {
	iAnimal
	walk()
}

// nonAquatic struct embedds animal struct
type nonAquatic struct {
	animal
}

// nonAquatic struct implements iNonAquatic interface
func (a *nonAquatic) walk() {
	fmt.Println("Non-Aquatic walk")
}

type shark struct {
	aquatic
}
type lion struct {
	nonAquatic
}

func ExampleTypeHeirarchy() {
	shark := &shark{}
	checkAquatic(shark)
	checkAnimal(shark)
	lion := &lion{}
	checkNonAquatic(lion)
	checkAnimal(lion)
}
func checkAquatic(a iAquatic)       {}
func checkNonAquatic(a iNonAquatic) {}
func checkAnimal(a iAnimal)         {}
