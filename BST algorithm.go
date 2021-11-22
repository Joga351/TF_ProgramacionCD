package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
)

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
type Persona struct {
	Index            int
	CAI              string
	Edad             int
	Trabajo          int
	Vinculo          int
	TipoVio          int
	ConsumeAlcohol   int
	Fuma             int
	ConsumeDroga     int
	Adiccion         int
	RiesgoPresuntivo int
	Mes              int
}

var count int
var data []Data
var ejemplo []Persona

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
}
func (n *Data) Insert(k Data) {
	if n.Index < data[k.Index].Index {
		//move right
		if n.Right == nil {
			n.Right = &data[k.Index-1]
		} else {
			n.Right.Insert(k)
		}
	} else if n.Index > data[k.Index].Index {
		//move left
		if n.Left == nil {
			n.Left = &data[k.Index-1]
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Data) SearchIndex(k Data) bool {
	if n == nil {
		return false
	}
	if n.Edad < k.Edad {
		//move right
		return n.Right.SearchIndex(k)
	} else if n.Edad > k.Edad {
		//move left
		return n.Left.SearchIndex(k)
	}
	return true
}

func (n *Data) SearchEdad(k int) bool {
	if n.Edad == k {
		nuevo := Persona{
			Index:            n.Index,
			CAI:              n.CAI,
			Edad:             n.Edad,
			Trabajo:          n.Trabajo,
			Vinculo:          n.Vinculo,
			TipoVio:          n.TipoVio,
			ConsumeAlcohol:   n.ConsumeAlcohol,
			Fuma:             n.Fuma,
			ConsumeDroga:     n.ConsumeDroga,
			Adiccion:         n.Adiccion,
			RiesgoPresuntivo: n.RiesgoPresuntivo,
			Mes:              n.Mes,
		}
		ejemplo = append(ejemplo, nuevo)
	}
	if n.Edad < k {
		//move right
		return n.Right.SearchEdad(k)
	} else if n.Edad > k {
		//move left
		return n.Left.SearchEdad(k)
	}
	return true
}

func main() {

	cargarData()
	tree := data[0]
	for i := 0; i <= 30; i++ {
		tree.Insert(data[i+1])
	}
	fmt.Println(tree)
}
