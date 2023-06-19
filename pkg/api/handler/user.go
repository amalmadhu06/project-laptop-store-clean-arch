package handler

import (
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/api/handlerUtil"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// @title Ecommerce REST API
// @version 1.0
// @description Ecommerce REST API built using Go Lang, PSQL, REST API following Clean Architecture.
//              Hosted with Nginx, AWS EC2, and RDS.

// @contact
// name: Amal Madhu
// url: https://github.com/amalmadhu06
// email: madhuamal06@gmail.com

// @license
// name: MIT
// url: https://opensource.org/licenses/MIT

// @host localhost:3000

// @BasePath /
// @Accept json
// @Produce json
// @Router / [get]

// CreateUser
// @Summary Create a new user
// @ID create-user
// @Description Create a new user with the specified details.
// @Tags Users
// @Accept json
// @Produce json
// @Param user_details body model.UserDataInput true "User details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /signup [post]
func (cr *UserHandler) CreateUser(c *gin.Context) {
	//cancelling the request if it is taking more than one minute to send back a response
	//ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute)
	//defer cancel()
	// 1. receive data from request body
	var body model.UserDataInput
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err.Error())
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
	userData, err := cr.userUseCase.CreateUser(c.Request.Context(), body)
	if err != nil {
		// Return a 400 Bad request response if there is an error while creating the user.
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to create user",
			Data:       model.UserDataOutput{},
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
// @Param user_details body model.UserLoginEmail true "User details"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /login/email [post]
func (cr *UserHandler) LoginWithEmail(c *gin.Context) {
	//receive data from request body
	var body model.UserLoginEmail
	if err := c.BindJSON(&body); err != nil {
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
			Data:       model.UserDataOutput{},
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
// @Param user_details body model.UserLoginPhone true "User details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /login/phone [post]
func (cr *UserHandler) LoginWithPhone(c *gin.Context) {
	// receive data from request body
	var body model.UserLoginPhone
	if err := c.BindJSON(&body); err != nil {
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
			Data:       model.UserDataOutput{},
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
// @Param user_address body model.AddressInput true "User address"
// @Success 201 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /addresses/ [post]
func (cr *UserHandler) AddAddress(c *gin.Context) {
	// receive data from request body
	var body model.AddressInput
	if err := c.Bind(&body); err != nil {
		// Return a 421 response if the request body is malformed.
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "unable to process the request", Data: nil, Errors: err.Error()})
		return
	}
	userID, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 400, Message: "unable to fetch user id from context", Data: nil, Errors: err.Error()})
		return
	}
	address, err := cr.userUseCase.AddAddress(c.Request.Context(), body, userID)
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
// @Param user_address body model.AddressInput true "User address"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /addresses/ [put]
func (cr *UserHandler) UpdateAddress(c *gin.Context) {
	var body model.AddressInput
	if err := c.Bind(&body); err != nil {
		// Return a 421 response if the request body is malformed.
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "unable to process the request", Data: nil, Errors: err.Error()})
		return
	}

	userID, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 400, Message: "unable to fetch user id from context", Data: nil, Errors: err.Error()})
		return
	}

	address, err := cr.userUseCase.UpdateAddress(c.Request.Context(), body, userID)
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
// @Param page query int false "Page number for pagination"
// @Param limit query int false "Number of items to retrieve per page"
// @Param query query string false "Search query string"
// @Param filter query string false "Filter criteria for the users"
// @Param sort_by query string false "Sorting criteria for the users"
// @Param sort_desc query bool false "Sorting in descending order"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/users [get]
func (cr *UserHandler) ListAllUsers(c *gin.Context) {

	var viewUserInfo model.QueryParams

	viewUserInfo.Page, _ = strconv.Atoi(c.Query("page"))
	viewUserInfo.Limit, _ = strconv.Atoi(c.Query("limit"))
	viewUserInfo.Query = c.Query("query")
	viewUserInfo.Filter = c.Query("filter")
	viewUserInfo.SortBy = c.Query("sort_by")
	viewUserInfo.SortDesc, _ = strconv.ParseBool(c.Query("sort_desc"))

	users, err := cr.userUseCase.ListAllUsers(c.Request.Context(), viewUserInfo)
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
// @Router /admin/users/{id} [get]
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
// @Param user_id body model.BlockUser true "ID of the user to be blocked"
// @Success 200 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/users/block [put]
func (cr *UserHandler) BlockUser(c *gin.Context) {
	var blockUser model.BlockUser
	if err := c.Bind(&blockUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to read request body", Data: nil, Errors: err.Error()})
		return
	}

	adminID, err := handlerUtil.GetAdminIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 400, Message: "unable to fetch admin id from context", Data: nil, Errors: err.Error()})
		return
	}

	blockedUser, err := cr.userUseCase.BlockUser(c.Request.Context(), blockUser, adminID)
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
// @Router /admin/users/unblock/{id} [put]
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

// UserProfile
// @Summary User can view their profile
// @ID user-profile
// @Description Users can visit their profile
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /profile [get]
func (cr *UserHandler) UserProfile(c *gin.Context) {
	userID, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 400, Message: "unable to fetch user id from context", Data: nil, Errors: err.Error()})
		return
	}

	userProfile, err := cr.userUseCase.UserProfile(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to fetch user profile data", Data: nil, Errors: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully fetched user profile", Data: userProfile, Errors: nil})
}
