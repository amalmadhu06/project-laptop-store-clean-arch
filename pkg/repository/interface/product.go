package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
)

type ProductRepository interface {
	CreateCategory(ctx context.Context, newCategory string) (domain.ProductCategory, error)
	ViewAllCategories(ctx context.Context) ([]domain.ProductCategory, error)
	FindCategoryByID(ctx context.Context, id int) (domain.ProductCategory, error)
	UpdateCategory(ctx context.Context, info domain.ProductCategory) (domain.ProductCategory, error)
	DeleteCategory(ctx context.Context, categoryID int) (string, error)
}
