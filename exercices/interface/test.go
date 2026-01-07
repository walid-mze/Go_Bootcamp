package main

import (
	"fmt"
	"log"
)
func square(x int ) int {
	return x*x
}
func main(){
	var x any // empty interface 
	x=1
	//x="hh"
	xInt,ok:=x.(int) //check if x is int

	if !ok{
		log.Fatal("x is not an int")
	}
	fmt.Println(square(xInt))

	switch value:=x.(type){
	case int : 
		value+=10
	case string :
		value+="hello"
	case float64:
		value+=value*3.5
	default:
		log.Fatal("invalue")

	}
	

}
