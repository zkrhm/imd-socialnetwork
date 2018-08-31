package main

import (
	"github.com/zkrhm/imd-socialnetwork/db"
)

func main() {
	app := NewApp()
	db, err := db.NewCayleyStore()
	if err != nil {
		panic("DB initialization is failed")
	}
	app.UseDb(db)
	app.Initialize()
	
	app.Run("localhost:8000")
}
