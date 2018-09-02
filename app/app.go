package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/zkrhm/imd-socialnetwork/db"
	"github.com/zkrhm/imd-socialnetwork/model"
)

type App struct {
	Router *mux.Router
	DB     IFriendMgtStore
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
	app.DB = db
}

func (app *App) Initialize() {
	app.dataPreload()
	app.initRoutes()
}

func (app *App) dataPreload() {

	u := make(map[string]string)
	u["bob"] = "bob@example.com"
	u["alice"] = "alice@example.com"
	u["greg"] = "greg@example.com"
	u["fred"] = "fred@example.com"
	u["emily"] = "emily@example.com"
	u["dani"] = "dani@example.com"
	u["charlie"] = "charlie@example.com"
	u["jonathan"] = "jonathan@example.com"
	u["maria"] = "maria@example.com"

	for _, v := range u {
		app.DB.AddUser(model.User(v))
	}

	app.DB.ConnectAsFriend(model.User(u["bob"]), model.User(u["alice"]))
	app.DB.ConnectAsFriend(model.User(u["bob"]), model.User(u["charlie"]))
	app.DB.ConnectAsFriend(model.User(u["bob"]), model.User(u["dani"]))
	app.DB.ConnectAsFriend(model.User(u["bob"]), model.User(u["fred"]))
	app.DB.ConnectAsFriend(model.User(u["charlie"]), model.User(u["dani"]))
	app.DB.ConnectAsFriend(model.User(u["greg"]), model.User(u["dani"]))
	app.DB.ConnectAsFriend(model.User(u["fred"]), model.User(u["greg"]))
	app.DB.ConnectAsFriend(model.User(u["fred"]), model.User(u["emily"]))
}

func (app *App) initRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/connect", app.ConnectAsFriend)
	router.HandleFunc("/friend-list", app.GetFriendList)
	router.HandleFunc("/common-friends", app.GetCommonFriends)
	router.HandleFunc("/subscribe", app.Subscribe)
	router.HandleFunc("/list-subscriber", app.ListSubscribers)
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
