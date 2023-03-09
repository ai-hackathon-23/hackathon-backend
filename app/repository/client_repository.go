package repository

import (
	"database/sql"
	"log"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) ClientRepository {
	return ClientRepository{db: db}
}

func (r *ClientRepository) FindByID(id string) (*Client, error) {
	client := &Client{}
	err := r.db.QueryRow("SELECT * FROM Clients WHERE id = ?", id).Scan(&client.Id, &client.Name, &client.Age, &client.FamilyLivingTogethers)
	log.Print(id)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (r *ClientRepository) CreateClient(name string, age int, familyLivingTogethers []byte) (*Client, error) {
	ins, err := r.db.Prepare("INSERT INTO Clients(name,age,family_living_togethers) VALUES(?,?,?)")
	if err != nil {
		return nil, err
	}
	defer ins.Close()
	result, err := ins.Exec(name, age, familyLivingTogethers)
	if err != nil {
		return nil, err
	}
	lastId, _ := result.LastInsertId()
	return &Client{Id: int(lastId), Name: name, Age: age, FamilyLivingTogethers: string(familyLivingTogethers)}, nil
}

func (r *ClientRepository) IndexClients() (*[]Client, error) {

	stmt, err := r.db.Prepare("SELECT * FROM Clients")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	log.Print(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	clients := []Client{}
	for rows.Next() {
		carePlan := Client{}
		err := rows.Scan(
			&carePlan.Id,
			&carePlan.Name,
			&carePlan.Age,
			&carePlan.FamilyLivingTogethers,
		)
		clients = append(clients, carePlan)
		if err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &clients, nil

}


type Client struct {
	Id                    int    `json:"id"`
	Name                  string `json:"name"`
	Age                   int    `json:"age"`
	FamilyLivingTogethers string `json:"family_living_togethers"`
	CarePlans             []CarePlans `json:"care_plans"`
}

type CarePlans struct {
	Id                   int64  `json:"id"`
	Author               string `json:"author"`
	FacilityName         string `json:"facility_name"`
	ResultAnalyze        string `json:"result_analyze"`
	CareCommitteeOpinion string `json:"care_committee_opinion"`
	SpecifiedService     string `json:"specified_service"`
	CarePolicy           string `json:"care_policy"`
	UpdatedAt            string `json:"updated_at"`
	ClientId         string `json:"client_id"`
}