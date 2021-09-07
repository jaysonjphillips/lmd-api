package storage

import (
	"phillz-life-dashboard/db"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

type MovingBox struct {
	ID                 uuid.UUID             `json:"id" gorm:""`
	BoxNumber          int                   `json:"box_number"`
	BoxSize            string                `json:"box_size"`
	BoxWeight          string                `json:"box_weight"`
	BoxRoom            string                `json:"box_room"`
	BoxContents        string                `json:"box_contents"`
	BoxQRCode          string                `json:"box_qr_code,omitempty"`
	BoxPacked          bool                  `json:"box_packed"`
	BoxReceived        bool                  `json:"box_received"`
	BoxDamaged         bool                  `json:"box_damaged"`
	BoxHasDamagedItems bool                  `json:"box_damaged_items"`
	CreatedAt          time.Time             `json:"created_at,omitempty"`
	UpdatedAt          time.Time             `json:"updated_at,omitempty"`
	DeletedAt          soft_delete.DeletedAt `json:"deleted_at,omitempty"`
}

func CreateBox(c *fiber.Ctx) error {
	db := db.Conn
	box := new(MovingBox)
	if err := c.BodyParser(box); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	db.Create(&box)
	return c.Status(fiber.StatusAccepted).JSON(box)
}
