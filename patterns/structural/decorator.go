package structural

import (
	"fmt"
)

/*
MargerettaPizza & FarmhousePizza implements iPizza interface
*/
type iPizza interface {
	getPrice() int
}

type MargerettaPizza struct {
}

func (mp *MargerettaPizza) getPrice() (prc int) {
	prc = 100
	return
}

type FarmhousePizza struct {
}

func (fp *FarmhousePizza) getPrice() (prc int) {
	prc = 200
	return
}

/*
To add more features to these (MargerettaPizza & FarmhousePizza)
OR to decorate these and get new pizza types, we can do the following:
We have done 2 things:
- Implemented the iPizza interface
- ALSO embedded the same
This way, we have not changed the existing types (MargerettaPizza & FarmhousePizza),
Instead we hahve inhanced/decorated them
*/

type MargerettaPizzaWithToppings struct {
	iPizza iPizza
}

func (mpwt *MargerettaPizzaWithToppings) getPrice() (prc int) {
	prc = mpwt.iPizza.getPrice() + 20
	return
}

type FarmhousePizzaWithToppings struct {
	iPizza iPizza
}

func (fpwt *FarmhousePizzaWithToppings) getPrice() (prc int) {
	prc = fpwt.iPizza.getPrice() + 50
	return
}

func ExecuteDecorator() {
	mp := &MargerettaPizza{}
	fp := new(FarmhousePizza)

	mpwt := &MargerettaPizzaWithToppings{}
	mpwt.iPizza = mp

	fpwt := &FarmhousePizzaWithToppings{}
	fpwt.iPizza = fp

	fmt.Println("-- Without Toppings --")
	fmt.Println("MargerettaPizza : ", mp.getPrice())
	fmt.Println("FarmhousePizza : ", fp.getPrice())

	fmt.Println("-- With Toppings --")
	fmt.Println("MargerettaPizza : ", mpwt.getPrice())
	fmt.Println("FarmhousePizza : ", fpwt.getPrice())
}
