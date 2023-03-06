package usecase

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
)

type productUseCase struct {
	productRepo interfaces.ProductRepository
}

func NewProductUseCase(repo interfaces.ProductRepository) services.ProductUseCase {
	return &productUseCase{
		productRepo: repo,
	}
}

//Category management

func (c *productUseCase) CreateCategory(ctx context.Context, newCategory string) (domain.ProductCategory, error) {
	createdCategory, err := c.productRepo.CreateCategory(ctx, newCategory)
	return createdCategory, err
}

func (c *productUseCase) ViewAllCategories(ctx context.Context) ([]domain.ProductCategory, error) {
	allCategories, err := c.productRepo.ViewAllCategories(ctx)
	return allCategories, err
}

func (c *productUseCase) FindCategoryByID(ctx context.Context, id int) (domain.ProductCategory, error) {
	category, err := c.productRepo.FindCategoryByID(ctx, id)
	if category.CategoryName == "" {
		return category, fmt.Errorf("invalid cateogry id")
	}
	return category, err
}

func (c *productUseCase) UpdateCategory(ctx context.Context, info domain.ProductCategory) (domain.ProductCategory, error) {
	updatedInfo, err := c.productRepo.UpdateCategory(ctx, info)
	return updatedInfo, err
}

func (c *productUseCase) DeleteCategory(ctx context.Context, categoryID int) (string, error) {
	deleteCategoryName, err := c.productRepo.DeleteCategory(ctx, categoryID)
	return deleteCategoryName, err
}

func (c *productUseCase) CreateProduct(ctx context.Context, newProduct domain.Product) (domain.Product, error) {
	createdProduct, err := c.productRepo.CreateProduct(ctx, newProduct)
	return createdProduct, err
}

func (c *productUseCase) ViewAllProducts(ctx context.Context) ([]domain.Product, error) {
	allProducts, err := c.productRepo.ViewAllProducts(ctx)
	return allProducts, err
}

func (c *productUseCase) FindProductByID(ctx context.Context, id int) (domain.Product, error) {
	product, err := c.productRepo.FindProductByID(ctx, id)
	if product.Name == "" {
		return product, fmt.Errorf("invalid product id")
	}
	return product, err
}

func (c *productUseCase) UpdateProduct(ctx context.Context, info domain.Product) (domain.Product, error) {
	updatedProduct, err := c.productRepo.UpdateProduct(ctx, info)
	return updatedProduct, err
}

func (c *productUseCase) DeleteProduct(ctx context.Context, productID int) error {
	err := c.productRepo.DeleteProduct(ctx, productID)
	return err
}
