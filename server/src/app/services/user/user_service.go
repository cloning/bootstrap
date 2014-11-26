package user

import (
	"../../core/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	mongo.MongoServiceBase
}

const (
	collectionName = "users"
)

// Initialization of user service
// Usually occurs once per application lifecycle
func (this *UserService) Init() error {
	err := this.InitSession()

	if err != nil {
		return err
	}

	this.WithCollection(collectionName, func(c *mgo.Collection) {
		err = c.EnsureIndex(mgo.Index{
			Key:    []string{"email"},
			Unique: true,
		})
	})

	return err
}

func (this *UserService) Create(email, fullName string) (*User, error) {
	var created *User
	var err error

	this.WithCollection(collectionName, func(c *mgo.Collection) {
		err = c.Insert(&User{
			Email:    email,
			FullName: fullName,
		})

		// Duplicate key on email
		// Can't continue
		if mgo.IsDup(err) {
			err = EmailAlreadyExists{}
			return
		}

		if err != nil {
			return
		}

		// Fetch newly created user
		created, err = this.FindFromEmail(email)
	})

	return created, err
}

func (this *UserService) FindFromEmail(email string) (*User, error) {
	var user *User
	var err error

	this.WithCollection(collectionName, func(c *mgo.Collection) {
		err = c.Find(bson.M{
			"email": email,
		}).One(&user)
	})
	return user, err

}
