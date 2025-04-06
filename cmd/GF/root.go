package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-fast",
	Short: "A Laravel-like framework for Go",
	Long: `Go-Fast is a Laravel-like framework for Go that supports:
- Single-tenant single DB
- Multi-tenant single DB
- Multi-tenant multi DB
- Custom SSO (Single Sign-On)
- Modular code organization
- Migrations & seeding
- Command-line scaffolding`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Go-Fast!")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
