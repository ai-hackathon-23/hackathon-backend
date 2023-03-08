package repository

import (
	"database/sql"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) ClientRepository {
	return ClientRepository{db: db}
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
	stmt, err := r.db.Prepare("Select * from Clients")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	clients := []Client{}
	for rows.Next() {
		var id int
		var name string
		var age int
		var familyLivingTogethers string
		clients = append(clients,
			Client{Id: id, Name: name, Age: age, FamilyLivingTogethers: familyLivingTogethers})
	}
	return clients, nil
}

type Client struct {
	Id                    int
	Name                  string
	Age                   int
	FamilyLivingTogethers string
}
