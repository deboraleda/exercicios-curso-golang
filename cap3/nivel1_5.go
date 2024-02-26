package cap3

import "fmt"

var x1_5 hotdog

var y1_5 int

func main1_5() {
	fmt.Println(x1_5)
	fmt.Printf("%T\n", x1_5)
	x1_5 = 42
	fmt.Println(x1_5)

	//

	y1_5 = int(x1_5)
	fmt.Printf("%T\n", y1_5)
	fmt.Println(y1_5)
}
