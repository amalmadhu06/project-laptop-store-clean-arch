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

func (c *productUseCase) CreateBrand(ctx context.Context, newBrand domain.ProductBrand) (domain.ProductBrand, error) {
	createdBrand, err := c.productRepo.CreateBrand(ctx, newBrand)
	return createdBrand, err
}

func (c *productUseCase) UpdateBrand(ctx context.Context, brandInfo domain.ProductBrand) (domain.ProductBrand, error) {
	updatedBrand, err := c.productRepo.UpdateBrand(ctx, brandInfo)
	return updatedBrand, err
}

func (c *productUseCase) DeleteBrand(ctx context.Context, brandID int) (domain.ProductBrand, error) {
	deletedBrand, err := c.productRepo.DeleteBrand(ctx, brandID)
	return deletedBrand, err
}

func (c *productUseCase) ViewAllBrands(ctx context.Context) ([]domain.ProductBrand, error) {
	allBrands, err := c.productRepo.ViewAllBrands(ctx)
	return allBrands, err
}

func (c *productUseCase) ViewBrandByID(ctx context.Context, brandID int) (domain.ProductBrand, error) {
	brand, err := c.productRepo.ViewBrandByID(ctx, brandID)
	return brand, err
}

//Product Management

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

//Product Item Management

func (c *productUseCase) CreateProductItem(ctx context.Context, newProductItem domain.ProductItem) (domain.ProductItem, error) {
	createdProductItem, err := c.productRepo.CreateProductItem(ctx, newProductItem)
	return createdProductItem, err
}

func (c *productUseCase) ViewAllProductItems(ctx context.Context) ([]domain.ProductItem, error) {
	allProductItems, err := c.productRepo.ViewAllProductItems(ctx)
	return allProductItems, err
}

func (c *productUseCase) FindProductItemByID(ctx context.Context, id int) (domain.ProductItem, error) {
	productItem, err := c.productRepo.FindProductItemByID(ctx, id)
	if productItem.Model == "" {
		return productItem, fmt.Errorf("invalid product item id")
	}
	return productItem, err
}

func (c *productUseCase) UpdateProductItem(ctx context.Context, info domain.ProductItem) (domain.ProductItem, error) {
	updatedProductItem, err := c.productRepo.UpdateProductItem(ctx, info)
	return updatedProductItem, err
}

func (c *productUseCase) DeleteProductItem(ctx context.Context, productItemID int) error {
	err := c.productRepo.DeleteProductItem(ctx, productItemID)
	return err
}
