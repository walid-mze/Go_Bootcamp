package main

import "fmt"

func main(){
	var arr = [5]int{1,2,3,4,5}
	var arr2=[]int{}
	for i,x:=range arr{
		fmt.Println(i,x)
	}
	arr2=append(arr2,6)
	fmt.Println(arr)

	arr3:=arr[1:4]// len 4-1 cap end-1
	fmt.Println(len(arr3),cap(arr3))

	arr3=arr[:3]
	fmt.Println(len(arr3),cap(arr3))




}