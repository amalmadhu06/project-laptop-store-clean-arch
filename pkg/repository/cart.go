package repository

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
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

	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}

	if cartID == 0 {
		//If user has no cart, creating one
		err := tx.Raw("INSERT INTO carts (user_id, total) VALUES ($1,0) RETURNING id", userID).Scan(&cartID).Error

		if err != nil {
			tx.Rollback()
			return domain.CartItems{}, err
		}
	}

	//checking if productItem is already present in the cart
	var cartItem domain.CartItems
	err = tx.Raw("SELECT id, quantity FROM cart_items WHERE cart_id = $1 AND product_item_id = $2 LIMIT 1", cartID, productItemID).Scan(&cartItem).Error

	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}
	//if item is not present in the cart
	if cartItem.ID == 0 {
		err := tx.Raw("INSERT INTO cart_items (cart_id, product_item_id, quantity) VALUES ($1, $2, 1) RETURNING *;", cartID, productItemID).Scan(&cartItem).Error

		if err != nil {
			tx.Rollback()
			return domain.CartItems{}, err
		}
	} else {
		//	if item is already present in the cart
		err := tx.Raw("UPDATE cart_items SET quantity = $1 WHERE id = $2 RETURNING *;", cartItem.Quantity+1, cartItem.ID).Scan(&cartItem).Error

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

	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}
	//fetch current total from cart table
	err = tx.Raw("SELECT total FROM carts WHERE id = $1", cartItem.CartID).Scan(&currentTotal).Error

	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}
	//add price of new product item to the current total and update it in the cart table
	newTotal := currentTotal + itemPrice

	err = tx.Exec("UPDATE carts SET total = $1", newTotal).Error
	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}
	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}
	return cartItem, nil
}

func (c *cartDatabase) RemoveFromCart(ctx context.Context, userID int, productItemID int) error {
	tx := c.DB.Begin()
	//find cart_id from carts table
	var cartID int
	err := tx.Raw("SELECT id FROM carts WHERE user_id = $1", userID).Scan(&cartID).Error

	if err != nil {
		tx.Rollback()
		return err
	}
	//	find the quantity
	var quantity int
	err = tx.Raw("SELECT quantity FROM cart_items WHERE cart_id = $1 AND product_item_id = $2", cartID, productItemID).Scan(&quantity).Error

	if err != nil {
		tx.Rollback()
		return err
	}
	//	if quantity is 1, delete the row
	if quantity == 0 {
		tx.Rollback()
		return fmt.Errorf("nothing to remove")
	} else if quantity == 1 {
		err := tx.Exec("DELETE FROM cart_items WHERE cart_id = $1 AND product_item_id = $2", cartID, productItemID).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		err := tx.Exec("UPDATE cart_items SET quantity = quantity - 1 WHERE cart_id = $1 AND product_item_id = $2", cartID, productItemID).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	//fetch price from product_items table
	var itemPrice float64
	//var currentTotal float64
	err = tx.Raw("SELECT price FROM product_items WHERE id = $1", productItemID).Scan(&itemPrice).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	result := tx.Exec("UPDATE carts SET total = total-$1 WHERE user_id = $2", itemPrice, userID)

	if result.Error != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *cartDatabase) ViewCart(ctx context.Context, userID int) (modelHelper.ViewCart, error) {

	tx := c.DB.Begin()
	//find cart_id from carts table
	var cartDetails struct {
		ID    int
		Total float64
	}

	err := tx.Raw("SELECT id, total FROM carts WHERE user_id = $1", userID).Scan(&cartDetails).Error

	if err != nil {
		tx.Rollback()
		return modelHelper.ViewCart{}, err
	}

	var allItems []modelHelper.DisplayCart
	joinQuery := `	SELECT pi.id as product_item_id, b.brand, p.name, pi.model, ci.quantity, pi.product_item_image, pi.price, (ci.quantity * pi.price) AS total
					FROM cart_items ci 
					JOIN product_items pi
					ON ci.product_item_id = pi.id
					JOIN products p
					ON p.id = pi.product_id
					JOIN product_brands b
					ON b.id = p.brand_id
					WHERE ci.cart_id = $1
					`

	rows, err := tx.Raw(joinQuery, cartDetails.ID).Rows()
	if err != nil {
		tx.Rollback()
		return modelHelper.ViewCart{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var item modelHelper.DisplayCart
		err := rows.Scan(&item.ProductItemID, &item.Brand, &item.Name, &item.Model, &item.Quantity, &item.ProductItemImage, &item.Price, &item.Total)
		if err != nil {
			tx.Rollback()
			return modelHelper.ViewCart{}, err
		}
		allItems = append(allItems, item)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return modelHelper.ViewCart{}, err
	}
	var finalCart modelHelper.ViewCart

	finalCart.CartTotal = cartDetails.Total
	finalCart.CartItems = allItems

	return finalCart, nil
}

func (c *cartDatabase) EmptyCart(ctx context.Context, userID int) error {
	tx := c.DB.Begin()

	//set cart total as 0 and return cart_id
	var cartID int
	updateCartQuery := `UPDATE carts SET total = 0 WHERE user_id = $1 RETURNING id;`

	err := tx.Raw(updateCartQuery, userID).Scan(&cartID).Error
	fmt.Println("fetched cart id :", cartID)
	if err != nil {
		tx.Rollback()
		return err
	}

	deleteRowsQuery := `DELETE FROM cart_items WHERE cart_id = $1;`
	err = tx.Exec(deleteRowsQuery, cartID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
