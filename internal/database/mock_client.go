package database

import "github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"

type MockClient struct {
	OnGetUsers   func(tableName string) ([]types.User, error)
	OnCreateUser func(tableName string) (types.User, error)
}

func (m *MockClient) GetUsers(tableName string) ([]types.User, error) {
	if m.OnGetUsers != nil {
		return m.OnGetUsers(tableName)
	}
	return nil, nil
}

func (m *MockClient) CreateUser(tableName string) (types.User, error) {
	if m.OnCreateUser != nil {
		return m.OnCreateUser(tableName)
	}
	return types.User{}, nil
}
