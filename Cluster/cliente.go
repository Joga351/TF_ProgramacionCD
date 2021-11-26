package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type INFO struct {
	Origen  string
	Datos   string
	Persona []PERSONA
}
type PERSONA struct {
	Nombre string
	Edad   int
}

func Send(info INFO) {
	C, err := net.Dial("tcp", ":9000")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(C).Encode(info)
	if err != nil {
		fmt.Println(err)
	}
	C.Close()
}
func Cliente() {
	S, err := net.Listen("tcp", ":9001")
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
	var info INFO
	err := gob.NewDecoder(C).Decode(&info)

	if err != nil {
		fmt.Println(err)
		return
	} else {

		fmt.Println(info)
		info.Origen = "Cliente"
		info.Persona[0].Nombre = "PersonaCliente"
		info.Persona = append(info.Persona, PERSONA{Nombre: "PersonaPrueba", Edad: 30})
		go Send(info)
	}
}

func main() {

	go Cliente()
	var input string

	fmt.Scanln(&input)
}
