package main

import (
	"net/http"

	"github.com/zkrhm/imd-socialnetwork/helper"
)

type connectFriendReq struct {
	Friends []string `json:"friends"`
}

type connectFriendRes struct {
	Success bool `json:"success"`
}

func (app *App) ConnectAsFriend(res http.ResponseWriter, req *http.Request) {
	reqObj := connectFriendReq{}
	helper.GetRequest(req, &reqObj)

	if len(reqObj.Friends) != 2 {
		helper.Error(res, "Parameter of friends should exactly has 2 element", 422)
	} else {
		helper.WriteReponse(res, connectFriendRes{Success: true})
	}
}

func (app *App) GetFriendList(w http.ResponseWriter, r *http.Request) {
	panic("handler not implemented")
}

func (app *App) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	panic("handler not implemented")
}

func (app *App) Subsribe(w http.ResponseWriter, r *http.Request) {
	panic("handler not implemented")
}

func (app *App) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	panic("handler not implemented")
}

func (app *App) GetSubscribers(w http.ResponseWriter, r *http.Request) {
	panic("handler not implemented")
}
