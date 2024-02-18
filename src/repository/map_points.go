package repository

import (
	"github.com/google/uuid"
	"github.com/uzdemirdl/ready2play/src/entity"
	"gorm.io/gorm"
)

type mapPointsRepository struct{}

func (m mapPointsRepository) GetByMapUUID(db *gorm.DB, mapUUID uuid.UUID) ([]entity.MapPoint, error) {
	var result []entity.MapPoint
	err := db.Where("map_uuid = ?", mapUUID).Find(&result).Error
	return result, err
}
