package wine

import (
	"phillz-life-dashboard/db"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Wine struct {
	ID        uuid.UUID             `json:"id" gorm:"type:uuid;primary_key"`
	UserID    uuid.UUID             `json:"user_id" gorm:"type:uuid;foreign_key"`
	Name      string                `json:"name"`
	Vineyard  string                `json:"vineyard"`
	Location  string                `json:"location"`
	Year      int                   `json:"year"`
	Reserve   bool                  `json:"reserve"`
	Stored    bool                  `json:"stored" gorm:"default:true"`
	Packed    bool                  `json:"packed" gorm:"default:false"`
	Received  bool                  `json:"received" gorm:"default:false"`
	Consumed  bool                  `json:"consumed" gorm:"default:false"`
	BoxNumber int                   `json:"box_number,omitempty"`
	CreatedAt time.Time             `json:"created_at,omitempty"`
	UpdatedAt time.Time             `json:"updated_at,omitempty"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at,omitempty"`
}

func init() {
	db := db.Conn
	db.AutoMigrate(&Wine{})
}

func (wine *Wine) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	wine.ID = uuid
	return
}

func GetWines(c *fiber.Ctx) error {
	db := db.Conn
	var wine []Wine
	db.Find(&wine)
	return c.JSON(wine)
}

func GetWine(c *fiber.Ctx) error {
	db := db.Conn
	var wine Wine

	db.First(&wine, "id = ?", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(wine)
}

func AddWine(c *fiber.Ctx) error {
	db := db.Conn
	wine := new(Wine)
	if err := c.BodyParser(wine); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	db.Create(&wine)
	return c.Status(fiber.StatusAccepted).JSON(wine)
}

func UpdateWine(c *fiber.Ctx) error {
	return c.SendString("Update a Wine")
}

func DeleteWine(c *fiber.Ctx) error {
	db := db.Conn
	var wine Wine

	db.Where("id = ?", c.Params("id")).Delete(&wine)
	return c.Status(fiber.StatusOK).SendString("")
}
