package repository

import (
	"database/sql"
	"encoding/json"
	"log"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) CreateClient(client Client) (*Client,error) {
	log.Print(client)
	ins, err := r.db.Prepare("INSERT INTO Clients(name,age,family_living_togethers) VALUES(?,?,?)")
    if err != nil {
        return nil,err
    }
	defer ins.Close()
	familyFamilyLivingTogethersJson,_ := json.Marshal(client.FamilyLivingTogethers)
    result, err := ins.Exec(client.Name,client.Age,familyFamilyLivingTogethersJson)
	if err != nil {
		return nil,err
	}
	lastId,_ := result.LastInsertId()
	return &Client{Id: int(lastId),Name: client.Name,Age: client.Age,FamilyLivingTogethers: client.FamilyLivingTogethers},nil
}

type Client struct {
	Id int
	Name string 
	Age int 
	FamilyLivingTogethers string
}