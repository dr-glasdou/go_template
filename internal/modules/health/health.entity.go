package health

// Entity represents the health entity.
type Entity struct {
	ID      string
	Healthy bool
}

// IsHealthy returns true if the entity is healthy.
func (e *Entity) IsHealthy() bool {
	return e.Healthy
}
