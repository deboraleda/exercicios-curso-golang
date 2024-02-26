package cap3

import "fmt"

type hotdog int

var x1_4 hotdog

func main1_4() {
	fmt.Println(x1_4)
	fmt.Printf("%T", x1_4)
	x1_4 = 42
	fmt.Println()
	fmt.Println(x1_4)
}
