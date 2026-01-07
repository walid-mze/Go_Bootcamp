package main

import "fmt"
func incrementer() func() int {
	i:=0
	return func() int{
		i++
		return i
	}

}
func main(){
	inc:=incrementer()
	fmt.Println(inc())
	fmt.Println(inc())


}