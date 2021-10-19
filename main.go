package main

import "fmt"

func main() {
	names := [3]int{22, 24, 25}

	updateName(names)

	fmt.Println(names)
}

func updateName(names [3]int) {
	names[1] = 50
	fmt.Println(names)
}
