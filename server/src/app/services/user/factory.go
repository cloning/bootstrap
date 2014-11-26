package user

import (
	"../../core/mongo"
)

func NewUserService(connectionstring, database string) (*UserService, error) {
	userService := &UserService{
		mongo.MongoServiceBase{
			ConnectionString: connectionstring,
			Database:         database,
		},
	}

	err := userService.Init()

	return userService, err
}
