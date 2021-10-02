package Easy_Problems_to_Get_Started

import "fmt"

func main(){
	var x int
	fmt.Scan(x)

	five := x%5==0
	eleven := x%11==0

	if five && eleven{
		fmt.Println("TWO")
	} else if five || eleven{
		fmt.Println("ONE")
	} else {
		fmt.Println("NONE")
	}
}