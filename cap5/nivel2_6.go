package cap5

package main

import "fmt"

func main() {
	const (
		_ = 2024 + iota
		ano_2025
		ano_2026
		ano_2027
		ano_2028
	)
	fmt.Printf("%v\n%v\n%v\n%v", ano_2025, ano_2026, ano_2027, ano_2028)
}

