package main
import ("fmt"
"math/rand"
)
func main(){
	n := rand.Intn(100) 
	var limit int 
	fmt.Print("Eneter your  : ")
	fmt.Scanf("%d",&guess)
	var guess int
	i:=0
	for {
		fmt.Print("Eneter your guess : ")
		fmt.Scanf("%d",&guess)
		if guess>100 || guess<0{
			fmt.Println("your guess is out of range ")

		}else if(guess==n){
			fmt.Println("correct guess! ")
			fmt.Println("you made ",i+1,"attempts")
			break
		}else if guess<n{
			fmt.Println(" your guess is too low ")
			i++
		}else{
			fmt.Println(" your guess is too high ")
			i++
		}
	}

}