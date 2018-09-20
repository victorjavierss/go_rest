package router

import (

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/user/app/bootstrap"
	"github.com/user/app/index"
)

type routerHandler interface {
    Handle()
}


func InitRoutes (deps bootstrap.Dependecies) *mux.Router {
	r := mux.NewRouter()
	index := index.Index{}

	r.HandleFunc("/", index.Handle)
	r.HandleFunc("/users", usersHandler)

	return r;
}


func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm a user Handler REST API :-)")
}