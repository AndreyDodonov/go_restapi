package store

// store interface
type Store interface {
	User() UserRepository
}