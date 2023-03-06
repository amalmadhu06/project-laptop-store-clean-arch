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
