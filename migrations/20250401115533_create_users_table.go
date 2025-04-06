package migrations

import (
	"time"
)

// User represents the users table
type User struct {
	CustomIdColumn
	FirstName       string `gorm:"not null"`
	LastName        string `gorm:"null"`
	Name            string `gorm:"not null"`
	Email           string `gorm:"unique;null"`
	EmailVerifiedAt time.Time
	CustomActorAndActingTimeColumns
}
