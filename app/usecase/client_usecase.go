package usecase

import (
	"hackathon/repository"
	"strconv"
	"log"
)

type ClientUseCase struct {
	repository repository.ClientRepository
	carePlanRepository *repository.CarePlanRepository
}

func NewClientUseCase(repository repository.ClientRepository, carePlanRepository *repository.CarePlanRepository) ClientUseCase {
	return ClientUseCase{repository, carePlanRepository}
}

func (usecase *ClientUseCase) CreateClient(name string, age string, familyLivingTogethers string) (*repository.Client, error) {
	int_age, _ := strconv.Atoi(age)
	client, err := usecase.repository.CreateClient(name, int_age, []byte(familyLivingTogethers))
	return client, err
}

func (usecase *ClientUseCase) IndexClients() (*[]repository.Client, error) {
	clients, err := usecase.repository.IndexClients()
	var updatedClients []repository.Client
	log.Print(clients)
	for _,client := range *clients {
		carePlans, _ := usecase.carePlanRepository.GetCarePlansByClientId(strconv.Itoa(client.Id))
		client.CarePlans = carePlans
		updatedClients = append(updatedClients, client)
	}

	return &updatedClients, err
}
