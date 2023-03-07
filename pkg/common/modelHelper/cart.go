package modelHelper

import "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"

type ViewCart struct {
	cartItems []domain.CartItems
	Total     float64
}
