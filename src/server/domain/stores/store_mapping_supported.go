package domain

// IStoreMappingSupported represents an entity which supports store mapping
type IStoreMappingSupported interface {
	GetLimitedToStores() bool
	SetLimitedToStores(limitedToStores bool)
}
