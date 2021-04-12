package database

import "github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"

type DatabaseClient interface {
	GetUsers(tableName string) ([]types.User, error)
	CreateUser(tableName string, user types.User) (types.User, error)
	DeleteUser(tableName string, user types.User) error
}
