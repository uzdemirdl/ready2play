package entity

import (
	"github.com/google/uuid"
	"github.com/uzdemirdl/ready2play/src/types"
	"time"
)

type Device struct {
	UUID      uuid.UUID
	Type      types.DeviceType
	CreatedAt time.Time
	UpdatedAt time.Time
}
