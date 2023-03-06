package handler

import (
	"encoding/json"
	"fmt"
	rp "hackathon/repository"
	"log"
	"net/http"
)

type CarePlanHandler struct {
	rp *rp.CarePlanRepository
}

func NewCarePlanHandler(repository *rp.CarePlanRepository) CarePlanHandler {
	return CarePlanHandler{repository}
}

func (hd *CarePlanHandler) HandleGetCarePlan(w http.ResponseWriter, r *http.Request) error {
	clientId := r.FormValue("client_id")

	care_plan, err := hd.rp.CreateCarePlan(clientId)
	if err != nil {
		log.Print(err)
	} else {
		jsonData, _ := json.Marshal(care_plan)
		fmt.Fprintf(w, string(jsonData))
	}
	return nil
}
