package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/zkrhm/imd-socialnetwork/db"
)

type App struct {
	Router *mux.Router
	DB     *IFriendMgtStore
}

func NewApp() *App {
	return &App{}
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		w.Header().Set("Content-type", "application/json")
	})
}

func (app *App) UseDb(db IFriendMgtStore) {
	app.DB = &db
}

func (app *App) Initialize() {
	app.initRoutes()
}

func (app *App) initRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/connect", app.ConnectAsFriend)
	router.HandleFunc("/friend-list", app.GetFriendList)
	router.HandleFunc("/common-friends", app.GetCommonFriends)
	router.HandleFunc("/subscribe", app.Subsribe)
	router.HandleFunc("/block", app.Block)
	router.HandleFunc("/post-update", app.PostUpdate)
	router.Use(jsonMiddleware)

	app.Router = router
}

func (app *App) Run(addr string) {
	server := http.Server{
		Handler: app.Router,
		Addr:    addr,
	}

	fmt.Println("running server on ", addr)

	server.ListenAndServe()
}
