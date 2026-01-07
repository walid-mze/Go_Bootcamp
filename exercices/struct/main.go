package main

type Person struct{
	Name string 
	email string 

}

func main(){
	p:=Person{"walid","walid@gmail.com"}
	fmt.Println(p.email,p.name)
	anotherStruct := struct {//use it once 
		firstName string 
		lastName string 
	}{"walid","mamze"}

}
