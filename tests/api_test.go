package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/niconosenzo/devopsapi/pkg/app"
	"github.com/niconosenzo/devopsapi/pkg/app/model"
)

var a *app.App

func TestMain(m *testing.M) {
	a = &app.App{}
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

//TestGetCorrectUser check if a particular user is retrieved
func TestGetUser(t *testing.T) {

	request, err := http.NewRequest(http.MethodGet, "/user/1", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	response := executeRequest(request)

	// Check request went fine
	checkResponseCode(t, http.StatusOK, response.Code)

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	var got model.User

	err = json.Unmarshal(b, &got)
	if err != nil {
		t.Fatalf("could not unmarshall response %v", err)
	}

	want := model.User{ID: "1", Name: "Jose", Surname: "Perez"}

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

//TestGetNonExistenttUser check response when requested a non-existent user.
func TestGetNonExistenttUser(t *testing.T) {

	request, err := http.NewRequest(http.MethodGet, "/user/5", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	response := executeRequest(request)

	//Check request returns 404 (NotFound)
	checkResponseCode(t, http.StatusNotFound, response.Code)

}

func TestCreateUser(t *testing.T) {

	var jsonUser = []byte(`{"id": "4", "name": "Pp", "surname": "Mm"}`)
	request, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	response := executeRequest(request)

	// Check request went fine
	checkResponseCode(t, http.StatusOK, response.Code)

	// Retrieve the new user and test
	var got model.User
	err = json.Unmarshal(response.Body.Bytes(), &got)

	if err != nil {
		t.Fatalf("could not unmarshall response %v", err)
	}

	want := model.User{ID: "4", Name: "Pp", Surname: "Mm"}

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}

func TestDeleteUser(t *testing.T) {

	request, err := http.NewRequest(http.MethodDelete, "/user/2", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	response := executeRequest(request)

	// Check request went fine
	checkResponseCode(t, http.StatusOK, response.Code)

	// Check user no longer exists
	request, _ = http.NewRequest(http.MethodGet, "/user/2", nil)
	response = executeRequest(request)

	// Check request went fine
	checkResponseCode(t, http.StatusNotFound, response.Code)

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
