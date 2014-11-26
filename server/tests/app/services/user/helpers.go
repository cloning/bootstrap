package user

import (
	"../../../../src/app/services/user"
	"gopkg.in/mgo.v2"
	"testing"
)

const (
	connectionstring = "localhost"
	database         = "intellizen-test-users"
)

type userServiceCallback (func(*user.UserService))

func withUserService(t *testing.T, callback userServiceCallback) {
	// Clean user databse after each test
	defer cleanDb(t)
	userService, err := user.NewUserService(connectionstring, database)

	if err != nil {
		t.Errorf("Error while creating user service %v", err)
		return
	}
	defer userService.Close()

	callback(userService)
}

func cleanDb(t *testing.T) {
	s, err := mgo.Dial(connectionstring)
	if err != nil {
		t.Errorf("Error in test cleanup %v", err)
		return
	}
	s.DB(database).DropDatabase()
}
