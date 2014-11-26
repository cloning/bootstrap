package mongo

import (
	"gopkg.in/mgo.v2"
)

type MongoServiceBase struct {
	/*
	   Connectionstring to be used throughout
	*/
	ConnectionString string
	/*
	   Database to be used throughout
	*/
	Database string

	/*
	   The session that will be cloned
	   whenever a new database call is made
	*/
	session *mgo.Session
}

/*
   Initialize the session that will be used throughout the
   service lifecycle.
*/
func (this *MongoServiceBase) InitSession() error {
	var err error
	this.session, err = mgo.Dial(this.ConnectionString)
	return err
}

func (this *MongoServiceBase) Close() {
	this.session.Close()
}

/*
   Defines the callback
*/
type withCollectionCallback func(*mgo.Collection)

/*
   Returns a mgo.Collection reference in the callback
   and makes sure that the connection is properly closed
*/
func (this *MongoServiceBase) WithCollection(collectionName string, callback withCollectionCallback) {
	s := this.session.Copy()
	defer s.Close()
	callback(s.DB(this.Database).C(collectionName))
}
