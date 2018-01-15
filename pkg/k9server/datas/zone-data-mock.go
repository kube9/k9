package datas

import (
	"sync"

	"github.com/kube9/k9/pkg/gen/k9server/models"
	"github.com/kube9/k9/pkg/k9server/errors"
	uuid "github.com/satori/go.uuid"
)

var data models.ZoneList

// ZoneDataMock stores zone data in memory
type ZoneDataMock struct {
	sync.RWMutex
}

// NewZoneDataMock ctor function
func NewZoneDataMock() *ZoneDataMock {
	mock := ZoneDataMock{}
	go mock.init()
	return &mock
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

func (z *ZoneDataMock) init() {
	z.Lock()
	defer z.Unlock()

	{
		name := "zone 1"
		id := uuid.NewV4().String()

		item1 := models.Zone{
			NewZone: models.NewZone{
				Name: &name,
			},
			ZoneAllOf1: models.ZoneAllOf1{
				ID: &id,
			},
		}

		data = append(data, &item1)
	}

	{
		name := "zone 2"
		id := uuid.NewV4().String()

		item1 := models.Zone{
			NewZone: models.NewZone{
				Name: &name,
			},
			ZoneAllOf1: models.ZoneAllOf1{
				ID: &id,
			},
		}

		data = append(data, &item1)
	}

	{
		name := "zone 3"
		id := uuid.NewV4().String()

		item1 := models.Zone{
			NewZone: models.NewZone{
				Name: &name,
			},
			ZoneAllOf1: models.ZoneAllOf1{
				ID: &id,
			},
		}

		data = append(data, &item1)
	}

	{
		name := "zone 4"
		id := uuid.NewV4().String()

		item1 := models.Zone{
			NewZone: models.NewZone{
				Name: &name,
			},
			ZoneAllOf1: models.ZoneAllOf1{
				ID: &id,
			},
		}

		data = append(data, &item1)
	}
}
