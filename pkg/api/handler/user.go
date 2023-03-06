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
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
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
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /login-email [post]
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
// @Failure 500 {object} response.Response
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
