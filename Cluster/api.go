package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Persona struct {
	Nombre string
	Origen string
	Email  []string
}

func Api(persona Persona) {
	C, err := net.Dial("tcp", ":9000")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(C).Encode(persona)
	if err != nil {
		fmt.Println(err)
	}
	C.Close()
}

func main() {

	persona := Persona{
		Nombre: "Denilson",
		Origen: "API",
		Email: []string{
			"Prueba01",
			"Pruena02",
			"",
			""
		},
	}

	go Api(persona)
	var input string

	fmt.Scanln(&input)
}
