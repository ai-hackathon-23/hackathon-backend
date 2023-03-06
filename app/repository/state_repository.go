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

func (r *StateRepository) CreateState(state State,clientId string) (*State,error) {
	ins, err := r.db.Prepare("INSERT INTO States(disease,treatments,medicines,treatment_policy) VALUES(?,?,?,?)")
    if err != nil {
        return nil,err
    }
	defer ins.Close()
	treatmentsJson,_ := json.Marshal(state.Treatments)
	medicinesJson,_ := json.Marshal(state.Medicines)
    result, err := ins.Exec(state.Disease,treatmentsJson,medicinesJson,"oiuerfcdou")
	if err != nil {
		log.Print(err)
		return nil,err
	}
	lastId,_ := result.LastInsertId()
	ins, err = r.db.Prepare("INSERT INTO StateRecords(client_id,state_id) VALUES(?,?)")
	if err != nil {
        return nil,err
    }
	result, err = ins.Exec(clientId,lastId)
	if err != nil {
		log.Print(err)
		return nil,err
	}
	defer ins.Close()
	return &State{Id: int(lastId),Disease: state.Disease,Treatments: state.Treatments,Medicines: state.Medicines},nil
}

type State struct {
	Id int
    Disease string 
	Treatments []string
	Medicines []string
}