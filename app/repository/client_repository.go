package repository

import (
	"database/sql"
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
	ins, err := r.db.Prepare("INSERT INTO Clients(name,age,living_info) VALUES(?,?,?)")
    if err != nil {
        return nil,err
    }
	defer ins.Close()
    result, err := ins.Exec(client.Name,client.Age,"")
	if err != nil {
		return nil,err
	}
	lastId,_ := result.LastInsertId()
	return &Client{Id: int(lastId),Name: client.Name,Age: client.Age,LivingInfo: client.LivingInfo},nil
}

type Client struct {
	Id int
	Name string 
	Age int 
	LivingInfo string
}