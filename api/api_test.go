package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(getUser))
	defer ts.Close()
	var tt = []struct {
		name         string
		body         *bytes.Buffer
		expectedUser types.User
		assertion    func(*testing.T, *http.Response, types.User, error)
	}{
		{
			name: "when user is found",
			expectedUser: types.User{
				Id:       "123",
				Name:     "seila",
				Password: "456",
			},
			body: func() *bytes.Buffer {
				user := types.User{
					Id: "123",
				}
				bbytes, _ := json.Marshal(user)
				return bytes.NewBuffer(bbytes)
			}(),
			assertion: func(t *testing.T, resp *http.Response, expectedUser types.User, err error) {
				assert.NoError(t, err)
				var user types.User
				err = json.NewDecoder(resp.Body).Decode(&user)
				assert.NoError(t, err)
				assert.Equal(t, expectedUser, user)
				assert.Equal(t, 200, resp.StatusCode)
			},
		},
		{
			name: "when user not found",
			body: func() *bytes.Buffer {
				user := types.User{
					Id: "000",
				}
				bbytes, _ := json.Marshal(user)
				return bytes.NewBuffer(bbytes)
			}(),
			assertion: func(t *testing.T, resp *http.Response, expectedUser types.User, err error) {
				assert.NoError(t, err)
				var user types.User
				err = json.NewDecoder(resp.Body).Decode(&user)
				assert.Error(t, err, "EOF")
				assert.Equal(t, expectedUser, user)
				assert.Equal(t, 404, resp.StatusCode)
			},
		},
		{
			name: "when server receive gibberish",
			body: func() *bytes.Buffer {
				user := "ggff"
				bbytes, _ := json.Marshal(user)
				return bytes.NewBuffer(bbytes)
			}(),
			assertion: func(t *testing.T, resp *http.Response, expectedUser types.User, e error) {
				assert.NoError(t, e)
				body, err := ioutil.ReadAll(resp.Body)
				assert.NoError(t, err)
				assert.Equal(t, "json: cannot unmarshal string into Go value of type types.User\n", string(body))
				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
			},
		},
	}
	for _, testcase := range tt {
		resp, err := ts.Client().Post(ts.URL, "application/json", testcase.body)
		testcase.assertion(t, resp, testcase.expectedUser, err)

	}
}

func TestUpdateUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(updateUser))
	defer ts.Close()
	var tt = []struct {
		name         string
		userId       string
		expectedUser types.User
		assertion    func(*testing.T, *http.Response, types.User, error)
	}{
		{
			name: "when user is updated",
			expectedUser: types.User{
				Id:       "123",
				Name:     "seila",
				Password: "456",
			},
			userId: "123",
			assertion: func(t *testing.T, resp *http.Response, expectedUser types.User, err error) {
				assert.NoError(t, err)
				var user types.User
				err = json.NewDecoder(resp.Body).Decode(&user)
				assert.NoError(t, err)
				assert.Equal(t, expectedUser, user)
				assert.Equal(t, 200, resp.StatusCode)
			},
		},
	}
	for _, testcase := range tt {
		user := types.User{
			Id: testcase.userId,
		}
		bbytes, err := json.Marshal(user)
		assert.NoError(t, err)
		bb := bytes.NewBuffer(bbytes)
		resp, err := ts.Client().Post(ts.URL, "application/json", bb)
		testcase.assertion(t, resp, testcase.expectedUser, err)

	}
}

func TestCreateUser(t *testing.T) {
	tp := httptest.NewServer(http.HandlerFunc(createUser))
	defer tp.Close()
	var tt = []struct {
		name         string
		userId       string
		expectedUser types.User
		assertion    func(*testing.T, *http.Response, types.User, error)
	}{
		{
			name: "when user is create",
			expectedUser: types.User{
				Id:       "123",
				Name:     "clemente",
				Password: "456",
			},
			userId: "123",
			assertion: func(t *testing.T, resp *http.Response, expectedUser types.User, err error) {
				assert.NoError(t, err)
				var user types.User
				err = json.NewDecoder(resp.Body).Decode(&user)
				assert.NoError(t, err)
				assert.Equal(t, expectedUser, user)
				assert.Equal(t, 200, resp.StatusCode)
			},
		},
		{
			name:         "when user is empty",
			expectedUser: types.User{},
			assertion: func(t *testing.T, resp *http.Response, expectedUser types.User, err error) {
				assert.NoError(t, err)
				var user types.User
				err = json.NewDecoder(resp.Body).Decode(&user)
				assert.Error(t, err)
				assert.Equal(t, expectedUser, user)
				assert.Equal(t, 404, resp.StatusCode)
			},
		},
	}
	for _, testcase := range tt {
		bbytes, err := json.Marshal(testcase.expectedUser)
		assert.NoError(t, err)
		bb := bytes.NewBuffer(bbytes)
		resp, err := tp.Client().Post(tp.URL, "application/json", bb)
		assert.Equal(t, 200, resp.StatusCode)
		testcase.assertion(t, resp, testcase.expectedUser, err)
	}
}

func TestDeleteUser(t *testing.T) {
	tp := httptest.NewServer(http.HandlerFunc(deleteUser))
	defer tp.Close()
	var tt = []struct {
		name         string
		userId       string
		body         *bytes.Buffer
		expectedUser types.User
		assertion    func(*testing.T, *http.Response, types.User, error)
	}{
		{
			name:         "when user is successfully deleted",
			expectedUser: types.User{},
			body: func() *bytes.Buffer {
				user := types.User{
					Id: "123",
				}
				bbytes, _ := json.Marshal(user)
				return bytes.NewBuffer(bbytes)
			}(),
			assertion: func(t *testing.T, resp *http.Response, expectedUser types.User, err error) {
				assert.NoError(t, err)
				var user types.User
				err = json.NewDecoder(resp.Body).Decode(&user)
				assert.Error(t, err)
				assert.Equal(t, expectedUser, user)
				assert.Equal(t, 200, resp.StatusCode)
			},
		},
		{
			name:         "when user not deleted",
			expectedUser: types.User{},
			assertion: func(t *testing.T, resp *http.Response, expectedUser types.User, err error) {
				assert.NoError(t, err)
				var user types.User
				err = json.NewDecoder(resp.Body).Decode(&user)
				assert.NoError(t, err)
				assert.Equal(t, expectedUser, user)
				assert.Equal(t, 404, resp.StatusCode)
			},
		},
		{
			name: "when server receive gibberish",
			body: nil,
			assertion: func(t *testing.T, resp *http.Response, expectedUser types.User, err error) {
				assert.NoError(t, err)
			},
		},
	}
	for _, testcase := range tt {
		resp, err := tp.Client().Post(tp.URL, "application/json", testcase.body)
		testcase.assertion(t, resp, testcase.expectedUser, err)

	}
}
