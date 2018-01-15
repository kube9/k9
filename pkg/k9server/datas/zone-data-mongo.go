package datas

import (
	"github.com/kube9/k9/pkg/gen/k9server/models"
	mgo "gopkg.in/mgo.v2"
)

// ZoneDataMongo stores zone data in mongodb
type ZoneDataMongo struct {
	session *mgo.Session
}

// NewZoneDataMongo ctor function
func NewZoneDataMongo(session *mgo.Session) *ZoneDataMongo {
	return &ZoneDataMongo{
		session,
	}
}

// CreateOne ...
func (z *ZoneDataMongo) CreateOne(newZone models.Zone) error {
	return nil
}

// FindAll returns all zones
func (z *ZoneDataMongo) FindAll() (*models.ZoneList, error) {
	return nil, nil
}

// DeleteOneByID ...
func (z *ZoneDataMongo) DeleteOneByID(id string) error {
	return nil
}
