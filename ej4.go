package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const normalType = "normal"
const jubiladoType = "jubilado"
const invitadoType = "invitado"

// Client : Datos del cliente
type Client struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

func (p Client) toString() string {
	return toJSON(p)
}

func toJSON(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

func main() {

	var price float64 = 80
	var total float64

	clients := getClients()

	f, err := os.Create("./error.log")
	check(err)
	defer f.Close()

	for _, client := range clients {
		// fmt.Println(client.toString())

		switch client.Type {
		case normalType:
			total += price
		case jubiladoType:
			total += price / 2
		case invitadoType:
		default:
			{
				f.Write([]byte("Error al procesar cliente " + strconv.Itoa(client.ID) + ", tipo no valido: " + client.Type + "\n"))
			}
		}
	}

	fmt.Printf("La recaudacion total es de %.2f", total)
}

func getClients() []Client {
	raw, err := ioutil.ReadFile("./data.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Client
	json.Unmarshal(raw, &c)
	return c
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
