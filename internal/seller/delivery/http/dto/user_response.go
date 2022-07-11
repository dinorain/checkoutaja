package dto

import (
	"time"

	"github.com/google/uuid"

	"github.com/dinorain/checkoutaja/internal/models"
)

type SellerResponseDto struct {
	SellerID  uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      string    `json:"role"`
	Avatar    *string   `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func SellerResponseFromModel(user *models.Seller) *SellerResponseDto {
	return &SellerResponseDto{
		SellerID:  user.SellerID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
