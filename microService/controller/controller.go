package controller

import (
	"net/http"
)

func (ms *MicroService) RegisterRoutes(mux *http.ServeMux) {

}

func (ms *MicroService) AddUser(user string) string {
	ms.app.AddUser()
	return user
}

func (ms *MicroService) GetUser(user string) string {
	ms.app.GetUser()
	return user
}

func (ms *MicroService) UpdateUser(user string) string {
	ms.app.UpdateUser()
	return user
}

func (ms *MicroService) DeleteUser(user string) string {
	ms.app.DeleteUser()
	return user
}
