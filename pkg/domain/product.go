package domain

type ProductCategory struct {
	ID           uint   `gorm:"primaryKey,uniqueIndex" json:"id"`
	CategoryName string `gorm:"not null,index,unique" json:"category_name"`
}

type ProductBrand struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Brand string `gorm:"not null,index,unique" json:"brand" validate:"required"`
}

type Product struct {
	ID                uint            `gorm:"primaryKey" json:"id"`
	ProductCategoryID uint            `gorm:"not null" json:"product_category_id" validate:"required"`
	ProductCategory   ProductCategory `gorm:"foreignKey:ProductCategoryID" json:"-"`
	Name              string          `gorm:"not null,uniqueIndex" json:"name" validate:"required"`
	BrandID           uint            `gorm:"not null" json:"brand_id" validate:"required"`
	ProductBrand      ProductBrand    `gorm:"foreignKey:BrandID" json:"-"`
	Description       string          `json:"description"`
	ProductImage      string          `json:"product_image"`
}

type ProductItem struct {
	ID               uint    `gorm:"primaryKey" json:"id"`
	ProductID        uint    `gorm:"not null" json:"product_id" validate:"required"`
	Product          Product `gorm:"foreignKey:ProductID" json:"-"`
	Model            string  `gorm:"not null" json:"model" validate:"required"`
	Processor        string  `gorm:"not null" json:"processor" validate:"required"`
	Ram              string  `gorm:"not null" json:"ram" validate:"required"`
	Storage          string  `gorm:"not null" json:"storage" validate:"required"`
	DisplaySize      string  `gorm:"not null" json:"display_size" validate:"required"`
	GraphicsCard     string  `json:"graphics_card"`
	OS               string  `gorm:"not null" json:"os" validate:"required"`
	SKU              string  `gorm:"not null" json:"sku" validate:"required"`
	QntyInStock      int     `gorm:"not null" json:"qnty_in_stock" validate:"required"`
	ProductItemImage string  `json:"product_item_image"`
	Price            float64 `gorm:"not null" json:"price" validate:"required"`
}
