package model

type NewCategory struct {
	CategoryName string `json:"category_name"`
}

type CategoryID struct {
	CategoryID int `json:"id"`
}

type ProductID struct {
	ID int `json:"id"`
}

type ProductItemID struct {
	ID int `json:"id"`
}

type QueryParams struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Query    string `json:"query"`
	Filter   string `json:"filter"`
	SortBy   string `json:"sort_by"`
	SortDesc bool   `json:"sort_desc"`
}
