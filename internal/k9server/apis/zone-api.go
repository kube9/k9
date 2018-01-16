package apis

import (
	"github.com/kube9/k9/internal/k9server/datas"
	"github.com/kube9/k9/pkg/gen/k9server/models"
	"github.com/prometheus/client_golang/prometheus"
	uuid "github.com/satori/go.uuid"
)

var zoneCreateCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "zones_created_in_system",
		Help: "Counts the zones created ",
	},
)

func init() {
	prometheus.MustRegister(zoneCreateCounter)
}

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
	if err != nil {
		return nil, err
	}

	zoneCreateCounter.Inc()

	return &newZone, nil
}

// FindAll returns all zones
func (z *ZoneAPI) FindAll() (*models.ZoneList, error) {
	return z.data.FindAll()
}

// DeleteOneByID ...
func (z *ZoneAPI) DeleteOneByID(id string) error {
	err := z.data.DeleteOneByID(id)
	if err != nil {
		return err
	}

	return nil
}
