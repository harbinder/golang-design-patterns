package interfaces

import (
	"fmt"
)

/*
Reference Link: https://golangbyexample.com/interface-in-golang/

Interface Variable

	___________|_______
	|				  |

Interface Type		Interface Value

			___________|_______
			|				  |
	Concrete Type		Concrete Value

%T : used to print the concrete type of the interface value
%v : used to print the concrete value of the interface value.

### Empty Interface ###
An empty interface has no methods , hence by default all concrete types implement the empty interface.
If you write a function that accepts an empty interface then you can pass any type to that function.

##### WHY Interfaces ?
Below are some benefits of using interface.
1. Helps write more modular and decoupled code between different parts of codebase –
It can help reduce dependency between different parts of codebase and provide loose coupling.
2. Interface can be used to achieve run time polymorphism in golang.
RunTime Polymorphism means that a call is resolved at runtime.

###### Pointer vs Value Receiver in methods while implementing an interface in Go ######

A method of a type can either have a pointer receiver or a value receiver.
There is a caveat while pointer vs value receiver for methods of a type when that type implements an interface

1. If a concrete type implements all methods of an interface(Method-set) using value receiver,
then BOTH value and pointer of that type can be used
while assigning to that interface variable or while passing to a function which accept an argument as that interface.

2. If a type implements all methods of an interface(Method-set) using pointer receiver,
then the ONLY pointer of that type can be used
while assigning to that interface variable or while passing to a function that accepts an argument as that interface.

-- Access Concrete/Underlying types of interface variable --
The underlying variable of an interface can be accessed in two ways

 1. Type Assertion
    if val, ok := interfaceVariable.(concreteType); ok == true {

    }

 2. Type Switch
    func print(i interfaceVariable) {
    switch v := i.(type) {
    case concreteType_1:
    fmt.Println("Type: concreteType_1")
    case concreteType_2:
    fmt.Println("Type: concreteType_2")
    default:
    fmt.Printf("Unknown Type %T", v)
    }
    }
*/
type IPerson interface {
	PrintPerson()
	SetPerson(name, gender string)
}
type SportsPerson struct {
	Name   string
	Gender string
}

func (sp *SportsPerson) SetPerson(name, gender string) {
	sp.Name = name
	sp.Gender = gender
}
func (sp *SportsPerson) PrintPerson() {
	fmt.Printf("Name: %s | Gender: %s", sp.Name, sp.Gender)
}

type ICricketer interface {
	PrintCricketer(p IPerson)
	SetCricketer(name, gender, team string)
}
type Cricketer struct {
	IPerson // embedded interface
	Team    string
}

func (c Cricketer) PrintCricketer(p IPerson) {
	p.PrintPerson()
	//c.IPerson.PrintPerson()
	fmt.Printf("Team: %s\n", c.Team)
}

func (c Cricketer) PrintCricketer1() {
	//p.PrintPerson()
	c.IPerson.PrintPerson()
	fmt.Printf("Team: %s\n", c.Team)
}

func (c *Cricketer) SetCricketer(name, gender, team string) {
	c.IPerson.SetPerson(name, gender) // or c.SetPerson(name, gender)
	c.Team = team
}

func ExampleInterfaceInStruct() {

	/*
		create an object of SportsPerson type
		BUT assigned to an interface type which is implemented by SportsPerson
		We can assign it in 2 ways
		1. Value Type
			ip = SportsPerson{}
		2. Pointer Type
			ip = &SportsPerson{}
		If assigned as Value DataType(ip=sp),
			ip will NOT BE able to access Pointer Receiver Methods of SportsPerson struct
		If assigned as Pointer DataType(ip=&sp),
			ip will BE able to access BOTH Pointer & Value Receiver Methods of SportsPerson struct

	*/
	var ip IPerson
	// Memory Map -> ip : [nil,nil]
	sp := SportsPerson{Name: "Sachin Tendulkar", Gender: "Male"}
	// Memory Map -> sp : [Name, Gender]
	//ip = sp
	/*
		Memory Map : ip =
		[
			Pointer -> sp,
			Pointer  -> SportsPerson.PrintPerson()
		]
	*/

	ip = &sp
	fmt.Printf("Underlying Type: %T\n", ip)
	fmt.Printf("Underlying Value: %v\n", ip)
	/*
		Memory Map : ip =
		[
			Pointer -> Pointer -> sp,
			Pointer  -> SportsPerson.PrintPerson()
		]
	*/

	// 2 Ways to Create an object of Cricketer type
	/*
		1. cr := new(interfaces.Cricketer) - cr is Pointer type struct variable
		   Memory Map : cr -> [Name: "", Gender: "", Team: ""]
		2. cr := interfaces.Cricketer{} - cr is Value type struct variable
	*/

	//cr := new(interfaces.Cricketer)
	cr := Cricketer{}
	// Memory Map : cr = [Name: "", Gender: "", Team: ""]
	cr.Team = "Mumbai Indians"
	// Memory Map : cr = [Name: "", Gender: "", Team: "Mumbai Indians"]

	/*
		cr.PrintPerson()

		Cricketer struct can access PROMOTED Method -> PrintPerson()  of SportsPerson struct,
		without Cricketer struct implementing methods of IPerson interface,
		which would have to be defined, in case Cricketer struct implicitly implements Iperson interface by defining all methods of IPerson


	*/
	fmt.Println("cr.PrintCricketer() call using object of type -> Cricketer struct ")
	cr.PrintCricketer(ip)

	var crI ICricketer
	/*
		Memory Map : crI =
		[
			Pointer -> nil,
			Pointer -> nil
		]
	*/
	crI = &cr
	fmt.Printf("Underlying Type: %T\n", crI)
	fmt.Printf("Underlying Value: %v\n", crI)
	// crI = cr
	//  - Compile Error :
	//  cr (Cricketer) does not implement ICricketer
	//  Because Cricketer struct has implemented interface method via Pointer Receiver AND
	// a Value Type struct object CANNOT access Pointer Receiver Methods HOWEVER vice versa is accessible i.e
	// a Pointer Type struct object CAN  access Value Receiver Methods
	/*
		Memory Map : crI =
		[
			Pointer -> cr,
			Pointer -> cr.PrintCricketer()
		]
	*/
	fmt.Println("crI.PrintCricketer() call using variable of type -> ICricketer interface ")
	crI.PrintCricketer(ip)

	cr1 := new(Cricketer)
	// cr1 is Pointer type struct (Cricketer) object, AND is able to access BOTH
	// Value & Pointer Receiver Methods of struct (Cricketer)

	cr1.IPerson = ip
	cr1.PrintCricketer(ip)
	cr1.SetCricketer("Virat Kohli", "Male", "RCB")
	cr1.PrintCricketer(ip)

	cr1.PrintCricketer1()
}
