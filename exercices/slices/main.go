package main
import "fmt"

func main(){
	a := []int {1,2,3,4}
	for i,x:=range(a){
		fmt.Println(i,x)
	}	
}