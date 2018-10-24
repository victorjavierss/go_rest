package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gorilla/mux"

	"github.com/user/app/bootstrap"
	"github.com/user/app/collection"
)

type UserCtrl struct {
	deps bootstrap.Dependecies
}

func NewCtrl(d bootstrap.Dependecies) UserCtrl {
	ctrl := UserCtrl{
		deps: d,
	}

	return ctrl
}

func (u UserCtrl) Get(w http.ResponseWriter, r *http.Request) {

	var usersJSON []byte

	userID := mux.Vars(r)["id"]

	cacheKey := fmt.Sprintf("users_%s", userID)
	userItem, err := u.deps.Cache.Get(cacheKey)

	if err != nil {
		if err == memcache.ErrCacheMiss {
			log.Printf("Cache MISS")

			userCollection := collection.NewUserCollection(u.deps.Database)
			var users []collection.User
			if userID != "" {
				users = userCollection.Find(userID)
			} else {
				users = userCollection.FindAll()
			}

			usersJSON, err = json.Marshal(users)

			err = u.deps.Cache.Set(&memcache.Item{
				Key:        cacheKey,
				Value:      usersJSON,
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
		usersJSON = userItem.Value
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(usersJSON))
}

func (u UserCtrl) Delete(w http.ResponseWriter, r *http.Request) {
	//res := col.Find("id", 4)
	// res.Delete()
}

func (u UserCtrl) Post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	userCollection := collection.NewUserCollection(u.deps.Database)

	err := json.NewDecoder(r.Body).Decode(&userCollection)

	if err != nil {
		log.Print("Unable to read body")
		http.Error(w, "{\"status\":\"fail\"}", http.StatusBadRequest)
		return
	}

	userCollection.Save()

	fmt.Printf("%+v", userCollection)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"status\":\"ok\"}")
}
