//go:generate mockgen -source usecase.go -destination mock/usecase.go -package mock
package product

import (
	"context"

	"github.com/google/uuid"

	"github.com/dinorain/checkoutaja/internal/models"
	"github.com/dinorain/checkoutaja/pkg/utils"
)

//  Product UseCase interface
type ProductUseCase interface {
	Create(ctx context.Context, user *models.Product) (*models.Product, error)
	FindAll(ctx context.Context, pagination *utils.Pagination) ([]models.Product, error)
	FindAllBySellerId(ctx context.Context, sellerId string, pagination *utils.Pagination) ([]models.Product, error)
	FindById(ctx context.Context, userID uuid.UUID) (*models.Product, error)
	CachedFindById(ctx context.Context, userID uuid.UUID) (*models.Product, error)
	UpdateById(ctx context.Context, user *models.Product) (*models.Product, error)
	DeleteById(ctx context.Context, userID uuid.UUID) error
}