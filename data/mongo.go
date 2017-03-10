package data

import (
	"gopkg.in/mgo.v2"
)

const (
	db         = "twitch"
	collection = "fragments"
)

type MongoStorage struct {
	session *mgo.Session
}

func NewMongoStorage(session *mgo.Session) Storage {
	return &MongoStorage{
		session: session,
	}
}

func (s *MongoStorage) GetUploadedImages() ([]UploadedFragment, error) {
	result := []UploadedFragment{}

	err := s.session.DB(db).C(collection).Find(nil).Sort("-time").All(&result)

	return result, err
}