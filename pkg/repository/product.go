package repository

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
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
//Todo : brand management

//product management

func (c *productDatabase) CreateProduct(ctx context.Context, newProduct domain.Product) (domain.Product, error) {
	var createdProduct domain.Product
	productCreateQuery := `INSERT INTO products(product_category_id, name, brand_id, description, product_image)
							VALUES($1,$2,$3,$4,$5)
							RETURNING *`
	err := c.DB.Raw(productCreateQuery, newProduct.ProductCategoryID, newProduct.Name, newProduct.BrandID, newProduct.Description, newProduct.ProductImage).Scan(&createdProduct).Error
	return createdProduct, err
}

func (c *productDatabase) ViewAllProducts(ctx context.Context) ([]domain.Product, error) {
	var allProducts []domain.Product

	findAllQuery := `SELECT * FROM products;`
	rows, err := c.DB.Raw(findAllQuery).Rows()
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

func (c *productDatabase) ViewAllProductItems(ctx context.Context) ([]domain.ProductItem, error) {
	var allProductItems []domain.ProductItem

	findAllQuery := `SELECT * FROM product_items;`
	rows, err := c.DB.Raw(findAllQuery).Rows()
	if err != nil {
		return allProductItems, err
	}
	defer rows.Close()

	for rows.Next() {
		var productItem domain.ProductItem

		err := rows.Scan(&productItem.ID, &productItem.ProductID, &productItem.Model, &productItem.Processor, &productItem.Ram, &productItem.Storage, &productItem.DisplaySize, &productItem.GraphicsCard, &productItem.OS, &productItem.SKU, &productItem.QntyInStock, &productItem.ProductItemImage, &productItem.Price)
		if err != nil {
			return allProductItems, err
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
