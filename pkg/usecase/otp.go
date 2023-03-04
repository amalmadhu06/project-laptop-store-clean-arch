package usecase

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/config"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
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

func (c *otpUseCase) SendOtp(ctx context.Context, phone modelHelper.OTPData) error {
	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.cfg.TWILIOACCOUNTSID,
		Password: c.cfg.TWILIOAUTHTOKEN,
	})
	params := &openapi.CreateVerificationParams{}
	params.SetTo("+91" + phone.Phone)
	params.SetChannel("sms")

	_, err := client.VerifyV2.CreateVerification(c.cfg.TWILIOSERVICESID, params)

	return err
}

func (c *otpUseCase) ValidateOtp(ctx context.Context, data modelHelper.VerifyData) (*openapi.VerifyV2VerificationCheck, error) {
	var resp *openapi.VerifyV2VerificationCheck
	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.cfg.TWILIOACCOUNTSID,
		Password: c.cfg.TWILIOAUTHTOKEN,
	})
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo("+91" + data.Phone.Phone)
	params.SetCode(data.Otp)
	resp, err := client.VerifyV2.CreateVerificationCheck(c.cfg.TWILIOSERVICESID, params)

	//update database on successful phone number verification
	fmt.Println("No error till approving otp")
	if *resp.Status == "approved" {
		//Todo : Update in database
		err = c.otpRepo.UpdateAsVerified(ctx, data.Phone.Phone)
		fmt.Println("update usecase called")
	}

	return resp, err
}
