package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type OtpUseCase interface {
	SendOtp(ctx context.Context, input modelHelper.OTPData) error
	ValidateOtp(ctx context.Context, data modelHelper.VerifyData) (*openapi.VerifyV2VerificationCheck, modelHelper.UserDataOutput, string, error)
}
