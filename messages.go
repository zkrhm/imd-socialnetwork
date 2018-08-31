package main

type ConnectFriendRequest struct {
	Friends []string `json:"friends"`
}

type ConnectFriendResponse struct {
	Success bool `json:"success"`
}

type FriendListRequest struct {
	Email string `json:"email"`
}

type FriendListResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends:`
	Count   int      `json:"count"`
}

type CommonFriendRequest struct {
	Friends []string `json:"friends"`
}

type CommonFriendResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   bool     `json:"count"`
}

type SubscribeRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

type SubscribeResponse struct {
	Success bool `json:"success"`
}

type BlockRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

type BlockResponse struct {
	Success bool `json:"success"`
}

type UpdateRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

type UpdateResponse struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}
