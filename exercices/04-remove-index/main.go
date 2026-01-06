package main

import (
	"errors"
	"fmt"
)




func remove(s []int , index  int ) ([]int, error) {
	if index>len(s)-1 || index<0 {
		return []int{},errors.New("index out of range")
	}else{
		var arr=[]int{}
		for i:=0;i<index;i++{
			arr=append(arr, s[i])
		}
		for i:=index+1;i<len(s);i++{
			arr=append(arr, s[i])
		}
		return arr,nil
	}
}
func main(){
	var arr = []int{1,2,3,4,5}
		fmt.Println(arr)

	arr,err:=remove(arr,-1)
	fmt.Println(arr)
	fmt.Println(err)

	m:= map[string]int{
		"one":1,
		"two":2,
		"three":3,
	}
	fmt.Println(m)

}