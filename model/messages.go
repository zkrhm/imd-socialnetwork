package model

import (
	"encoding/json"
)


type ConnectFriendRequest struct {
	Friends []string `json:"friends"`
}
func UnmarshalConnectFriendRequest(data []byte) (ConnectFriendRequest, error){
	o := ConnectFriendRequest{}
	err := json.Unmarshal(data,&o)
	return o, err 
}
func (r *ConnectFriendRequest) Marshal() ([]byte, error){
	return json.Marshal(r)
}

type ConnectFriendResponse struct {
	Success bool `json:"success"`
	Message string `json:"message,omitempty"`
	Code int `json:"code,omitempty"`
}
func UnmarshalConnectFriendResponse(data []byte) (ConnectFriendResponse, error){
	o := ConnectFriendResponse{}
	err := json.Unmarshal(data,&o)
	return o, err 
}
func (r *ConnectFriendResponse) Marshal() ([]byte, error){
	return json.Marshal(r)
}

type FriendListRequest struct {
	Email string `json:"email"`
}

func UnmarshalFriendListRequest(data []byte) (FriendListRequest, error){
	o := FriendListRequest{}
	err := json.Unmarshal(data,&o)
	return o, err 
}
func (r *FriendListRequest) Marshal() ([]byte, error){
	return json.Marshal(r)
}

type FriendListResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends,omitempty"`
	Count   int      `json:"count,omitempty"`
	Message string `json:"message,omitempty"`
	Code int `json:"code,omitempty"`
}

func UnmarshalFriendListResponse(data []byte) (FriendListResponse, error){
	o := FriendListResponse{}
	err := json.Unmarshal(data,&o)
	return o, err 
}
func (r *FriendListResponse) Marshal() ([]byte, error){
	return json.Marshal(r)
}

type CommonFriendRequest struct {
	Friends []string `json:"friends"`
}

func UnmarshalCommonFriendRequest(data []byte) (CommonFriendRequest, error){
	o := CommonFriendRequest{}
	err := json.Unmarshal(data,&o)
	return o, err 
}
func (r *CommonFriendRequest) Marshal() ([]byte, error){
	return json.Marshal(r)
}

type CommonFriendResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends,omitempty"`
	Count   int     `json:"count,omitempty"`
	Message string `json:"message,omitempty"`
	Code int `json:"code,omitempty"`
}

func UnmarshalCommonFriendResponse(data []byte) (CommonFriendResponse, error){
	o := CommonFriendResponse{}
	err := json.Unmarshal(data,&o)
	return o, err 
}
func (r *CommonFriendResponse) Marshal() ([]byte, error){
	return json.Marshal(r)
}

type SubscribeRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

func UnmarshalSubscribeRequest(data []byte) (SubscribeRequest, error){
	o := SubscribeRequest{}
	err := json.Unmarshal(data,&o)
	return o, err 
}
func (r *SubscribeRequest) Marshal() ([]byte, error){
	return json.Marshal(r)
}

type SubscribeResponse struct {
	Success bool `json:"success"`
	Message string `json:"message,omitempty"`
	Code int `json:"code,omitempty"`
}

func UnmarshalSubscribeResponse(data []byte) (SubscribeResponse, error){
	o := SubscribeResponse{}
	err := json.Unmarshal(data,&o)
	return o, err 
}
func (r *SubscribeResponse) Marshal() ([]byte, error){
	return json.Marshal(r)
}

type SubscriberListRequest struct {
	Email string `json:"email"`
}

func UnmarshalSubscriberListRequest(data []byte) (SubscriberListRequest, error){
	o := SubscriberListRequest{}
	err := json.Unmarshal(data,&o)
	return o, err 
}
func (r *SubscriberListRequest) Marshal() ([]byte, error){
	return json.Marshal(r)
}

type SubscriberListResponse struct {
	Success bool `json:"success"`
	Subscribers []string `json:"subscribers"`
	Count int `json:"count,omitempty"`
	Code int `json:"code,emitempty"`
	Message string `json:"message,omitempty"`
}

func UnmarshalSubscribeListResponse(data []byte) (SubscriberListResponse, error){
	o := SubscriberListResponse{}
	err := json.Unmarshal(data,&o)
	return o, err 
}
func (r *SubscriberListResponse) Marshal() ([]byte, error){
	return json.Marshal(r)
}

type BlockRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

func UnmarshalBlockRequest(data []byte) (BlockRequest, error){
	o := BlockRequest{}
	err := json.Unmarshal(data,&o)
	return o, err 
}
func (r *BlockRequest) Marshal() ([]byte, error){
	return json.Marshal(r)
}

type BlockResponse struct {
	Success bool `json:"success"`
	Message string `json:"message,omitempty"`
	Code int `json:"code,omitempty"`
}

func UnmarshalBlockResponse(data []byte) (BlockResponse, error){
	o := BlockResponse{}
	err := json.Unmarshal(data,&o)
	return o, err 
}
func (r *BlockResponse) Marshal() ([]byte, error){
	return json.Marshal(r)
}

type UpdateRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

func UnmarshalUpdateRequest(data []byte) (UpdateRequest, error){
	o := UpdateRequest{}
	err := json.Unmarshal(data,&o)
	return o, err 
}
func (r *UpdateRequest) Marshal() ([]byte, error){
	return json.Marshal(r)
}

type UpdateResponse struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
	Message string `json:"message,omitempty"`
	Code int `json:"code,omitempty"`
}

func UnmarshalUpdateResponse(data []byte) (UpdateResponse, error){
	o := UpdateResponse{}
	err := json.Unmarshal(data,&o)
	return o, err 
}
func (r *UpdateResponse) Marshal() ([]byte, error){
	return json.Marshal(r)
}
