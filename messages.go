package main


type ConnectFriendRequest struct {
	Friends []string `json:"friends"`
}

type ConnectFriendResponse struct {
	Success bool `json:"success"`
	Message string `json:"message,omitempty"`
	Code int `json:"code,omitempty"`
}

type FriendListRequest struct {
	Email string `json:"email"`
}

type FriendListResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends,omitempty"`
	Count   int      `json:"count,omitempty"`
	Message string `json:"message,omitempty"`
	Code int `json:"code,omitempty"`
}

type CommonFriendRequest struct {
	Friends []string `json:"friends"`
}

type CommonFriendResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends,omitempty"`
	Count   int     `json:"count,omitempty"`
	Message string `json:"message,omitempty"`
	Code int `json:"code,omitempty"`
}

type SubscribeRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

type SubscribeResponse struct {
	Success bool `json:"success"`
	Message string `json:"message,omitempty"`
	Code int `json:"code,omitempty"`
}

type SubscriberListRequest struct {
	Email string `json:"email"`
}

type SubscriberListResponse struct {
	Success bool `json:"success"`
	Subscribers []string `json:"subscribers"`
	Count int `json:"count,omitempty"`
	Code int `json:"code,emitempty"`
	Message string `json:"message,omitempty"`
}

type BlockRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

type BlockResponse struct {
	Success bool `json:"success"`
	Message string `json:"message,omitempty"`
	Code int `json:"code,omitempty"`
}

type UpdateRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

type UpdateResponse struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
	Message string `json:"message,omitempty"`
	Code int `json:"code,omitempty"`
}
