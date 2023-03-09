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

func (r *ClientRepository) IndexClients() ([]Client, error) {
	stmt, err := r.db.Prepare(`
        SELECT c.id, c.name, c.age, c.family_living_togethers, cp.author, cp.facility_name, cp.result_analyze, cp.care_committee_opinion, cp.specified_service, cp.care_policy, cp.updated_at 
        FROM Clients c 
        LEFT JOIN CarePlans cp ON c.id = cp.client_id
    `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	clients := map[int]Client{}
	for rows.Next() {
		var id int
		var name *string
		var age int
		var familyLivingTogethers string
		var author, facilityName, resultAnalyze, careCommitteeOpinion, specifiedService, carePolicy *string
		var updatedAt *string
		if err := rows.Scan(&id, &name, &age, &familyLivingTogethers, &author, &facilityName, &resultAnalyze, &careCommitteeOpinion, &specifiedService, &carePolicy, &updatedAt); err != nil {
			return nil, err
		}
		client, ok := clients[id]
		if !ok {
			client = Client{Id: id, Name: *name, Age: age, FamilyLivingTogethers: familyLivingTogethers, CarePlans: []CarePlans{}}
		}
		carePlan := CarePlans{Author: *author, FacilityName: *facilityName, ResultAnalyze: *resultAnalyze, CareCommitteeOpinion: *careCommitteeOpinion, SpecifiedService: *specifiedService, CarePolicy: *carePolicy, UpdatedAt: *updatedAt}
		client.CarePlans = append(client.CarePlans, carePlan)
		clients[id] = client
	}

	result := make([]Client, 0)
	for _, client := range clients {
		result = append(result, client)
	}


	return result, nil
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
	Client               Client `json:"client"`
}