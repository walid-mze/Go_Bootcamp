package main

import (
	"fmt"
	"math"
)


type Shape interface{
	GetArea() float64
	GetPerimeter() float64

}
type Rectangle struct{
	Width,Height float64
}
type Circle struct{
	Radius float64
}
type Triangle struct{
	SideA,SideB,SideC float64
}
func (r Rectangle) GetArea() float64{
	return r.Width*r.Height
}
func (r Rectangle) GetPerimeter() float64{
	return 2*(r.Width+r.Height)
}

func (c Circle) GetArea() float64{
	return math.Pi*c.Radius*c.Radius
}
func (c Circle) GetPerimeter() float64{
	return 2*math.Pi*c.Radius
}


func (t Triangle) GetArea() float64{
	return 0.5*t.SideA*t.SideB
}
func (t Triangle) GetPerimeter() float64{
	return t.SideA+t.SideB+t.SideC
}
func PrintShapeInfo(s interface{}) {
	switch value := s.(type) {
	case Shape:
		fmt.Println("Area:", value.GetArea())
		fmt.Println("Perimeter:", value.GetPerimeter())
	default:
		fmt.Println("error :Type does not implement the Shape interface")
	}
}



func main(){
	p:=Rectangle{3,4}
	c:=Circle{3}
	t:=Triangle{1,2,3}
	fmt.Println("Rectangle  ")
	PrintShapeInfo(p)
	fmt.Println("Circle  ")
	PrintShapeInfo(c)
	fmt.Println("Triangle  ")
	PrintShapeInfo(t)

}



