package main

import "fmt"

func main(){
	var g1 int
	var g2 int
	var g3 int
	var g4 int
	var g5 int

	fmt.Print("Enter the first value: ")
	fmt.Scanf("%d",&g1)

	fmt.Print("Enter the second value: ")
	fmt.Scanf("%d",&g2)

	fmt.Print("Enter the third value: ")
	fmt.Scanf("%d",&g3)

	fmt.Print("Enter the fourth value: ")
	fmt.Scanf("%d",&g4)

	fmt.Print("Enter the fifth value: ")
	fmt.Scanf("%d",&g5)


	avg:=(float64)(g1+g2+g3+g4+g5)/5
	fmt.Println("the average grade is",avg)


	switch {
	case avg>=90: fmt.Println("The letter is A")
	case avg>=80: fmt.Println("The letter is B")
	case avg>=70: fmt.Println("The letter is C")
	case avg>=60: fmt.Println("The letter is D")
	case avg<60 : fmt.Println("The letter is F")
	default : 	fmt.Println("The letter is A")

	}



}



