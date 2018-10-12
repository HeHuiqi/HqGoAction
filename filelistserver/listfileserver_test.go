package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
	"os"
	"github.com/pkg/errors"
)

func errPanic(w http.ResponseWriter, r *http.Request) error{

	panic(123)
	//return nil
}


type testUserError string

func (e testUserError)Error() string  {
	return e.Message()
}
func (e testUserError) Message() string  {
	return string(e)
}
func errUserError(w http.ResponseWriter, r *http.Request) error{

	return testUserError("user error")
}

func errNotFound(w http.ResponseWriter, r *http.Request) error{

	return os.ErrNotExist
}
func errNotPermission(w http.ResponseWriter, r *http.Request) error{

	return os.ErrPermission
}
func errUnknown(w http.ResponseWriter, r *http.Request) error{

	return errors.New("Unknown error")
}
func noError(w http.ResponseWriter, r *http.Request) error{

	return nil
}

var tests = []struct{
	h appHandler
	code int
	message string
}{
	{errPanic,500,"Internal server error"},
	{errUserError,400,"user error"},
	{errNotFound,404,"NotFound"},
	{errNotPermission,403,"forbidden"},
	{errUnknown,500,"Internal server error"},
	{noError,200,"no error"},
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {


	b,_ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b),"\n")
	if resp.StatusCode != expectedCode ||
		body != expectedMsg{
		t.Errorf("expect(%d,%s);"+"got (%d,%s)",expectedCode,expectedMsg,resp.StatusCode,body)
	}

}

func TestErrWrapper(t *testing.T)  {

	for _,tt := range tests{
		f := errWrapper(tt.h)
		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet,"http://www.baidu.com",nil)

		f(resp,req)
		//b,_ := ioutil.ReadAll(resp.Body)
		//body := strings.Trim(string(b),"\n")
		//if resp.Code != tt.code ||
		//	body != tt.message{
		//	t.Errorf("expect(%d,%s);"+"got (%d,%s)",tt.code,tt.message,resp.Code,body)
		//}
		verifyResponse(resp.Result(),tt.code,tt.message,t)
	}
}

func TestErrWrapperInServer(t *testing.T)  {
	for _,tt := range tests{
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp,_:= http.Get(server.URL)

		//b,_ := ioutil.ReadAll(resp.Body)
		//body := strings.Trim(string(b),"\n")
		//if resp.StatusCode != tt.code ||
		//	body != tt.message{
		//	t.Errorf("expect(%d,%s);"+"got (%d,%s)",tt.code,tt.message,resp.StatusCode,body)
		//}

		verifyResponse(resp,tt.code,tt.message,t)

	}
}