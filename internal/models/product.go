package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// Product model
type Product struct {
	ProductID   uuid.UUID `json:"product_id" db:"product_id" validate:"omitempty"`
	Name        string    `json:"name" db:"name" validate:"required,lte=30"`
	Description string    `json:"description" db:"description" validate:"required,lte=250"`
	Price       float64   `json:"price" db:"price" validate:"required"`
	SellerID    uuid.UUID `json:"seller_id" db:"seller_id" validate:"omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (p *Product) PrepareCreate() error {
	p.Name = strings.TrimSpace(p.Name)
	p.Description = strings.TrimSpace(p.Description)
	return nil
}
