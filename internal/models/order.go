package models

import (
	"time"

	"github.com/google/uuid"
)

// Order model
type Order struct {
	OrderID                    uuid.UUID `json:"order_id" db:"order_id" validate:"omitempty"`
	UserID                     uuid.UUID `json:"user_id" db:"user_id" validate:"omitempty"`
	SellerID                   uuid.UUID `json:"seller_id" db:"seller_id" validate:"omitempty"`
	Item                       Product   `json:"item" db:"item" validate:"required"`
	Quantity                   uint64    `json:"quantity" db:"quantity" validate:"required"`
	Price                      float64   `json:"price" db:"price" validate:"required"`
	TotalPrice                 float64   `json:"total_price" db:"total_price" validate:"required"`
	Status                     string    `json:"status" db:"status" validate:"required"`
	DeliverySourceAddress      string    `json:"delivery_source_address" db:"delivery_source_address" validate:"required,lte=250"`
	DeliveryDestinationAddress string    `json:"delivery_destination_address" db:"delivery_destination_address" validate:"required,lte=250"`
	CreatedAt                  time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
