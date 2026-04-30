package health

// Repository defines the interface for the health repository.
type Repository interface {
	Ping() error
}

// InMemRepository is an in-memory implementation of the health repository.
type InMemRepository struct{}

// NewRepository returns a new in-memory health repository.
func NewRepository() *InMemRepository {
	return &InMemRepository{}
}

// Ping implements the Repository interface for InMemRepository.
func (r *InMemRepository) Ping() error {
	return nil
}

// NewRepositoryFromEnv returns a new health repository based on the environment.
func NewRepositoryFromEnv() Repository {
	return NewRepository()
}
