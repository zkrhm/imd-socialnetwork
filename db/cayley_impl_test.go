
package db

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/zkrhm/imd-socialnetwork/model"
	"strings"
)

func populate(cayleyStore *CayleyStore) {
	cayleyStore.AddUser("john@example.com")
	cayleyStore.AddUser("mia@example.com")
	cayleyStore.AddUser("kelsey@example.com")
	cayleyStore.AddUser("mike@example.com")
	cayleyStore.AddUser("deva@example.com")
}

func connectFriends(store *CayleyStore) {
	err1 := store.ConnectAsFriend("john@example.com", "mia@example.com")
	err2 := store.ConnectAsFriend("john@example.com", "mike@example.com")
	err3 := store.ConnectAsFriend("john@example.com", "deva@example.com")
	Expect(err1).ShouldNot(HaveOccurred())
	Expect(err2).ShouldNot(HaveOccurred())
	Expect(err3).ShouldNot(HaveOccurred())
}

func storeAndConnect(store *CayleyStore) *CayleyStore {
	store.AddUser("alice@example.com")
	store.AddUser("bob@example.com")
	store.AddUser("charlie@example.com")
	store.AddUser("dani@example.com")
	store.AddUser("greg@example.com")
	store.AddUser("fred@example.com")
	store.AddUser("emily@example.com")

	store.ConnectAsFriend("alice@example.com", "bob@example.com")
	store.ConnectAsFriend("charlie@example.com", "bob@example.com")
	store.ConnectAsFriend("dani@example.com", "bob@example.com")
	store.ConnectAsFriend("fred@example.com", "bob@example.com")

	store.ConnectAsFriend("charlie@example.com", "dani@example.com")
	store.ConnectAsFriend("dani@example.com", "greg@example.com")
	store.ConnectAsFriend("greg@example.com", "fred@example.com")
	store.ConnectAsFriend("fred@example.com", "emily@example.com")

	return store
}

func subscribe(store *CayleyStore){
	err := store.SubscribeTo("alice@example.com","bob@example.com")
	Expect(err).ShouldNot(HaveOccurred())
	err = store.SubscribeTo("charlie@example.com","bob@example.com")
	Expect(err).ShouldNot(HaveOccurred())
	err = store.SubscribeTo("dani@example.com","bob@example.com")
	Expect(err).ShouldNot(HaveOccurred())
}

var _ = Describe("CayleyImpl - Implementation of db", func() {
	var cayleyStore *CayleyStore
	BeforeEach(func() {
		var err error
		cayleyStore, err = NewCayleyStore()
		if err != nil {
			panic("store creation failed")
		}
		populate(cayleyStore)
		connectFriends(cayleyStore)
		storeAndConnect(cayleyStore)
	})

	Describe("SPEC 1 : Connect as friend  ", func() {
		Context("Normal Flow - User has been stored but not been friend", func() {

			It("should be a friend", func() {
				By("Connecting john with Mia, Mike and Deva. and check who's john friend")

				friends, err := cayleyStore.GetFriendList("john@example.com")
				if err != nil {
					Expect(err).To(BeEmpty())
				}
				// Expect(errors.New("Hello")).ShouldNot(HaveOccurred())

				// Expect(friends).Should(BeEmpty())

				Expect(friends).Should(BeEquivalentTo([]User{"mia@example.com", "mike@example.com", "deva@example.com"}))

				By("checking who is mia's friend")

				friends, err = cayleyStore.GetFriendList("mia@example.com")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(friends).Should(BeEquivalentTo([]User{"john@example.com"}))

				By("checking who is mike's friend")

				friends, err = cayleyStore.GetFriendList("mike@example.com")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(friends).Should(BeEquivalentTo([]User{"john@example.com"}))

				By("checking who is deva's friend")

				friends, err = cayleyStore.GetFriendList("deva@example.com")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(friends).Should(BeEquivalentTo([]User{"john@example.com"}))
			})

		})

		Context("Alternate Flow - connecting unregisted user", func() {
			It("Should throw error", func() {
				By("Passing unknown user on first param")
				err := cayleyStore.ConnectAsFriend("who@example.com", "john@example.com")
				Expect(err).Should(HaveOccurred())

				By("Passing unknown user on second param")
				err = cayleyStore.ConnectAsFriend("john@example.com", "who@example.com")
				Expect(err).Should(HaveOccurred())
			})
		})

		PContext("Alternate Flow - connecting blocking user", func() {
			It("Should complains that the user is blocked and cannot be connected as a friend", func() {
				err := cayleyStore.ConnectAsFriend("john@example.com", "blockinguser@example.com")
				Expect(err).Should(HaveOccurred())
			})

		})
	})

	Describe("SPEC 2: Get Friend List", func() {

		Context("Normal Flow - Getting registered connected people ", func() {

			It("All right", func() {

				By("checking john's friend")
				friends, err := cayleyStore.GetFriendList("john@example.com")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(friends) > 0).To(BeTrue())

				By("Checking mia's friend to make sure friendship is not a digraph relation")

				friends, err = cayleyStore.GetFriendList("mia@example.com")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(friends) > 0).To(BeTrue())
			})
		})

		Context("Alt Flow - get friends of unregistered user", func() {
			It("complains that user is not registered", func() {
				friends, err := cayleyStore.GetFriendList("nonexistent@example.com")
				Expect(err).Should(HaveOccurred())
				Expect(friends).To(BeEmpty())
			})
		})

		Context("Alt Flow - get non connected people", func() {
			It("returns empty list", func() {
				friends, err := cayleyStore.GetFriendList("kelsey@example.com")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(friends).To(BeEmpty())
			})
		})
	})

	Describe("SPEC 3: Get Common Friends", func() {
		Context("NF - Having Common Friends", func() {
			It("return common friend between them", func() {

				store := cayleyStore

				fmt.Println("All relations : ", store.ShowAllRelations())

				charliesFriend, _ := store.GetFriendList("charlie@example.com")
				fredsFriend, _ := store.GetFriendList("fred@example.com")

				fmt.Println("charlies friend : ", charliesFriend)

				sCharliesFriend := ConvertToStringArray(charliesFriend)
				sFredFriend := ConvertToStringArray(fredsFriend)

				By(fmt.Sprintf("using followings ==> charlie's friends: %s, fred's friends: : %s ", strings.Join(sCharliesFriend[:], ","), strings.Join(sFredFriend[:], ",")))

				By("checking charlie and fred's common friends: should be bob")
				commonFriends, err := store.CommonFriends("charlie@example.com", "fred@example.com")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(commonFriends).Should(BeEquivalentTo([]User{"bob@example.com"}))

				By("checking bob and emily's common friends: should be fred")
				commonFriends, err = store.CommonFriends("bob@example.com", "emily@example.com")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(commonFriends).Should(BeEquivalentTo([]User{"fred@example.com"}))

			})
		})

		Context("Alt Flow - Have no Common Friends", func() {
			It("return empty list", func() {
				By("Lookup to emily and charlie's friend")
				commonFriends, err := cayleyStore.CommonFriends("charlie@example.com", "emily@example.com")
				sCommonFriends := ConvertToStringArray(commonFriends)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(commonFriends) == 0).Should(BeTrue())
				By(fmt.Sprintln("common friends : ", strings.Join(sCommonFriends, ",")))
				Expect(commonFriends).Should(BeEquivalentTo([]User{}))
			})
		})

		Context("Misc - Make sure email address is registered ", func() {

			It("complains if user not regisred", func() {
				By("Passing non registered user as first param")
				friends, err := cayleyStore.CommonFriends("unknown@example.com", "emily@example.com")
				Expect(err).Should(HaveOccurred())
				Expect(friends).Should(BeEquivalentTo([]User{}))
				By("Passing non registered user as second param")
				friends, err = cayleyStore.CommonFriends("charlie@example.com", "unknown@example.com")
				Expect(err).Should(HaveOccurred())
				Expect(friends).Should(BeEquivalentTo([]User{}))
			})
		})

		Context("Misc - make sure passed email address have correct format", func() {
			It("complains if format is not correct", func() {
				By("passing incorrect format as first param")
				_, err := cayleyStore.CommonFriends("notanemailaddr", "emily@example.com")
				Expect(err).Should(HaveOccurred())
				By("passing incorrect format as second param")
				_, err = cayleyStore.CommonFriends("charlie@xample.com", "notanemailaddr")
				Expect(err).Should(HaveOccurred())
			})
		})

	})

	Describe("SPEC 4 : Follow Updates", func() {
		Context("NF - folow [action] successful", func() {
			It("shows that user in followed users's follower list", func() {
				store := cayleyStore
				subscribe(store)

				relations := store.ShowAllRelations()
				By(fmt.Sprintln("showing relations : ", relations))

				By("subscribing to one who already friend ")

				subscribers, err := cayleyStore.GetSubscribers("bob@example.com")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(subscribers).Should(ContainElement(User("alice@example.com")))

				By("subscribing to one who has not already a friend")
				cayleyStore.SubscribeTo("emily@example.com","bob@example.com")
				subscribers, err = cayleyStore.GetSubscribers("bob@example.com")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(subscribers).Should(ContainElement(User("emily@example.com")))
			})
		})

		Context("AF = user not registerd", func() {
			It("complains that user is not registered", func() {
				store := cayleyStore
				By("passing unregistered user as followed")
				err := store.SubscribeTo("alice@example.com","unknown@example.com")
				Expect(err).Should(HaveOccurred())
				By("passing unregistered user as follower")
				err = store.SubscribeTo("unknown@example.com","alice@example.com")
				Expect(err).Should(HaveOccurred())
			})
		})

		Context("AF - incorrect user's email address format", func() {
			It("Complains about email address format", func() {
				store := cayleyStore
				By("passing wrong format as followed")
				err := store.SubscribeTo("alice@example.com","notanemail")
				Expect(err).Should(HaveOccurred())
				By("passing wrong format as follower")
				err = store.SubscribeTo("alice@example.com","notanemail")
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe("SPEC 5 : Block Updates", func() {
		Context("NF - block [action] successfull", func() {
			Context("users already friend, then blocking his friend", func(){
				It("does not show the blocker user in blocked user's subscriber list ", func() {
					store := cayleyStore
					subscribe(store)

					By("person who already friend do blocking")
					err := store.BlockUpdate("alice@example.com","bob@example.com")
					Expect(err).ShouldNot(HaveOccurred())

					subscribers, err := store.GetSubscribers("bob@example.com")
					Expect(err).ShouldNot(HaveOccurred())
					Expect(subscribers).ShouldNot(ContainElement(User("alice@example.com")))
				})
			})

			Context("user has not been a friend, then blocking.. ", func(){
				It("complains when user want to connect", func(){
					store := cayleyStore
					subscribe(store)

					err := store.BlockUpdate("emily@example.com","charlie@example.com")
					Expect(err).ShouldNot(HaveOccurred())

					err = store.ConnectAsFriend("emily@example.com","charlie@example.com")
					Expect(err).Should(HaveOccurred())
				})
			})
			
		})
		Context("AF - format & registered user check", func() {
			It("Complains that user is not registered, when using unregistered user", func() {
				store := cayleyStore
				By("passing unregister user as blocker")
				err := store.BlockUpdate("unreg@example.com","charlie@example.com")
				Expect(err).Should(HaveOccurred())
				By("passing unregister user as blocked")
				err = store.BlockUpdate("unreg@example.com","charlie@example.com")
				Expect(err).Should(HaveOccurred())
			})

			It("Complains that email address is not correct", func() {
				store := cayleyStore
				By("passing incorrect format as blocker")
				err := store.BlockUpdate("notanemail","charlie@example.com")
				Expect(err).Should(HaveOccurred())
				By("passing incorrect format as blocked")
			})

		})
	})

	Describe("SUPPORTING FUNCTION test ", func(){
		Context("get mentioned users in a status update", func(){
			It("list email address mentioned in a message", func(){
				mentioned, err := cayleyStore.getMentions("yoohoo greg@example.com , how are you kylie@example.com xoxo \n maria@panda.co.cn \n john@jw-org.co dani@example.com")
				Expect(err).ShouldNot(HaveOccurred())
				By("mentioning non user email, should not be listed as recipients")
				Expect(mentioned).ShouldNot(ContainElement(User("kylie@example.com")))
				By("mentioning registered user email, should be listed as recipients")
				Expect(mentioned).Should(ContainElement(User("greg@example.com")))
				Expect(mentioned).Should(ContainElement(User("dani@example.com")))
			})
		})
	})

	FDescribe("SPEC 6 : List Post recipients", func() {

		Context("NF - show list of follower", func() {

			It("displays friends", func() {
				store := cayleyStore
				subscribe(store)

				recipients, err := store.DoUpdate(User("bob@example.com"),"hello world")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(recipients).Should(ConsistOf([]User{"alice@example.com","charlie@example.com","dani@example.com","fred@example.com"}))
			})

			It("displays users who follows", func() {
				By("not friend user <emily> subscribing to <bob>, emily should be a recipient of bob updates")
				store := cayleyStore
				subscribe(store)
				err := store.SubscribeTo("emily@example.com","bob@example.com")
				Expect(err).ShouldNot(HaveOccurred())

				recipients , err := store.DoUpdate("bob@example.com","Hello world")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(recipients).Should(ContainElement(User("emily@example.com")))

			})

			It("contains users who mentioned at post", func() {
				By("not friend user <greg> mentioned in the updates by <bob>. <greg> should be a recipient ")
				store := cayleyStore
				subscribe(store)

				recipient , err := store.DoUpdate("bob@example.com","you rocks greg@example.com. cc: kylie-jenner@kadarshian.com")
				Expect(err).ShouldNot(HaveOccurred())
				sRecipients := ConvertToStringArray(recipient)
				By(fmt.Sprintln("showing who are the recipient ",strings.Join(sRecipients,",")))
				Expect(recipient).Should(ContainElement(User("greg@example.com")))
				Expect(recipient).ShouldNot(ContainElement(User("kylie-jenner@kadarshian.com")))
			})

			It("not displaying user who is blocking ", func() {
				store := cayleyStore
				subscribe(store)

				err := store.BlockUpdate("alice@example.com","bob@example.com")
				Expect(err).ShouldNot(HaveOccurred())
				By("not mentioning user who blocking")
				recipients, err := store.DoUpdate("bob@example.com","Hello World")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(recipients).ShouldNot(ContainElement(User("alice@example.com")))

				By("mentioning user who blocking")
				recipients, err = store.DoUpdate("bob@example.com","Hello alice@example.com")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(recipients).ShouldNot(ContainElement(User("alice@example.com")))
			})
		})

		Context("AF - format & registered user check", func() {
			It("rejects empty message", func() {
				_, err := cayleyStore.DoUpdate("bob@example.com","")
				Expect(err).Should(HaveOccurred())
			})
			It("complains that user is not registered", func() {
				_, err := cayleyStore.DoUpdate("unreg@example.com","hello...")
				Expect(err).Should(HaveOccurred())
			})
			It("complains that email addres format is not correct", func() {
				_, err := cayleyStore.DoUpdate("invalidemailaddr","hello...")
				Expect(err).Should(HaveOccurred())
			})
		})
	})

})
