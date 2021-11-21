package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Data struct {
	//numero del dato
	Index int `json:"Numero de orden"`
	//features de la data
	CAI              string `json:"Distrito"`
	Edad             int    `json:"Edad"`
	Trabajo          int    `json:"Trabaja"`
	Vinculo          int    `json:"Cantidad de veces vinculado con delitos"`
	TipoVio          int    `json:"Tipo de Violencia"`
	ConsumeAlcohol   int    `json:"Cantidad de veces detenido por consumo de alcohol"`
	Fuma             int    `json:"Veces detenido por fumar en lugares cerrados"`
	ConsumeDroga     int    `json:"Veces detenido por consumo de drogas"`
	Adiccion         int    `json:"Adiccion no convencional"`
	RiesgoPresuntivo int    `json:"Riesgo de presunto delito"`
	Mes              int    `json:"Mes de ocurrencia"`
}

//globar
var data []Data

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
	url := "https://raw.githubusercontent.com/abad2016/TFConcurrente/main/Assets/casos_cai_2019.csv"
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
		file.CAI = rec[1]
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

//Funciones resuleve los requests
func listarData(res http.ResponseWriter, req *http.Request) {
	log.Println("Llamada al endpoint /listarData")
	//estableceer el tipo de contenido que se devuelve
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	//logica de manejo de listar
	//codificar a json
	jsonBytes, _ := json.MarshalIndent(data, "", " ")
	io.WriteString(res, string(jsonBytes))
}

func buscarDato(res http.ResponseWriter, req *http.Request) {
	log.Println("Llamada al endpoint /buscarDato")
	//estableceer el tipo de contenido que se devuelve
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	//recuperar el parametro enviado
	sIndex := req.FormValue("Index")
	iIndex, _ := strconv.Atoi(sIndex)

	//logica de la funcion
	for _, dat := range data {
		if dat.Index == iIndex {
			//codificacion
			ojsonBytes, _ := json.MarshalIndent(dat, "", " ")
			io.WriteString(res, string(ojsonBytes))
		}
	}
}

func home(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res,
		`<doctype html>
	<html>
	<head>
		<title>Trabajo Final Programacion Concurrente y Distribuida</title>
	</head>
	<body>
		<link rel= "stylesheet" type="text/css" href="front.css">
		<h2>Delitos registrados en Lima</h2>
		<h3> Integrantes del equipo </h3>
		<h4> -Fernando </h4>
		<h4> -Juben </h4>
		<h4> -Josue </h4>
		<form>
			<button> Ver lista de datos </button>
			<button > Buscar dato </button>
			<button> Quienes somos </button>
		</form>

		<a href="/busca">Busca</a>
		
	</body>	
	</html>
	`)
}

func handleRequest() {
	//declarar los endpoints
	http.HandleFunc("/listarData", listarData)
	http.HandleFunc("/buscarDato", buscarDato)
	http.HandleFunc("/home", home)

	//puesto esucha del servicio
	log.Fatal(http.ListenAndServe(":3505", nil))
}

func main() {

	cargarData()
	handleRequest()
	//Ingresar el siguiente URL en el buscador: http://localhost:3505/home

}
