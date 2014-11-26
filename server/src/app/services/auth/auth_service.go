package auth

import (
	"../../core/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CredentialsCollection = "credentials"
	TokenCollection       = "tokens"
)

type AuthService struct {
	mongo.MongoServiceBase
	adminUsers map[string]string
}

func (this *AuthService) Init() error {
	// Initialize base session
	err := this.InitSession()

	if err != nil {
		return err
	}

	this.WithCollection(CredentialsCollection, func(c *mgo.Collection) {
		err = c.EnsureIndex(mgo.Index{
			Key:    []string{"email"},
			Unique: true,
		})
	})

	if err != nil {
		return err
	}

	this.WithCollection(TokenCollection, func(c *mgo.Collection) {
		err = c.EnsureIndex(mgo.Index{
			Key:    []string{"key"},
			Unique: true,
		})
	})

	return err
}

func (this *AuthService) Register(email, password string) error {
	hashedPassword := this.HashPassword(password)

	var err error

	this.WithCollection(CredentialsCollection, func(c *mgo.Collection) {
		err = c.Insert(&Credentials{
			Email:    email,
			Password: hashedPassword,
		})
	})

	return err
}

func (this *AuthService) Login(email, password string) (*Token, error) {

	if this.validate(email, password) == true {
		return this.createAndPresistToken(email)
	}

	return nil, nil
}

func (this *AuthService) ValidateToken(key string) (bool, string) {
	var token *Token
	var err error

	this.WithCollection(TokenCollection, func(c *mgo.Collection) {
		err = c.Find(bson.M{"key": key}).One(&token)
	})

	if token == nil {
		return false, ""
	}
	return true, token.Email
}

func (this *AuthService) HashPassword(password string) string {
	return hashPassword(password)
}

func (this *AuthService) validate(email, password string) bool {
	hashed := this.HashPassword(password)

	if this.adminUsers[email] == hashed {
		return true
	}

	var err error
	var credentials *Credentials

	this.WithCollection(CredentialsCollection, func(c *mgo.Collection) {
		err = c.Find(bson.M{"email": email}).One(&credentials)
	})

	return credentials != nil && credentials.Password == hashed
}

func (this *AuthService) createAndPresistToken(email string) (*Token, error) {

	var err error

	token := &Token{
		Email:   email,
		Key:     generateTokenKey(),
		Expires: this.getExpiresDate(),
	}

	this.WithCollection(TokenCollection, func(c *mgo.Collection) {
		err = c.Insert(token)
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (this *AuthService) getExpiresDate() time.Time {
	// TODO: Move to configuration?
	return time.Now().Add(10 * time.Hour)
}
