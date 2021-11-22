package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

//globar
var data []Data

type Data struct {
	//numero del dato
	Index int `json:"index"`
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
}

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

//Funciones resuleve los requests
func listarData(res http.ResponseWriter, req *http.Request) {

	enableCors(&res)
	log.Println("Llamada al endpoint /listarData")
	//estableceer el tipo de contenido que se devuelve
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	//logica de manejo de listar
	//codificar a json
	jsonBytes, _ := json.MarshalIndent(data, "", " ")
	io.WriteString(res, string(jsonBytes))
}

func buscarDato(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
	log.Println("Llamada al endpoint /buscarDato")
	//estableceer el tipo de contenido que se devuelve
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	//recuperar el parametro enviado
	sEdad := req.FormValue("Edad")
	iEdad, _ := strconv.Atoi(sEdad)

	//logica de la funcion
	for _, dat := range data {
		if dat.Edad == iEdad {
			//codificacion
			ojsonBytes, _ := json.MarshalIndent(dat, "", " ")
			io.WriteString(res, string(ojsonBytes))
		}
	}
}

func home(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
	log.Println("Llamada al endpoint /home")
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
			<a href="/buscarDato">Busca</a>
			<a href="/listarData">Listar</a>
		</form>

		<a href="/buscarDato">Busca</a>
		
	</body>	
	</html>

	`)
}
func enableCors(res *http.ResponseWriter) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
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
