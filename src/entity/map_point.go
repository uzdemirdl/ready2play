package entity

import (
	"github.com/google/uuid"
	"time"
)

type MapPoint struct {
	UUID       uuid.UUID
	DeviceUUID uuid.UUID
	MapUUID    uuid.UUID
	X          uint
	y          uint
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
