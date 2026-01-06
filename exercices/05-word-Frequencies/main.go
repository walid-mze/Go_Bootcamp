package main
import "fmt"


func countFreq(arr []string) map[string]int {
   var res =map[string]int{}
	for i:=range arr{
		_,ok:=res[arr[i]]
		if ok==true {
			res[arr[i]]++
		}else{
			res[arr[i]]=1
		}
	}
	return res

}


func main() {
		var arr = []string{"walid","mamze","walid","mm","mamze","walid"}
		res:=countFreq(arr)
		fmt.Println(res)
}