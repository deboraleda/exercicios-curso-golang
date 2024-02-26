package cap5

package main

import "fmt"

func main() {
	x := 10
	y := 15
	igual := x == y
	diferente := x != y
	maior := x > y
	maior_igual := x >= y
	menor := x < y
	menor_igual := x <= y
	fmt.Printf("%t\n", igual)
	fmt.Printf("%t\n", diferente)
	fmt.Printf("%t\n", maior)
	fmt.Printf("%t\n", maior_igual)
	fmt.Printf("%t\n", menor)
	fmt.Printf("%t\n", menor_igual)
}
