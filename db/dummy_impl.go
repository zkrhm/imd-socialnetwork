package db

type DummyStore struct {
}

func NewDummyStore() *DummyStore {
	return &DummyStore{}
}

func (store *DummyStore) ConnectAsFriend(user1, user2 string) {

}
