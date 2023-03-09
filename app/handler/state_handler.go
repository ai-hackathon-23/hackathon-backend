package handler

import (
	"encoding/json"
	"fmt"
	rp "hackathon/repository"
	"log"
	"net/http"
	"strings"
)

type StateHandler struct {
	rp *rp.StateRepository
}

func NewStateHandler(repository *rp.StateRepository) StateHandler {
	return StateHandler{repository}
}

func (hd *StateHandler) HandleCreateState(w http.ResponseWriter, r *http.Request) error {
	disease := r.FormValue("disease")
	treatments := r.FormValue("treatments")
	medicines := r.FormValue("medicines")
	treatmentPolicy := r.FormValue("treatment_policy")
	clientId := r.FormValue("client_id")
	treatmentsArray := strings.Split(treatments, ",")
	medicinesArray := strings.Split(medicines, ",")
	state, err := hd.rp.CreateState(disease, treatmentsArray, medicinesArray, treatmentPolicy, clientId)
	if err != nil {
		log.Print(err)
	} else {
		jsonStr, _ := json.Marshal(state)
		fmt.Fprintf(w, string(jsonStr))
	}
	return nil
}

type State struct {
	Id              int      `json:"id"`
	Disease         string   `json:"disease"`
	Treatments      []string `json:"treatments"`
	Medicines       []string `json:"medicines"`
	TreatmentPolicy string   `json:"treatment_policy"`
	ClientId        string   `json:"client_id"`
}
