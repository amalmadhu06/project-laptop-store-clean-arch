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

	//Todo : brand management

	CreateProduct(ctx context.Context, newProduct domain.Product) (domain.Product, error)
	ViewAllProducts(ctx context.Context) ([]domain.Product, error)
	FindProductByID(ctx context.Context, id int) (domain.Product, error)
	UpdateProduct(ctx context.Context, info domain.Product) (domain.Product, error)
	DeleteProduct(ctx context.Context, productID int) error

	CreateProductItem(ctx context.Context, newProductItem domain.ProductItem) (domain.ProductItem, error)
	ViewAllProductItems(ctx context.Context) ([]domain.ProductItem, error)
	FindProductItemByID(ctx context.Context, id int) (domain.ProductItem, error)
	UpdateProductItem(ctx context.Context, info domain.ProductItem) (domain.ProductItem, error)
	DeleteProductItem(ctx context.Context, productItemID int) error
}
