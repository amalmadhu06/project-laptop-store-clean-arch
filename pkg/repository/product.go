package repository

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"gorm.io/gorm"
	"strings"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB}
}

//product category management

func (c *productDatabase) CreateCategory(ctx context.Context, newCategory string) (domain.ProductCategory, error) {
	var createdCategory domain.ProductCategory
	categoryCreateQuery := `INSERT INTO product_categories(category_name)
							VALUES($1)
							RETURNING id,category_name`
	err := c.DB.Raw(categoryCreateQuery, newCategory).Scan(&createdCategory).Error
	return createdCategory, err
}

func (c *productDatabase) ViewAllCategories(ctx context.Context) ([]domain.ProductCategory, error) {
	// Declare an empty array to store all the product categories.
	var allCategories []domain.ProductCategory

	// Construct the SQL query to fetch all the categories from the product_categories table.
	findAllQuery := `SELECT * FROM product_categories;`

	// Execute the query and get a reference to the result set.
	rows, err := c.DB.Raw(findAllQuery).Rows()
	if err != nil {
		// If an error occurs while executing the query, return an empty array and the error.
		return allCategories, err
	}
	// Close the result set when we're done with it.
	defer rows.Close()

	// Iterate through each row of the result set.
	for rows.Next() {
		// Declare a new ProductCategory struct to hold the data from the current row.
		var category domain.ProductCategory

		// Scan the values from the current row into the fields of the ProductCategory struct.
		err := rows.Scan(&category.ID, &category.CategoryName)
		if err != nil {
			// If an error occurs while scanning the row, return the categories we have so far and the error.
			return allCategories, err
		}
		// Add the ProductCategory struct to the array of all categories.
		allCategories = append(allCategories, category)
	}
	// Return the array of all categories.
	return allCategories, nil
}

func (c *productDatabase) FindCategoryByID(ctx context.Context, id int) (domain.ProductCategory, error) {
	var category domain.ProductCategory
	fetchCategoryQuery := ` SELECT * FROM product_categories
							WHERE id = $1`
	err := c.DB.Raw(fetchCategoryQuery, id).Scan(&category).Error
	return category, err
}

func (c *productDatabase) UpdateCategory(ctx context.Context, info domain.ProductCategory) (domain.ProductCategory, error) {
	var updatedCategory domain.ProductCategory
	updateCategoryQuery := `UPDATE product_categories
							SET category_name = $1
							WHERE id = $2
							RETURNING id, category_name`

	err := c.DB.Raw(updateCategoryQuery, info.CategoryName, info.ID).Scan(&updatedCategory).Error

	return updatedCategory, err
}

func (c *productDatabase) DeleteCategory(ctx context.Context, categoryID int) (string, error) {
	var deletedCategory string
	deleteCategoryQuery := `DELETE  FROM product_categories
							WHERE id = $1`
	err := c.DB.Exec(deleteCategoryQuery, categoryID).Error
	fmt.Println(categoryID)
	return deletedCategory, err
}

//brand management

func (c *productDatabase) CreateBrand(ctx context.Context, newBrand domain.ProductBrand) (domain.ProductBrand, error) {
	var createdBrand domain.ProductBrand
	createBrandQuery := `INSERT INTO product_brands (brand) VALUES($1) RETURNING *;`
	err := c.DB.Raw(createBrandQuery, newBrand.Brand).Scan(&createdBrand).Error
	return createdBrand, err
}

func (c *productDatabase) UpdateBrand(ctx context.Context, brandInfo domain.ProductBrand) (domain.ProductBrand, error) {
	var updatedBrand domain.ProductBrand
	updateBrandQuery := `UPDATE product_brands SET brand = $1 WHERE id = $2 RETURNING *;`
	err := c.DB.Raw(updateBrandQuery, brandInfo.Brand, brandInfo.ID).Scan(&updatedBrand).Error
	return updatedBrand, err
}

func (c *productDatabase) DeleteBrand(ctx context.Context, brandID int) (domain.ProductBrand, error) {
	var deletedBrand domain.ProductBrand
	selectBrandQuery := `SELECT * FROM product_brands WHERE id = $1;`
	err := c.DB.Raw(selectBrandQuery, brandID).Scan(&deletedBrand).Error
	if deletedBrand.ID == 0 || err != nil {
		return domain.ProductBrand{}, fmt.Errorf("no brand found")
	}
	deleteBrandQuery := `DELETE FROM product_brands WHERE id = $1`
	err = c.DB.Exec(deleteBrandQuery, brandID).Error
	return deletedBrand, err
}

func (c *productDatabase) ViewAllBrands(ctx context.Context) ([]domain.ProductBrand, error) {
	var allBrands []domain.ProductBrand
	fetchBrandsQuery := `SELECT * FROM product_brands;`
	err := c.DB.Raw(fetchBrandsQuery).Scan(&allBrands).Error
	return allBrands, err
}

func (c *productDatabase) ViewBrandByID(ctx context.Context, brandID int) (domain.ProductBrand, error) {
	var brand domain.ProductBrand
	fetchBrandQuery := `SELECT * FROM product_brands WHERE id = $1;`
	err := c.DB.Raw(fetchBrandQuery, brandID).Scan(&brand).Error
	if brand.ID == 0 {
		return domain.ProductBrand{}, fmt.Errorf("no brand found")
	}
	return brand, err
}

//product management

func (c *productDatabase) CreateProduct(ctx context.Context, newProduct domain.Product) (domain.Product, error) {
	var createdProduct domain.Product
	productCreateQuery := `INSERT INTO products(product_category_id, name, brand_id, description, product_image)
							VALUES($1,$2,$3,$4,$5)
							RETURNING *`
	err := c.DB.Raw(productCreateQuery, newProduct.ProductCategoryID, newProduct.Name, newProduct.BrandID, newProduct.Description, newProduct.ProductImage).Scan(&createdProduct).Error
	return createdProduct, err
}

func (c *productDatabase) ViewAllProducts(ctx context.Context, queryParams model.QueryParams) ([]domain.Product, error) {

	findQuery := "SELECT * FROM products"
	if queryParams.Query != "" && queryParams.Filter != "" {
		findQuery = fmt.Sprintf("%s WHERE LOWER(%s) LIKE '%%%s%%'", findQuery, queryParams.Filter, strings.ToLower(queryParams.Query))
	}
	if queryParams.SortBy != "" {
		if queryParams.SortDesc {
			findQuery = fmt.Sprintf("%s ORDER BY %s DESC", findQuery, queryParams.SortBy)
		} else {
			findQuery = fmt.Sprintf("%s ORDER BY %s ASC", findQuery, queryParams.SortBy)
		}
	}
	if queryParams.Limit != 0 && queryParams.Page != 0 {
		findQuery = fmt.Sprintf("%s LIMIT %d OFFSET %d", findQuery, queryParams.Limit, (queryParams.Page-1)*queryParams.Limit)
	}
	if queryParams.Limit == 0 || queryParams.Page == 0 {
		findQuery = fmt.Sprintf("%s LIMIT 10 OFFSET 0", findQuery)
	}

	var allProducts []domain.Product
	rows, err := c.DB.Raw(findQuery).Rows()
	if err != nil {
		return allProducts, err
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.Product

		err := rows.Scan(&product.ID, &product.ProductCategoryID, &product.Name, &product.BrandID, &product.Description, &product.ProductImage)
		if err != nil {
			return allProducts, err
		}
		allProducts = append(allProducts, product)
	}
	return allProducts, nil
}

func (c *productDatabase) FindProductByID(ctx context.Context, id int) (domain.Product, error) {
	var product domain.Product
	fetchProductQuery := ` SELECT * FROM products
							WHERE id = $1`
	err := c.DB.Raw(fetchProductQuery, id).Scan(&product).Error
	return product, err
}

func (c *productDatabase) UpdateProduct(ctx context.Context, info domain.Product) (domain.Product, error) {
	var updatedProduct domain.Product
	updateProductQuery := `	UPDATE products
							SET 
								product_category_id = $1,
								name = $2,
								brand_id = $3,
								description = $4,
								product_image = $5
							WHERE id = $6
							RETURNING id,product_category_id,name,brand_id,description,product_image`
	//Todo : fix scanning bug
	err := c.DB.Raw(updateProductQuery, info.ProductCategoryID, info.Name, info.BrandID, info.Description, info.ProductImage, info.ID).Scan(&updatedProduct).Error
	return updatedProduct, err
}

func (c *productDatabase) DeleteProduct(ctx context.Context, productID int) error {
	deleteProductQuery := `DELETE  FROM products
							WHERE id = $1`
	err := c.DB.Exec(deleteProductQuery, productID).Error
	return err
}

//product item management

func (c *productDatabase) CreateProductItem(ctx context.Context, newProductItem domain.ProductItem) (domain.ProductItem, error) {
	var createdProductItem domain.ProductItem
	productItemCreateQuery := `INSERT INTO product_items(product_id, model, processor, ram, storage, display_size, graphics_card, os, sku, qnty_in_stock, product_item_image, price)
							VALUES( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
							RETURNING *`
	err := c.DB.Raw(productItemCreateQuery, newProductItem.ProductID, newProductItem.Model, newProductItem.Processor, newProductItem.Ram, newProductItem.Storage, newProductItem.DisplaySize, newProductItem.GraphicsCard, newProductItem.OS, newProductItem.SKU, newProductItem.QntyInStock, newProductItem.ProductItemImage, newProductItem.Price).Scan(&createdProductItem).Error
	return createdProductItem, err
}

func (c *productDatabase) ViewAllProductItems(ctx context.Context, queryParams model.QueryParams) ([]domain.ProductItem, error) {
	// Building query based on query params received.
	findQuery := "SELECT * FROM product_items"
	if queryParams.Query != "" && queryParams.Filter != "" {
		findQuery = fmt.Sprintf("%s WHERE LOWER(%s) LIKE '%%%s%%'", findQuery, queryParams.Filter, strings.ToLower(queryParams.Query))
	}
	if queryParams.SortBy != "" {
		if queryParams.SortDesc {
			findQuery = fmt.Sprintf("%s ORDER BY %s DESC", findQuery, queryParams.SortBy)
		} else {
			findQuery = fmt.Sprintf("%s ORDER BY %s ASC", findQuery, queryParams.SortBy)
		}
	}
	if queryParams.Limit != 0 && queryParams.Page != 0 {
		findQuery = fmt.Sprintf("%s LIMIT %d OFFSET %d", findQuery, queryParams.Limit, (queryParams.Page-1)*queryParams.Limit)
	}
	if queryParams.Limit == 0 || queryParams.Page == 0 {
		findQuery = fmt.Sprintf("%s LIMIT 10 OFFSET 0", findQuery)
	}

	var allProductItems []domain.ProductItem
	rows, err := c.DB.Raw(findQuery).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var productItem domain.ProductItem

		err := rows.Scan(&productItem.ID, &productItem.ProductID, &productItem.Model, &productItem.Processor, &productItem.Ram, &productItem.Storage, &productItem.DisplaySize, &productItem.GraphicsCard, &productItem.OS, &productItem.SKU, &productItem.QntyInStock, &productItem.ProductItemImage, &productItem.Price)
		if err != nil {
			return nil, err
		}
		allProductItems = append(allProductItems, productItem)
	}
	return allProductItems, nil
}

func (c *productDatabase) FindProductItemByID(ctx context.Context, id int) (domain.ProductItem, error) {
	var productItem domain.ProductItem
	fetchProductItemQuery := ` SELECT * FROM product_items
							WHERE id = $1`
	err := c.DB.Raw(fetchProductItemQuery, id).Scan(&productItem).Error
	return productItem, err
}

func (c *productDatabase) UpdateProductItem(ctx context.Context, info domain.ProductItem) (domain.ProductItem, error) {
	var updatedProductItem domain.ProductItem
	updateProductItemQuery := `	UPDATE product_items
								SET
									product_id = $1,
									model = $2, 
									processor = $3, 
									ram = $4, 
									storage = $5, 
									display_size = $6, 
									graphics_card = $7, 
									os = $8,
									sku = $9, 
									qnty_in_stock = $10, 
									product_item_image = $11, 
									price = $12
								WHERE id = $13
								RETURNING id, product_id, model, processor, ram, storage, display_size, graphics_card, os, sku, qnty_in_stock, product_item_image, price`
	//Todo : fix scanning bug
	err := c.DB.Raw(updateProductItemQuery, info.ProductID, info.Model, info.Processor, info.Ram, info.Storage, info.DisplaySize, info.GraphicsCard, info.OS, info.SKU, info.QntyInStock, info.ProductItemImage, info.Price, info.ID).Scan(&updatedProductItem).Error
	return updatedProductItem, err
}

func (c *productDatabase) DeleteProductItem(ctx context.Context, productItemID int) error {
	deleteProductItemQuery := `DELETE  FROM product_items
							WHERE id = $1`
	err := c.DB.Exec(deleteProductItemQuery, productItemID).Error
	return err
}

// coupon management

func (c *productDatabase) CreateCoupon(ctx context.Context, newCoupon model.CreateCoupon) (domain.Coupon, error) {
	var createdCoupon domain.Coupon
	createCouponQuery := `	INSERT INTO coupons(code, min_order_value, discount_percent, discount_max_amount, valid_till) 
							VALUES($1, $2, $3, $4, $5) 
							RETURNING *;`
	err := c.DB.Raw(createCouponQuery, newCoupon.Code, newCoupon.MinOrderValue, newCoupon.DiscountPercent, newCoupon.DiscountMaxAmount, newCoupon.ValidTill).Scan(&createdCoupon).Error
	if err != nil {
		return domain.Coupon{}, err
	}
	if createdCoupon.ID == 0 {
		return domain.Coupon{}, fmt.Errorf("failed to create new coupon")
	}
	return createdCoupon, nil
}

func (c *productDatabase) UpdateCoupon(ctx context.Context, couponInfo model.UpdateCoupon) (domain.Coupon, error) {
	var updatedCoupon domain.Coupon
	updateCouponQuery := `	UPDATE coupons SET 
								code = $1,
								min_order_value = $2,
								discount_percent = $3,
								discount_max_amount = $4,
								valid_till = $5
							WHERE id = $6
							RETURNING *;`
	err := c.DB.Raw(updateCouponQuery, couponInfo.Code, couponInfo.MinOrderValue, couponInfo.DiscountPercent, couponInfo.DiscountMaxAmount, couponInfo.ValidTill, couponInfo.ID).Scan(&updatedCoupon).Error

	if err != nil {
		return domain.Coupon{}, err
	}

	if updatedCoupon.ID == 0 {
		return domain.Coupon{}, fmt.Errorf("no product found")
	}
	return updatedCoupon, nil
}

func (c *productDatabase) DeleteCoupon(ctx context.Context, couponID int) error {
	var fetchedID int
	findCouponQuery := `SELECT id FROM coupons WHERE id = $1`
	err := c.DB.Raw(findCouponQuery, couponID).Scan(&fetchedID).Error
	if err != nil {
		return err
	}
	if fetchedID == 0 {
		return fmt.Errorf("no such coupon found")
	}

	deleteCouponQuery := `DELETE FROM coupons WHERE id = $1;`
	err = c.DB.Exec(deleteCouponQuery, couponID).Error
	return err

}

func (c *productDatabase) ViewCouponByID(ctx context.Context, couponID int) (domain.Coupon, error) {
	var coupon domain.Coupon
	fetchCouponQuery := `SELECT * FROM coupons WHERE id = $1;`
	err := c.DB.Raw(fetchCouponQuery, couponID).Scan(&coupon).Error
	if err != nil {
		return domain.Coupon{}, err
	}
	if coupon.ID == 0 {
		return domain.Coupon{}, fmt.Errorf("no coupon found")
	}
	return coupon, nil
}

func (c *productDatabase) ViewAllCoupons(ctx context.Context) ([]domain.Coupon, error) {
	var allCoupons []domain.Coupon
	fetchAllCouponsQuery := `SELECT * FROM coupons WHERE valid_till > NOW();`
	err := c.DB.Raw(fetchAllCouponsQuery).Scan(&allCoupons).Error
	if err != nil {
		return allCoupons, err
	}
	if len(allCoupons) == 0 {
		return allCoupons, fmt.Errorf("no coupons found")
	}
	return allCoupons, err
}

func (c *productDatabase) CouponUsed(ctx context.Context, userID, couponID int) (bool, error) {
	var isUsed bool
	checkQuery := `	SELECT 
					EXISTS(
							SELECT 1 FROM orders 
							WHERE user_id = $1 AND 
							coupon_id = $2
					);`
	err := c.DB.Raw(checkQuery, userID, couponID).Scan(&isUsed).Error
	return isUsed, err
}
