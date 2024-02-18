package entity

import (
	"github.com/google/uuid"
	"github.com/uzdemirdl/ready2play/src/types/money"
	"time"
)

type Map struct {
	UUID      uuid.UUID
	Width     uint
	Height    uint
	Name      string
	Price     money.Money
	CreatedAt time.Time
	UpdatedAt time.Time
}
