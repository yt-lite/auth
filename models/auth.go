package models

// import uuid
import (
	"github.com/google/uuid"
)

type Auth struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email string    `gorm:"type:varchar(100);unique_index;column:email"`
}
