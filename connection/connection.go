package connection

import (
	"labix.org/v2/mgo"
)

var (
	mgoSession   *mgo.Session
	databaseName = "msgme"
)

func GetDBSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial("localhost")
		if err != nil {
			panic(err) // no, not really
		}
	}
	return mgoSession.Clone()
}

func WithCollection(collection string, s func(*mgo.Collection) error) error {
	dbSession := GetDBSession()
	defer dbSession.Close()
	c := dbSession.DB(databaseName).C(collection)
	return s(c)
}
