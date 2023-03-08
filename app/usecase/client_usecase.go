package usecase

import (
	"hackathon/repository"
	"strconv"
)

type ClientUseCase struct {
	repository repository.ClientRepository
}

func NewClientUseCase(repository repository.ClientRepository) ClientUseCase {
	return ClientUseCase{repository}
}

func (usecase *ClientUseCase) CreateClient(name string, age string, familyLivingTogethers string) (*repository.Client, error) {
	int_age, _ := strconv.Atoi(age)
	client, err := usecase.repository.CreateClient(name, int_age, []byte(familyLivingTogethers))
	return client, err
}

func (usecase *ClientUseCase) IndexClients() ([]repository.Client, error) {
	clients, err := usecase.repository.IndexClients()
	return clients, err
}
