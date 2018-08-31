package db

import (
	"encoding/json"
	"errors"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/quad"
	. "github.com/zkrhm/imd-socialnetwork/model"
	"regexp"
	"github.com/deckarep/golang-set"
)

var (
	rSubscribeTo = quad.String("subscribeTo")
	rFriendWith  = quad.String("friendWith")
	rBlock       = quad.String("block")
)

type CayleyStore struct {
	UserStore map[User]int
	quadStore *graph.Handle
	UserCount int
}

type Status struct {
	Message  string
	Mentions []User
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

func (s *CayleyStore) ShowAllRelations() []interface{} {
	var ret []interface{}
	cayley.StartPath(s.quadStore, nil).InPredicates().Iterate(nil).EachValue(nil, func(value quad.Value) {

		ret = append(ret, quad.NativeOf(value))
	})

	return ret
}

func (s *CayleyStore) AddUser(user User) {
	s.UserCount += 1
	s.UserStore[user] = s.UserCount
}

func (s *CayleyStore) getUserId(user User) (int, error) {
	id := s.UserStore[user]
	if id == 0 {
		return id, errors.New("User not registered")
	}
	return id, nil
}

func (s *CayleyStore) isBlocking(user1, user2 User) (bool, error) {

	path := cayley.StartPath(s.quadStore, quad.String(user1)).Has(rBlock, quad.String(user2))
	values, err := path.Iterate(nil).All()
	if err != nil {
		return false, err
	}
	if len(values) == 0 {
		return false, nil
	} else {
		return true, nil
	}

	return false, nil
}

func (s *CayleyStore) whoisBlocking(user User) ([]User, error){
	path := cayley.StartPath(s.quadStore, quad.String(user)).In(rBlock)
	var blockingUser []User
	blockingUser = []User{}
	path.Iterate(nil).EachValue(nil, func(val quad.Value){
		blockingUser = append(blockingUser, User(quad.NativeOf(val).(string)))
	})
	return blockingUser, nil
}

func (s *CayleyStore) ConnectAsFriend(user1, user2 User) error {
	_, err := s.getUserId(user1)
	if err != nil {
		return err
	}

	_, err2 := s.getUserId(user2)
	if err2 != nil {
		return errors.New("user 2 not registered")
	}

	// cayley.StartPath(s.quadStore, nil).Has(rBlock, quad.IRI(user1))
	blocking, err := s.isBlocking(user1, user2)
	if err != nil {
		return err
	}
	if blocking {
		return errors.New("cannot connect as friend user 1 is blocking user 2")
	}
	blocking, err = s.isBlocking(user2, user1)
	if err != nil {
		return err
	}
	if blocking {
		return errors.New("cannot connect as friend, user2 is blocking user1")
	}

	err = s.SubscribeTo(user1,user2)
	if err !=nil {
		return err
	}
	err = s.SubscribeTo(user2,user1)
	if err != nil {
		return err
	}

	s.quadStore.AddQuad(quad.Make(quad.String(user1), rFriendWith, quad.String(user2), nil))
	s.quadStore.AddQuad(quad.Make(quad.String(user2), rFriendWith, quad.String(user1), nil))

	return nil
}

func (s *CayleyStore) GetFriendList(user1 User) ([]User, error) {
	_, err := s.getUserId(user1)
	if err != nil {
		return nil, err
	}
	var friends []User
	p := cayley.StartPath(s.quadStore, quad.String(user1)).Out(rFriendWith)
	err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		friends = append(friends, User(quad.NativeOf(value).(string)))
	})

	if err != nil {
		return nil, err
	}
	return friends, nil
}
func (s *CayleyStore) CommonFriends(user1, user2 User) ([]User, error) {
	var resCommonFriends []User
	resCommonFriends = []User{}

	_, err := s.getUserId(user1)
	if err != nil {
		return resCommonFriends, err
	}
	_, err = s.getUserId(user2)
	if err != nil {
		return resCommonFriends, err
	}

	friendsOfUser1 := cayley.StartPath(s.quadStore, quad.String(user1)).Out(rFriendWith)
	friendsOfUser2 := cayley.StartPath(s.quadStore, quad.String(user2)).Out(rFriendWith)

	commonFriends := friendsOfUser1.And(friendsOfUser2)

	err = commonFriends.Iterate(nil).EachValue(nil, func(value quad.Value) {
		resCommonFriends = append(resCommonFriends, User(quad.NativeOf(value).(string)))
	})

	return resCommonFriends, nil
}
func (s *CayleyStore) SubscribeTo(subscriber, subscribed User) error {
	_, err := s.getUserId(subscriber)
	if err != nil {
		return err
	}
	_, err = s.getUserId(subscribed)
	if err != nil {
		return err
	}

	s.quadStore.AddQuad(quad.Make(quad.String(subscriber), rSubscribeTo, quad.String(subscribed), nil))

	return nil
}
func (s *CayleyStore) BlockUpdate(blocker, blocked User) error {
	var err error
	_, err = s.getUserId(blocker)
	if err != nil {
		return err
	}

	_, err = s.getUserId(blocked)
	if err != nil {
		return err
	}

	s.quadStore.RemoveQuad(quad.Make(quad.String(blocker), rSubscribeTo, quad.String(blocked), nil))
	s.quadStore.AddQuad(quad.Make(quad.String(blocker), rBlock, quad.String(blocked), nil))

	return nil
}

func (s *CayleyStore) getMentions(message string) ([]User, error) {
	re := regexp.MustCompile("[\\w][\\w\\-]+@[\\w\\-]+(\\.[\\w]{2,})+")
	foundEmails := re.FindAllString(message,-1)
	
	var mentionedUsers []User
	for _, email := range foundEmails {
		user := User(string(email))
		_, err := s.getUserId(user)

		if err == nil {
			mentionedUsers = append(mentionedUsers, user)
		}
	}
	return mentionedUsers, nil
}

func (s *CayleyStore) DoUpdate(sender User, message string) ([]User, error) {

	if len(message) <= 0 {
		return nil, errors.New("empty message is not allowed")
	}
	mentionedUsers, err := s.getMentions(message)
	status := &Status{
		Message:  message,
		Mentions: mentionedUsers,
	}
	b, err := json.Marshal(status)
	if err != nil {

	}

	subscribers, err := s.GetSubscribers(sender)
	if err != nil {
		return nil, err
	}

	var recipients []User

	recipients = append(recipients, mentionedUsers...)
	recipients = append(recipients, subscribers...)
	blockings, err := s.whoisBlocking(sender)

	
	recipientsSet := mapset.NewSetFromSlice(ConvertToInterfaceArary(recipients...) )

	blockingSet := mapset.NewSetFromSlice(ConvertToInterfaceArary(blockings...))

	recipientSlice := recipientsSet.Difference(blockingSet).ToSlice()
	returnValue := make([]User, len(recipientSlice))
	for i := range recipientSlice{
		returnValue[i] = recipientSlice[i].(User)
	}
	s.quadStore.AddQuad(quad.Make(sender, "updates", quad.String(string(b)), nil))

	return returnValue , nil
}

func (s *CayleyStore) GetSubscribers(user User) ([]User, error) {
	var subscribers []User
	_, err := s.getUserId(user)
	if err != nil {
		return nil, err
	}
	p := cayley.StartPath(s.quadStore, quad.String(user)).In(rSubscribeTo)
	
	subscribers = []User{}
	p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		subsciber := quad.NativeOf(value)
		subscribers = append(subscribers, User(subsciber.(string)))
	})
	return subscribers, nil
}
