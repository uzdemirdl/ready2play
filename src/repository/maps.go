package repository

import (
	"github.com/google/uuid"
	"github.com/uzdemirdl/ready2play/src/entity"
	"gorm.io/gorm"
)

type mapsRepository struct{}

func (m mapsRepository) GetByUUID(db *gorm.DB, mapUUID uuid.UUID) (entity.Map, error) {
	var result entity.Map
	err := db.Where("uuid = ?", mapUUID).Take(&result).Error
	return result, err
}
