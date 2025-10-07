package main

import (
	"fmt"
	"strings"
)

func main() {
	// String Manipulação
	var nome string = "caue souza"
	fmt.Println(nome)
	// Upper
	fmt.Println("Grande:",strings.ToUpper(nome))
	// Lower
	fmt.Println("Pequeno:",strings.ToLower(nome))

	// titlecase
	fmt.Println("Titulo:", strings.Title(nome))

	// trim
	fmt.Println("Trim:", nome, "Algo")
	fmt.Println("")
	fmt.Println("Trim:", strings.TrimRight(nome, ""), "Algo")
	fmt.Println("")
	fmt.Println("Trim:", strings.TrimLeft(nome, ""), "Algo")


}