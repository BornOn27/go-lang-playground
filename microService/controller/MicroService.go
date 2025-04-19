package controller

import "net/http"

type MicroService struct {
	app Application
}

func StartControllers(dbUrl string, mux *http.ServeMux) *MicroService {
	app := NewMicroServiceController(dbUrl)

	return &MicroService{app: app}
}
