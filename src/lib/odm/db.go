package odm

import (
	"gopkg.in/mgo.v2"
)

type M map[string]interface{}

type DB struct {
	dbName  string
	Session *mgo.Session
}

func New(url, dbname string) (*DB, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	return &DB{
		dbName:  dbname,
		Session: session,
	}, nil
}
