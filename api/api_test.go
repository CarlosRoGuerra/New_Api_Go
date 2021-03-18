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

func TestGetUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(getUser))
	defer ts.Close()
	var tt = []struct {
		name         string
		userId       string
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
			name: "when user not found",
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

// func TestUpdateUser(t *testing.T) {

// 	tp := httptest.NewServer(http.HandlerFunc(updateUser))
// 	for _, testcase := range tt {
// 		bbytes, err := json.Marshal(testcase.user)
// 		assert.NoError(t, err)
// 		bb := bytes.NewBuffer(bbytes)
// 		gs, err := tp.Client().Post(tp.URL, "application/json", bb)
// 		response, err := ioutil.ReadAll(gs.Body)
// 		if err != nil {
// 			panic(err)
// 		}
// 		var jsond []types.User
// 		err = json.Unmarshal(response, &jsond)
// 		fmt.Println(jsond)
// 		assert.Equal(t, 200, gs.StatusCode)
// 	}
// }

// func TestGetUser(t *testing.T) {
// 	tp := httptest.NewServer(http.HandlerFunc(getUser))
// 	for _, testcase := range tt {
// 		bbytes, err := json.Marshal(testcase.user)
// 		assert.NoError(t, err)
// 		bb := bytes.NewBuffer(bbytes)
// 		gs, err := tp.Client().Post(tp.URL, "application/json", bb)
// 		assert.Equal(t, 200, gs.StatusCode)
// 	}
// }
// func TestDeleteUser(t *testing.T) {

// 	tp := httptest.NewServer(http.HandlerFunc(updateUser))
// 	for _, testcase := range tt {
// 		bbytes, err := json.Marshal(testcase.user)
// 		assert.NoError(t, err)
// 		bb := bytes.NewBuffer(bbytes)
// 		gs, err := tp.Client().Post(tp.URL, "application/json", bb)
// 		// if gs.StatusCode != 301 {
// 		// 	t.Errorf("Returned wrong status code: got %v want %v", gs.StatusCode, 301)
// 		// }
// 		assert.Equal(t, 301, gs.StatusCode)
// 	}
// }
