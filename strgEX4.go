package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		f string
	)
	fmt.Print("Digite um frase")
	fmt.Scan(&f)
	fmt.Println(strings.ToUpper(f))
}
