package domain

// ISlugSupported represents an entity which supports slug (SEO friendly one-word URLs)
type ISlugSupported interface {
	GetSlug() string
	SetSlug(slug string)
}
