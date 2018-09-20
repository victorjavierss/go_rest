package main

import (
	"log"
	"net/http"

    "github.com/user/app/router"
    "github.com/user/app/bootstrap"
)

func main() {
    deps := bootstrap.Init()
	r := router.InitRoutes(deps)

	log.Println("Running api server in dev mode")


    defer deps.Database.Close()
	http.ListenAndServe(":8081", r)
}