package main

import (
	"fmt"
	"strconv"
	"os"
	"github.com/zkrhm/imd-socialnetwork/db"
	"github.com/zkrhm/imd-socialnetwork/app"
)

func main() {
	port := os.Getenv("APP_PORT")

	if port == "" {
		port = ":8000"
	}
	iport, err := strconv.Atoi(port)
	if err == nil {
		port = fmt.Sprintf(":%d",iport)
	}

	app := app.NewApp()
	db, err := db.NewCayleyStore()
	if err != nil {
		panic("DB initialization is failed")
	}
	app.UseDb(db)
	app.Initialize()
	
	app.Run(port)
}
