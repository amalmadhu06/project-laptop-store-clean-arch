package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type OtpUseCase interface {
	SendOtp(ctx context.Context, input model.OTPData) error
	ValidateOtp(ctx context.Context, data model.VerifyData) (*openapi.VerifyV2VerificationCheck, model.UserDataOutput, string, error)
}
