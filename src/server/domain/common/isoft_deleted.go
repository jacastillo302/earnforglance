package domain

// ISoftDeletedEntity represents a soft-deleted (without actually deleting from storage) entity
type ISoftDeletedEntity interface {
	// GetDeleted returns whether the entity has been deleted
	GetDeleted() bool

	// SetDeleted sets whether the entity has been deleted
	SetDeleted(deleted bool)
}
