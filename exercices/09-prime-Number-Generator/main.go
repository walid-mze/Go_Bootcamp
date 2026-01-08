package main
import "fmt"
import "math"


func isPrime(a int) bool{
	if a<=1{
		return false
	}
	for i:=2;i<=int(math.Sqrt(float64(a)));i++{
		if a%i==0{
			return false
		}
	}
	return true
}

func generatePrime(n int, c chan<- int){
	defer close(c)

	for i:=0;i<n;i++ {
		if isPrime(i){
			c<-i
		}
	}

}
func printPrime(ch <-chan int,done chan<- bool){


	for prime :=range ch{
		fmt.Println("Prime: ",prime)
	}
	done<-true


	
}


func main(){
	ch:=make(chan int)
	done:=make(chan bool)
	

	go generatePrime(50,ch)
	go printPrime(ch,done)


	_=<-done

}
