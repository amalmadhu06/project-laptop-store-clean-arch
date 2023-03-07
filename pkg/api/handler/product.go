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
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "unable to process the request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	//call the CreateCategory use case to create a new category
	createdCategory, err := cr.productUseCase.CreateCategory(c.Request.Context(), category.CategoryName)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to create new category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "Successfully created category",
		Data:       createdCategory,
		Errors:     nil,
	})
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
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "failed to fetch categories",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Successfully fetched all categories",
		Data:       categories,
		Errors:     nil,
	})

}

// FindCategoryByID
// @Summary
// @ID
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
func (cr *ProductHandler) FindCategoryByID(c *gin.Context) {
	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "failed to parse category id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	category, err := cr.productUseCase.FindCategoryByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "unable to find category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Successfully fetched category",
		Data:       category,
		Errors:     nil,
	})

}

// UpdateCategory
// @Summary
// @ID
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
func (cr *ProductHandler) UpdateCategory(c *gin.Context) {
	var updateInfo domain.ProductCategory
	if err := c.Bind(&updateInfo); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	updatedCategory, err := cr.productUseCase.UpdateCategory(c.Request.Context(), updateInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 401,
			Message:    "unable to update category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{
		StatusCode: 202,
		Message:    "Successfully updated categories",
		Data:       updatedCategory,
		Errors:     nil,
	})
}

// DeleteCategory
// @Summary
// @ID
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
func (cr *ProductHandler) DeleteCategory(c *gin.Context) {
	var category modelHelper.CategoryID
	if err := c.Bind(&category); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	deletedCategory, err := cr.productUseCase.DeleteCategory(c.Request.Context(), category.CategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 401,
			Message:    "unable to delete category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{
		StatusCode: 202,
		Message:    "Successfully deleted category",
		Data:       deletedCategory,
		Errors:     nil,
	})
}

//----------------------------------------------------------------------------------------------------------------------
//Todo : brand management

//----------------------------------------------------------------------------------------------------------------------
//product management

// CreateProduct
// @Summary
// @ID
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
func (cr *ProductHandler) CreateProduct(c *gin.Context) {
	var newProduct domain.Product
	if err := c.Bind(&newProduct); err != nil {
		// Return a 422 Bad request response if the request body is malformed.
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "unable to process the request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	//call the CreateCategory use case to create a new category
	createdProduct, err := cr.productUseCase.CreateProduct(c.Request.Context(), newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to add new product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "Successfully added new product",
		Data:       createdProduct,
		Errors:     nil,
	})
}

// ViewAllProducts
// @Summary
// @ID
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
func (cr *ProductHandler) ViewAllProducts(c *gin.Context) {
	products, err := cr.productUseCase.ViewAllProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "failed to fetch products",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Successfully fetched all products",
		Data:       products,
		Errors:     nil,
	})
}

// FindProductByID
// @Summary
// @ID
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
func (cr *ProductHandler) FindProductByID(c *gin.Context) {
	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "failed to parse product id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	product, err := cr.productUseCase.FindProductByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "unable to find product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Successfully fetched product",
		Data:       product,
		Errors:     nil,
	})

}

// UpdateProduct
// @Summary
// @ID
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
func (cr *ProductHandler) UpdateProduct(c *gin.Context) {
	var updateInfo domain.Product
	if err := c.Bind(&updateInfo); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	updatedProduct, err := cr.productUseCase.UpdateProduct(c.Request.Context(), updateInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 401,
			Message:    "unable to update product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{
		StatusCode: 202,
		Message:    "Successfully updated product",
		Data:       updatedProduct,
		Errors:     nil,
	})
}

// DeleteProduct
// @Summary
// @ID
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
func (cr *ProductHandler) DeleteProduct(c *gin.Context) {
	var product modelHelper.ProductID
	if err := c.Bind(&product); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err := cr.productUseCase.DeleteProduct(c.Request.Context(), product.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 401,
			Message:    "unable to delete product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{
		StatusCode: 202,
		Message:    "Successfully deleted product",
		Data:       nil,
		Errors:     nil,
	})
}

//----------------------------------------------------------------------------------------------------------------------
//Product Item Management

// CreateProductItem
// @Summary
// @ID
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
func (cr *ProductHandler) CreateProductItem(c *gin.Context) {
	var newProductItem domain.ProductItem
	if err := c.Bind(&newProductItem); err != nil {
		// Return a 422 Bad request response if the request body is malformed.
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "unable to process the request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	createdProductItem, err := cr.productUseCase.CreateProductItem(c.Request.Context(), newProductItem)
	fmt.Println(createdProductItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to add new product item",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "Successfully added new product item",
		Data:       createdProductItem,
		Errors:     nil,
	})
}

// ViewAllProductItems
// @Summary
// @ID
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
func (cr *ProductHandler) ViewAllProductItems(c *gin.Context) {
	productItems, err := cr.productUseCase.ViewAllProductItems(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "failed to fetch product items",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Successfully fetched all product items",
		Data:       productItems,
		Errors:     nil,
	})
}

// FindProductItemByID
// @Summary
// @ID
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
func (cr *ProductHandler) FindProductItemByID(c *gin.Context) {
	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "failed to parse product item id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	productItem, err := cr.productUseCase.FindProductItemByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "unable to find product item",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Successfully fetched product item",
		Data:       productItem,
		Errors:     nil,
	})

}

// UpdateProductItem
// @Summary
// @ID
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
func (cr *ProductHandler) UpdateProductItem(c *gin.Context) {
	var updateInfo domain.ProductItem
	if err := c.Bind(&updateInfo); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	updatedProductItem, err := cr.productUseCase.UpdateProductItem(c.Request.Context(), updateInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "unable to update product item",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{
		StatusCode: 202,
		Message:    "Successfully updated product item",
		Data:       updatedProductItem,
		Errors:     nil,
	})
}

// DeleteProductItem
// @Summary
// @ID
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
func (cr *ProductHandler) DeleteProductItem(c *gin.Context) {
	var productItem modelHelper.ProductItemID
	if err := c.Bind(&productItem); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err := cr.productUseCase.DeleteProductItem(c.Request.Context(), productItem.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 401,
			Message:    "unable to delete product item",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{
		StatusCode: 202,
		Message:    "Successfully deleted product item",
		Data:       nil,
		Errors:     nil,
	})
}
