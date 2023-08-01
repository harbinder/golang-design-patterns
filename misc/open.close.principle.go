package misc

/*
Refer: https://towardsdev.com/golang-solid-principles-fd7bf513874d
*/

type square struct {
	side float32
}

type triangle struct {
	base   float32
	height float32
}

func (s square) area() float32 {
	return s.side * s.side
}

func (t triangle) area() float32 {
	return 1 / 2 * t.base * t.height
}

/*
calculate type doesnt use Open/Close Principle
If any new type is introduced and its area needs to be calculated, then we will be required to
MODIFY the existing calculate struct, which shoud not be the case according to SOLID principles
*/
type caclulator struct {
	area float32
}

func (c *caclulator) calculateArea(shapes interface{}) {
	switch shapes.(type) {
	case square:
		c.area = shapes.(square).area()
	case triangle:
		c.area = shapes.(triangle).area()
	}
}

// Instead we should use the below code to calculate area to satisfy SOLID princliple Open/Close

// Both square and triangle Structs types implement shape interface
type shape interface {
	area() float32
}

func (c *caclulator) calculateAreaSolid(shapes shape) {
	c.area = shapes.area()
}
