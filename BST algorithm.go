package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var count int
var data []Data

type Data struct {
	//numero del dato
	Index int `json:index`
	//features de la data
	Cai              string `json:cai`
	Edad             int    `json:edad`
	Trabajo          int    `json:trabajo`
	Vinculo          int    `json:vinculo`
	TipoVio          int    `json:tipovio`
	ConsumeAlcohol   int    `json:cAlcohol`
	Fuma             int    `json:fuma`
	ConsumeDroga     int    `json:cDroga`
	Adiccion         int    `json:adiccion`
	RiesgoPresuntivo int    `json:riesgo`
	Mes              int    `json:mes`
	Left             *Data
	Right            *Data
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// leemos csv del github
func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func cargarData() {
	// Lee csv del github :D
	url := "https://raw.githubusercontent.com/Joga351/TF_ProgramacionCD/main/Assets/casos_cai_2019.csv"
	datos, err := readCSVFromUrl(url)
	if err != nil {
		panic(err)
	}

	var file Data

	for idx, rec := range datos {
		// skipea la primera linea (header)
		if idx == 0 {
			continue
		}
		file.Index, _ = strconv.Atoi(rec[0])
		file.Cai = rec[1]
		file.Edad, _ = strconv.Atoi(rec[2])
		file.Trabajo, _ = strconv.Atoi(rec[3])
		file.Vinculo, _ = strconv.Atoi(rec[4])
		file.TipoVio, _ = strconv.Atoi(rec[5])
		file.ConsumeAlcohol, _ = strconv.Atoi(rec[6])
		file.Fuma, _ = strconv.Atoi(rec[7])
		file.ConsumeDroga, _ = strconv.Atoi(rec[8])
		file.Adiccion, _ = strconv.Atoi(rec[9])
		file.RiesgoPresuntivo, _ = strconv.Atoi(rec[10])
		file.Mes, _ = strconv.Atoi(rec[11])
		data = append(data, file)
	}
	log.Println(data)
}

func (n *Data) Insert(k int) {
	if n.Index < k {
		//move right
		if n.Right == nil {
			n.Right = &Data{Index: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Index > k {
		//move left
		if n.Left == nil {
			n.Left = &Data{Index: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Data) Search(k int) bool {
	count++
	if n == nil {
		return false
	}
	if n.Index < k {
		//move right
		return n.Right.Search(k)
	} else if n.Index > k {
		//move left
		return n.Left.Search(k)
	}
	return true
}

func main() {

	tree := &Data{Index: 5}
	//en tree en vez de index: 5 tiene que jalar el 5to objeto del data set
	tree.Insert(2)
	tree.Insert(6)
	fmt.Println(tree.Search(2))

}
