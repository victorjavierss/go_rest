package collection

import (
	"log"

	"upper.io/db.v3/lib/sqlbuilder"
)

const userTbl = "user"

type User struct {
	ID   uint64 `db:"id,omitempty" json:"pid"`
	Name string `db:"name" json:"name"`

	db sqlbuilder.Database
}

func NewUserCollection(db sqlbuilder.Database) User {
	user := User{
		db: db,
	}

	return user
}

func (u User) FindAll() []User {
	var users []User

	userCollection := u.db.Collection(userTbl)
	res := userCollection.Find()
	err := res.All(&users)

	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}

	return users
}

func (u User) Find(id string) []User {
	var users []User

	userCollection := u.db.Collection(userTbl)
	res := userCollection.Find("id", id)
	err := res.All(&users)

	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
		return nil
	}

	return users
}

func (u User) Save() (int64, error) {
	userCollection := u.db.Collection(userTbl)
	newID, err := userCollection.Insert(u)
	return newID.(int64), err
}
