// Package app provides the application entry point and routing setup.
package app

import (
	"github.com/gin-gonic/gin"

	"glasdou.wtf/gotemplate/internal/modules/health"
)

// New returns a new gin engine with the application routing setup.
func New() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")

	// register modules
	health.NewModule().Register(v1)

	return r
}
