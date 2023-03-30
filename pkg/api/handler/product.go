package handler

import (
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/response"
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
// @Tags Product Category
// @Accept json
// @Produce json
// @Param category_name body model.NewCategory true "New category name"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/categories/ [post]
func (cr *ProductHandler) CreateCategory(c *gin.Context) {
	var category model.NewCategory
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
// @Tags Product Category
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/categories/ [get]
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
// @Tags Product Category
// @Accept json
// @Produce json
// @Param category_id path string true "category id"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/categories/{id} [get]
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
// @Tags Product Category
// @Accept json
// @Produce json
// @Param category_details body domain.ProductCategory true "category info"
// @Success 202 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/categories/ [put]
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
// @Tags Product Category
// @Accept json
// @Produce json
// @Param category_id path string true "category_id"
// @Success 202 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response
// @ Router /admin/categories/{id} [delete]
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
// @Tags  Product Brand
// @Accept json
// @Produce json
// @Param new_brand_details body domain.ProductBrand true "new brand details"
// @Success 201 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/brands/ [post]
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
// @Tags Product Brand
// @Accept json
// @Produce json
// @Param brand_details body domain.ProductBrand true "brand details"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/brands/ [put]
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
// @Tags Product Brand
// @Accept json
// @Produce json
// @Param brand_id path string true "brand id"
// @Success 202 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/brands/{id} [delete]
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
// @Tags Product Brand
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/brands/ [get]
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
// @Tags Product Brand
// @Accept json
// @Produce json
// @Param brand_id path string true "brand id"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/brands/{id} [get]
func (cr *ProductHandler) ViewBrandByID(c *gin.Context) {
	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to parse brand id", Data: nil, Errors: err.Error()})
		return
	}

	category, err := cr.productUseCase.ViewBrandByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "unable to find brand", Data: nil, Errors: err.Error()})
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
// @Router /admin/products/ [post]
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
// @ID admin-view-all-products
// @Description Admins and users can ses all available products
// @Tags Product
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param limit query int false "Number of items to retrieve per page"
// @Param query query string false "Search query string"
// @Param filter query string false "Filter criteria for the products"
// @Param sort_by query string false "Sorting criteria for the products"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/products/ [get]

// ViewAllProducts
// @Summary Admins and users can see all available products
// @ID user-view-all-products
// @Description Admins and users can ses all available products
// @Tags Product
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param limit query int false "Number of items to retrieve per page"
// @Param query query string false "Search query string"
// @Param filter query string false "Filter criteria for the product items"
// @Param sort_by query string false "Sorting criteria for the product items"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /products/ [get]
func (cr *ProductHandler) ViewAllProducts(c *gin.Context) {

	//fetch query parameters
	var viewProduct model.QueryParams

	viewProduct.Page, _ = strconv.Atoi(c.Query("page"))
	viewProduct.Limit, _ = strconv.Atoi(c.Query("limit"))
	viewProduct.Query = c.Query("query")
	viewProduct.Filter = c.Query("filter")
	viewProduct.SortBy = c.Query("sort_by")
	viewProduct.SortDesc, _ = strconv.ParseBool(c.Query("sort_desc"))

	products, err := cr.productUseCase.ViewAllProducts(c.Request.Context(), viewProduct)
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
// @ Router /admin/products/{id} [get]
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
// @Router /admin/products/ [put]
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
// @Failure 400 {object} response.Response "Invalid product ID"
// @Failure 422 {object} response.Response "Unable to delete product"
// @Router /admin/products/{product_id} [delete]
func (cr *ProductHandler) DeleteProduct(c *gin.Context) {
	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to parse product id", Data: nil, Errors: err.Error()})
		return
	}
	err = cr.productUseCase.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "unable to delete product", Data: nil, Errors: err.Error()})
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
// @Tags Product Item
// @Accept json
// @Produce json
// @Param product_item body domain.ProductItem true "Product item details"
// @Success 201 {object} response.Response "Successfully added new product item"
// @Failure 400 {object} response.Response "Failed to add new product item"
// @Failure 422 {object} response.Response "Unable to process the request"
// @Router /admin/product-items/ [post]
func (cr *ProductHandler) CreateProductItem(c *gin.Context) {
	var newProductItem domain.ProductItem
	if err := c.Bind(&newProductItem); err != nil {
		// Return a 422 Bad request response if the request body is malformed.
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "unable to process the request", Data: nil, Errors: err.Error()})
		return
	}
	createdProductItem, err := cr.productUseCase.CreateProductItem(c.Request.Context(), newProductItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to add new product item", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response.Response{StatusCode: 201, Message: "Successfully added new product item", Data: createdProductItem, Errors: nil})
}

// ViewAllProductItems
// @Summary Handler function to view all product items
// @ID admin-view-all-product-items
// @Description view all product items
// @Tags Product Item
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param limit query int false "Number of items to retrieve per page"
// @Param query query string false "Search query string"
// @Param filter query string false "Filter criteria for the product items"
// @Param sort_by query string false "Sorting criteria for the product items"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/product-items/ [get]

// ViewAllProductItems for user
// @Summary Handler function to view all product items
// @ID user-view-all-product-items
// @Description view all product items for user
// @Tags Product Item
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param limit query int false "Number of items to retrieve per page"
// @Param query query string false "Search query string"
// @Param filter query string false "Filter criteria for the product items"
// @Param sort_by query string false "Sorting criteria for the product items"
// @Param sort_desc query bool false "Sorting in descending order"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /product-items/ [get]
func (cr *ProductHandler) ViewAllProductItems(c *gin.Context) {
	//fetch query parameters
	var viewProductItem model.QueryParams

	viewProductItem.Page, _ = strconv.Atoi(c.Query("page"))
	viewProductItem.Limit, _ = strconv.Atoi(c.Query("limit"))
	viewProductItem.Query = c.Query("query")
	viewProductItem.Filter = c.Query("filter")
	viewProductItem.SortBy = c.Query("sort_by")
	viewProductItem.SortDesc, _ = strconv.ParseBool(c.Query("sort_desc"))

	productItems, err := cr.productUseCase.ViewAllProductItems(c.Request.Context(), viewProductItem)
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
// @Tags Product Item
// @Accept json
// @Produce json
// @Param id path string true "Product item ID"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/product-items/{id} [get]
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
// @Tags Product Item
// @Accept json
// @Produce json
// @Param product_item body domain.ProductItem true "Product item information to update"
// @Success 202 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/product-items/ [put]
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
// @Tags Product Item
// @Accept json
// @Produce json
// @Param product_item_id path string true "ID of the product item to be deleted"
// @Success 202 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/product-items/{id} [delete]
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

//Coupon Management

// CreateCoupon
// @Summary Admin can create new coupon
// @ID create-coupon
// @Description Admin can create new coupons
// @Tags Coupon
// @Accept json
// @Produce json
// @Param new_coupon_details body model.CreateCoupon true "details of new coupon to be created"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/coupons/ [post]
func (cr *ProductHandler) CreateCoupon(c *gin.Context) {
	var newCoupon model.CreateCoupon

	if err := c.Bind(&newCoupon); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to read request body", Data: nil, Errors: err})
		return
	}

	createdCoupon, err := cr.productUseCase.CreateCoupon(c.Request.Context(), newCoupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to create new coupon", Data: nil, Errors: err})
		return
	}

	c.JSON(http.StatusCreated, response.Response{StatusCode: 201, Message: "successfully created new coupon", Data: createdCoupon, Errors: nil})
}

// UpdateCoupon
// @Summary Admin can update existing coupon
// @ID update-coupon
// @Description Admin can update existing coupon
// @Tags Coupon
// @Accept json
// @Produce json
// @Param coupon_details body model.UpdateCoupon true "details of coupon to be updated"
// @Success 202 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/coupons/ [put]
func (cr *ProductHandler) UpdateCoupon(c *gin.Context) {
	var updateCoupon model.UpdateCoupon
	if err := c.Bind(&updateCoupon); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to read request body", Data: nil, Errors: err.Error()})
		return
	}

	updatedCoupon, err := cr.productUseCase.UpdateCoupon(c.Request.Context(), updateCoupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to update the coupon", Data: nil, Errors: err})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{StatusCode: 202, Message: "successfully updated coupon", Data: updatedCoupon, Errors: nil})

}

// DeleteCoupon
// @Summary Admin can delete existing coupon
// @ID delete-coupon
// @Description Admin can delete existing coupon
// @Tags Coupon
// @Accept json
// @Produce json
// @Param coupon_id path string true "details of coupon to be updated"
// @Success 202 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/coupons/{coupon_id} [delete]
func (cr *ProductHandler) DeleteCoupon(c *gin.Context) {
	paramsID := c.Param("id")
	couponID, err := strconv.Atoi(paramsID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to parse coupon id", Data: nil, Errors: err.Error()})
		return
	}

	err = cr.productUseCase.DeleteCoupon(c.Request.Context(), couponID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to delete the coupon", Data: nil, Errors: err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{StatusCode: 202, Message: "successfully deleted coupon", Data: nil, Errors: nil})

}

// ViewCouponByID
// @Summary Admins and users can see coupon with coupon id
// @ID view-coupon-by-id
// @Description Admins and users can see coupon with id
// @Tags Coupon
// @Accept json
// @Produce json
// @Param coupon_id path string true "coupon_id"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/coupons/{coupon_id} [get]
func (cr *ProductHandler) ViewCouponByID(c *gin.Context) {
	paramsID := c.Param("coupon_id")
	couponID, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to parse coupon id", Data: nil, Errors: err.Error()})
		return
	}
	coupon, err := cr.productUseCase.ViewCouponByID(c.Request.Context(), couponID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to fetch coupon details", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully deleted coupon", Data: coupon, Errors: nil})

}

// ViewAllCoupons
// @Summary Admins and users can see all available coupons
// @ID view-coupons
// @Description Admins and users can see all available coupons
// @Tags Coupon
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/coupons/ [get]
func (cr *ProductHandler) ViewAllCoupons(c *gin.Context) {
	coupons, err := cr.productUseCase.ViewAllCoupons(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to fetch coupons", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully fetched coupons", Data: coupons, Errors: nil})
}
