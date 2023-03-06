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

func (c *productDatabase) CreateCategory(ctx context.Context, newCategory string) (domain.ProductCategory, error) {
	var createdCategory domain.ProductCategory
	categoryCreateQuery := `INSERT INTO product_categories(category_name)
							VALUES($1)
							RETURNING id,category_name`
	err := c.DB.Raw(categoryCreateQuery, newCategory).Scan(&createdCategory).Error
	return createdCategory, err
}

// ViewAllCategories fetches all the product categories from the database and returns them as an array of ProductCategory structs.
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
