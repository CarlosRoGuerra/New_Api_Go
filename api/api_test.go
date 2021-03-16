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

var tt = []struct {
	name     string
	password string
	user     types.User
	expected error
}{
	{
		name: "when create user is sucessfull",

		user: types.User{
			Id:       "1",
			Name:     "test-carlos",
			Password: "1234",
		},
	},
}

func TestCreatUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(createUser))
	defer ts.Close()

	for _, testcase := range tt {
		bbytes, err := json.Marshal(testcase.user)
		assert.NoError(t, err)
		bb := bytes.NewBuffer(bbytes)
		gs, err := ts.Client().Post(ts.URL, "application/json", bb)
		if gs.StatusCode != http.StatusCreated {
			t.Errorf("Handler returned wrong status code: got %v want %v", gs.StatusCode, http.StatusCreated)
		}
	}
}

func TestUpdateUser(t *testing.T) {

	tp := httptest.NewServer(http.HandlerFunc(updateUser))
	for _, testcase := range tt {
		bbytes, err := json.Marshal(testcase.user)
		assert.NoError(t, err)
		bb := bytes.NewBuffer(bbytes)
		gs, err := tp.Client().Post(tp.URL, "application/json", bb)
		if gs.StatusCode != 301 {
			t.Errorf("Handler returned wrong status code: got %v want %v", gs.StatusCode, 301)
		}
	}
}
func TestDeleteUser(t *testing.T) {

	tp := httptest.NewServer(http.HandlerFunc(updateUser))
	for _, testcase := range tt {
		bbytes, err := json.Marshal(testcase.user)
		assert.NoError(t, err)
		bb := bytes.NewBuffer(bbytes)
		gs, err := tp.Client().Post(tp.URL, "application/json", bb)
		if gs.StatusCode != 301 {
			t.Errorf("Handler returned wrong status code: got %v want %v", gs.StatusCode, 301)
		}
	}
}
