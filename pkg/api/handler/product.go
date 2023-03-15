package handler

import (
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/response"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	productUseCase services.ProductUseCase
}

func NewProductHandler(usecase services.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase: usecase,
	}
}

//----------------------------------------------------------------------------------------------------------------------
//Category management

// CreateCategory
// @Summary Create new product category
// @ID create-category
// @Description Admin can create new category from admin panel
// @Tags Product
// @Accept json
// @Produce json
// @Param category_name body modelHelper.NewCategory true "New category name"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /adminPanel/create-category [post]
func (cr *ProductHandler) CreateCategory(c *gin.Context) {
	var category modelHelper.NewCategory
	if err := c.Bind(&category); err != nil {
		// Return a 422 Bad request response if the request body is malformed.
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "unable to process the request", Data: nil, Errors: err.Error()})
		return
	}
	//call the CreateCategory use case to create a new category
	createdCategory, err := cr.productUseCase.CreateCategory(c.Request.Context(), category.CategoryName)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to create new category", Data: nil, Errors: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.Response{StatusCode: 201, Message: "Successfully created category", Data: createdCategory, Errors: nil})
}

// ViewAllCategories
// @Summary View all available categories
// @ID view-all-categories
// @Description Admin, users and unregistered users can see all the available categories
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /view-all-categories [get]
func (cr *ProductHandler) ViewAllCategories(c *gin.Context) {
	categories, err := cr.productUseCase.ViewAllCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to fetch categories", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "Successfully fetched all categories", Data: categories, Errors: nil})

}

// FindCategoryByID
// @Summary Fetch details of a specific category using category id
// @ID find-category-by-id
// @Description Users and admins can fetch details of a specific category using id
// @Tags Product
// @Accept json
// @Produce json
// @Param category_id path string true "category id"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /find-category-by-id/{id} [get]
func (cr *ProductHandler) FindCategoryByID(c *gin.Context) {
	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to parse category id", Data: nil, Errors: err.Error()})
		return
	}

	category, err := cr.productUseCase.FindCategoryByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "unable to find category", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "Successfully fetched category", Data: category, Errors: nil})

}

// UpdateCategory
// @Summary Admin can update category details
// @ID update-category
// @Description Admin can update category details
// @Tags Product
// @Accept json
// @Produce json
// @Param category_details body domain.ProductCategory true "category info"
// @Success 202 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /adminPanel/update-category [put]
func (cr *ProductHandler) UpdateCategory(c *gin.Context) {
	var updateInfo domain.ProductCategory
	if err := c.Bind(&updateInfo); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to read request body", Data: nil, Errors: err.Error()})
		return
	}
	updatedCategory, err := cr.productUseCase.UpdateCategory(c.Request.Context(), updateInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 401, Message: "unable to update category", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{StatusCode: 202, Message: "Successfully updated categories", Data: updatedCategory, Errors: nil})
}

// DeleteCategory
// @Summary Admin can delete a category
// @ID delete-category
// @Description Admin can delete a category
// @Tags Product
// @Accept json
// @Produce json
// @Param category_id path string true "category_id"
// @Success 202 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response
// @ Router adminPanel/delete-category/{id}
func (cr *ProductHandler) DeleteCategory(c *gin.Context) {

	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to parse category id", Data: nil, Errors: err.Error()})
		return
	}
	deletedCategory, err := cr.productUseCase.DeleteCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 401, Message: "unable to delete category", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{StatusCode: 202, Message: "Successfully deleted category", Data: deletedCategory, Errors: nil})
}

// ----------------------------------------------------------------------------------------------------------------------
// Brand management

// CreateBrand
// @Summary Admin can create new brand
// @ID create-brand
// @Description Admin can create new brand
// @Tags  Product
// @Accept json
// @Produce json
// @Param new_brand_details body domain.ProductBrand true "new brand details"
// @Success 201 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /adminPanel/create-brand [post]
func (cr *ProductHandler) CreateBrand(c *gin.Context) {
	var newBrand domain.ProductBrand
	if err := c.Bind(&newBrand); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to read request body", Data: nil, Errors: err.Error()})
		return
	}
	createdBrand, err := cr.productUseCase.CreateBrand(c.Request.Context(), newBrand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to create new brand", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response.Response{StatusCode: 201, Message: "successfully created new brand", Data: createdBrand, Errors: nil})
}

// UpdateBrand
// @Summary Admin can update brand details
// @ID update-brand
// @Description Admin can update brand details
// @Tags Product
// @Accept json
// @Produce json
// @Param brand_details body domain.ProductBrand true "brand details"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /adminPanel/update-brand [put]
func (cr *ProductHandler) UpdateBrand(c *gin.Context) {
	var updateBrand domain.ProductBrand
	if err := c.Bind(&updateBrand); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to read request body", Data: nil, Errors: err.Error()})
		return
	}
	updatedBrand, err := cr.productUseCase.UpdateBrand(c.Request.Context(), updateBrand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to update the brand", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully updated brand", Data: updatedBrand, Errors: nil})
}

// DeleteBrand
// @Summary Admin can delete a brand
// @ID delete-brand
// @Description Admin can delete a brand
// @Tags Product
// @Accept json
// @Produce json
// @Param brand_id path string true "brand id"
// @Success 202 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /adminPanel/delete-brand/{id} [delete]
func (cr *ProductHandler) DeleteBrand(c *gin.Context) {
	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to parse brand id", Data: nil, Errors: err.Error()})
		return
	}
	deletedBrand, err := cr.productUseCase.DeleteBrand(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 401, Message: "unable to delete brand", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{StatusCode: 202, Message: "Successfully deleted brand", Data: deletedBrand, Errors: nil})
}

// ViewAllBrands
// @Summary Admin and users can all brands
// @ID view-all-brands
// @Description Admins and users can view all brands
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /view-all-brands [get]
func (cr *ProductHandler) ViewAllBrands(c *gin.Context) {
	brands, err := cr.productUseCase.ViewAllBrands(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to fetch brands", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "Successfully fetched all brands", Data: brands, Errors: nil})

}

// ViewBrandByID
// @Summary Admins and users can view a specific brand details with brand id
// @ID view-brand-by-id
// @Description Admins and users can view a specific brand details with brand id
// @Tags Product
// @Accept json
// @Produce json
// @Param brand_id path string true "brand id"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /view-brand-by-id/{id} [get]
func (cr *ProductHandler) ViewBrandByID(c *gin.Context) {
	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to parse brand id", Data: nil, Errors: err.Error()})
		return
	}

	category, err := cr.productUseCase.ViewBrandByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "unable to find category", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "Successfully fetched category", Data: category, Errors: nil})

}

//----------------------------------------------------------------------------------------------------------------------
//product management

// CreateProduct
// @Summary Admin can create new product listings
// @ID create-product
// @Description Admins can create new product listings
// @Tags Product
// @Accept json
// @Produce json
// @Param new_product_details body domain.Product true "new product details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /adminPanel/create-product [post]
func (cr *ProductHandler) CreateProduct(c *gin.Context) {
	var newProduct domain.Product
	if err := c.Bind(&newProduct); err != nil {
		// Return a 422 Bad request response if the request body is malformed.
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "unable to process the request", Data: nil, Errors: err.Error()})
		return
	}
	//call the CreateCategory use case to create a new category
	createdProduct, err := cr.productUseCase.CreateProduct(c.Request.Context(), newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to add new product", Data: nil, Errors: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.Response{StatusCode: 201, Message: "Successfully added new product", Data: createdProduct, Errors: nil})
}

// ViewAllProducts
// @Summary Admins and users can see all available products
// @ID view-all-products
// @Description Admins and users can ses all available products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /view-all/products [get]
func (cr *ProductHandler) ViewAllProducts(c *gin.Context) {
	products, err := cr.productUseCase.ViewAllProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to fetch products", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "Successfully fetched all products", Data: products, Errors: nil})
}

// FindProductByID
// @Summary Admins and users can see products with product id
// @ID find-product-by-id
// @Description Admins and users can see products with product id
// @Tags Product
// @Accept json
// @Produce json
// @Param product_id path string true "product id"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @ Router /find-product-by-id/{id} [get]
func (cr *ProductHandler) FindProductByID(c *gin.Context) {
	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to parse product id", Data: nil, Errors: err.Error()})
		return
	}

	product, err := cr.productUseCase.FindProductByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "unable to find product", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "Successfully fetched product", Data: product, Errors: nil})

}

// UpdateProduct
// @Summary Admin can update product details
// @ID update-product
// @Description This endpoint allows an admin user to update a product's details.
// @Tags Product
// @Accept json
// @Produce json
// @Param updated_product_details body domain.Product true "Updated product details"
// @Success 202 {object} response.Response "Successfully updated product"
// @Failure 400 {object} response.Response "Unable to update product"
// @Failure 422 {object} response.Response "Failed to read request body"
// @Router adminPanel/update-product [put]
func (cr *ProductHandler) UpdateProduct(c *gin.Context) {
	var updateInfo domain.Product
	if err := c.Bind(&updateInfo); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to read request body", Data: nil, Errors: err.Error()})
		return
	}
	updatedProduct, err := cr.productUseCase.UpdateProduct(c.Request.Context(), updateInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 401, Message: "unable to update product", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{StatusCode: 202, Message: "Successfully updated product", Data: updatedProduct, Errors: nil})
}

// DeleteProduct
// @Summary Deletes a product by ID
// @ID delete-product
// @Description This endpoint allows an admin user to delete a product by ID.
// @Tags Product
// @Accept json
// @Produce json
// @Param product_id path int true "Product ID to delete"
// @Success 202 {object} response.Response "Successfully deleted product"
// @Failure 401 {object} response.Response "Invalid product ID"
// @Failure 422 {object} response.Response "Unable to delete product"
// @Router /adminPanel/delete-product/{product_id} [delete]
func (cr *ProductHandler) DeleteProduct(c *gin.Context) {
	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to parse product id", Data: nil, Errors: err.Error()})
		return
	}
	err = cr.productUseCase.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 401, Message: "unable to delete product", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{StatusCode: 202, Message: "Successfully deleted product", Data: nil, Errors: nil})
}

//----------------------------------------------------------------------------------------------------------------------
//Product Item Management

// CreateProductItem
// @Summary Creates a new product item
// @ID create-product-item
// @Description This endpoint allows an admin user to create a new product item.
// @Tags Product
// @Accept json
// @Produce json
// @Param product_item body domain.ProductItem true "Product item details"
// @Success 201 {object} response.Response "Successfully added new product item"
// @Failure 400 {object} response.Response "Failed to add new product item"
// @Failure 422 {object} response.Response "Unable to process the request"
// @Router /adminPanel/create-product-item [post]
func (cr *ProductHandler) CreateProductItem(c *gin.Context) {
	var newProductItem domain.ProductItem
	if err := c.Bind(&newProductItem); err != nil {
		// Return a 422 Bad request response if the request body is malformed.
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "unable to process the request", Data: nil, Errors: err.Error()})
		return
	}
	createdProductItem, err := cr.productUseCase.CreateProductItem(c.Request.Context(), newProductItem)
	fmt.Println(createdProductItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to add new product item", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response.Response{StatusCode: 201, Message: "Successfully added new product item", Data: createdProductItem, Errors: nil})
}

// ViewAllProductItems
// @Summary Handler function to view all product items
// @ID view-all-product-items
// @Description view all product items
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /view-all-product-items [get]
func (cr *ProductHandler) ViewAllProductItems(c *gin.Context) {
	productItems, err := cr.productUseCase.ViewAllProductItems(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to fetch product items", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "Successfully fetched all product items", Data: productItems, Errors: nil})
}

// FindProductItemByID
// @Summary Retrieve a product item by ID
// @ID find-product-item-by-id
// @Description Retrieve a product item by its ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product item ID"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /view-product-item/{id} [get]
func (cr *ProductHandler) FindProductItemByID(c *gin.Context) {
	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to parse product item id", Data: nil, Errors: err.Error()})
		return
	}

	productItem, err := cr.productUseCase.FindProductItemByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "unable to find product item", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "Successfully fetched product item", Data: productItem, Errors: nil})

}

// UpdateProductItem updates a product item in the database.
// @Summary Update a product item
// @ID update-product-item
// @Description Update an existing product item with new information.
// @Tags Product
// @Accept json
// @Produce json
// @Param product_item body domain.ProductItem true "Product item information to update"
// @Success 202 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /adminPanel/update-product-item [put]
func (cr *ProductHandler) UpdateProductItem(c *gin.Context) {
	var updateInfo domain.ProductItem
	if err := c.Bind(&updateInfo); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to read request body", Data: nil, Errors: err.Error()})
		return
	}
	updatedProductItem, err := cr.productUseCase.UpdateProductItem(c.Request.Context(), updateInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "unable to update product item", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{StatusCode: 202, Message: "Successfully updated product item", Data: updatedProductItem, Errors: nil})
}

// DeleteProductItem
// @Summary Deletes a product item from the system
// @ID delete-product-item
// @Description Deletes a product item from the system
// @Tags Product
// @Accept json
// @Produce json
// @Param product_it path string true "ID of the product item to be deleted"
// @Success 202 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /adminPanel/delete-product-item/{id} [delete]
func (cr *ProductHandler) DeleteProductItem(c *gin.Context) {
	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to parse product item id", Data: nil, Errors: err.Error()})
		return
	}
	err = cr.productUseCase.DeleteProductItem(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 401, Message: "unable to delete product item", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{StatusCode: 202, Message: "Successfully deleted product item", Data: nil, Errors: nil})
}
