package main

import "fmt"

func main2(){
	var g1 int

	i:=0
	sum:=0
	for{
		
		fmt.Print("enter next value: ");
		fmt.Scanf("%d",&g1)
		if g1==-1{
			break
		}
		sum+=g1
		i++;
	}
	avg:=(float64)(sum)/5
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



