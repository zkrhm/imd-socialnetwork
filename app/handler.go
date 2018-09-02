package app

import (
	"fmt"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/zkrhm/imd-socialnetwork/errors"
	"github.com/zkrhm/imd-socialnetwork/helper"
	. "github.com/zkrhm/imd-socialnetwork/model"
)

func (app *App) ConnectAsFriend(res http.ResponseWriter, req *http.Request) {
	reqObj := ConnectFriendRequest{}
	helper.GetRequest(req, &reqObj)

	fmt.Println("reqObj : ", reqObj)
	fmt.Println("len of friends field : ", len(reqObj.Friends))

	if len(reqObj.Friends) != 2 {
		helper.WriteReponse(res, ConnectFriendResponse{
			Code:    http.StatusUnprocessableEntity,
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
			Success: false,
			Code:    http.StatusUnprocessableEntity,
			Message: "Invalid parameter. should be valid email address ",
		})

		return
	}

	err := app.DB.ConnectAsFriend(User(friends[0]), User(friends[1]))
	if err != nil {
		helper.WriteReponse(res, ConnectFriendResponse{
			Success: false,
			Code:    err.(*errors.ErrorWithCode).Code(),
			Message: err.Error(),
		})

		return
	}
	
	helper.WriteReponse(res, ConnectFriendResponse{Success: true})
}

func (app *App) GetFriendList(w http.ResponseWriter, r *http.Request) {
	reqObj := FriendListRequest{}
	helper.GetRequest(r, &reqObj)

	friends, err := app.DB.GetFriendList(User(reqObj.Email))

	if err != nil {
		helper.WriteReponse(w, FriendListResponse{
			Success: false,
			Code:    err.(*errors.ErrorWithCode).Code(),
			Message: err.Error(),
		})
		return
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
			Code:    http.StatusUnprocessableEntity,
			Message: "invalid email address format",
			Success: false,
		})
		return
	}

	if len(params) != 2 {
		helper.WriteReponse(w, CommonFriendResponse{
			Success: false,
			Code:    http.StatusUnprocessableEntity,
			Message: "number of friends parameter should be exactly two",
		})

		return
	}

	friends, err := app.DB.CommonFriends(User(params[0]), User(params[1]))

	if err != nil {
		helper.WriteReponse(w, CommonFriendResponse{
			Success: false,
			Code:    err.(*errors.ErrorWithCode).Code(),
			Message: err.Error(),
		})

		return
	}

	helper.WriteReponse(w, CommonFriendResponse{
		Success: true,
		Friends: ConvertToStringArray(friends),
		Count:   len(friends),
	})
}

func (app *App) Subscribe(w http.ResponseWriter, r *http.Request) {
	reqObj := &SubscribeRequest{}
	helper.GetRequest(r, reqObj)

	err1 := checkmail.ValidateFormat(reqObj.Requestor)
	err2 := checkmail.ValidateFormat(reqObj.Target)

	if err1 != nil || err2 != nil {
		helper.WriteReponse(w, CommonFriendResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "invalid email address format",
			Success: false,
		})
		return 
	}

	err := app.DB.SubscribeTo(User(reqObj.Requestor), User(reqObj.Target))

	if err != nil {
		helper.WriteReponse(w, CommonFriendResponse{
			Code:    err.(*errors.ErrorWithCode).Code(),
			Message: err.Error(),
			Success: false,
		})

		return
	}

	helper.WriteReponse(w, SubscribeResponse{
		Success: true,
	})
}

func (app *App) ListSubscribers(w http.ResponseWriter, r *http.Request){
	reqObj := &SubscriberListRequest{}
	err := helper.GetRequest(r, reqObj)

	if err!= nil {
		helper.WriteReponse(w,SubscriberListResponse{
			Code: http.StatusUnprocessableEntity,
			Message : err.Error(),
			Success: false,
		})

		return
	}

	err = checkmail.ValidateFormat(reqObj.Email)
	if err!= nil{
		helper.WriteReponse(w,SubscriberListResponse{
			Code : http.StatusUnprocessableEntity,
			Message : err.Error(),
			Success: false,
		})

		return 
	}

	subscribers, err := app.DB.GetFriendList(User(reqObj.Email))

	if err != nil {
		helper.WriteReponse(w, SubscriberListResponse{
			Code: err.(*errors.ErrorWithCode).Code(),
			Message: err.Error(),

		})

		return
	}

	helper.WriteReponse(w, SubscriberListResponse{
		Subscribers: ConvertToStringArray(subscribers),
		Success: true,
		Count : len(subscribers),
	})
}

func (app *App) Block(w http.ResponseWriter, r *http.Request) {
	reqObj := &BlockRequest{}
	err := helper.GetRequest(r, reqObj)

	if err != nil {
		helper.WriteReponse(w, BlockResponse{
			Code : http.StatusUnprocessableEntity,
			Message : err.Error(),
			Success : false,
		})
		return 
	}

	err1 := checkmail.ValidateFormat(reqObj.Requestor)
	err2 := checkmail.ValidateFormat(reqObj.Target)
	
	if err1 != nil || err2 != nil {
		helper.WriteReponse(w, BlockResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "invalid email format at parameter requestor or target",
			Success: false,
		})
		return
	}

	err = app.DB.BlockUpdate(User(reqObj.Requestor), User(reqObj.Target))
	if err != nil {
		helper.WriteReponse(w, BlockResponse{
			Code:    err.(*errors.ErrorWithCode).Code(),
			Message: err.Error(),
			Success: false,
		})
		return
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
			Code:    http.StatusUnprocessableEntity,
			Message: "Invalid email format",
			Success: false,
		})
		return 
	}

	recipients, err := app.DB.DoUpdate(User(reqObj.Sender), reqObj.Text)

	if err != nil {
		helper.WriteReponse(w, UpdateResponse{
			Code:    err.(*errors.ErrorWithCode).Code(),
			Message: err.Error(),
			Success: false,
		})
		return 
	}

	helper.WriteReponse(w, UpdateResponse{
		Success:    true,
		Recipients: ConvertToStringArray(recipients),
	})
}
