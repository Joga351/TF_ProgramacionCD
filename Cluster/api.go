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

func Enviar(info INFO) {
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
func Api() {
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
		go ManejadorApi(C)
	}
}

func ManejadorApi(C net.Conn) {
	var info INFO
	err := gob.NewDecoder(C).Decode(&info)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(info)
	}
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
	go Api()
	var input string
	fmt.Scanln(&input)
}
