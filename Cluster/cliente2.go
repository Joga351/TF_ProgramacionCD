package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Data struct {
	Nombre   string
	Origen   string
	Busqueda []string
	Datos    []int
}

func Send(data Data) {
	C, err := net.Dial("tcp", ":9000")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(C).Encode(data)
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
	var data Data
	err := gob.NewDecoder(C).Decode(&data)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(data)
		data.Nombre = "Cliente"
		data.Origen = "Cliente02"
		// llamamos al algoitmo y le enviamos los datos de busqueda
		//
		go Send(data)
	}
}

func main() {

	go Cliente()
	var input string

	fmt.Scanln(&input)
}
