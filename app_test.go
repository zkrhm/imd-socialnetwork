package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/zkrhm/imd-socialnetwork"
	"github.com/zkrhm/imd-socialnetwork/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// . "github.com/Benjamintf1/ExpandedUnmarshalledMatchers"
)

var _ = Describe("Friend Management Specs", func() {
	Describe("R1 - As a user I need to create friend connection between two email addresses", func() {

		app := NewApp()
		store, _ := db.NewCayleyStore()
		app.UseDb(store)
		app.Initialize()
		handler := http.HandlerFunc(app.ConnectAsFriend)

		Context("Normal Flow - Connecting two user", func() {
			It("return success status", func() {
				rr := httptest.NewRecorder()
				const reqBody = `
					{
						"friends":["jonathan@example.com","maria@example.com"]
					}`
				req, err := http.NewRequest("POST", "/connect", strings.NewReader(reqBody))
				if err != nil {

				}
				handler.ServeHTTP(rr, req)

				const expectedResponse = `{"success":true}`
				Expect(rr.Code).To(Equal(200))
				Expect(rr.Body.String()).Should(MatchJSON(expectedResponse))
			})
		})

		Context("Alternate Flow - Passing only one user", func() {
			It("Complain about parameters", func() {
				rr := httptest.NewRecorder()
				reqBody := `{"friends":["johnathan@example.com"]}`
				req, err := http.NewRequest("POST", "/connect", strings.NewReader(reqBody))
				if err != nil {

				}
				handler.ServeHTTP(rr, req)

				const expectedResponse = `{
					"success": false,
					"message": "Parameter of friends should exactly has 2 element",
					"code": 422
				  }`
				Expect(rr.Code).To(Equal(422))
				Expect(rr.Body.String()).Should(MatchJSON(expectedResponse))
			})
		})

		Context("Alternate Flow - Passing more than two user", func() {
			It("Complains about parameters", func() {
				rr := httptest.NewRecorder()
				const reqBody = `
					{
						"friends":["johnathan@example.com","maria@example.com","mercedes@example.com"]
					}`
				req, err := http.NewRequest("POST", "/connect", strings.NewReader(reqBody))
				if err != nil {

				}
				handler.ServeHTTP(rr, req)

				const expectedResponse = `{
					"success": false,
					"message":"Parameter of friends should exactly has 2 element",
					"code": 422
					}`
				Expect(rr.Code).To(Equal(422))
				Expect(rr.Body.String()).Should(MatchJSON(expectedResponse))
			})
		})

		PContext("Connecting already connected two user", func() {
			It("Complains that users already friends", func() {

			})
		})

		Context("Connecting non-existent user", func() {
			It("Complains that non-existent user are not available on the system", func() {
				rr := httptest.NewRecorder()
				const reqBody = `
					{
						"friends" : ["bob@example.com", "nonexistent@example.com"]
					}`

				req, err := http.NewRequest("POST", "/connect", strings.NewReader(reqBody))
				Expect(err).ShouldNot(HaveOccurred())
				handler.ServeHTTP(rr, req)

				const expectedResponse = `{
					"success": false,
					"message":"(404) user 2 not registered",
					"code": 404}`
				By(fmt.Sprintln("returned body : ", rr.Body.String()))
				Expect(rr.Code).To(Equal(404))
				Expect(rr.Body.String()).Should(MatchJSON(expectedResponse))

			})
		})

		PContext("Connecting blocked users", func() {

			It("Throws error that the user cannot be connected because of blockage", func() {

				rr := httptest.NewRecorder()
				const reqBody = `{
					"friends": ["bob@example.com","maria@example.com"]
				}`
				req, err := http.NewRequest("POST", "/connect", strings.NewReader(reqBody))
				Expect(err).ShouldNot(HaveOccurred())
				handler.ServeHTTP(rr, req)

				const expectedResponse = `{

				}`
				Expect(rr.Code).To(Equal(http.StatusForbidden))
			})
		})
	})

	Describe("R2 - I need to retrieve friend list of an email user", func() {

		app := NewApp()
		store, _ := db.NewCayleyStore()
		app.UseDb(store)
		app.Initialize()
		handler := http.HandlerFunc(app.GetFriendList)

		Context("Normal Flow - It fetch friend list of email address who already has friend", func() {
			It("returns list of friends", func() {
				rr := httptest.NewRecorder()
				const reqBody = `{
					"email": "bob@example.com"
				}`
				req, err := http.NewRequest("POST", "/friend-list", strings.NewReader(reqBody))
				Expect(err).ShouldNot(HaveOccurred())
				handler.ServeHTTP(rr, req)

				const expectedResponse = `{
					"success": true,
					"friends": [
						"alice@example.com",
						"charlie@example.com",
						"dani@example.com",
						"fred@example.com"
					],
					"count": 4
				}`
				Expect(rr.Code).To(Equal(http.StatusOK))
				Expect(rr.Body.String()).Should(MatchJSON(expectedResponse))

			})
		})

		Context("Alternate Flow - email address with no friend", func() {
			It("returns list with zero entry", func() {
				rr := httptest.NewRecorder()
				const reqBoby = `{
					"email":"maria@example.com"
				}`
				req, err := http.NewRequest("POST", "/friend-list", strings.NewReader(reqBoby))
				Expect(err).ShouldNot(HaveOccurred())
				handler.ServeHTTP(rr, req)

				const expResponse = `{
					"success": false,
					"code":404,
					"message": "(404) User has no friend"
				}`
				Expect(rr.Code).To(Equal(404))
				Expect(rr.Body.String()).Should(MatchJSON(expResponse))
			})
		})

		Context("Alternate Flow - request non-existing email address", func() {
			It("Throws error that email address is not registered", func() {
				rr := httptest.NewRecorder()
				const reqBoby = `{
					"email": "nonexist@example.com"
				}`
				req, err := http.NewRequest("POST", "/friend-list", strings.NewReader(reqBoby))
				Expect(err).ShouldNot(HaveOccurred())
				handler.ServeHTTP(rr, req)

				const expResponse = `{
					"success": false,
					"message": "(404) User not registered",
					"code":404
				}`
				Expect(rr.Code).To(Equal(http.StatusNotFound))
				Expect(rr.Body.String()).Should(MatchJSON(expResponse))
			})
		})

	})

	Describe("R3 - I need to retrieve common friends between two email address", func() {
		app := NewApp()
		store, _ := db.NewCayleyStore()
		app.UseDb(store)
		app.Initialize()
		handler := http.HandlerFunc(app.GetCommonFriends)
		Context("NF ", func() {
			It("returns common friends of two email address", func() {
				rr := httptest.NewRecorder()
				const reqBody = `{
					"friends": [
						"alice@example.com",
						"dani@example.com"
					]
				}`
				req, err := http.NewRequest("POST", "/common-friends", strings.NewReader(reqBody))
				Expect(err).ShouldNot(HaveOccurred())
				handler.ServeHTTP(rr, req)

				const expRes = `{
					"success": true,
					"friends":[
						"bob@example.com"
					],
					"count": 1
				}`
				Expect(rr.Code).To(Equal(http.StatusOK))
				Expect(rr.Body.String()).Should(MatchJSON(expRes))

			})
		})

		Context("AF1 - two email address have no common friends", func() {
			It("returns zero records without error", func() {
				rr := httptest.NewRecorder()
				const reqBody = `{
					"friends": [
						"alice@example.com",
						"emily@example.com"
					]
				}`
				req, err := http.NewRequest("POST", "/common-friends", strings.NewReader(reqBody))
				Expect(err).ShouldNot(HaveOccurred())
				handler.ServeHTTP(rr, req)

				const expRes = `{
					"success": false,
					"message": "(404) No common friend found",
					"code": 404
				}`
				Expect(rr.Code).To(Equal(http.StatusNotFound))
				Expect(rr.Body.String()).Should(MatchJSON(expRes))
			})
		})

		Context("AF2 - email address is not registered ", func() {
			It("throws error that email address is not registred", func() {
				rr := httptest.NewRecorder()
				const reqBody = `{
					"friends": [
						"unreg@example.com",
						"emily@example.com"
					]
				}`
				req, err := http.NewRequest("POST", "/common-friends", strings.NewReader(reqBody))
				Expect(err).ShouldNot(HaveOccurred())
				handler.ServeHTTP(rr, req)

				const expRes = `{
					"success": false,
					"message": "(404) User not registered",
					"code": 404
				}`
				Expect(rr.Code).To(Equal(http.StatusNotFound))
				Expect(rr.Body.String()).Should(MatchJSON(expRes))
			})
		})
	})

	Describe("R4 - I need API to subscribe to updates from email address", func() {

		app := NewApp()
		store, _ := db.NewCayleyStore()
		app.UseDb(store)
		app.Initialize()
		handler := http.HandlerFunc(app.Subscribe)

		Context("Normal Flow", func() {
			It("returns status 'succeess'", func() {
				rr := httptest.NewRecorder()
				const reqBody = `{
					"requestor": "maria@example.com",
					"target": "bob@example.com"
				}`
				req, err := http.NewRequest("POST", "/subscribe", strings.NewReader(reqBody))
				Expect(err).ShouldNot(HaveOccurred())
				handler.ServeHTTP(rr, req)

				const expRes = `{
					"success":true
				}`
				Expect(rr.Code).To(Equal(http.StatusOK))
				Expect(rr.Body.String()).Should(MatchJSON(expRes))
			})
		})

		Context("Alt Flow - not registered users ", func(){
			It("Should be complain when passing nonexistent user", func() {

				By("passing non existent user as requestor")
				rr := httptest.NewRecorder()
				const reqBody = `{
					"requestor": "nonexistent@example.com",
					"target": "bob@example.com"
				}`
				req, err := http.NewRequest("POST", "/subscribe", strings.NewReader(reqBody))
				Expect(err).ShouldNot(HaveOccurred())
				handler.ServeHTTP(rr, req)

				const expRes = `{
					"success":false,
					"message": "(404) User (subscriber) not registered",
					"code": 404
				}`
				Expect(rr.Code).To(Equal(http.StatusNotFound))
				Expect(rr.Body.String()).Should(MatchJSON(expRes))

				By("passing non existent user as target")
				rr = httptest.NewRecorder()
				const reqBody2 = `{
					"requestor": "emily@example.com",
					"target": "nonexistent@example.com"
				}`
				req, err = http.NewRequest("POST", "/subscribe", strings.NewReader(reqBody2))
				Expect(err).ShouldNot(HaveOccurred())
				handler.ServeHTTP(rr, req)

				const expRes2 = `{
					"success":false,
					"message": "(404) User (subscribed) not registered",
					"code": 404
				}`
				Expect(rr.Code).To(Equal(http.StatusNotFound))
				Expect(rr.Body.String()).Should(MatchJSON(expRes2))
			})
		})
	})

	Describe("R5 - I need an API to block updates from an email address", func() {

		app := NewApp()
		store, _ := db.NewCayleyStore()
		app.UseDb(store)
		app.Initialize()
		handler := http.HandlerFunc(app.Block)

		Context("NF", func(){
			It("should be success", func(){
					By("blocking non-friend user")
					rr := httptest.NewRecorder()
					const reqBody = `{"requestor":"maria@example.com", "target":"bob@example.com"}`
					req, err := http.NewRequest("POST", "/block", strings.NewReader(reqBody))
					Expect(err).ShouldNot(HaveOccurred())
					handler.ServeHTTP(rr, req)
	
					fmt.Println("http response:",rr.Body.String())

					Expect(err).ShouldNot(HaveOccurred())
					Expect(rr.Code).To(Equal(200))
			})
			It("should be success", func(){
				By("blocking already friend user")
				rr := httptest.NewRecorder()
				const reqBody = `{"requestor":"alice@example.com", "target":"bob@example.com"}`
				req, err := http.NewRequest("POST", "/block", strings.NewReader(reqBody))
				Expect(err).ShouldNot(HaveOccurred())
				handler.ServeHTTP(rr, req)

				fmt.Println("http response:",rr.Body.String())

				Expect(err).ShouldNot(HaveOccurred())
				Expect(rr.Code).To(Equal(200))
			})
		})

		
	})

	Describe("R6 - I need an API to retrieve all email address that can receive update from an email address", func() {
		app := NewApp()
		store, _ := db.NewCayleyStore()
		app.UseDb(store)
		app.Initialize()
		handler := http.HandlerFunc(app.PostUpdate)

		Context("Normal Flow", func() {
			It("returns list of subscrbers", func() {
				By("blocking non-friend user")
					rr := httptest.NewRecorder()
					const reqBody = `{
						"sender":"bob@example.com", 
						"text":"Hello Folks!"
					}`
					req, err := http.NewRequest("POST", "/post-update", strings.NewReader(reqBody))
					Expect(err).ShouldNot(HaveOccurred())
					handler.ServeHTTP(rr, req)
	
					fmt.Println("http response:",rr.Body.String())

					Expect(err).ShouldNot(HaveOccurred())
					Expect(rr.Code).To(Equal(200))
			})

			PIt("Includes not-blocked user in the result", func() {

			})

			PIt("Includes subscribed user in the result", func() {

			})

			PIt("Includes user who is mentioned in an update", func() {

			})
		})

		PContext("Alternate Flow - user with no subscribers", func() {
			It("returns empty set of subscriber", func() {

			})
		})

		PContext("Alternate FLow - requesting subscribers of non existing email address", func() {
			It("Throws error that the requested email address is non existent", func() {

			})
		})
	})
})
