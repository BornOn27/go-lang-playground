package controller

import (
	"GoPlayground/microService/repository"
)

type Application interface {
	AddUser()
	GetUser()
	UpdateUser()
	DeleteUser()
}


type base struct {
	master repository.MasterRepository
	replica repository.ReplicaRepository
}

func (b *base) AddUser() {
	b.replica.Read()
}

func (b *base) GetUser() {

}

func (b *base) UpdateUser() {
	b.master.Update()
}

func (b *base) DeleteUser() {
	b.master.Delete()
}

func NewMicroServiceController(dbUrl string) Application {
	return &base{
		master: repository.NewMasterRepository(dbUrl),
		replica: repository.NewReplicaRepository(dbUrl),
	}
}
