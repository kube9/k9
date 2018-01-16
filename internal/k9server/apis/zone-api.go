package apis

import (
	"github.com/kube9/k9/internal/k9server/datas"
	"github.com/kube9/k9/pkg/gen/k9server/models"
	uuid "github.com/satori/go.uuid"
)

// ZoneAPI ...
type ZoneAPI struct {
	data datas.ZoneData
}

// NewZoneAPI ctor function
func NewZoneAPI(data datas.ZoneData) *ZoneAPI {
	return &ZoneAPI{
		data,
	}
}

// CreateOne ...
func (z *ZoneAPI) CreateOne(data models.NewZone) (*models.Zone, error) {
	id := uuid.NewV4().String()
	newZone := models.Zone{
		NewZone: data,
		ZoneAllOf1: models.ZoneAllOf1{
			ID: &id,
		},
	}
	err := z.data.CreateOne(newZone)
	return &newZone, err
}

// FindAll returns all zones
func (z *ZoneAPI) FindAll() (*models.ZoneList, error) {
	return z.data.FindAll()
}

// DeleteOneByID ...
func (z *ZoneAPI) DeleteOneByID(id string) error {
	return z.data.DeleteOneByID(id)
}
