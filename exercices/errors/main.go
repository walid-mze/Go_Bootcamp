package main


import (
	"log"
	"fmt"
	"errors"
)
func dividde(a,b float64)(float64,error){
	if b==0{
		return 0,errors.New("error occured")
	}else {
			return a/b,nil
	}
}

func main(){
	result,err:=dividde(5,0)
	if(err!=nil){
		log.Fatal(err)
	}
	fmt.Println("Result:",result)
}