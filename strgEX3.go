package main

import (
	"fmt"
)

func main() {
	var (
		f, c     string
		contador int
	)
	fmt.Println("Digite uma frase :")
	fmt.Scan(&f)
	fmt.Println("Digite um caractere para verificar se esta na frase :")
	fmt.Scan(&c)

	for i := 0; i < len(f); i++ {
		if c == string(f[i]) {
			contador++
		} else {
			continue
		}

	}
	fmt.Print(contador)
}
