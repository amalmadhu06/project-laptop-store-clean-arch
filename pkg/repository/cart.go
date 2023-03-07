package repository

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type cartDatabase struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &cartDatabase{DB}
}

func (c *cartDatabase) AddToCart(ctx context.Context, userID int, productItemID int) (domain.CartItems, error) {
	//Begin transaction
	tx := c.DB.Begin()

	//checking is user has a cart
	var cartID int
	cartCheckQuery := `	SELECT id
						FROM carts
						WHERE user_id = ?
						LIMIT 1`
	err := tx.Raw(cartCheckQuery, userID).Scan(&cartID).Error
	fmt.Println("1. cart ID fetched ", cartID)
	fmt.Println("1. Error if any", err)
	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}
	if cartID == 0 {
		//If user has no cart, creating one
		err := tx.Raw("INSERT INTO carts (user_id, total) VALUES ($1,0) RETURNING id", userID).Scan(&cartID).Error
		fmt.Println("2. cart ID fetched ", cartID)
		fmt.Println("2. Error if any", err)
		if err != nil {
			tx.Rollback()
			return domain.CartItems{}, err
		}
	}

	//checking if productItem is already present in the cart
	var cartItem domain.CartItems
	err = tx.Raw("SELECT id, quantity FROM cart_items WHERE cart_id = $1 AND product_item_id = $2 LIMIT 1", cartID, productItemID).Scan(&cartItem).Error
	fmt.Println("3. cart item id fetched ", cartItem.ID)
	fmt.Println("3. Error if any", err)
	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}
	//if item is not present in the cart
	if cartItem.ID == 0 {
		err := tx.Raw("INSERT INTO cart_items (cart_id, product_item_id, quantity) VALUES ($1, $2, 1) RETURNING *;", cartID, productItemID).Scan(&cartItem).Error
		fmt.Println("4. cart Item id fetched ", cartItem.CartID)
		fmt.Println("4. Error if any", err)

		if err != nil {
			tx.Rollback()
			return domain.CartItems{}, err
		}
	} else {
		//	if item is already present in the cart
		err := tx.Raw("UPDATE cart_items SET quantity = $1 WHERE id = $2 RETURNING *;", cartItem.Quantity+1, cartItem.ID).Scan(&cartItem).Error
		fmt.Println("5. cart ID fetched ", cartItem)
		fmt.Println("5. Error if any", err)

		if err != nil {
			tx.Rollback()
			return domain.CartItems{}, err
		}
	}

	//update total in cart table
	//product_item_id is known, quantity is known, cart_id is known
	//fetch price from product_items table
	var currentTotal, itemPrice float64
	err = tx.Raw("SELECT price FROM product_items WHERE id = $1", productItemID).Scan(&itemPrice).Error
	fmt.Println("item price :", itemPrice)
	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}
	//fetch current total from cart table
	err = tx.Raw("SELECT total FROM carts WHERE id = $1", cartItem.CartID).Scan(&currentTotal).Error
	fmt.Println("current total :", currentTotal)

	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}
	//add price of new product item to the current total and update it in the cart table
	newTotal := currentTotal + itemPrice
	fmt.Println("new total :", newTotal)

	err = tx.Exec("UPDATE carts SET total = $1", newTotal).Error
	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}
	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return domain.CartItems{}, err
	}
	return cartItem, nil
}

func (c *cartDatabase) RemoveFromCart(ctx context.Context, userID int, productItemID int) error {
	tx := c.DB.Begin()
	//find cart_id from carts table
	var cartID int
	err := tx.Raw("SELECT id FROM carts WHERE user_id = $1", userID).Scan(&cartID).Error
	fmt.Println("fetched cart id", cartID)
	if err != nil {
		tx.Rollback()
		return err
	}
	//	find the quantity
	var quantity int
	err = tx.Raw("SELECT quantity FROM cart_items WHERE cart_id = $1 AND product_item_id = $2", cartID, productItemID).Scan(&quantity).Error
	fmt.Println("fetched quantity", quantity)
	if err != nil {
		fmt.Println("triggering error 1")
		tx.Rollback()
		return err
	}
	//	if quantity is 1, delete the row
	if quantity == 0 {
		fmt.Println("triggering error 2")
		tx.Rollback()
		return fmt.Errorf("nothing to remove")
	} else if quantity == 1 {
		err := tx.Exec("DELETE FROM cart_items WHERE cart_id = $1 AND product_item_id = $2", cartID, productItemID).Error
		if err != nil {
			fmt.Println("triggering error 4")
			tx.Rollback()
			return err
		}
	} else {
		err := tx.Exec("UPDATE cart_items SET quantity = quantity - 1 WHERE cart_id = $1 AND product_item_id = $2", cartID, productItemID).Error
		if err != nil {
			fmt.Println("triggering error 5")
			tx.Rollback()
			return err
		}
	}

	//fetch price from product_items table
	var currentTotal, itemPrice float64
	err = tx.Raw("SELECT price FROM product_items WHERE id = $1", productItemID).Scan(&itemPrice).Error
	fmt.Println("item price :", itemPrice)
	if err != nil {
		tx.Rollback()
		return err
	}
	//fetch current total from cart table
	err = tx.Raw("SELECT total FROM carts WHERE id = $1", cartID).Scan(&currentTotal).Error
	fmt.Println("current total :", currentTotal)

	if err != nil {
		tx.Rollback()
		return err
	}
	//subtract price of removed product item from the current total and update it in the cart table
	newTotal := currentTotal - itemPrice
	fmt.Println("new total :", newTotal)

	fmt.Println("final cart id check", cartID)
	//todo: fix update total in cart issue
	result := tx.Exec("UPDATE carts SET total = $1 WHERE user_id = $2", newTotal, userID)
	fmt.Println("final result :", result)

	if result.Error != nil {
		tx.Rollback()
		return err
	}
	return nil
}
