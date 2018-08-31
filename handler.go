package main

import (
	"fmt"
	"net/http"

	"github.com/zkrhm/imd-socialnetwork/helper"
	"github.com/badoux/checkmail"
	. "github.com/zkrhm/imd-socialnetwork/model"
)

func (app *App) ConnectAsFriend(res http.ResponseWriter, req *http.Request) {
	reqObj := ConnectFriendRequest{}
	helper.GetRequest(req, &reqObj)
	
	fmt.Println("reqObj : ", reqObj)
	fmt.Println("len of friends field : ", len(reqObj.Friends))

	if len(reqObj.Friends) != 2 {
		helper.WriteReponse(res, ConnectFriendResponse{
			Code: http.StatusUnprocessableEntity,
			Message: "Parameter of friends should exactly has 2 element",
			Success: false,
		})

		return
	}

	friends := reqObj.Friends
	fmt.Println("friends: ", friends)
	err1 := checkmail.ValidateFormat(friends[0])
	err2 := checkmail.ValidateFormat(friends[1])
	if err1 != nil || err2 != nil {
		helper.WriteReponse(res, ConnectFriendResponse{
			Success:false,
			Code: http.StatusUnprocessableEntity,
			Message: "Invalid parameter. should be valid email address ",
		})
	}

	err := app.DB.ConnectAsFriend(User(friends[0]), User(friends[1]))
	if err != nil {
		helper.WriteReponse(res, ConnectFriendResponse{
			Success: false,
			Code : http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	helper.WriteReponse(res, ConnectFriendResponse{Success: true})
	
}

func (app *App) GetFriendList(w http.ResponseWriter, r *http.Request) {
	reqObj := FriendListRequest{}
	helper.GetRequest(r, &reqObj)

	friends , err := app.DB.GetFriendList(User(reqObj.Email))

	if err != nil {
		helper.WriteReponse(w, FriendListResponse{
			Success : false,
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	helper.WriteReponse(w, FriendListResponse{
		Success: true,
		Friends: ConvertToStringArray(friends),
		Count:   len(friends),
	})
}

func (app *App) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	reqObj := &CommonFriendRequest{}
	helper.GetRequest(r, reqObj)

	params := reqObj.Friends
	err1 := checkmail.ValidateFormat(params[0])
	err2 := checkmail.ValidateFormat(params[1])

	if err1 != nil || err2 != nil {
		helper.WriteReponse(w, CommonFriendResponse{
			Code: http.StatusUnprocessableEntity,
			Message: "invalid email address format",
			Success: false,
		})
	}

	if len(params) != 2 {
		helper.WriteReponse(w,CommonFriendResponse{
			Success: false,
			Code: http.StatusUnprocessableEntity,
			Message : "number of friends parameter should be exactly two",
		})
	}
	
	friends, err := app.DB.CommonFriends(User(params[0]), User(params[1]))

	if err != nil {
		helper.WriteReponse(w, CommonFriendResponse{
			Success: false,
			Code : http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	helper.WriteReponse(w, CommonFriendResponse{
		Success: true,
		Friends: ConvertToStringArray(friends),
		Count : len(friends),
	})
}

func (app *App) Subsribe(w http.ResponseWriter, r *http.Request) {
	reqObj := &SubscribeRequest{}
	helper.GetRequest(r, reqObj)

	err1 := checkmail.ValidateFormat(reqObj.Requestor) 
	err2 := checkmail.ValidateFormat(reqObj.Target) 

	if err1 != nil || err2 != nil {
		helper.WriteReponse(w, CommonFriendResponse{
			Code: http.StatusUnprocessableEntity,
			Message: "invalid email address format",
			Success: false,
		})
	}

	err := app.DB.SubscribeTo(User(reqObj.Requestor), User(reqObj.Target))

	if err != nil {
		helper.WriteReponse(w, CommonFriendResponse{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
			Success: false,
		})
	}

	helper.WriteReponse(w, SubscribeResponse{
		Success : true,
	})

}

func (app *App) Block(w http.ResponseWriter, r *http.Request) {
	reqObj := &BlockRequest{}
	helper.GetRequest(r, reqObj)

	// check format
	err1 := checkmail.ValidateFormat(reqObj.Requestor)
	err2 := checkmail.ValidateFormat(reqObj.Target)
	if err1 != nil || err2 != nil {
		helper.WriteReponse(w, BlockResponse{
			Code: http.StatusUnprocessableEntity,
			Message: "invalid email format at parameter requestor or target",
			Success: false,
		})
	}

	err := app.DB.BlockUpdate(User(reqObj.Requestor), User(reqObj.Target))
	if err != nil {
		helper.WriteReponse(w, BlockResponse{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
			Success: false,
		})
	}

	helper.WriteReponse(w, BlockResponse{
		Success: true,
	})
}

func (app *App) PostUpdate(w http.ResponseWriter, r *http.Request) {
	reqObj := &UpdateRequest{}
	helper.GetRequest(r, reqObj)

	err := checkmail.ValidateFormat(reqObj.Sender) 
	if err != nil {
		helper.WriteReponse(w, UpdateResponse{
			Code: http.StatusUnprocessableEntity,
			Message: "Invalid email format",
			Success: false,
		})
	}

	recipients, err := app.DB.DoUpdate(User(reqObj.Sender), reqObj.Text)

	if err != nil {
		helper.WriteReponse(w, UpdateResponse{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
			Success: false,
		})
	}

	helper.WriteReponse(w, UpdateResponse{
		Success: true,
		Recipients: ConvertToStringArray(recipients),
	})
}
