package repository

import (
	"database/sql"
	"encoding/json"
	"log"
)

type StateRepository struct {
	db *sql.DB
}

func NewStateRepository(db *sql.DB) *StateRepository {
	return &StateRepository{db: db}
}

func (r *StateRepository) CreateState(disease string, treatments []string, medicines []string, treatmentPolicy string, clientId string) (*State, error) {
	log.Print(treatmentPolicy)
	ins, err := r.db.Prepare("INSERT INTO States(disease,treatments,medicines,treatment_policy, client_id) VALUES(?,?,?,?,?)")
	if err != nil {
		return nil, err
	}
	defer ins.Close()
	treatmentsJson, _ := json.Marshal(treatments)
	medicinesJson, _ := json.Marshal(medicines)
	result, err := ins.Exec(disease, treatmentsJson, medicinesJson, treatmentPolicy, clientId)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	lastId, _ := result.LastInsertId()
	return &State{Id: int(lastId), Disease: disease, Treatments: treatments, Medicines: medicines, ClientId: clientId, TreatmentPolicy: treatmentPolicy}, nil
}

type State struct {
	Id              int      `json:"id"`
	Disease         string   `json:"disease"`
	Treatments      []string `json:"treatments"`
	Medicines       []string `json:"medicines"`
	TreatmentPolicy string   `json:"treatment_policy"`
	ClientId        string   `json:"client_id"`
}
