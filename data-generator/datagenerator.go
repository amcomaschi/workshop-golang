package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
)

const file = "./data.json"

var clientTypes = []string{"jubilado", "invitado", "normal", "incorrecto", "vip"}

// Client : Datos del cliente
type Client struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

func main() {

	var clients [500]Client

	dataFile, error := os.Create(file)
	if error != nil {
		log.Fatal(error)
	}

	defer dataFile.Close()

	for i := 0; i < 500; i++ {

		client := Client{i + 1, clientTypes[rand.Intn(5)]}
		clients[i] = client

		//fmt.Printf("Cliente: %d, %s", client.ID, client.Type)
	}

	if b, error := json.MarshalIndent(clients, "", " "); error == nil {
		dataFile.Write(b)
	}

}
