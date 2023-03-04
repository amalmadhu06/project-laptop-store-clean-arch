package interfaces

import "context"

type OtpRepository interface {
	UpdateAsVerified(ctx context.Context, phone string) error
}
