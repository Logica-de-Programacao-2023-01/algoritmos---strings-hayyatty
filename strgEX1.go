package main

import (
	"fmt"
)

func main() {
	var t1, t2 string
	fmt.Print("Digite a primeira parte")
	fmt.Scan(&t1)
	fmt.Println("Digite a segunda parte")
	fmt.Scan(&t2)

	t3 := t1 + " " + t2
	fmt.Print(t3)
}
