package main

import "fmt"
import "math"



type Shape interface {
	GetArea() float64
	GetPerimeter() float64

}

type Rectangle struct{
	Width, Height float64
}
type Circle struct{
	Radius float64
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
	return 2*c.Radius*math.Pi
}
func PrintShapeInf(s Shape){
	fmt.Println("Area: ",s.GetArea())
	fmt.Println("Perimeter: ",s.GetPerimeter())

}

func main(){

	c:=Circle{2}
	r:=Rectangle{3,4}
	PrintShapeInf(c)
	PrintShapeInf(r)



}


