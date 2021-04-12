package database

import "github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"

type MockClient struct {
	OnGetUsers   func(tableName string) ([]types.User, error)
	OnCreateUser func(tableName string, user types.User) (types.User, error)
	OnDeleteUser func(tableName string, user types.User) error
}

func (m *MockClient) GetUsers(tableName string) ([]types.User, error) {
	if m.OnGetUsers != nil {
		return m.OnGetUsers(tableName)
	}
	return nil, nil
}

func (m *MockClient) CreateUser(tableName string, user types.User) (types.User, error) {
	if m.OnCreateUser != nil {
		return m.OnCreateUser(tableName, user)
	}
	return types.User{}, nil
}

func (m *MockClient) DeleteUser(tableName string, user types.User) error {
	if m.OnDeleteUser != nil {
		return m.OnDeleteUser(tableName, user)
	}
	return nil
}
