// main is the entry point for the application.
package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"glasdou.wtf/gotemplate/configs"
	"glasdou.wtf/gotemplate/internal/modules/health"
)

func main() {
	config, err := configs.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	r := gin.Default()
	v1 := r.Group("api/v1")

	healthModule := health.NewModule()
	healthModule.Register(v1)

	addr := fmt.Sprintf(":%d", config.Port)
	log.Printf("starting server on %s", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
