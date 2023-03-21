package usecase

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
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

func (c *cartUseCase) AddToCart(ctx context.Context, cookie string, productItemID int) (domain.CartItems, error) {
	var addedProduct domain.CartItems
	userID, err := FindUserID(cookie)
	if err != nil {
		return addedProduct, err
	}
	addedProduct, err = c.cartRepo.AddToCart(ctx, userID, productItemID)
	return addedProduct, err
}

func (c *cartUseCase) RemoveFromCart(ctx context.Context, cookie string, productItemID int) error {
	userID, err := FindUserID(cookie)
	if err != nil {
		return err
	}
	err = c.cartRepo.RemoveFromCart(ctx, userID, productItemID)
	return err
}

func (c *cartUseCase) ViewCart(ctx context.Context, cookie string) (modelHelper.ViewCart, error) {
	var cart modelHelper.ViewCart
	userID, err := FindUserID(cookie)
	if err != nil {
		return cart, err
	}
	cart, err = c.cartRepo.ViewCart(ctx, userID)
	return cart, err
}

func (c *cartUseCase) EmptyCart(ctx context.Context, cookie string) error {
	userID, err := FindUserID(cookie)
	if err != nil {
		return err
	}
	err = c.cartRepo.EmptyCart(ctx, userID)
	return err
}

func (c *cartUseCase) AddCouponToCart(ctx context.Context, userID int, couponID int) (modelHelper.ViewCart, error) {

	//checking is coupon is already used
	isUsed, err := c.productRepo.CouponUsed(ctx, userID, couponID)
	if err != nil {
		return modelHelper.ViewCart{}, err
	}
	if isUsed {
		return modelHelper.ViewCart{}, fmt.Errorf("user already used the coupon")
	}

	//fetching coupon details
	couponInfo, err := c.productRepo.ViewCouponByID(ctx, couponID)
	if err != nil {
		return modelHelper.ViewCart{}, err
	}
	//if no coupon found
	if couponInfo.ID == 0 {
		return modelHelper.ViewCart{}, fmt.Errorf("invalid coupon id")
	}
	//checking coupon validity
	currentTime := time.Now()
	if couponInfo.ValidTill.Before(currentTime) {
		return modelHelper.ViewCart{}, fmt.Errorf("coupon expired")
	}

	//	fetch cart total
	cartInfo, err := c.cartRepo.ViewCart(ctx, userID)

	fmt.Println("cart sub total ", cartInfo.SubTotal, "coupon min order value", couponInfo.MinOrderValue)

	if cartInfo.SubTotal < couponInfo.MinOrderValue {
		return modelHelper.ViewCart{}, fmt.Errorf("cart total not enogh for applying the coupon")
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
	fmt.Println("find user id : ", id)

	return id, nil
}
