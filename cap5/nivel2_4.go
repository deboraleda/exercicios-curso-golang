package cap5

package main

import "fmt"

var x int = 10

func main() {
	fmt.Printf("%d\n%b\n%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\n%b\n%#x", y, y, y)
}
