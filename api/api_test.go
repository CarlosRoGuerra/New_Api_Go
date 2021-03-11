package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestCreatUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(createUser))
	defer ts.Close()

	tt := []struct {
		name     string
		password string
		user     types.User
		expected error
	}{
		{
			name: "when create user is sucessfull",

			user: types.User{
				Name:     "test",
				Password: "1234",
			},
		},
	}

	for _, testcase := range tt {
		bbytes, err := json.Marshal(testcase.user)
		assert.NoError(t, err)
		bb := bytes.NewBuffer(bbytes)
		_, err = ts.Client().Post(ts.URL, "application/json", bb)
		assert.NoError(t, err)
	}
}

func TestUpdateUser(t *testing.T) {
	tu := httptest.NewServer(http.HandlerFunc(updateUser))
	defer tu.Close()
	var originalUser map[string]interface{}
	req, _ := http.NewRequest("GET", "/users", nil)
	responde := executeRequest(req)
	json.Unmarshal(responde.Body.Bytes(), &originalUser)
	puser := []byte(`{"name": "test-update","password","123"}`)
	req, _ = http.NewRequest("PUT", "/users", bytes.NewBuffer(puser))
	req.Header.Set("Content-Type", "application/json")

	responde = executeRequest(req)

	checkRespondeCode(t, http.StatusOK, responde.Code)

	var m map[string]interface{}
	json.Unmarshal(responde.Body.Bytes(), &m)

	if m["name"] == originalUser["name"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalUser["name"], m["name"], m["name"])
	}
}

func TestDeleteUser(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users", nil)
	responde := executeRequest(req)
	checkRespondeCode(t, http.StatusOK, responde.Code)

	req, _ = http.NewRequest("DELETE", "/users", nil)
	response := executeRequest(req)
	checkRespondeCode(t, http.StatusOK, response.Code)

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	New().ServeHTTP(rr, req)
	return rr
}

func checkRespondeCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
