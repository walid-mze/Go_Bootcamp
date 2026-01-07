package main
import "fmt"
import "math"



type Circle struct {
	Radius float64
}

func (c Circle) calculateArea() float64{
	return math.Pi*c.Radius*c.Radius
}

 func (c Circle ) calculateCircumference() float64{
	return 2*math.Pi*c.Radius
}

func (c *Circle) scale(factor int) float64{
	return c.Radius*float64(factor)
}

func main(){
	circle:=Circle{3.0}
	fmt.Println(circle.calculateArea())
	fmt.Println(circle.calculateCircumference())


}