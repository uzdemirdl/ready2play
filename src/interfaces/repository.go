package interfaces

import (
	"github.com/google/uuid"
	"github.com/uzdemirdl/ready2play/src/entity"
	"gorm.io/gorm"
)

type DevicesRepository interface {
	GetByUUIDs(db *gorm.DB, deviceUUIDs []uuid.UUID) ([]entity.Device, error)
}

type MapsRepository interface {
	GetByUUID(db *gorm.DB, mapUUID uuid.UUID) (entity.Map, error)
}

type MapPointsRepository interface {
	GetByMapUUID(db *gorm.DB, mapUUID uuid.UUID) ([]entity.MapPoint, error)
}
