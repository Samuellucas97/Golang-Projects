package main

import (
	"errors"
	"fmt"
	"strconv"
)

// The Golang is strongly typed

func main() {
	var firstName string  // Creating the variable
	firstName="Samuel" // Atributing to the variable

	lastName := "Ferino" // Creating and atributing

	looping()

	creatingStruct()

	fmt.Println("Hello " + firstName + " " + lastName)
}

func looping() {

	/** Creating infinite loop
	for {
		fmt.Println("Infinite loop....")
	} */

	for i:=0;i<10;i++ {
		iString := strconv.Itoa(i)
		fmt.Print("[i=" + iString + "] " )
	}

	fmt.Println()
}

func creatingStruct() {
	pessoa := Pessoa{
		"Samuel",
		210,
	}

	fmt.Println(pessoa)

	idadeString := strconv.Itoa(pessoa.idade)
	fmt.Println(pessoa.nome + " " + idadeString)

	andou(pessoa)

	pessoa2 := Pessoa{
		"Daniel",
		210,
	}

	andou(pessoa2)
}

func andou(pessoa Pessoa) (string, error) {
	if pessoa.nome != "Samuel" {
		return "", errors.New("Pessoa não é Samuel")
	}

	return pessoa.nome + " andou", nil
}


type Pessoa struct {
	nome string
	idade int
}


