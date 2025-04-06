package cmd

import (
	"fmt"
	"github.com/TheSaifZaman/GoFast/config"
	"github.com/TheSaifZaman/GoFast/internal/database"
	"github.com/spf13/cobra"
	"log"

	"github.com/TheSaifZaman/GoFast/internal/migration"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate [type]",
	Short: "Run database migrations",
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

		if err := migration.CreateMigrationsTable(db); err != nil {
			log.Fatal(err)
		}

		var anyApplied bool
		for _, m := range config.LandlordMigrationList {
			alreadyApplied, err := migration.IsMigrationApplied(db, m.Name, migrationType)
			if err != nil {
				log.Printf("Error checking migration %s: %v", m.Name, err)
				continue
			}
			if alreadyApplied {
				fmt.Printf("Migration %s already applied, skipping\n", m.Name)
				continue
			}

			if err := db.AutoMigrate(m.Structure); err != nil {
				log.Printf("Error running Up() for %s: %v", m.Name, err)
				continue
			}

			if err := migration.RecordMigration(db, m.Name, migrationType); err != nil {
				log.Printf("Error recording migration %s: %v", m.Name, err)
				continue
			}

			fmt.Printf("Migration %s applied successfully\n", m.Name)
			anyApplied = true
		}

		if !anyApplied {
			fmt.Println("No new migrations to apply.")
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
