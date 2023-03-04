package handler

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/response"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// CreateUser
// @Summary Create a new user
// @ID create-user
// @Description Create a new user with the specified details.
// @Tags Users
// @Accept json
// @Produce json
// @Param user_details body modelHelper.UserDataInput true "User details"
// @Success 201 {object} response.response
// @Failure 400 {object} response.response
// @Failure 500 {object} response.response
// @Router /signup [post]
func (cr UserHandler) CreateUser(c *gin.Context) {
	//cancelling the request if it is taking more than one minute to send back a response
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute)
	defer cancel()
	// 1. receive data from request body
	var body modelHelper.UserDataInput
	if err := c.Bind(&body); err != nil {
		// Return a 422 Bad request response if the request body is malformed.
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "unable to process the request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	// Call the CreateUser method of the userUseCase to create the user.
	userData, err := cr.userUseCase.CreateUser(ctx, body)
	if err != nil {
		// Return a 400 Bad request response if there is an error while creating the user.
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to create user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	//Return a 201 Created response if the user is successfully created.
	c.Status(http.StatusCreated)
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "User created successfully",
		Data:       userData,
		Errors:     nil,
	})
}

// LoginWithEmail
// @Summary User Login
// @ID user-login-email
// @Description Login as a user to access the ecommerce site
// @Tags Users
// @Accept json
// @Produce json
// @Param user_details body modelHelper.UserLoginEmail true "User details"
// @Success 200
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /login [post]
func (cr *UserHandler) LoginWithEmail(c *gin.Context) {
	// receive data from request body
	var body modelHelper.UserLoginEmail
	if err := c.Bind(&body); err != nil {
		// Return a 421 response if the request body is malformed.
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "unable to process the request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	// Call the UserLogin method of the userUseCase to login as a user.
	ss, user, err := cr.userUseCase.LoginWithEmail(c.Request.Context(), body)
	if err != nil {
		// Return a 400 Bad Request response if there is an error while creating the user.
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("UserAuth", ss, 3600*24*30, "", "", false, true)
	// Return a 201 Created response if the user is successfully logged in.
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Successfully logged in",
		Data:       user,
		Errors:     nil,
	})
}

// LoginWithPhone
// @Summary User Login-Phone
// @ID user-login-phone
// @Description Login as a user to access the ecommerce site
// @Tags Users
// @Accept json
// @Produce json
// @Param user_details body modelHelper.UserLoginPhone true "User details"
// @Success 200
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /login [post]
func (cr *UserHandler) LoginWithPhone(c *gin.Context) {
	// receive data from request body
	var body modelHelper.UserLoginPhone
	if err := c.Bind(&body); err != nil {
		// Return a 421 response if the request body is malformed.
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "unable to process the request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	// Call the UserLogin method of the userUseCase to login as a user.
	ss, user, err := cr.userUseCase.LoginWithPhone(c.Request.Context(), body)
	if err != nil {
		// Return a 400 Bad Request response if there is an error while creating the user.
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("UserAuth", ss, 3600*24*30, "", "", false, true)
	// Return a 201 Created response if the user is successfully logged in.
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Successfully logged in",
		Data:       user,
		Errors:     nil,
	})
}

// UserLogout
// @Summary User Logout
// @ID user-logout
// @Description Logs out a logged-in user from the E-commerce web api
// @Tags Users
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Failure 500
// @Router /logout [get]
func (cr *UserHandler) UserLogout(c *gin.Context) {
	// Set the user authentication cookie's expiration to -1 to invalidate it.
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") //indicates to the client that it should not cache any response data and should always revalidate it with the server
	c.SetSameSite(http.SameSiteLaxMode)                                           //sets the SameSite cookie attribute to "Lax" for the response. This attribute restricts the scope of cookies and helps prevent cross-site request forgery attacks
	c.SetCookie("UserAuth", "", -1, "", "", false, true)                          //Immediately by setting the maxAge to -1, and marks the cookie as secure and HTTP-only
	c.Status(http.StatusOK)
}

// @Summary Find by user ID
// @ID find-by-id
// @Description Find user details by user ID.
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} domain.Users
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /adminpanel/users/:id [get]
//func (cr *UserHandler) FindByID(c *gin.Context) {
//	paramsID := c.Param("id")
//	id, err := strconv.Atoi(paramsID)
//
//	// Return error response if user ID is not a valid integer
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, response.Response{
//			StatusCode: 500,
//			Errors:     err.Error(),
//			Message:    "Failed to parse user id",
//			Data:       nil,
//		})
//		return
//	}
//
//	// Retrieve user details by ID using FindByID method from user use case
//	user, err := cr.userUseCase.FindByID(c.Request.Context(), uint(id))
//	// Return error response if user is not found
//	if err != nil {
//		c.JSON(http.StatusNotFound, response.Response{
//			StatusCode: 404,
//			Errors:     err.Error(),
//			Message:    "No user found",
//			Data:       nil,
//		})
//		return
//	}
//	// Return user details as a successful response
//	c.JSON(http.StatusOK, user)
//}

// @Summary Find all users
// @ID find-all
// @Description Retrieve a list of all users registered in the system.
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {array} domain.Users
// @Failure 500 {object} response.Response
// @Router /adminpanel/users [get]
//func (cr *UserHandler) FindAll(c *gin.Context) {
//	users, err := cr.userUseCase.FindAll(c.Request.Context())
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, response.Response{
//			StatusCode: 500,
//			Errors:     err.Error(),
//			Message:    "Something went wrong/ Internal server error",
//			Data:       nil,
//		})
//		return
//	}
//	c.JSON(http.StatusOK, users)
//}

// @Summary Block user with User ID
// @ID block-user
// @Description Block a user with user ID, which restricts the user from accessing the ecommerce platform
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id body string true "user id"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /adminpanel/block [patch]
//func (cr *UserHandler) BlockUser(c *gin.Context) {
//	// Define a struct to represent the expected JSON request body
//	var body struct {
//		Email string
//	}
//	// Parse the JSON request body into the 'body' struct
//	err := c.Bind(&body)
//	if err != nil {
//		// If the request body is malformed, return a bad request error
//		c.JSON(http.StatusBadRequest, response.Response{
//			StatusCode: 400,
//			Errors:     err.Error(),
//			Message:    "Invalid request body",
//			Data:       nil,
//		})
//		return
//	}
//	// Call the 'BlockUser' function from the userUseCase
//	user, err := cr.userUseCase.BlockUser(c.Request.Context(), body.Email)
//	if err != nil {
//		// If the 'BlockUser' function returns an error, return an internal server error
//		c.JSON(http.StatusInternalServerError, response.Response{
//			StatusCode: 500,
//			Errors:     err.Error(),
//			Message:    "Unable to block user",
//			Data:       nil,
//		})
//		return
//	}
//	// If the 'BlockUser' function returns successfully, return an HTTP 200 OK response
//	c.JSON(http.StatusOK, response.Response{
//		StatusCode: 200,
//		Errors:     nil,
//		Message:    "Successfully blocked user",
//		Data:       user.UserID,
//	})
//}

// @Summary Unblock user with User ID
// @ID unblock-user
// @Description Unblock a user with user ID, which restores the access for user to the ecommerce platform
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id body string true "user id"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /adminpanel/unblock [patch]
//func (cr *UserHandler) UnblockUser(c *gin.Context) {
//	// define a struct to receive the values from request body
//	var body struct {
//		Email string
//	}
//	// Parse the JSON request body into the 'body' struct
//	err := c.Bind(&body)
//	if err != nil {
//		// If the request body is malformed, return a bad request error
//		c.JSON(http.StatusBadRequest, response.Response{
//			StatusCode: 400,
//			Errors:     err.Error(),
//			Message:    "Invalid request body",
//			Data:       nil,
//		})
//		return
//	}
//	// Call the UnblockUser method from userUseCase
//	user, err := cr.userUseCase.UnblockUser(c.Request.Context(), body.Email)
//
//	if err != nil {
//		// If the 'UnblockUser' function returns an error, return an internal server error
//		c.JSON(http.StatusInternalServerError, response.Response{
//			StatusCode: 500,
//			Errors:     err.Error(),
//			Message:    "Unable to unblock the user",
//			Data:       nil,
//		})
//		return
//	}
//	// If the 'UnblockUser' function returns successfully, return an HTTP 200 OK response
//	c.JSON(http.StatusOK, response.Response{
//		StatusCode: 200,
//		Errors:     nil,
//		Message:    "Successfully unblocked user",
//		Data:       user.UserID,
//	})
//}
