package auth

import (
	"../../core/mongo"
)

func NewAuthService(accountsFile,
	connectionString,
	database string) (*AuthService, error) {

	users, err := parseAccountsFile(accountsFile)

	if err != nil {
		return nil, err
	}

	authService := &AuthService{
		mongo.MongoServiceBase{
			ConnectionString: connectionString,
			Database:         database,
		},
		users,
	}

	authService.Init()

	return authService, nil
}
