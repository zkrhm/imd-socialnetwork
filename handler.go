package main

import (
	"net/http"

	"github.com/zkrhm/imd-socialnetwork/helper"
)

func (app *App) ConnectAsFriend(res http.ResponseWriter, req *http.Request) {
	reqObj := ConnectFriendRequest{}
	helper.GetRequest(req, &reqObj)

	if len(reqObj.Friends) != 2 {
		helper.Error(res, "Parameter of friends should exactly has 2 element", 422)
	} else {
		helper.WriteReponse(res, ConnectFriendResponse{Success: true})
	}
}

func (app *App) GetFriendList(w http.ResponseWriter, r *http.Request) {
	reqObj := FriendListRequest{}
	helper.GetRequest(r, &reqObj)

	helper.WriteReponse(w, FriendListResponse{
		Success: true,
		Friends: []string{"who", "who"},
		Count:   0,
	})
}

func (app *App) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	panic("handler not implemented")
}

func (app *App) Subsribe(w http.ResponseWriter, r *http.Request) {
	panic("handler not implemented")
}

func (app *App) Block(w http.ResponseWriter, r *http.Request) {
	panic("handler not implemented")
}

func (app *App) PostUpdate(w http.ResponseWriter, r *http.Request) {
	panic("handler not implemented")
}
