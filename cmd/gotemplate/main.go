// main is the entry point for the application.
package main

import (
	"fmt"
	"log"

	"glasdou.wtf/gotemplate/configs"
	"glasdou.wtf/gotemplate/internal/app"
)

func main() {
	config, err := configs.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	app := app.New()

	addr := fmt.Sprintf(":%d", config.Port)
	log.Printf("starting server on %s", addr)

	if err := app.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
