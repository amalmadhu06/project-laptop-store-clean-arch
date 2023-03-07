package usecase

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/golang-jwt/jwt/v4"
)

type cartUseCase struct {
	cartRepo interfaces.CartRepository
}

func NewCartUseCase(repo interfaces.CartRepository) services.CartUseCases {
	return &cartUseCase{
		cartRepo: repo,
	}
}

func (c *cartUseCase) AddToCart(ctx context.Context, cookie string, productItemID int) (domain.CartItems, error) {
	var addedProduct domain.CartItems
	userID, err := findUserID(cookie)
	if err != nil {
		return addedProduct, err
	}
	addedProduct, err = c.cartRepo.AddToCart(ctx, userID, productItemID)
	return addedProduct, err
}

func (c *cartUseCase) RemoveFromCart(ctx context.Context, cookie string, productItemID int) error {
	userID, err := findUserID(cookie)
	if err != nil {
		return err
	}
	err = c.cartRepo.RemoveFromCart(ctx, userID, productItemID)
	return err
}

//func (c *cartUseCase) ViewCart(ctx context.Context, cookie string) (modelHelper.ViewCart, error) {
//	var cart modelHelper.ViewCart
//	userID, err := findUserID(cookie)
//	if err != nil {
//		return cart, err
//	}
//	cart, err = c.cartRepo.ViewCart(ctx, userID)
//	return cart, err
//}

func findUserID(cookie string) (int, error) {
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
