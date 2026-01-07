package main

import (
	"fmt"
	"time"
)

func doNothing() {
}

func main() {
	t1:=time.Now()
	for range(1000000000){
		doNothing()
	}
	t2:=time.Now()
	fmt.Println(t2.Sub(t1))


	

}