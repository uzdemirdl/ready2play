package repository

import (
	"context"
	"github.com/uzdemirdl/ready2play/src/interfaces"
	"gorm.io/gorm"
)

type Repository interface {
	Conn(ctx context.Context) *gorm.DB
	Devices() interfaces.DevicesRepository
	Maps() interfaces.MapsRepository
	MapPoints() interfaces.MapPointsRepository
}

type repository struct {
	maps      *mapsRepository
	devices   *deviceRepository
	mapPoints *mapPointsRepository
	DB        *gorm.DB
}

func New(db *gorm.DB) Repository {
	return repository{
		maps:      &mapsRepository{},
		devices:   &deviceRepository{},
		mapPoints: &mapPointsRepository{},
		DB:        db,
	}
}
func (r repository) Conn(ctx context.Context) *gorm.DB {
	return r.DB.WithContext(ctx)
}

func (r repository) Devices() interfaces.DevicesRepository {
	return r.devices
}

func (r repository) Maps() interfaces.MapsRepository {
	return r.maps
}

func (r repository) MapPoints() interfaces.MapPointsRepository {
	return r.mapPoints
}
