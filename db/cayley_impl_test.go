package db

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/zkrhm/imd-socialnetwork/model"
)

func populate(cayleyStore *CayleyStore){
	cayleyStore.AddUser("john@example.com")
	cayleyStore.AddUser("mia@example.com")
	cayleyStore.AddUser("kelsey@example.com")
	cayleyStore.AddUser("mike@example.com")
	cayleyStore.AddUser("deva@example.com")
}


var _ = Describe("CayleyImpl - Implementation of db", func() {
	var cayleyStore *CayleyStore
	BeforeEach(func(){
		var err error
		cayleyStore, err = NewCayleyStore()
		if err != nil {
			panic("store creation failed")
		}
		populate(cayleyStore)
	})
	
	Describe("Connect as friend  ", func() {
		Context("Normal Flow - User has been stored but not been friend", func() {

			It("should be a friend", func(){
				By("Connecting john with Mia, Mike and Deva. and check who's john friend")
				cayleyStore.ConnectAsFriend("john@example.com", "mia@example.com")
				cayleyStore.ConnectAsFriend("john@example.com", "mike@example.com")
				cayleyStore.ConnectAsFriend("john@example.com", "deva@example.com")
				
				friends, err := cayleyStore.GetFriendList("john@example.com")
				if err != nil {
					Expect(err).To(BeEmpty())
				}
				// Expect(errors.New("Hello")).ShouldNot(HaveOccurred())

				// Expect(friends).Should(BeEmpty())
				
				Expect(friends).Should(BeEquivalentTo([]User{"mia@example.com","mike@example.com","deva@example.com"}))

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
	})
})