package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
)

type ProductUseCase interface {
	CreateCategory(ctx context.Context, newCategory string) (domain.ProductCategory, error)
	ViewAllCategories(ctx context.Context) ([]domain.ProductCategory, error)
	FindCategoryByID(ctx context.Context, id int) (domain.ProductCategory, error)
	UpdateCategory(ctx context.Context, info domain.ProductCategory) (domain.ProductCategory, error)
	DeleteCategory(ctx context.Context, categoryID int) (string, error)

	CreateBrand(ctx context.Context, newBrand domain.ProductBrand) (domain.ProductBrand, error)
	UpdateBrand(ctx context.Context, brandInfo domain.ProductBrand) (domain.ProductBrand, error)
	DeleteBrand(ctx context.Context, brandID int) (domain.ProductBrand, error)
	ViewAllBrands(ctx context.Context) ([]domain.ProductBrand, error)
	ViewBrandByID(ctx context.Context, brandID int) (domain.ProductBrand, error)

	CreateProduct(ctx context.Context, newProduct domain.Product) (domain.Product, error)
	ViewAllProducts(ctx context.Context, viewProductInfo model.QueryParams) ([]domain.Product, error)
	FindProductByID(ctx context.Context, id int) (domain.Product, error)
	UpdateProduct(ctx context.Context, info domain.Product) (domain.Product, error)
	DeleteProduct(ctx context.Context, productID int) error

	CreateProductItem(ctx context.Context, newProductItem domain.ProductItem) (domain.ProductItem, error)
	ViewAllProductItems(ctx context.Context, viewProductItemInfo model.QueryParams) ([]domain.ProductItem, error)
	FindProductItemByID(ctx context.Context, id int) (domain.ProductItem, error)
	UpdateProductItem(ctx context.Context, info domain.ProductItem) (domain.ProductItem, error)
	DeleteProductItem(ctx context.Context, productItemID int) error

	CreateCoupon(ctx context.Context, newCoupon model.CreateCoupon) (domain.Coupon, error)
	UpdateCoupon(ctx context.Context, couponInfo model.UpdateCoupon) (domain.Coupon, error)
	DeleteCoupon(ctx context.Context, couponID int) error
	ViewCouponByID(ctx context.Context, couponID int) (domain.Coupon, error)
	ViewAllCoupons(ctx context.Context) ([]domain.Coupon, error)
}
