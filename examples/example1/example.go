package main

import (
	db "github.com/kiarash8112/query-controller/examples"
)

type User struct {
	nam1 string
}

func main() {
	users := []User{{nam1: "admin"}, {nam1: "guest"}}
	db := &db.GormDB{}

	for _, u := range users {
		GetUser(db, u.nam1)
	}

}

func GetUser(db *db.GormDB, u string) {
	db.Where("it is", u).Find(nil)
}

func getID() int {
	return getID() // Recursive call
}

func main1(db *db.GormDB) {
	id := getID()
	db.Where("id = ?", id) // Sink
}
