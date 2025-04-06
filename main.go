package main

import (
	"github.com/TheSaifZaman/GoFast/cmd/GF"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
