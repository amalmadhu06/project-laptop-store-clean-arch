package handler

import (
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/response"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

func NewAdminHandler(adminUseCase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: adminUseCase,
	}
}

// CreateAdmin
// @Summary Create a new admin from admin panel
// @ID create-admin
// @Description Super admin can create a new admin from admin panel.
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin_details body modelHelper.NewAdminInfo true "New Admin details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /adminPanel/create-admin [post]
func (cr *AdminHandler) CreateAdmin(c *gin.Context) {
	var newAdminInfo modelHelper.NewAdminInfo
	if err := c.Bind(&newAdminInfo); err != nil {
		//if request body is malformed, return 422
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "unable to read the request body", Data: nil, Errors: err.Error()})
		return
	}
	cookie, err := c.Cookie("AdminAuth")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to find cookie", Data: nil, Errors: err.Error()})
		return
	}

	adminID, err := cr.adminUseCase.FindAdminID(c.Request.Context(), cookie)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to fetch admin id", Data: nil, Errors: err.Error()})
		return
	}

	//	call CreateAdmin method from Admin Usecase
	newAdminOutput, err := cr.adminUseCase.CreateAdmin(c.Request.Context(), newAdminInfo, adminID)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to create admin", Data: nil, Errors: err.Error()})
		return
	}

	//	return 201 status created is new admin is created successfully
	c.JSON(http.StatusCreated, response.Response{StatusCode: 201, Message: "admin created successfully", Data: newAdminOutput, Errors: nil})

}

// AdminLogin
// @Summary Admin Login
// @ID admin-login
// @Description Admin login
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin_credentials body modelHelper.AdminLogin true "Admin login credentials"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /adminPanel/login [post]
func (cr *AdminHandler) AdminLogin(c *gin.Context) {
	// receive data from request body
	var body modelHelper.AdminLogin
	if err := c.Bind(&body); err != nil {
		// Return a 421 response if the request body is malformed.
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "unable to process the request", Data: nil, Errors: err.Error()})
		return
	}
	// Call the UserLogin method of the userUseCase to login as a user.
	ss, admin, err := cr.adminUseCase.AdminLogin(c.Request.Context(), body)
	if err != nil {
		// Return a 400 Bad Request response if there is an error while creating the user.
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to login", Data: nil, Errors: err.Error()})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AdminAuth", ss, 3600*24*30, "", "", false, true)
	// Return a 201 Created response if the user is successfully logged in.
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "Successfully logged in", Data: admin, Errors: nil})
}

// AdminLogout
// @Summary Admin Logout
// @ID admin-logout
// @Description Logs out a logged-in admin from the E-commerce web api admin panel
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Failure 500
// @Router /adminPanel/logout [get]
func (cr *AdminHandler) AdminLogout(c *gin.Context) {
	// Set the user authentication cookie's expiration to -1 to invalidate it.
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") //indicates to the client that it should not cache any response data and should always revalidate it with the server
	c.SetSameSite(http.SameSiteLaxMode)                                           //sets the SameSite cookie attribute to "Lax" for the response. This attribute restricts the scope of cookies and helps prevent cross-site request forgery attacks
	c.SetCookie("AdminAuth", "", -1, "", "", false, true)                         //Immediately by setting the maxAge to -1, and marks the cookie as secure and HTTP-only
	c.Status(http.StatusOK)
}

// BlockAdmin
// @Summary Block an admin
// @ID block-admin
// @Description Super-admin can block admins
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin_id path string true "ID of the admin to be blocked"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /adminPanel/block-admin/:admin-id [put]
func (cr *AdminHandler) BlockAdmin(c *gin.Context) {
	//	get the id of the admin to be blocked
	paramsID := c.Param("admin-id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to read admin id from path parameter", Data: nil, Errors: err})
		return
	}
	cookie, err := c.Cookie("AdminAuth")
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 401, Message: "failed to verify admin", Data: nil, Errors: err})
		return
	}

	blockedAdmin, err := cr.adminUseCase.BlockAdmin(c.Request.Context(), id, cookie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to block admin", Data: nil, Errors: err})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully blocked admin", Data: blockedAdmin, Errors: nil})
}

// UnblockAdmin
// @Summary Unblock a blocked admin
// @ID unblock-admin
// @Description Super-admin can unblock a blocked admin
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin_id path string true "ID of the admin to be unblocked"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /adminPanel/unblock-admin/:admin-id [put]
func (cr *AdminHandler) UnblockAdmin(c *gin.Context) {
	//	get the id of the admin to be blocked
	paramsID := c.Param("admin-id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to read admin id from path parameter", Data: nil, Errors: err})
		return
	}
	cookie, err := c.Cookie("AdminAuth")
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 401, Message: "failed to verify admin", Data: nil, Errors: err})
		return
	}

	unblockedAdmin, err := cr.adminUseCase.UnblockAdmin(c.Request.Context(), id, cookie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to unblock admin", Data: nil, Errors: err})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully unblocked admin", Data: unblockedAdmin, Errors: nil})
}

// AdminDashboard
// @Summary Admin Dashboard
// @ID admin-dashboard
// @Description Admin can access dashboard and view details regarding orders, users, products, etc.
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /adminPanel/dashboard [get]
func (cr *AdminHandler) AdminDashboard(c *gin.Context) {
	dashboard, err := cr.adminUseCase.AdminDashboard(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to fetch admin dashboard data", Data: nil, Errors: err})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully fetched dashboard", Data: dashboard, Errors: nil})
}
