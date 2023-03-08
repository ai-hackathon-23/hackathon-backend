package handler

import (
	"encoding/json"
	"fmt"
	usecase "hackathon/usecase"
	"log"
	"net/http"
)

type ClientHandler struct {
	usecase usecase.ClientUseCase
}

func NewClientHandler(usecase usecase.ClientUseCase) ClientHandler {
	return ClientHandler{usecase}
}

func (hd *ClientHandler) HandleCreateClient(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	age := r.FormValue("age")
	livingInfo := r.FormValue("family_living_togethers")

	client, err := hd.usecase.CreateClient(name, age, livingInfo)
	if err != nil {
		log.Print(err)
	} else {
		jsonData, _ := json.Marshal(client)
		fmt.Fprintf(w, string(jsonData))
	}
	return nil
}

func (hd *ClientHandler) IndexClients(w http.ResponseWriter, r *http.Request) error {
	clients, err := hd.usecase.IndexClients()
	if err != nil {
		log.Print(err)
	} else {
		jsonData, _ := json.Marshal(clients)
		fmt.Fprintf(w, string(jsonData))
	}
	return nil
}
