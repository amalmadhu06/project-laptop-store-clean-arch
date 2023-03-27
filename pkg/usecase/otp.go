package usecase

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/config"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
	"time"
)

type otpUseCase struct {
	otpRepo interfaces.OtpRepository
	cfg     config.Config
}

func NewOtpUseCase(repo interfaces.OtpRepository, cfg config.Config) services.OtpUseCase {
	return &otpUseCase{
		otpRepo: repo,
		cfg:     cfg,
	}
}

func (c *otpUseCase) SendOtp(ctx context.Context, phone model.OTPData) error {

	//check if user exists
	isPresent, err := c.otpRepo.CheckWithMobile(ctx, phone.Phone)
	//if user is not present, return error
	if !isPresent {
		return fmt.Errorf("user not registered")
	}
	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.cfg.TWILIOACCOUNTSID,
		Password: c.cfg.TWILIOAUTHTOKEN,
	})
	params := &openapi.CreateVerificationParams{}
	params.SetTo("+91" + phone.Phone)
	params.SetChannel("sms")

	_, err = client.VerifyV2.CreateVerification(c.cfg.TWILIOSERVICESID, params)

	return err
}

func (c *otpUseCase) ValidateOtp(ctx context.Context, data model.VerifyData) (*openapi.VerifyV2VerificationCheck, model.UserDataOutput, string, error) {
	var resp *openapi.VerifyV2VerificationCheck
	var userData model.UserDataOutput

	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.cfg.TWILIOACCOUNTSID,
		Password: c.cfg.TWILIOAUTHTOKEN,
	})
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo("+91" + data.Phone.Phone)
	params.SetCode(data.Otp)
	resp, err := client.VerifyV2.CreateVerificationCheck(c.cfg.TWILIOSERVICESID, params)
	if err != nil {
		return resp, userData, "", err
	}
	//update database on successful phone number verification
	if *resp.Status == "approved" {
		//update as user verified
		err = c.otpRepo.UpdateAsVerified(ctx, data.Phone.Phone)
		//	login user

		// 1. Find the userData with given email
		user, err := c.otpRepo.FindByPhone(ctx, data.Phone.Phone)
		if err != nil {
			return resp, userData, "", fmt.Errorf("error finding userData")
		}
		if user.Email == "" {
			return resp, userData, "", fmt.Errorf("no such user found")
		}

		// 3. Check whether the userData is blocked by admin
		if user.IsBlocked {
			return resp, userData, "", fmt.Errorf("userData account is blocked")
		}

		// 4. Create JWT Token
		// creating jwt token and sending it in cookie
		claims := jwt.MapClaims{
			"id":  user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// singed string
		ss, err := token.SignedString([]byte("secret"))

		// 4. Send back the created token

		//user data for sending back as response
		userData.ID, userData.FName, userData.LName, userData.Email, userData.Phone = user.ID, user.FName, user.LName, user.Email, user.Phone

		return resp, userData, ss, err

	}
	return resp, userData, "", err
}
