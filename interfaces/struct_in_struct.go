package interfaces

import "fmt"

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
