package Easy_Problems_to_Get_Started

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	arr := make([]int, 0)
	l := 1
	for i:=1;i<=int(n/2);i++{
		if n%i==0{
			l += 1
			arr = append(arr, i)
		}
	}
	fmt.Println(l)
	for _, e := range(arr){
		fmt.Print(e, " ")
	}
	fmt.Print(n)
}