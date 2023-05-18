package pattern

import "fmt"



type Vehicle interface {
	Name()
	Accept(Visitor)
}

type Car struct {
}

func (c *Car) Name() string {
	return "Its car"
}

func (c *Car) Accept(v Visitor) {
	v.VisitCar(c)
}

type Plane struct {
}

func (p *Plane) Name() string {
	return "Its plane"
}

func (p *Plane) Accept(v Visitor) {
	v.VisitPlane(p)
}

type Visitor interface {
	VisitCar(c *Car)
	VisitPlane(p *Plane)
}

type MoveChecker struct {
}

func (m *MoveChecker) VisitCar(c *Car) {
	fmt.Println("the car is driving")
}

func (m *MoveChecker) VisitPlane(p *Plane) {
	fmt.Println("the plane is flying")
}

// func main() {
// 	car := &Car{}
// 	car.Accept(&MoveChecker{})

// 	plane := &Plane{}
// 	plane.Accept(&MoveChecker{})
// }
