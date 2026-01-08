package main
import "fmt"
import "sync"

var mu sync.Mutex

func squareNumber(a int) int {

	result:=a*a
	fmt.Printf("The square of %d is %d\n",a,result)
	return result
}

func main(){
	list:=[]int{1,2,3,4,5,6,7,8,9,10}
	squares:=make([]int,10)
	var wg sync.WaitGroup

	for i:=range list{
		wg.Add(1)

		go func(){
			defer wg.Done()
			squares[i]=squareNumber(list[i])
		}()
	}
	wg.Wait()
	fmt.Println(squares)

}
