package routes

import (
	"kasir/database"
	"kasir/models"
	"kasir/utils"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Login(c *fiber.Ctx) error {

	loginRequest := new(LoginRequest)
	log.Println(loginRequest)
	if err := c.BodyParser(&loginRequest); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	//Check availabe user
	var user models.User
	err := database.DB.First(&user, "email= ?", loginRequest.Email).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not Found",
		})
	}

	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["phone"] = user.Phone

	//time token
	// claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, errGenerateToken := utils.GenerateToken(&claims)

	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}
	// Check availabe password
	return c.JSON(fiber.Map{
		"token": token,
	})
}
