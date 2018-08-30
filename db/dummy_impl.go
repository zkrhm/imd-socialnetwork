package db

import (
	"container/list"
	"errors"

	. "github.com/zkrhm/imd-socialnetwork/model"
)

type Relation int
type GraphAdjacent map[int]*list.List

const (
	FRIEND      Relation = 1
	SUBSCRIBING Relation = 2
	BLOCKING    Relation = 3
)

type DummyStore struct {
	Nodes      map[User]int
	Friends    GraphAdjacent
	Subscibing GraphAdjacent
	Blocking   GraphAdjacent
	LastId     int
}

func NewDummyStore() *DummyStore {
	return &DummyStore{LastId: 0}
}

func (s *DummyStore) AddUser(user User) error {
	id := s.Nodes[user]
	if id < 0 {
		return errors.New("Already Existed")
	}
	s.Nodes[user] = s.LastId
	s.LastId += 1
	return nil
}

func (s *DummyStore) AddUsers(users []User) error {
	for _, user := range users {
		err := s.AddUser(user)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DummyStore) getRelation(r Relation) GraphAdjacent {
	return map[Relation]GraphAdjacent{
		FRIEND:      s.Friends,
		SUBSCRIBING: s.Subscibing,
		BLOCKING:    s.Blocking,
	}[r]
}

func (s *DummyStore) SetRelation(r Relation, user1 User, user2 User) error {
	relation := s.getRelation(r)
	fUserId := s.getUserId(user1)
	sUserId := s.getUserId(user2)

	relation[fUserId].PushBack(sUserId)

	return nil
}

func (s *DummyStore) getUserId(user User) int {
	return s.Nodes[user]
}

func (s *DummyStore) ConnectAsFriend(user1, user2 User) error {
	userId1 := s.getUserId(user1)
	userId2 := s.getUserId(user2)

	user1Friends := s.Friends[userId1]
	for e := user1Friends.Front(); e != nil; e = e.Next() {
		if e.Value == userId2 {
			return errors.New("Already Friend")
		}
	}
	user2Friends := s.Friends[userId2]
	for e := user2Friends.Front(); e != nil; e = e.Next() {
		if e.Value == userId1 {
			return errors.New("Already Friend")
		}
	}

	s.Friends[userId1].PushBack(userId2)
	s.Friends[userId2].PushBack(userId1)

	return nil
}

func (store *DummyStore) GetFriendList(user1 User) ([]User, error) {
	return nil, nil
}

func (store *DummyStore) CommonFriends(user1, user2 User) ([]User, error) {
	return nil, nil
}
func (store *DummyStore) SubscribeUpdate(user1, user2 User) error {
	return nil
}
func (store *DummyStore) BlockUpdate(user1, user2 User) error {
	return nil
}
func (store *DummyStore) GetSubscibers(user1 User) ([]User, error) {
	return nil, nil
}
