package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler handles health check requests.
type Handler struct {
	svc *Service
}

// NewHandler returns a new health handler.
func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

// Health handles the health check request.
func (h *Handler) Health(ctx *gin.Context) {
	resp, err := h.svc.Check()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// Register registers the health handler with the given router.
func (h *Handler) Register(r *gin.RouterGroup) {
	r.GET("/health", h.Health)
}
