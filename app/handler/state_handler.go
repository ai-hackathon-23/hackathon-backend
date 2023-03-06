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

func (hd *StateHandler) HandleCreateState(w http.ResponseWriter,r *http.Request) error {
    disease := r.FormValue("disease")
    treatments := r.FormValue("treatments")
    medicines := r.FormValue("medicines")
	clientId := r.FormValue("client_id")
	treatmentsArray := strings.Split(treatments, ",")
	medicinesArray := strings.Split(medicines, ",")
	state,err := hd.rp.CreateState(rp.State{Disease: disease,Treatments: treatmentsArray,Medicines: medicinesArray},clientId)
    if err != nil {
		log.Print(err)
	} else {
		jsonData, _ := json.Marshal(state)
		fmt.Fprintf(w, string(jsonData))
	}
	return nil
}