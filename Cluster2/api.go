package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type INFO struct {
	Nombre   string
	Origen   string
	Busqueda []string
	Datos    []int
}

func Api(persona INFO) {
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

func Send(data INFO) {
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
func Receptor() {
	S, err := net.Listen("tcp", ":9003")
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
		go ManejadorReceptor(C)
	}
}

func ManejadorReceptor(C net.Conn) {
	var data Data
	err := gob.NewDecoder(C).Decode(&data)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(data)
		// llamamos al algoitmo y le enviamos los datos de busqueda

	}
}

func main() {

	data := INFO{
		Nombre: "Denilson",
		Origen: "API",
		Busqueda
		Datos:  []int{},
	}

	go Api(data)

	go Receptor()

	var input string

	fmt.Scanln(&input)
}
