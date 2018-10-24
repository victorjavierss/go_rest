package router

import (
	"github.com/gorilla/mux"

	"github.com/user/app/bootstrap"
	"github.com/user/app/index"
	"github.com/user/app/user"
)

func InitRoutes(deps bootstrap.Dependecies) *mux.Router {
	r := mux.NewRouter()

	indexCtrl := index.NewCtrl()
	userCtrl := user.NewCtrl(deps)

	r.HandleFunc("/", indexCtrl.Get)
	r.HandleFunc("/users", userCtrl.Get).Methods("GET")
	r.HandleFunc("/users/{id}", userCtrl.Get).Methods("GET")
	r.HandleFunc("/users/{id}", userCtrl.Delete).Methods("DELETE")
	r.HandleFunc("/users", userCtrl.Post).Methods("POST")

	return r
}
