package cmd

import (
	"github.com/TheSaifZaman/GoFast/internal/router"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Go-Fast server",
	Run: func(cmd *cobra.Command, args []string) {
		mux := router.NewRouter() // our custom router from the router package

		log.Printf("Starting server on :9090")
		if err := http.ListenAndServe(":9090", mux); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
