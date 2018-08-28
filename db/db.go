package db

type IFriendMgtStore interface {
	ConnectAsFriend(user1, user2 User) error
	GetFriendList(user1 User) []User
	CommonFriends(user1, user2 User) []User
	Subscribe(user1, user2 User) error
	Unsubscribe(user1, user2 User) error
	GetSubscibers(user1 User) []User
}
