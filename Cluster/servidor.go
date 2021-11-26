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
	var info INFO
	err := gob.NewDecoder(C).Decode(&info)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		//fmt.Println(persona)
		if info.Origen == "Cliente" {
			fmt.Println(info)
			go EnviarAPI(info)

		} else {
			fmt.Println(info)
			go Enviar(info)
		}
	}
}

func Enviar(info INFO) {
	C, err := net.Dial("tcp", ":9001")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(C).Encode(info)
	if err != nil {
		fmt.Println(err)
	}
	C.Close()

	C1, err := net.Dial("tcp", ":9002")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(C1).Encode(info)
	if err != nil {
		fmt.Println(err)
	}
	C1.Close()

}

func EnviarAPI(info INFO) {
	C, err := net.Dial("tcp", ":9003")

	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(C).Encode(info)
	C.Close()
}
func main() {

	info := INFO{
		Origen: "API",
		Datos:  "Prueba",
		Persona: []PERSONA{
			{Nombre: "Persona01", Edad: 25},
			{Nombre: "Persona02", Edad: 30},
		},
	}
	go Enviar(info)
	go Servidor()

	var input string
	fmt.Scanln(&input)
}
