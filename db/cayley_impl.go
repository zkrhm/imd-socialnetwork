package db

import (
	"errors"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/quad"
	. "github.com/zkrhm/imd-socialnetwork/model"
)

type CayleyStore struct {
	UserStore map[User]int
	quadStore *graph.Handle
	UserCount int
}

func NewCayleyStore() (*CayleyStore, error) {
	memgraph, err := cayley.NewMemoryGraph()
	if err != nil {
		return nil, err
	}
	userStore := make(map[User]int)
	return &CayleyStore{
		quadStore: memgraph,
		UserStore: userStore,
		UserCount: 0,
	}, nil
}

func (s *CayleyStore) AddUser(user User) {
	s.UserCount += 1
	s.UserStore[user] = s.UserCount
}

func (s *CayleyStore) getUserId(user User) int {
	return s.UserStore[user]
}

func (s *CayleyStore) ConnectAsFriend(user1, user2 User) error {
	uid1 := s.getUserId(user1)
	if uid1 == 0 {
		return errors.New("user 1 not registered")
	}

	uid2 := s.getUserId(user2)
	if uid2 == 0 {
		return errors.New("user 2 not registered")
	}

	s.quadStore.AddQuad(quad.Make(user1, "friendsOf", user2, nil))
	s.quadStore.AddQuad(quad.Make(user2, "friendsOf", user1, nil))

	return nil
}

func (s *CayleyStore) GetFriendList(user1 User) ([]User, error) {
	var friends []User
	p := cayley.StartPath(s.quadStore,quad.String(user1)).Out("friendsOf")
	err := p.Iterate(nil).EachValue(nil, func(value quad.Value){
		
		friends = append(friends, User(quad.NativeOf(value).(string)))
	})

	if err != nil {
		return nil, err
	}
	return friends, nil
}
func (s *CayleyStore) CommonFriends(user1, user2 User) ([]User, error) {
	return nil, nil
}
func (s *CayleyStore) SubscribeUpdate(user1, user2 User) error {
	return nil
}
func (s *CayleyStore) BlockUpdate(user1, user2 User) error {
	return nil
}
func (s *CayleyStore) GetSubscibers(user1 User) ([]User, error) {
	return nil, nil
}
