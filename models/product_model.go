package models

type Product struct {
	ImageUrl string `JSON:"imageUrl, omitempty" validate:"required"`
	ProductName string `JSON:"productName, omitempty" validate:"required"`
	ProductPrice string `JSON:"productPrice, omitempty" validate:"required"`
	ProductType string `JSON:"productType, omitempty" validate:"required"`
	ProductId string `JSON:"productId, omitempty" validate:"required"`
}