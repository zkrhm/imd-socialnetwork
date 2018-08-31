package db

import . "github.com/zkrhm/imd-socialnetwork/model"

type IFriendMgtStore interface {
	ConnectAsFriend(user1, user2 User) error
	GetFriendList(user1 User) ([]User, error)
	CommonFriends(user1, user2 User) ([]User, error)
	SubscribeTo(user1, user2 User) error
	BlockUpdate(user1, user2 User) error
	GetSubscribers(user1 User) ([]User, error)
	DoUpdate(user User, message string) ([]User, error)
}
