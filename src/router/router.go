package router

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/bradfitz/gomemcache/memcache"

	"github.com/user/app/index"
	"github.com/user/app/bootstrap"
	"github.com/user/app/collection"
)

var dependecies bootstrap.Dependecies

func InitRoutes (deps bootstrap.Dependecies) *mux.Router {
	r := mux.NewRouter()
	index := index.Index{}

	dependecies = deps

	r.HandleFunc("/", index.Get)
	r.HandleFunc("/users", usersHandler)

	return r;
}


func usersHandler(w http.ResponseWriter, r *http.Request) {

	userItem, err := dependecies.Cache.Get("users")
	var usersJson []byte

	 if err != nil {
		if err == memcache.ErrCacheMiss {
			log.Printf("Cache MISS")

			userCollection := dependecies.Database.Collection("users")

			res := userCollection.Find()

			var users []collection.User

			err := res.All(&users)

			if err != nil {
				log.Fatalf("res.All(): %q\n", err)
			}
			
			usersJson, err = json.Marshal(users)
			err = dependecies.Cache.Set(&memcache.Item{
				Key: "users", 
				Value: usersJson,
				Expiration: 180,
			})
			if err != nil {
				log.Println("Error setting cache #%v", err)
			}
		} else {
			log.Println("Error from cache #%v", err)
		}
	} else {
		log.Printf("Cache HIT")
		usersJson = userItem.Value;
	} 


	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(usersJson))
}