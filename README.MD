# Project Laptop Store | Go | PostgresSQL | Clean Architecture

## Project Flow
1. Database design and implementation:
    - Design and implement the database schema in PSQL to store all necessary data for your web API.
tset rest
2. User authentication:
    - Implement user sign up with validation to ensure only valid users can create accounts.
    - Implement user login to allow registered users to access their account.
    - Implement OTP login to provide an additional layer of security for user authentication.

3. Product management:
    - Implement category management, allowing admins to add, edit, and delete product categories.
    - Implement product management, allowing admins to add, edit, and delete products.

4. Cart management:
    - Implement add to cart, list products in cart, and remove products from cart features.

5. Order management:
    - Implement order management, including order placement, order cancellation in both admin and user sides, order history, and order status.

6. Payment methods:
    - Integrate payment methods like Razorpay and Paypal to facilitate secure and reliable payment processing.

7. User management:
    - Implement user management, allowing admins to list all users and manage user accounts, such as blocking or unblocking accounts.

8. User profile and settings:
    - Implement a user profile page where users can view and edit their personal information.
    - Implement user settings, such as notifications, preferences, and privacy settings.

9. Additional features:
    - Implement wishlist, sort, filter, search, product return, profile page, edit profile, dashboard, and sales reports with downloads.
    - Implement image cropping and zooming to enhance user experience.
    - Implement resend OTP and pagination for enhanced user experience.
    - Implement an offer module for product offers, category offers, referral offers, and coupon codes.
    - Implement return option for delivered products.
    - Implement sales report for daily, weekly, and yearly performance tracking.
    - Implement an admin dashboard and chart/graph reports for improved data visualization.

10. Load balancing and project hosting:
    - Implement load balancing for improved system performance.
    - Host the project.



## site_user
| id  | email_address           | phone_number   | password| created_at |
|----|-------------------------|----------------|----------|------------|
| 1  | john.doe@example.com    | 123-456-7890   | abc123   |            |
| 2  | jane.doe1@example.com   | 456-789-0123   | xyz456   |            |
| 3  | johndoe123@example.com  | 555-555-5555   | qwerty   |            |
| 4  | janedoe456@example.com  | 222-222-2222   | asdfgh   |            |
| 5  | johndoe456@example.com  | 333-333-3333   | zxcvbn   |            |
| 6  | janedoe123@example.com  | 444-444-4444   | 123qwe   |            |
| 7  | testuser@example.com    | 555-123-4567   | 456asd   |            |
| 8  | user123@example.com     | 123-555-7890   | qwe123   |            |
| 9  | exampleuser@example.com | 999-999-9999   | test123  |            |
| 10 | testuser123@example.com | 111-111-1111   | test456  |            |

## user_address
| user_id | address_id | is_default |
|---------|------------|------------|
| 1       | 1          | 1          |
| 1       | 2          | 0          |
| 2       | 3          | 1          |
| 2       | 4          | 0          |
| 3       | 5          | 1          |
| 3       | 6          | 0          |
| 4       | 7          | 1          |
| 4       | 8          | 0          |
| 5       | 9          | 1          |
| 5       | 10         | 0          |

## user_info

| id  | userID | first_name | last_name | isBlocked | blockedAt           | blockedBy | reasonForBlocking |
|-----| ------ |------------|-----------|-----------|---------------------|-----------|-------------------|
| 1   | 1      | sample     | user1     | false     | null                | null      | null              |
| 2   | 2      | sample     | user2     | true      | 2022-02-28 10:30:00 | 1         | Abusive behavior  |
| 3   | 3      | sample     | user3     | false     | null                | null      | null              |
| 4   | 4      | sample     | user4     | false     | null                | null      | null              |
| 5   | 5      | sample     | user5     | true      | 2022-02-28 12:30:00 | 1         | Spamming          |


## address
| id  | house_number | street           | city            | district       | pincode   | landmark      |
|-----|--------------|------------------|-----------------|----------------|-----------|---------------|
| 1   | 123          | Main St          | New York        | Manhattan      | 10001     | Empire State  |
| 2   | 456          | Broadway         | New York        | Manhattan      | 10002     | Times Square  |
| 3   | 789          | Wall St          | New York        | Manhattan      | 10003     | Charging Bull |
| 4   | 111          | Market St        | San Francisco   | California     | 94103     | Golden Gate   |
| 5   | 222          | Montgomery St    | San Francisco   | California     | 94104     | Coit Tower    |
| 6   | 333          | Castro St        | San Francisco   | California     | 94114     | Golden Gate   |
| 7   | 444          | Beach St         | San Francisco   | California     | 94133     | Fisherman's   |
| 8   | 555          | Valencia St      | San Francisco   | California     | 94110     | Golden Gate   |
| 9   | 666          | Lombard St       | San Francisco   | California     | 94123     | Golden Gate   |
| 10  | 777          | Hayes St         | San Francisco   | California     | 94102     | Golden Gate   |

## user_payment_method
| id  | user_id | payment_type_id | provider   | account_number    | expiry_date | is_default |
|-----|---------|-----------------|------------|-------------------|-------------|------------|
| 1   | 1       | 1               | Visa       | 1234567890123456  | 12/25       | t          |
| 2   | 1       | 2               | Mastercard | 2345678901234567  | 01/27       | f          |
| 3   | 2       | 1               | Visa       | 3456789012345678  | 11/23       | t          |
| 4   | 2       | 3               | PayPal     | example@gmail.com | 12/25       | f          |
| 5   | 3       | 2               | Mastercard | 4567890123456789  | 09/24       | t          |
| 6   | 3       | 4               | Google Pay | example@gmail.com | 06/26       | f          |
| 7   | 4       | 1               | Visa       | 5678901234567890  | 05/22       | t          |
| 8   | 4       | 3               | PayPal     | example@gmail.com | 07/24       | f          |
| 9   | 5       | 1               | Visa       | 6789012345678901  | 03/23       | t          |
| 10  | 5       | 2               | Mastercard | 7890123456789012  | 08/25       | f          |


## payment_type
| id | value       |
|----|-------------|
| 1  | Credit Card |
| 2  | Debit Card  |
| 3  | PayPal      |
| 4  | Google Pay  |
| 5  | Apple Pay   |
| 6  | Bitcoin     |
| 7  | Ethereum    |
| 8  | Litecoin    |
| 9  | Dogecoin    |
| 10 | Dash        |


## Table: product_item

| id  | product_id | Sku                   | Qty_in_stock | Product_image        | Price |
|-----|------------|-----------------------| ------------ |----------------------| ----- |
| 1   | 1          | HP-Pavilion-15-001    | 10           | [product image URL]  | 800   |
| 2   | 1          | HP-Pavilion-15-002    | 15           | [product image URL]  | 850   |
| 3   | 2          | Dell-Inspiron-14-001  | 20           | [product image URL]  | 700   |
| 4   | 2          | Dell-Inspiron-14-002  | 25           | [product image URL]  | 750   |
| 5   | 3          | Lenovo-IdeaPad-3-001  | 30           | [product image URL]  | 600   |
| 6   | 3          | Lenovo-IdeaPad-3-002  | 35           | [product image URL]  | 650   |
| 7   | 4          | Apple-MacBook-Air-001 | 40           | [product image URL]  | 1000  |
| 8   | 4          | Apple-MacBook-Air-002 | 45           | [product image URL]  | 1100  |
| 9   | 5          | Asus-VivoBook-15-001  | 50           | [product image URL]  | 750   |
| 10  | 5          | Asus-VivoBook-15-002  | 55           | [product image URL]  | 800   |

## Table: product

| Id  | category_id | name                | Description                                                                                                       | Product_image        |
|-----|-------------| ------------------- |-------------------------------------------------------------------------------------------------------------------|----------------------|
| 1   | 1           | HP Pavilion 15      | This laptop is ideal for work and entertainment. It comes with a 10th Gen Intel Core i5 processor and 8GB of RAM. | [product image URL]  |
| 2   | 1           | Dell Inspiron 14    | This laptop is perfect for everyday use. It has a 10th Gen Intel Core i3 processor and 4GB of RAM.                | [product image URL]  |
| 3   | 1           | Lenovo IdeaPad 3    | This laptop is great for both work and play. It comes with a 10th Gen Intel Core i7 processor and 16GB of RAM.    | [product image URL]  |
| 4   | 2           | Apple MacBook Air   | This laptop is designed for Apple enthusiasts. It comes with an M1 chip and 8GB of RAM.                           | [product image URL]  |
| 5   | 3           | Asus VivoBook 15    | This laptop is ideal for students and professionals. It comes with an AMD Ryzen 5 processor and 8GB of RAM.       | [product image URL]  |

## Table: product_category

| Id  | Parent_category_id | Category_name |
|-----|--------------------| ------------- |
| 1   | NULL               | Laptops       |
| 2   | 1                  | Apple         |
| 3   | 1                  | Asus          |
| 4   | 1                  | Dell          |
| 5   | 1                  | HP            |
| 6   | 1                  | Lenovo        |


## Table: variation

| Id  | Category_id | name         |
|-----|-------------| ------------ |
| 1   | 1           | Color        |
| 2   | 1           | RAM          |
| 3   | 1           | Storage      |

## Table: variation_option

| Id  | Variation_id | Value  |
|-----|--------------|--------|
| 2   | 1            | Silver |
| 3   | 2            | 4GB    |
| 4   | 2            | 8GB    |
| 5   | 2            | 16GB   |
| 6   | 3            | 256GB  |
| 7   | 3            | 512GB  |
| 8   | 3            | 1TB    |

## Table: product_configuration

| Product_item_id | Variation_option_id |
| --------------- |---------------------|
| 1               | 1                   |
| 1               | 4                   |
| 1               | 6                   |
| 2               | 2                   |
| 2               | 4                   |
| 2               | 7                   |
| 3               | 1                   |
| 3               | 5                   |
| 3               | 8                   |
| 4               | 2                   |
| 4               | 6                   |
| 4               | 7                   |

## Table: shopping_cart

| Id  | User_id | total |
|-----| ------- | ----- |
| 1   | 1       | 1250  |
| 2   | 2       | 2560  |
| 3   | 3       | 1870  |
| 4   | 4       | 530   |
| 5   | 5       | 2120  |
| 6   | 6       | 990   |
| 7   | 7       | 1670  |
| 8   | 8       | 410   |
| 9   | 9       | 2950  |
| 10  | 10      | 1820  |

## Table: shopping_cart_item

| Id  | Cart_id | Product_item_id | quantity |
|-----| ------- | --------------- | -------- |
| 1   | 1       | 1               | 2        |
| 2   | 1       | 3               | 1        |
| 3   | 1       | 5               | 1        |
| 4   | 2       | 2               | 1        |
| 5   | 2       | 4               | 3        |
| 6   | 3       | 6               | 2        |
| 7   | 4       | 1               | 1        |
| 8   | 4       | 7               | 2        |
| 9   | 4       | 8               | 1        |
| 10  | 5       | 5               | 4        |
| 11  | 5       | 9               | 1        |
| 12  | 6       | 2               | 1        |
| 13  | 7       | 4               | 2        |
| 14  | 7       | 7               | 1        |
| 15  | 7       | 8               | 2        |


## Table: Shop_order

| Id  | User_id | Order_date       | payment_method_id | Shipping_address | Shipping_method | order_total | order_status |
|-----| ------- | ---------------- | ----------------- | ---------------- | --------------- | ----------- | ------------ |
| 1   | 1       | 2022-02-28 10:30 | 1                 | 1                | 1               | 1120.00     | 1            |
| 2   | 2       | 2022-02-28 11:30 | 2                 | 2                | 2               | 2500.00     | 2            |
| 3   | 3       | 2022-02-28 12:30 | 3                 | 3                | 1               | 1850.00     | 3            |
| 4   | 4       | 2022-02-28 13:30 | 1                 | 4                | 2               | 500.00      | 1            |
| 5   | 5       | 2022-02-28 14:30 | 2                 | 5                | 1               | 2100.00     | 2            |
| 6   | 6       | 2022-02-28 15:30 | 3                 | 6                | 2               | 990.00      | 3            |
| 7   | 7       | 2022-02-28 16:30 | 1                 | 7                | 1               | 1670.00     | 1            |
| 8   | 8       | 2022-02-28 17:30 | 2                 | 8                | 2               | 410.00      | 2            |
| 9   | 9       | 2022-02-28 18:30 | 3                 | 9                | 1               | 2950.00     | 3            |
| 10  | 10      | 2022-02-28 19:30 | 1                 | 10               | 2               | 1820.00     | 1            |

## Table: Shipping_method

| Id  | Name                | price  |
|-----|---------------------|--------|
| 1   | Standard Shipping   | 50.00  |
| 2   | Express Shipping    | 100.00 |
| 3   | Same Day Delivery   | 200.00 |

## Table: Order_status

| Id  | Status        |
|-----| ------------- |
| 1   | Pending       |
| 2   | Processing    |
| 3   | Shipped       |
| 4   | Delivered     |
| 5   | Cancelled     |

## Table: order_line

| Id | Product_item_id | Order_id  | Quantity | price   |
|----|-----------------|-----------|----------|---------|
| 1  | 1               | 1         | 2        | 450.00  |
| 2  | 3               | 1         | 1        | 370.00  |
| 3  | 5               | 1         | 1        | 300.00  |
| 4  | 2               | 2         | 1        | 850.00  |
| 5  | 4               | 2         | 	3	      | 350.00  |
| 6	 | 6	              | 3	        | 2	       | 495.00  |
| 7	 | 7	              | 3	        | 1	       | 775.00  |
| 8	 | 10              | 	3        | 	3	      | 600.00  |
| 9	 | 8	              | 4	        | 1	       | 410.00  |
| 10	 | 9	              | 4	        | 2	       | 1475.00 |

| id | username   | password | email                  | isSuperAdmin | isBlocked | created_at          | updated_at          |
|----|------------|----------|------------------------|--------------|-----------|---------------------|---------------------|
| 1  | superadmin | *******  | superadmin@example.com | true         | false     | 2022-01-01 12:00:00 | 2022-01-01 12:00:00 |
| 2  | admin1     | *******  | admin1@example.com     | false        | false     | 2022-01-02 12:00:00 | 2022-01-02 12:00:00 |
| 3  | admin2     | *******  | admin2@example.com     | false        | false     | 2022-01-03 12:00:00 | 2022-01-03 12:00:00 |
| 4  | admin3     | *******  | admin3@example.com     | false        | true      | 2022-01-04 12:00:00 | 2022-01-04 12:00:00 |
| 5  | admin4     | *******  | admin4@example.com     | false        | false     | 2022-01-05 12:00:00 | 2022-01-05 12:00:00 |

