package wine

import (
	"fmt"
	"phillz-life-dashboard/db"
	"reflect"
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
	Varietal  string                `json:"varietal"`
	Blend     bool                  `json:"is_blend"`
	Reserve   bool                  `json:"is_reserve"`
	Stored    bool                  `json:"is_stored" gorm:"default:true"`
	Packed    bool                  `json:"is_packed" gorm:"default:false"`
	Received  bool                  `json:"is_received" gorm:"default:false"`
	Consumed  bool                  `json:"is_consumed" gorm:"default:false"`
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
	return c.JSON(fiber.Map{
		"results": len(wine),
		"data":    wine,
	})
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
	type UpdateBody struct {
		Name      string `json:"name"`
		Vineyard  string `json:"vineyard"`
		Location  string `json:"location"`
		Year      int    `json:"year"`
		Varietal  string `json:"varietal"`
		Blend     bool   `json:"is_blend"`
		Reserve   bool   `json:"is_reserve"`
		Stored    bool   `json:"is_stored" gorm:"default:true"`
		Packed    bool   `json:"is_packed" gorm:"default:false"`
		Received  bool   `json:"is_received" gorm:"default:false"`
		Consumed  bool   `json:"is_consumed" gorm:"default:false"`
		BoxNumber int    `json:"box_number,omitempty"`
	}

	db := db.Conn
	var data UpdateBody
	var wine Wine

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	db.First(&wine, "id = ?", c.Params("id"))

	v := reflect.ValueOf(wine)
	dv := reflect.ValueOf(data)

	for i := 0; i < v.NumField(); i++ {
		fmt.Println(dv.Elem().FieldByName(v.Type().Field(i).Name))
		fmt.Println("\t", dv.Field(i))
	}

	return c.SendString("Update a Wine")
}

func DeleteWine(c *fiber.Ctx) error {
	db := db.Conn
	var wine Wine

	db.Where("id = ?", c.Params("id")).Delete(&wine)
	return c.Status(fiber.StatusOK).SendString("")
}
