package handler

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/response"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"net/http"
	"strconv"
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
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
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
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /login-email [post]
func (cr *UserHandler) LoginWithEmail(c *gin.Context) {
	//receive data from request body
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
	// Return a 200 success ok response if the user is successfully logged in.
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
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /login-phone [post]
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

// AddAddress
// @Summary User can add address
// @ID add-address
// @Description Add address
// @Tags Users
// @Accept json
// @Produce json
// @Param user_address body modelHelper.AddressInput true "User address"
// @Success 201 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /add-address [post]
func (cr *UserHandler) AddAddress(c *gin.Context) {
	// receive data from request body
	var body modelHelper.AddressInput
	if err := c.Bind(&body); err != nil {
		// Return a 421 response if the request body is malformed.
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "unable to process the request", Data: nil, Errors: err.Error()})
		return
	}
	cookie, err := c.Cookie("UserAuth")
	address, err := cr.userUseCase.AddAddress(c.Request.Context(), body, cookie)
	if err != nil {
		// Return a 400 Bad Request response if there is an error while creating the user.
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to add address", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201, Message: "Successfully added address", Data: address, Errors: nil,
	})
}

// UpdateAddress
// @Summary User can update existing address
// @ID update-address
// @Description Update address
// @Tags Users
// @Accept json
// @Produce json
// @Param user_address body modelHelper.AddressInput true "User address"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /update-address [post]
func (cr *UserHandler) UpdateAddress(c *gin.Context) {
	var body modelHelper.AddressInput
	if err := c.Bind(&body); err != nil {
		// Return a 421 response if the request body is malformed.
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "unable to process the request", Data: nil, Errors: err.Error()})
		return
	}
	cookie, err := c.Cookie("UserAuth")
	address, err := cr.userUseCase.UpdateAddress(c.Request.Context(), body, cookie)
	if err != nil {
		// Return a 400 Bad Request response if there is an error while creating the user.
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to update address", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200, Message: "Successfully updated address", Data: address, Errors: nil,
	})
}

// ListAllUsers
// @Summary Admin can list all registered users
// @ID list-all-users
// @Description Admin can list all registered users
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /adminPanel/users [get]
func (cr *UserHandler) ListAllUsers(c *gin.Context) {
	users, err := cr.userUseCase.ListAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed fetch users", Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200, Message: "Successfully fetched all users", Data: users, Errors: nil,
	})
}

// FindUserByID
// @Summary Admin can fetch a specific user details using user id
// @ID find-user-by-id
// @Description Admin can fetch a specific user details using user id
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id path string true "ID of the user to be fetched"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /adminPanel/users/:id [get]
func (cr *UserHandler) FindUserByID(c *gin.Context) {
	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to parse user id", Data: nil, Errors: err.Error()})
		return
	}
	user, err := cr.userUseCase.FindUserByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed fetch user", Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200, Message: "Successfully fetched user details", Data: user, Errors: nil,
	})

}

// BlockUser
// @Summary Admin can block a user
// @ID block-user
// @Description Admin can block a registered user
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id body modelHelper.BlockUser true "ID of the user to be blocked"
// @Success 200 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /adminPanel/users/block-user [put]
func (cr *UserHandler) BlockUser(c *gin.Context) {
	var blockUser modelHelper.BlockUser
	if err := c.Bind(&blockUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to read request body", Data: nil, Errors: err.Error()})
		return
	}
	//cookie for finding admin id
	cookie, err := c.Cookie("AdminAuth")
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 401, Message: "failed to fetch cookie from request", Data: nil, Errors: err.Error()})
		return
	}
	blockedUser, err := cr.userUseCase.BlockUser(c.Request.Context(), blockUser, cookie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to block user", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully blocked user", Data: blockedUser, Errors: nil})
}

// UnblockUser
// @Summary Admin can unblock a blocked user
// @ID unblock-user
// @Description Admin can unblock a blocked user
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id path string true "ID of the user to be unblocked"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /adminPanel/users/unblock-user/:id [put]
func (cr *UserHandler) UnblockUser(c *gin.Context) {
	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to parse user id", Data: nil, Errors: err.Error()})
		return
	}
	unblockedUser, err := cr.userUseCase.UnblockUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to unblock user", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully unblocked user", Data: unblockedUser, Errors: nil})
}
