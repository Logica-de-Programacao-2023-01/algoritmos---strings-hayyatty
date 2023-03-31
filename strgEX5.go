package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		f string
		n int
	)
	fmt.Println("Digite algo ")
	fmt.Scan(&f)
	fmt.Println("Digite um numero")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		f1 := strings.ToUpper(string(f[i]))
		fmt.Print(f1)
	}

}


