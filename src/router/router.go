package router

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	//"github.com/bradfitz/gomemcache/memcache"

	"github.com/user/app/index"
	"github.com/user/app/bootstrap"
	"github.com/user/app/collection"
	//"upper.io/db.v3/mysql"
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
	//_, err := dependecies.Cache.Get("mykey") 

	userCollection := dependecies.Database.Collection("users")

  	res := userCollection.Find()

	var users []collection.User

	err := res.All(&users)

	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}

	usersJson, err := json.Marshal(users)

	/* if err != nil {
		if err == memcache.ErrCacheMiss {
			fmt.Fprintf(w, "Cache MISS")

			err = dependecies.Cache.Set(&memcache.Item{Key: "mykey", Value: []byte("my value")})
			if err != nil {
				log.Println("Error setting cache #%v", err)
			}

		} else {
			log.Println("Error from cache #%v", err)
		}
		
	} else {
		fmt.Fprintf(w, "Cache HIT")
	} */



	fmt.Fprintf(w, string(usersJson))
}