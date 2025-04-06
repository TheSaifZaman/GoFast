package migration

import (
	"errors"
	"gorm.io/gorm"
)

type Migrations struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"uniqueIndex;not null"`
	Batch         int
	MigrationType string
}

func CreateMigrationsTable(db *gorm.DB) error {
	return db.AutoMigrate(&Migrations{})
}

func IsMigrationApplied(db *gorm.DB, name, migrationType string) (bool, error) {
	var m Migrations
	err := db.Where("name = ? AND migration_type = ?", name, migrationType).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return err == nil, err
}

func RecordMigration(db *gorm.DB, name, migrationType string) error {
	return db.Create(&Migrations{
		Name:          name,
		MigrationType: migrationType,
	}).Error
}

func RemoveMigration(db *gorm.DB, name, migrationType string) error {
	return db.Where("name = ? AND migration_type = ?", name, migrationType).
		Delete(&Migrations{}).Error
}

func GetLatestMigration(db *gorm.DB, migrationType string) (*Migrations, error) {
	var m Migrations
	err := db.Where("migration_type = ?", migrationType).
		Order("id DESC").
		First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}
