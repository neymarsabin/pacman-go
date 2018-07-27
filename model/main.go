package model

import (
	"fmt"
	pg "github.com/go-pg/pg"
)

type Player struct {
	Id       int    `sql:"id,pk"`
	Username string `sql:"name"`
	Score    int    `sql:"score"`
}

func Connection() *pg.DB {
	opts := &pg.Options{
		User:     "prasvin",
		Password: "gopacman123",
		Addr:     "localhost:5432",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		fmt.Println("Database was not connected.")
	}
	fmt.Println("Database connected.")
	closeErr := db.Close()

	if closeErr != nil {
		fmt.Println("Error while closing the database.")
	}
	fmt.Println("Connection closed Successfully")
	return db
}

func (player *Player) SavePlayer(db *pg.DB) error {
	insertError := db.Insert(player)
	if insertError != nil {
		fmt.Println("There was an error during the process.")
		return insertError
	}
	fmt.Println("Player was created")
	return nil
}
