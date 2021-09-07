package user

import (
	"phillz-life-dashboard/db"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	ID        uuid.UUID             `json:"id" gorm:"type:uuid;primary_key"`
	FirstName string                `json:"first_name"`
	LastName  string                `json:"last_name"`
	Username  string                `gorm:"unique_index;not null" json:"username"`
	Email     string                `gorm:"unique_index;not null" json:"email"`
	Password  string                `gorm:"not null" json:"password"`
	CreatedAt time.Time             `json:"created_at,omitempty"`
	UpdatedAt time.Time             `json:"updated_at,omitempty"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at,omitempty"`
}

func init() {
	db := db.Conn
	db.AutoMigrate(&User{})
}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validateToken(t *jwt.Token, id string) bool {
	claims := t.Claims.(jwt.MapClaims)
	uid := claims["user_id"].(string)

	return uid == id
}

// func validateUser(id string, p string) bool {
// 	db := db.Conn
// 	var user User
// 	db.First(&user, id)
// 	if user.Username == "" {
// 		return false
// 	}
// 	if CheckPasswordHash(p, user.Password) {
// 		return false
// 	}
// 	return true
// }

// GetUser get a user
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := db.Conn
	var user User
	db.Find(&user, "id = ?", id)
	if user.Username == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": user})
}

// CreateUser new user
func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	db := db.Conn
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})

	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	newUser := NewUser{
		Email:    user.Email,
		Username: user.Username,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

// UpdateUser update user
func UpdateUser(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{"status": "success", "message": "User successfully updated", "data": ""})
}

// DeleteUser delete user
func DeleteUser(c *fiber.Ctx) error {
	db := db.Conn
	var user User

	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validateToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})

	}

	db.First(&user, id)

	db.Delete(&user)
	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}
