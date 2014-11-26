package user

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Email    string        `json:"email"`
	FullName string        `json:"fullName"`
}
