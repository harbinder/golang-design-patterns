package misc

import "fmt"

/*
In the code above we can analyze the following:

The code defines the “vehicle” type that defines the interface method “getVehicleName”.
“Car” and “motorcycle” has access to “vehicle” using composition, which means that apart to get access to the properties,
the car and motorcycle have access to the “vehicle” methods.

The printer method will work with vehicles, cars, and motorcycles.

In a conclusion, we can say the following, the “Liskov substitution principle” is applied because the “car” type and “motorcycle”
type could be substituted by the “vehicle” type.

*/

type transport interface {
	getVehicleName() string
}

type vehicle struct {
	name string
}

// vehicle concrete type implements transport interface
func (v vehicle) getVehicleName() string {
	return v.name
}

// using composition, create new concrete vehicle types
type car struct {
	vehicle // direct embedding
	wheel   int
	door    int
}

type motorcycle struct {
	vehicle // direct embedding
	wheel   int
}

// common method to get vehicle name
type Printer struct{}

func (p Printer) printVehicleName(t transport) {
	fmt.Println("Vehicle Name: ", t.getVehicleName())
}

func LiskovSubstitutionExample() {
	car1 := car{
		vehicle: vehicle{name: "Alto"},
		wheel:   4,
		door:    4,
	}
	motorcycle1 := motorcycle{
		vehicle: vehicle{name: "royal enfield"},
		wheel:   2,
	}
	p := Printer{}

	// we should be able to use memthod of vehicle concrete type for car and motorcycle concrete types,
	// if they have the same(implement or access via embedding) type ie transport interface type
	p.printVehicleName(car1)
	p.printVehicleName(motorcycle1)
}
