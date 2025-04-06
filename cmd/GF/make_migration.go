package cmd

import (
	"fmt"
	"github.com/TheSaifZaman/GoFast/internal/helpers"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

var migrationTemplate = `package migrations

// {{.MigrationName}} is an example model using your custom columns.
// Adjust or rename fields as needed.
type {{.MigrationName}} struct {
	CustomIdColumn
	// Example columns
	Title    string ` + "`gorm:\"not null\"`" + `
	Subtitle string
	// Insert more fields here...

	CustomActorAndActingTimeColumns
}
`

var makeMigrationCmd = &cobra.Command{
	Use:   "make:migration [name]",
	Short: "Create a new Go-based migration file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		migrationName := args[0]

		// Convert user input (e.g. "user_profile") to PascalCase struct name (e.g. "UserProfile").
		structName := helpers.ToPlural(helpers.ToSnakeCase(migrationName))

		// Create a timestamp prefix like 20230401103045
		timestamp := time.Now().Format("20060102150405")

		// Build file name: "20230401103045_user_profile.go"
		fileName := fmt.Sprintf("%s_create_%s_table.go", timestamp, strings.ToLower(structName))
		path := filepath.Join("migrations", fileName)

		// Ensure "migrations" directory
		if err := os.MkdirAll("migrations", 0755); err != nil {
			fmt.Printf("Error creating migrations directory: %v\n", err)
			return
		}

		// Prepare data for the template
		data := struct {
			MigrationName string
		}{
			MigrationName: migrationName,
		}

		// Parse and execute the template
		tpl, err := template.New("migration").Parse(migrationTemplate)
		if err != nil {
			fmt.Printf("Error parsing template: %v\n", err)
			return
		}

		outFile, err := os.Create(path)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", path, err)
			return
		}
		defer func(outFile *os.File) {
			err := outFile.Close()
			if err != nil {
				fmt.Printf("Error closing file %s: %v\n", path, err)
			}
		}(outFile)

		if err := tpl.Execute(outFile, data); err != nil {
			fmt.Printf("Error executing template for %s: %v\n", path, err)
			return
		}

		fmt.Printf("Migration file created: %s\n", path)
	},
}

func init() {
	rootCmd.AddCommand(makeMigrationCmd)
}
