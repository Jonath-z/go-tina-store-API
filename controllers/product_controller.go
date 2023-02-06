package controllers

import (
	"context"
	"fiber-tina-store-API/configs"
	"fiber-tina-store-API/models"
	"fiber-tina-store-API/responses"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)


var productCollection *mongo.Collection =  configs.GetCollection(configs.DB,"products")

func CreateProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var product models.Product
	defer cancel()

	err := c.BodyParser(&product)
	if err != nil{
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	validateErr := validate.Struct(&product)
	if validateErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error:missing required properties", Data: &fiber.Map{"data": err.Error()}})
	}

	newProducts := models.Product{
		ImageUrl: product.ImageUrl,
		ProductName: product.ProductName,
		ProductPrice: product.ProductPrice,
		ProductType: product.ProductType,
		ProductId: uuid.NewString(),
	}

	result, err := productCollection.InsertOne(ctx, newProducts)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.Response{Status: http.StatusCreated, Message: "error", Data: &fiber.Map{"data": result}})
}