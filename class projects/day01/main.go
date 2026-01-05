package main
import ("fmt"
"math/rand"
)
func main(){
	var limit int 
	fmt.Print("Enter the number of tries : ")
	fmt.Scanf("%d\n",&limit)

	var guess int
	var bestScore int

	play:='y'

	for play=='y' {
	var diff string
	min:=1
	var max int 
	fmt.Print("Enter the difficulty(easy,medium,hard): ")
	fmt.Scanf("%s\n",&diff)
	if diff=="easy"{
		max=50
	}else if diff == "medium"{
		max=100
	}else if diff=="hard"{
		max=200
	}else {
		fmt.Print("invalid input")
		break
	}
	n :=min+ rand.Intn(max-min+1) 
	i:=0
	win:=0

	for i<limit {
		fmt.Println("Enter your guess : ")
		fmt.Scanf("%d\n",&guess)
		if guess>max || guess<min {
			fmt.Println("your guess is out of range ")

		}else if(guess==n){
			fmt.Println("You win , correct guess ")
			fmt.Println("you made ",i+1,"attempts")
			win=1
			if i+1>bestScore{
				bestScore=i+1
			}

			break
		}else if guess<n{
			fmt.Println(" your guess is too low ")
			i++
		}else{
			fmt.Println(" your guess is too high ")
			i++
		}
	}
	if win==0{
	fmt.Println("You lost")
	}
	fmt.Println("You want to play again?(y/n): ")
	fmt.Scanf("%c\n",&play)
	



}


}