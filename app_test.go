package main_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/zkrhm/imd-socialnetwork"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Friend Management Specs", func() {
	Describe("R1 - As a user I need to create friend connection between two email addresses", func() {

		app := NewApp()
		app.Initialize()
		rr := httptest.NewRecorder()
		handler := app.Router
		Context("Normal Flow - Connecting two user", func() {
			It("return success status", func() {
				const reqBody = `
					{
						"friends":["johnathan@example.com","maria@example.com"]
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
				reqBody := `{"friends":["johnathan@example.com"]}`
				req, err := http.NewRequest("POST", "/connect", strings.NewReader(reqBody))
				if err != nil {

				}
				handler.ServeHTTP(rr, req)

				const expectedResponse = `{"message":"Parameter of friends should exactly has 2 element"}`
				Expect(rr.Code).To(Equal(422))
				log.Print(rr.Body.String())
				Expect(rr.Body.String()).Should(MatchJSON(expectedResponse))
			})
		})

		Context("Alternate Flow - Passing more than two user", func() {
			It("Complains about parameters", func() {
				const reqBody = `
					{
						"friends":["johnathan@example.com","maria@example.com","mercedes@example.com"]
					}`
				req, err := http.NewRequest("POST", "/connect", strings.NewReader(reqBody))
				if err != nil {

				}
				handler.ServeHTTP(rr, req)

				const expectedResponse = `{"message":"Parameter of friends should exactly has 2 element"}`
				Expect(rr.Code).To(Equal(422))
				Expect(rr.Body.String()).Should(MatchJSON(expectedResponse))
			})
		})

		Context("Connecting already connected two user", func() {
			It("Complains that users already friends", func() {

			})
		})

		Context("Connecting non-existent user", func() {
			It("Complains that non-existent user are not available on the system", func() {

			})
		})

		Context("Connecting blocked users", func() {
			It("Throws error that the user cannot be connected because of blockage", func() {

			})
		})
	})

	Describe("R2 - I need to retrieve friend list of an email user", func() {
		Context("Normal Flow - It fetch friend list of email address who already has friend", func() {
			It("returns list of friends", func() {

			})
		})

		Context("Alternate Flow - email address with no friend", func() {
			It("returns list with zero entry", func() {

			})
		})

		Context("Alternate Flow - request non-existing email address", func() {
			It("Throws error that email address is not registered", func() {

			})
		})

	})

	Describe("R3 - I need to retrieve common friends between two email address", func() {
		Context("NF ", func() {
			It("returns common friends of two email address", func() {

			})
		})

		Context("AF1 - two email address have no common friends", func() {
			It("returns zero records without error", func() {

			})
		})

		Context("AF2 - email address is not registered ", func() {
			It("throws error that email address is not registred", func() {

			})
		})
	})

	Describe("R4 - I need API to subscribe to updates from email address", func() {
		Context("Normal Flow", func() {
			It("returns status 'succeess'", func() {

			})
		})
	})

	Describe("R5 - I need an API to block updates from an email address", func() {
		Context("Normal Flow", func() {
			It("returns status 'success'", func() {

			})
		})
	})

	Describe("R6 - I need an API to retrieve all email address that can receive update from an email address", func() {
		Context("Normal Flow", func() {
			It("returns list of subscrbers", func() {

			})

			It("Includes not-blocked user in the result", func() {

			})

			It("Includes subscribed user in the result", func() {

			})

			It("Includes user who is mentioned in an update", func() {

			})
		})

		Context("Alternate Flow - user with no subscribers", func() {
			It("returns empty set of subscriber", func() {

			})
		})

		Context("Alternate FLow - requesting subscribers of non existing email address", func() {
			It("Throws error that the requested email address is non existent", func() {

			})
		})
	})
})
