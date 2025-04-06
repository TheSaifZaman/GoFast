package migrations

import (
	"gorm.io/gorm"
	"time"
)

type UserIdColumn struct {
	UserId string `gorm:"column:user_id;size:26;index;not null"`
	User   User   `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CustomIdColumn struct {
	ID string `gorm:"size:26;primarykey"`
}

type CustomActorAndActingTimeColumns struct {
	CreatedAt time.Time      `gorm:"null;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"null;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"null"`
	CreatedBy string         `gorm:"size:26;null"`
	UpdatedBy string         `gorm:"size:26;null"`
	DeletedBy string         `gorm:"size:26;null"`
}
