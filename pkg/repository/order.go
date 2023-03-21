package repository

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type orderDatabase struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &orderDatabase{DB}
}

func (c *orderDatabase) BuyProductItem(ctx context.Context, userID int, orderInfo modelHelper.PlaceOrder) (domain.Order, error) {

	tx := c.DB.Begin()
	// finding product price and qnty
	var productItem struct {
		Price       float64
		QntyInStock int
	}

	fetchPriceQuery := `SELECT price, qnty_in_stock FROM product_items WHERE id = $1`

	err := tx.Raw(fetchPriceQuery, orderInfo.ProductItemID).Scan(&productItem).Error
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}

	//if stock is empty
	if productItem.QntyInStock < 1 {
		tx.Rollback()
		return domain.Order{}, fmt.Errorf("product item out of stock")
	}

	//fetch coupon details
	var couponInfo domain.Coupon
	fetchCouponQuery := `SELECT * FROM coupons WHERE id = $1;`
	err = tx.Raw(fetchCouponQuery, orderInfo.CouponID).Scan(&couponInfo).Error

	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}

	if productItem.Price < couponInfo.MinOrderValue {
		tx.Rollback()
		return domain.Order{}, fmt.Errorf("cannot apply coupon as order values is less than required")
	}

	discountAmount := productItem.Price * (couponInfo.DiscountPercent / 100)
	if discountAmount > couponInfo.DiscountMaxAmount {
		discountAmount = couponInfo.DiscountMaxAmount
	}
	orderTotal := productItem.Price - discountAmount

	var orderDetails domain.Order

	createOrderQuery := `	INSERT INTO orders (user_id,order_date,payment_method_id,shipping_address_id,order_total,order_status_id, coupon_id)
							VALUES($1, NOW(), $2, $3, $4, 1, $5) RETURNING *;`

	err = tx.Raw(createOrderQuery, userID, orderInfo.PaymentMethodID, orderInfo.ShippingAddressID, orderTotal, orderInfo.CouponID).Scan(&orderDetails).Error
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}

	createOrderLineQuery := `	INSERT INTO order_lines (product_item_id,order_id,quantity,price)
								VALUES ($1, $2, 1, $3);`

	err = tx.Exec(createOrderLineQuery, orderInfo.ProductItemID, orderDetails.ID, orderDetails.OrderTotal).Error
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}
	//reduce the stock quantity of the product item by 1
	updateQtyQuery := `	UPDATE product_items
						SET qnty_in_stock = qnty_in_stock -1 
						WHERE qnty_in_stock > 0 
						AND id = $1`
	err = tx.Exec(updateQtyQuery, orderInfo.ProductItemID).Error
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}

	//create an entry in the payment_details table
	createPaymentEntry := `	INSERT INTO payment_details (order_id, order_total,payment_method_id, payment_status_id, updated_at) 	
							VALUES ($1, $2,$3, 1,NOW());`
	err = tx.Exec(createPaymentEntry, orderDetails.ID, orderDetails.OrderTotal, orderDetails.PaymentMethodID).Error
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}

	tx.Commit()
	return orderDetails, nil
}

func (c *orderDatabase) BuyAll(ctx context.Context, userID int, orderInfo modelHelper.PlaceAllOrders) (domain.Order, error) {
	tx := c.DB.Begin()
	var cartDetails struct {
		ID    int
		Total float64
	}
	findCart := `SELECT id, total FROM carts WHERE user_id = $1`
	err := tx.Raw(findCart, userID).Scan(&cartDetails).Error

	if cartDetails.ID == 0 {
		tx.Rollback()
		return domain.Order{}, fmt.Errorf("no items in cart")
	}
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}
	var cartItems []domain.CartItems
	fetchCartItemsQuery := `SELECT * FROM cart_items WHERE cart_id = $1`
	err = tx.Raw(fetchCartItemsQuery, cartDetails.ID).Scan(&cartItems).Error

	if len(cartItems) == 0 {
		tx.Rollback()
		return domain.Order{}, fmt.Errorf("nothing in cart")
	}

	var createdOrder domain.Order
	createOrderQuery := `	INSERT INTO orders (user_id, order_date, payment_method_id, shipping_address_id, order_total, order_status_id)
							VALUES($1, NOW(), $2, $3, $4,1) RETURNING *;`
	err = tx.Raw(createOrderQuery, userID, orderInfo.PaymentMethodID, orderInfo.ShippingAddressID, cartDetails.Total).Scan(&createdOrder).Error
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}

	//update carts table
	updateCartQuery := `UPDATE carts SET total = 0 WHERE user_id = $1`
	fmt.Println("user id in repository for buy all : ", userID)
	err = tx.Exec(updateCartQuery, userID).Error
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}

	//update cart_items table
	deleteCartItemRowsQuery := `DELETE FROM cart_items WHERE cart_id = $1;`
	err = tx.Exec(deleteCartItemRowsQuery, cartDetails.ID).Error
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}

	//create an entry in the payment_details table
	createPaymentEntry := `	INSERT INTO payment_details (order_id, order_total,payment_method_id, payment_status_id, updated_at) 	
							VALUES ($1, $2,$3, 1,NOW());`
	err = tx.Exec(createPaymentEntry, createdOrder.ID, createdOrder.OrderTotal, createdOrder.PaymentMethodID).Error
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}

	createOrderLineQuery := `	INSERT INTO order_lines (product_item_id, order_id, quantity, price) VALUES($1, $2, $3, $4);`

	for i := range cartItems {
		//check if product is in stock and fetch product
		var productDetails struct {
			QntyInStock int
			Price       float64
		}

		fetchDetailsQuery := ` SELECT qnty_in_stock, price FROM product_items WHERE id = $1`
		err := tx.Raw(fetchDetailsQuery, cartItems[i].ProductItemID).Scan(&productDetails).Error
		if err != nil {
			tx.Rollback()
			return domain.Order{}, err
		}

		//if product is out of stock
		if productDetails.QntyInStock < int(cartItems[i].Quantity) {
			tx.Rollback()
			return domain.Order{}, fmt.Errorf("product item out of stock for id : %v ", cartItems[i].ProductItemID)
		}

		// creating order line
		productTotal := productDetails.Price * float64(cartItems[i].Quantity)
		err = tx.Exec(createOrderLineQuery, cartItems[i].ProductItemID, createdOrder.ID, cartItems[i].Quantity, productTotal).Error
		if err != nil {
			tx.Rollback()
			return domain.Order{}, err
		}

		//	reducing quantity in stock
		reduceQuantityQuery := ` 	UPDATE product_items SET qnty_in_stock = qnty_in_stock - $1 WHERE id = $2`
		err = tx.Exec(reduceQuantityQuery, cartItems[i].Quantity, cartItems[i].ProductItemID).Error
		if err != nil {
			tx.Rollback()
			return domain.Order{}, err
		}
	}
	tx.Commit()
	return createdOrder, nil
}

func (c *orderDatabase) ViewOrderById(ctx context.Context, userID int, orderID int) (domain.Order, error) {
	fmt.Println("user id :", userID, "order id  :", orderID)
	var order domain.Order
	viewOrderQuery := `SELECT * FROM orders WHERE user_id = $1 AND id = $2;`
	err := c.DB.Raw(viewOrderQuery, userID, orderID).Scan(&order).Error
	fmt.Println("order in repo : ", order)
	//if no order is found
	if order.ID == 0 {
		return domain.Order{}, fmt.Errorf("no order found")
	}

	return order, err
}

func (c *orderDatabase) ViewAllOrders(ctx context.Context, userID int) ([]domain.Order, error) {
	var orders []domain.Order
	viewAllOrdersQuery := `SELECT * FROM orders WHERE user_id = $1`
	err := c.DB.Raw(viewAllOrdersQuery, userID).Scan(&orders).Error
	return orders, err
}

func (c *orderDatabase) CancelOrder(ctx context.Context, userID int, orderID int) (domain.Order, error) {
	tx := c.DB.Begin()

	//find order details. If order is in pending status, user can cancel the order. If order is not in pending status, user cannot cancel the order
	var orderStatus int
	viewStatusQuery := `SELECT order_status_id FROM orders WHERE user_id = $1 AND id = $2`
	err := tx.Raw(viewStatusQuery, userID, orderID).Scan(&orderStatus).Error
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}
	//If no order if found
	if orderStatus == 0 {
		return domain.Order{}, fmt.Errorf("no such order found")
	}

	//if order is in pending status
	if orderStatus == 1 {

		var cancelledOrder domain.Order
		cancelOrderQuery := `UPDATE orders SET order_status_id = 2 WHERE user_id = $1 AND id = $2 RETURNING *;`
		err := tx.Raw(cancelOrderQuery, userID, orderID).Scan(&cancelledOrder).Error
		if err != nil {
			tx.Rollback()
			return domain.Order{}, err
		}

		//increase the quantity in product items table
		var orderLineItems []domain.OrderLine
		findOrderLineQuery := `SELECT * FROM order_lines WHERE order_id = $1;`
		err = tx.Raw(findOrderLineQuery, orderID).Scan(&orderLineItems).Error
		if err != nil {
			tx.Rollback()
			return domain.Order{}, err
		}

		qntyUpdateQuery := `UPDATE product_items SET qnty_in_stock = qnty_in_stock + $1 WHERE id = $2`
		for i := range orderLineItems {
			err := tx.Exec(qntyUpdateQuery, orderLineItems[i].Quantity, orderLineItems[i].ProductItemID).Error
			if err != nil {
				return domain.Order{}, err
			}
		}

		tx.Commit()
		return cancelledOrder, nil
	}
	//if order is already cancelled
	if orderStatus == 2 {
		tx.Rollback()
		return domain.Order{}, fmt.Errorf("order already cancelled")
	}
	tx.Rollback()
	return domain.Order{}, fmt.Errorf("order processed, cannot cancel")
}

func (c *orderDatabase) UpdateOrder(ctx context.Context, orderInfo modelHelper.UpdateOrder) (domain.Order, error) {
	var updatedOrder domain.Order
	updateStatusQuery := `UPDATE orders SET order_status_id = $1 WHERE id = $2 RETURNING *`
	err := c.DB.Raw(updateStatusQuery, orderInfo.OrderStatusID, orderInfo.OrderID).Scan(&updatedOrder).Error

	if updatedOrder.ID == 0 {
		return domain.Order{}, fmt.Errorf("no order found")
	}

	return updatedOrder, err
}
