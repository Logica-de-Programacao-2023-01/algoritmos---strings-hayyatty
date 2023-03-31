package main

import "fmt"

func main() {
	var f string
	fmt.Println("Digite uma frase :")
	fmt.Scan(&f)
	fmt.Println("A sua frase tem :", len(f), "caractéres")
}
