package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/internal/database"
	"github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	a := &Api{Router: mux.NewRouter(), Client: &database.MockClient{
		OnGetUsers: func(tableName string) ([]types.User, error) {
			return []types.User{
				{
					Id:       "123",
					Name:     "test",
					Password: "456",
				},
			}, nil

		},
	}}
	ts := httptest.NewServer(http.HandlerFunc(a.getUser))
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
				Name:     "test",
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
				assert.Equal(t, http.StatusOK, resp.StatusCode)
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
				assert.Equal(t, http.StatusNotFound, resp.StatusCode)
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
	a := &Api{Router: mux.NewRouter(), Client: &database.MockClient{
		OnGetUsers: func(tableName string) ([]types.User, error) {
			return []types.User{
				{
					Id:       "123",
					Name:     "test",
					Password: "456",
				},
			}, nil

		},
	}}
	ts := httptest.NewServer(http.HandlerFunc(a.updateUser))
	defer ts.Close()
	var tt = []struct {
		name         string
		userId       string
		body         *bytes.Buffer
		expectedUser types.User
		assertion    func(*testing.T, *http.Response, types.User, error)
	}{
		{
			name: "when user is updated",
			expectedUser: types.User{
				Id:       "123",
				Name:     "test",
				Password: "456",
			},
			body: func() *bytes.Buffer {
				user := types.User{
					Id:       "123",
					Name:     "test",
					Password: "456",
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
				assert.Equal(t, http.StatusOK, resp.StatusCode)
			},
		},
		{
			name: "When not found id",
			body: func() *bytes.Buffer {
				user := "000"
				bbytes, _ := json.Marshal(user)
				return bytes.NewBuffer(bbytes)
			}(),
			assertion: func(t *testing.T, resp *http.Response, expectedUser types.User, e error) {
				assert.NoError(t, e)
				_, err := ioutil.ReadAll(resp.Body)
				assert.NoError(t, err)
				assert.Equal(t, http.StatusNotFound, resp.StatusCode)
			},
		},
	}
	for _, testcase := range tt {
		resp, err := ts.Client().Post(ts.URL, "application/json", testcase.body)
		testcase.assertion(t, resp, testcase.expectedUser, err)
	}
}

func TestCreateUser(t *testing.T) {
	a := &Api{Router: mux.NewRouter(), Client: &database.MockClient{
		OnCreateUser: func(tableName string) (types.User, error) {
			return types.User{
				Id:       "123",
				Name:     "Carlos",
				Password: "456",
			}, nil
		},
	},
	}

	tp := httptest.NewServer(http.HandlerFunc(a.createUser))
	defer tp.Close()
	var tt = []struct {
		name         string
		userId       string
		body         *bytes.Buffer
		expectedUser types.User
		assertion    func(*testing.T, *http.Response, types.User, error)
	}{
		{
			name: "when user is created",
			expectedUser: types.User{
				Id:       "123",
				Name:     "test",
				Password: "456",
			},
			body: func() *bytes.Buffer {
				user := types.User{
					Id:       "123",
					Name:     "test",
					Password: "456",
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
				assert.Equal(t, http.StatusOK, resp.StatusCode)
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
		resp, err := tp.Client().Post(tp.URL, "application/json", testcase.body)
		testcase.assertion(t, resp, testcase.expectedUser, err)
	}
}

func TestDeleteUser(t *testing.T) {
	a := &Api{Router: mux.NewRouter(), Client: &database.MockClient{
		OnGetUsers: func(tableName string) ([]types.User, error) {
			return []types.User{
				{
					Id:       "123",
					Name:     "test",
					Password: "456",
				},
			}, nil

		},
	}}
	tp := httptest.NewServer(http.HandlerFunc(a.deleteUser))
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
				assert.Equal(t, http.StatusOK, resp.StatusCode)
			},
		},
		{
			name:         "when user is not found",
			expectedUser: types.User{},
			body: func() *bytes.Buffer {
				user := types.User{
					Id: "345",
				}
				bbytes, _ := json.Marshal(user)
				return bytes.NewBuffer(bbytes)
			}(),
			assertion: func(t *testing.T, resp *http.Response, expectedUser types.User, err error) {
				assert.NoError(t, err)
				var user types.User
				err = json.NewDecoder(resp.Body).Decode(&user)
				assert.Error(t, err)
				assert.Equal(t, http.StatusNotFound, resp.StatusCode)
			},
		},
	}
	for _, testcase := range tt {
		resp, err := tp.Client().Post(tp.URL, "application/json", testcase.body)
		testcase.assertion(t, resp, testcase.expectedUser, err)

	}
}
