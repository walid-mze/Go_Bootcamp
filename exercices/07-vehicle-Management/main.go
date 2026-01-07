package main
import "fmt"


type Vehicle struct{
	Make string 
	Model string
	Year  int 
}

type Insurable interface{
	calculateInsurance() int
}
type Printable interface{
	PrintDetails()
}
type Car struct{
	Vehicle 
	NumberOfDoors int
}
type Truck struct{
	Vehicle 
	PlayLoadCapacity int
}

func PrintAll(p []Printable){
	for i:=range p{
		p[i].PrintDetails()
	}
}
func (c Car) calculateInsurance() int{
	return 1000*c.NumberOfDoors
}
func (t Truck) calculateInsurance() int {
	return 2000*t.PlayLoadCapacity
}
func (c Car) PrintDetails(){
	fmt.Println("Car details: ")
	fmt.Println("Make: ",c.Make)
	fmt.Println("Year: ",c.Year)
	fmt.Println("Model: ",c.Model)
	fmt.Println("Insurance: ",c.calculateInsurance())




}
func (t Truck) PrintDetails(){
	fmt.Println("Truck details: ")
	fmt.Println("Make: ",t.Make)
	fmt.Println("Year: ",t.Year)
	fmt.Println("Model: ",t.Model)
	fmt.Println("Insurance: ",t.calculateInsurance())
}



func main(){
	t:=Truck{Vehicle:Vehicle{"Scania","V202",2003},PlayLoadCapacity:15}
	c:=Car{Vehicle:Vehicle{"Mercedes","Gle",2020},NumberOfDoors:4}
	
	vehicles:=[]Printable{t,c}
	PrintAll(vehicles)



}
