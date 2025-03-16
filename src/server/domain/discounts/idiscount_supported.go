package domain

// IDiscountSupported represents an entity which supports discounts
type IDiscountSupported[T any] interface {
	GetID() int
	SetID(id int)
}
