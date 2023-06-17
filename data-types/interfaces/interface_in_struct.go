package interfaces

import (
	"fmt"
)

/*
-- Pointer vs Value Receiver in methods while implementing an interface in Go --
A method of a type can either have a pointer receiver or a value receiver.
There is a caveat while pointer vs value receiver for methods of a type when that type implements an interface

1. If a type implements all methods of an interface using value receiver, then both value and pointer of that type can be used
while assigning to that interface variable or while passing to a function which accept an argument as that interface.

2. If a type implements all methods of an interface using pointer receiver, then the only pointer of that type can be used
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
	IPerson
	Team string
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
