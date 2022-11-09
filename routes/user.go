package routes

import (
	"errors"
	"kasir/database"
	"kasir/models"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID       int    `json:"id" gorm:"primariKey"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address"`
	Phone    string `json:"phone" validate:"required"`
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID:       int(userModel.ID),
		Name:     userModel.Name,
		Email:    userModel.Email,
		Address:  userModel.Address,
		Phone:    userModel.Phone,
		Password: userModel.Password,
	}
}

// Create User
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

// GetAllUser
func GetAllUser(c *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	database.DB.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

// UpdateUserByIdD
func UpdateUserByID(c *fiber.Ctx) error {
	var user models.User
	userId := c.Params("id")
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if database.DB.Where("id = ?", userId).Updates(&user).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "tidak dapat mengupdate data",
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil mengupdate data",
	})
}

// Delete User
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findUser(id, &user)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.DB.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted User")
}
