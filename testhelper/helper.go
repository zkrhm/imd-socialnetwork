package testhelper

import (
	"net/http"
	"net/http/httptest"
	"strings"
	. "github.com/onsi/gomega"
	. "github.com/zkrhm/imd-socialnetwork/model"
	"github.com/zkrhm/imd-socialnetwork/app"
)

func DoConnect(app *app.App, user1 string, user2 string )(res ConnectFriendResponse, err error){
	
	reqBody, err := (&ConnectFriendRequest{
		Friends : []string{user1, user2},
	}).Marshal()

	// reqBody := fmt.Sprint(reqTemplate, user1, user2)

	tr := NewHttpTest("POST","/connect-friend",string(reqBody), app.ConnectAsFriend)
	rr, err := tr.DoRequestTest()
	
	res, err = UnmarshalConnectFriendResponse([]byte(rr.Body.String()))

	return
}

func DoBlock(app *app.App, requestor, target string)(res BlockResponse, err error){
	blockRequest, err := (&BlockRequest{
		Requestor: requestor,
		Target : target , 
	}).Marshal()
	Expect(err).ShouldNot(HaveOccurred())
	
	tr := NewHttpTest("POST","/block",string(blockRequest),app.Block)
	rr , err := tr.DoRequestTest()
	//fmt.Println("body string response : ",rr.Body.String())
	res, err = UnmarshalBlockResponse([]byte(rr.Body.String()))
	return
}

func PostUpdating (app *app.App, user string, text string) (UpdateResponse){

	reqBody, err := (&UpdateRequest{
		Sender: user,
		Text: text,
	}).Marshal()
	Expect(err).ShouldNot(HaveOccurred())

	tr := NewHttpTest("POST","/post-update",string(reqBody),app.PostUpdate)
	rr , err := tr.DoRequestTest()
	
	res,err := UnmarshalUpdateResponse([]byte(rr.Body.String()))
	
	Expect(err).ShouldNot(HaveOccurred())
	return res
}

type HttpTestRequest struct {
	Method string
	Path string
	RequestBody string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
}

func NewHttpTest (method string, path string, requestBody string, handler func(http.ResponseWriter, *http.Request)) (*HttpTestRequest) {
	return &HttpTestRequest{
		Method: method,
		Path: path,
		RequestBody: requestBody,
		HandlerFunc: handler,
	}
}

func (t *HttpTestRequest)DoRequestTest() (rrReturn *httptest.ResponseRecorder, err error) {

	rr := httptest.NewRecorder()
	req, err := http.NewRequest(t.Method,t.Path,strings.NewReader(t.RequestBody))
	if err != nil {
		return nil, err
	}
	http.HandlerFunc(t.HandlerFunc).ServeHTTP(rr, req)

	return rr,  nil
}