package structs

import "fmt"

/*
#### Reference Links: ####
https://golangbyexample.com/
https://eli.thegreenplace.net/2020/embedding-in-go-part-1-structs-in-structs/
https://eli.thegreenplace.net/2018/beware-of-copying-mutexes-in-go/
https://eli.thegreenplace.net/2020/embedding-in-go-part-3-interfaces-in-structs/


####### Embedding :	Compostion rather than Inheritance ########

Embedding can be done in 2 ways:
1. Named or Direct
2. Un-named/Anonymous or Variable

type Interface interface {
   Method()
}

1. Direct :
- Maintains same access type for public/private methods & attributes
- Method set of interface is part of this struct.
type Type struct  {
  Interface
}

2. Variable :
-  Access type can be changed/managed for public/private methods & attributes.
type Type struct  {
  Itf Interface
}

########## Combination of Interfaces/Structs ############
1. structs in struct
https://eli.thegreenplace.net/2020/embedding-in-go-part-1-structs-in-structs/"
2. interfaces in interface
3. interfaces in struct
https://eli.thegreenplace.net/2020/embedding-in-go-part-3-interfaces-in-structs/"

Embedding :
Helps achieve composition easily

Inheritance :
Is relational in nature (is-a).
It has Diamond Problem (in Multiple Inheritance)

Composition :
Is not relational, looks at things as composition of stuff (has-a).
It has method forwarding drawback.

1. structs in structs
	Promoted Fields - Taken from Embedded Struct
	Embedding of Methods
	Shadowing of Embedded Fields
2. inteface in interface

3. interface in struct
	Although it is implementing interface
	However its different in terms of how it implements
	As we know that by default interfaces are implicit. ie - If any type defines all the methods of an interface, it implemets the interface
	However, in case interface is embedded in struct, all the methods of the interface are already  part of ANY type which implements that interface
	If required, only required methods can be overridden accordingly
*/

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (p *Person) PrintDetails() {
	fmt.Printf("%s : %v/%v", p.Name, p.Age, p.Gender)
}

/*
As this method is private, it cant be accessed by any struct embedding the Person struct
*/
func (p *Person) getDetails() (pr *Person) {
	pr = p
	return
}

type Employee struct {
	/*
		P is exportable as it is capital
		So we can control the access of Person Object within Employee with this P variable
		However its by default exportable in Manager struct.
		Whichever fields or methods are exportable in Person, will be exportable in Manager as well
	*/
	P Person
}
type PersonalSecretory struct {
	p Person
}

type Manager struct {
	Person
}

func ExampleBasics() {
	m := Manager{Person: Person{"Manager", 40, "Male"}}
	m.PrintDetails()

	e := Employee{P: Person{"Employee", 30, "Male"}}
	e.P.PrintDetails()

	s := PersonalSecretory{p: Person{"Secretory", 20, "Female"}}
	s.p.PrintDetails()
}
