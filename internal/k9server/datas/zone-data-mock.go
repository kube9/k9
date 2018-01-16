package datas

import (
	"sync"

	"github.com/kube9/k9/internal/k9server/errors"
	"github.com/kube9/k9/pkg/gen/k9server/models"
)

var data models.ZoneList

// ZoneDataMock stores zone data in memory
type ZoneDataMock struct {
	sync.RWMutex
}

// NewZoneDataMock ctor function
func NewZoneDataMock() *ZoneDataMock {
	return &ZoneDataMock{}
}

// CreateOne creates a new zone
func (z *ZoneDataMock) CreateOne(newZone models.Zone) error {
	z.RLock()
	defer z.RUnlock()

	data = append(data, &newZone)
	return nil
}

// FindAll returns all zones
func (z *ZoneDataMock) FindAll() (*models.ZoneList, error) {
	z.RLock()
	defer z.RUnlock()

	return &data, nil
}

// DeleteOneByID deletes a zone
func (z *ZoneDataMock) DeleteOneByID(id string) error {
	z.RLock()
	defer z.RUnlock()

	for index, value := range data {
		if *value.ID == id {
			data = append(data[:index], data[index+1:]...)
			return nil
		}
	}

	return errors.ErrIDNotFound
}
