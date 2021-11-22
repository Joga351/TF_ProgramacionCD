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

func Servidor() {
	S, err := net.Listen("tcp", ":9000")
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
		go Manejador(C)
	}
}

func Manejador(C net.Conn) {
	var persona Persona
	err := gob.NewDecoder(C).Decode(&persona)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		//fmt.Println(persona)
		if persona.Nombre == "Cliente" {
			fmt.Println(persona)
		} else {
			go Enviar(persona)
		}
	}
}

func Enviar(persona Persona) {
	C, err := net.Dial("tcp", ":9001")
	C1, err := net.Dial("tcp", ":9002")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(C).Encode(persona)
	if err != nil {
		fmt.Println(err)
	}
	err = gob.NewEncoder(C1).Encode(persona)
	if err != nil {
		fmt.Println(err)
	}
	C.Close()
	C1.Close()
}
func main() {
	go Servidor()

	var input string
	fmt.Scanln(&input)
}
