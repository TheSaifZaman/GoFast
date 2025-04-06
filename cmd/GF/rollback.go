package cmd

import (
	"errors"
	"fmt"
	"github.com/TheSaifZaman/GoFast/config"
	"github.com/TheSaifZaman/GoFast/internal/database"
	"github.com/spf13/cobra"
	"log"

	"github.com/TheSaifZaman/GoFast/internal/migration"
	"gorm.io/gorm"
	"strings"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback [type]",
	Short: "Rollback the last applied migration",
	Run: func(cmd *cobra.Command, args []string) {
		migrationType := "landlord"
		if len(args) > 0 {
			migrationType = args[0]
		}

		dbConfig, err := migration.LoadConfig()
		if err != nil {
			log.Fatal(err)
		}

		db, err := database.NewGormDB(
			dbConfig.DBHost,
			dbConfig.DBUser,
			dbConfig.DBPassword,
			dbConfig.DBName,
			dbConfig.DBPort,
		)
		if err != nil {
			log.Fatal(err)
		}

		latest, err := migration.GetLatestMigration(db, migrationType)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Printf("No migrations to rollback for type '%s'\n", migrationType)
				return
			}
			log.Printf("Error getting latest migration: %v", err)
			return
		}

		var m config.MigrationListStruct
		for i := range config.LandlordMigrationList {
			if strings.EqualFold(config.LandlordMigrationList[i].Name, latest.Name) {
				m = config.LandlordMigrationList[i]
				break
			}
		}
		if m.Name == "" {
			log.Printf("No known Down() function for migration: %s", latest.Name)
			return
		}
		if err := db.Migrator().DropTable(m.Structure); err != nil {
			log.Printf("Error running Down() for %s: %v", latest.Name, err)
			return
		}

		if err := migration.RemoveMigration(db, latest.Name, migrationType); err != nil {
			log.Printf("Error removing migration record %s: %v", latest.Name, err)
			return
		}

		fmt.Printf("Migration %s rolled back successfully\n", latest.Name)
	},
}

func init() {
	rootCmd.AddCommand(rollbackCmd)
}
