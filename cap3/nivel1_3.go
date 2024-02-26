package cap3

import "fmt"

var x1_3 int = 42
var y1_3 string = "James Bond"
var z1_3 bool = true

func main1_3() {
	s := fmt.Sprintf("%v %v %v", x1_3, y1_3, z1_3)
	fmt.Println(s)
}
