package health

import "time"

// Response represents the health response.
type Response struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}
