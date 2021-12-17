package main

import "fmt"

func main() {
	var demoA int = 21
	var demoB int = 10
	var demoC int

	demoC = demoA + demoB
	fmt.Printf("C : %d", demoC)
	fmt.Println()

	demoC = demoA * demoB
	fmt.Printf("C : %d", demoC)
	fmt.Println()

	demoC = demoA - demoB
	fmt.Printf("C : %d", demoC)
	fmt.Println()

	demoC = demoA / demoB
	fmt.Printf("C : %d", demoC)
	fmt.Println()

	demoC = demoA % demoB
	fmt.Printf("C : %d", demoC)
	fmt.Println()

	demoA++
	fmt.Printf("A : %d", demoA)
	fmt.Println()

	demoA--
	fmt.Printf("A : %d", demoA)

}
