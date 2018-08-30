package db

import . "github.com/zkrhm/imd-socialnetwork/model"

type IFriendMgtStore interface {
	ConnectAsFriend(user1, user2 User) error
	GetFriendList(user1 User) ([]User, error)
	CommonFriends(user1, user2 User) ([]User, error)
	SubscribeUpdate(user1, user2 User) error
	BlockUpdate(user1, user2 User) error
	GetSubscibers(user1 User) ([]User, error)
}