package main

import "fmt"

func main(){
a:=1
switch a {
case 1: 
fmt.Println("v is 1")
fallthrough
case 2: fmt.Println("v is 2")
default: fmt.Println("v is neither 1 nor 2")

}
}
