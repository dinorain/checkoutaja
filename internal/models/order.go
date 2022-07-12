package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

const (
	OrderStatusPending  = "pending"
	OrderStatusAccepted = "accepted"
)

// Order model
type Order struct {
	OrderID                    uuid.UUID `json:"order_id" db:"order_id"`
	UserID                     uuid.UUID `json:"user_id" db:"user_id"`
	SellerID                   uuid.UUID `json:"seller_id" db:"seller_id"`
	Item                       OrderItem `json:"item" db:"item"`
	Quantity                   uint64    `json:"quantity" db:"quantity"`
	TotalPrice                 float64   `json:"total_price" db:"total_price"`
	Status                     string    `json:"status" db:"status"`
	DeliverySourceAddress      string    `json:"delivery_source_address" db:"delivery_source_address"`
	DeliveryDestinationAddress string    `json:"delivery_destination_address" db:"delivery_destination_address"`
	CreatedAt                  time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type OrderItem Product

func (o *OrderItem) Scan(value interface{}) error {
	var item OrderItem
	if err := copier.Copy(&item, &value); err != nil {
		fmt.Errorf("copier.Copy %v", value)
	}
	*o = item
	return nil
}

func (o OrderItem) Value() (driver.Value, error) {
	valueJson, _ := json.Marshal(o)
	return string(valueJson), nil
}
