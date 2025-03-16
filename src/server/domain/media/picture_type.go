package domain

// PictureType represents a picture item type
type PictureType int

const (
	// Entity represents entities (products, categories, manufacturers)
	Entity PictureType = 1

	// Avatar represents avatar
	Avatar PictureType = 10
)
