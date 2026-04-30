// Package health provides the health module.
package health

import (
	"github.com/gin-gonic/gin"
)

// Module represents the health module.
type Module struct {
	handler *Handler
}

// NewModule creates a new health module.
func NewModule() *Module {
	repo := NewRepository()
	svc := NewService(repo)
	h := NewHandler(svc)

	return &Module{handler: h}
}

// Register registers the health module routes.
func (m *Module) Register(r *gin.RouterGroup) {
	m.handler.Register(r)
}
