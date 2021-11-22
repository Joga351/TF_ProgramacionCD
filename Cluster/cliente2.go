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

func Send(persona Persona) {
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
func Cliente() {
	S, err := net.Listen("tcp", ":9002")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		C, err := S.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go ManejadorCliente1(C)
	}
}

func ManejadorCliente1(C net.Conn) {
	var persona Persona
	err := gob.NewDecoder(C).Decode(&persona)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(persona)
		persona.Nombre = "Cliente"
		persona.Origen = "Cliente02"
		go Send(persona)
	}
}

func main() {

	go Cliente()
	var input string

	fmt.Scanln(&input)
}
