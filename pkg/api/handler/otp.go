package handler

import (
	//"context"
	// "net/http"

	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OtpHandler struct {
	otpUseCase services.OtpUseCase
	//cfg config.Config
}

func NewOtpHandler(otpUsecase services.OtpUseCase) *OtpHandler {
	return &OtpHandler{
		otpUseCase: otpUsecase,
		//cfg: cfg,
	}
}

// SendOtp
// @Summary Send OTP to user's mobile
// @ID send-otp
// @Description Send OTP to use's mobile
// @Tags Otp
// @Accept json
// @Produce json
// @Param user_mobile body model.OTPData true "User mobile number"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /send-otp [post]
func (cr *OtpHandler) SendOtp(c *gin.Context) {
	//declare a variable to receive data from request
	var phone model.OTPData
	if err := c.Bind(&phone); err != nil {
		//is request body is malformed, sent back 422 status code
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "Invalid OTP",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	//calling SendOtp method in otp use case
	err := cr.otpUseCase.SendOtp(c.Request.Context(), phone)
	if err != nil {
		//send back status code if unable to send otp
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Failed to send OTP",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	//If otp is successfully sent
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "OTP sent successfully",
		Data:       nil,
		Errors:     nil,
	})
}

// ValidateOtp
// @Summary Validate the OTP to user's mobile
// @ID validate-otp
// @Description Validate the  OTP sent to use's mobile
// @Tags Otp
// @Accept json
// @Produce json
// @Param otp body model.VerifyData true "OTP sent to user's mobile number"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /verify-otp [post]
func (cr *OtpHandler) ValidateOtp(c *gin.Context) {
	//declare a variable to receive request data
	var otpDetails model.VerifyData
	//receiving request data
	if err := c.Bind(&otpDetails); err != nil {
		//if request body is malformed, send back 422 status code
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "Unable to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	//call validateOtp method from otp use case
	resp, userData, ss, err := cr.otpUseCase.ValidateOtp(c.Request.Context(), otpDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Unable to verify the OTP",
			Data:       nil,
			Errors:     err.Error(),
		})
	} else if *resp.Status == "approved" {
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("UserAuth", ss, 3600*24*30, "", "", false, true)
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 200,
			Message:    "Successfully verified OTP and logged in",
			Data:       userData,
			Errors:     nil,
		})
	} else {
		c.JSON(http.StatusUnauthorized, response.Response{
			StatusCode: 401,
			Message:    "Invalid OTP",
			Data:       nil,
			Errors:     nil,
		})
	}
}
