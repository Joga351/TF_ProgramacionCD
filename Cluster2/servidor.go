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
	var data Data
	err := gob.NewDecoder(C).Decode(&data)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		//fmt.Println(persona)
		if data.Nombre == "Cliente" {
			go EnviarAPI(data)

		} else {
			fmt.Println(data)
			go Enviar(data)
		}
	}
}

func Enviar(data Data) {
	C, err := net.Dial("tcp", ":9001")
	C1, err := net.Dial("tcp", ":9002")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(C).Encode(data)
	if err != nil {
		fmt.Println(err)
	}
	err = gob.NewEncoder(C1).Encode(data)
	if err != nil {
		fmt.Println(err)
	}
	C.Close()
	C1.Close()
}

func EnviarAPI(data Data) {
	C, err := net.Dial("tcp", ":9003")

	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(C).Encode(data)
	C.Close()
}
func main() {
	go Servidor()

	var input string
	fmt.Scanln(&input)
}
