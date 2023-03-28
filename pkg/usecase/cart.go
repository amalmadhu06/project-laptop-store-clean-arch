package usecase

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type cartUseCase struct {
	cartRepo    interfaces.CartRepository
	productRepo interfaces.ProductRepository
}

func NewCartUseCase(cartRepo interfaces.CartRepository, productRepo interfaces.ProductRepository) services.CartUseCases {
	return &cartUseCase{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (c *cartUseCase) AddToCart(ctx context.Context, userID, productItemID int) (domain.CartItems, error) {
	var addedProduct domain.CartItems
	addedProduct, err := c.cartRepo.AddToCart(ctx, userID, productItemID)
	return addedProduct, err
}

func (c *cartUseCase) RemoveFromCart(ctx context.Context, userID, productItemID int) error {
	err := c.cartRepo.RemoveFromCart(ctx, userID, productItemID)
	return err
}

func (c *cartUseCase) ViewCart(ctx context.Context, userID int) (model.ViewCart, error) {
	var cart model.ViewCart
	cart, err := c.cartRepo.ViewCart(ctx, userID)
	return cart, err
}

func (c *cartUseCase) EmptyCart(ctx context.Context, userID int) error {
	err := c.cartRepo.EmptyCart(ctx, userID)
	return err
}

func (c *cartUseCase) AddCouponToCart(ctx context.Context, userID, couponID int) (model.ViewCart, error) {

	//checking is coupon is already used
	isUsed, err := c.productRepo.CouponUsed(ctx, userID, couponID)
	if err != nil {
		return model.ViewCart{}, err
	}
	if isUsed {
		return model.ViewCart{}, fmt.Errorf("user already used the coupon")
	}

	//fetching coupon details
	couponInfo, err := c.productRepo.ViewCouponByID(ctx, couponID)
	if err != nil {
		return model.ViewCart{}, err
	}
	//if no coupon found
	if couponInfo.ID == 0 {
		return model.ViewCart{}, fmt.Errorf("invalid coupon id")
	}
	//checking coupon validity
	currentTime := time.Now()
	if couponInfo.ValidTill.Before(currentTime) {
		return model.ViewCart{}, fmt.Errorf("coupon expired")
	}

	//	fetch cart total
	cartInfo, err := c.cartRepo.ViewCart(ctx, userID)

	if cartInfo.SubTotal < couponInfo.MinOrderValue {
		return model.ViewCart{}, fmt.Errorf("cart total not enogh for applying the coupon")
	}

	//	add coupon to the cart
	cart, err := c.cartRepo.AddCouponToCart(ctx, userID, couponID)
	return cart, err
}

func FindUserID(cookie string) (int, error) {
	//parses, validates, verifies the signature and returns the parsed token
	token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		//secret was used for signing the string
		return []byte("secret"), nil
	})
	if err != nil {
		return 0, err
	}
	var parsedID interface{}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		parsedID = claims["id"]
	}
	//type assertion
	value, ok := parsedID.(float64)
	if !ok {
		return 0, fmt.Errorf("expected an int value, but got %T", parsedID)
	}

	id := int(value)
	return id, nil
}
