package datas

import "github.com/kube9/k9/pkg/gen/k9server/models"

// ZoneData defines the storage interface
type ZoneData interface {
	CreateOne(newZone models.Zone) error
	FindAll() (*models.ZoneList, error)
	DeleteOneByID(id string) error
}
