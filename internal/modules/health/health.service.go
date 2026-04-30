package health

import (
	"time"
)

// Service represents the health service.
type Service struct {
	repo Repository
}

// NewService creates a new health service.
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Check returns the health status of the service.
func (s *Service) Check() (*Response, error) {
	status := "ok"
	if err := s.repo.Ping(); err != nil {
		status = "unhealthy"
	}

	return &Response{
		Status:    status,
		Timestamp: time.Now(),
	}, nil
}
