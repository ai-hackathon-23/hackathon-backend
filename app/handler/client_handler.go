package handler

import (
	"encoding/json"
	"fmt"
	rp "hackathon/repository"
	"log"
	"net/http"
	"strconv"
)

type ClientHandler struct {
    rp *rp.ClientRepository
}

func NewClientHandler(repository *rp.ClientRepository) ClientHandler {
    return ClientHandler{repository}
}

func (hd *ClientHandler) HandleCreateClient(w http.ResponseWriter,r *http.Request) error {
    name := r.FormValue("name")
    age, _ := strconv.Atoi(r.FormValue("age"))
    livingInfo := r.FormValue("family_living_togethers")
	
    client,err := hd.rp.CreateClient(rp.Client{Name: name, Age: age,FamilyLivingTogethers: livingInfo})
    if err != nil {
		log.Print(err)
	} else {
		jsonData, _ := json.Marshal(client)
		fmt.Fprintf(w, string(jsonData))
	}
	return nil
}
