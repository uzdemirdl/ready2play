package repository

import (
	"github.com/google/uuid"
	"github.com/uzdemirdl/ready2play/src/entity"
	"gorm.io/gorm"
)

type deviceRepository struct{}

func (d deviceRepository) GetByUUIDs(db *gorm.DB, deviceUUIDs []uuid.UUID) ([]entity.Device, error) {
	var result []entity.Device
	err := db.Where("uuid IN (?)", deviceUUIDs).Find(&result).Error
	return result, err
}
