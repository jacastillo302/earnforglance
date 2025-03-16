package domain

const (
	CollectionPictureHash = "picture_hashes"
)

// PictureHashItem is a helper class for making picture hashes from DB
type PictureHashItem struct {
	PictureID int    `bson:"picture_id"`
	Hash      []byte `bson:"hash"`
}

// CompareTo compares this instance to a specified object and returns an indication
func (p *PictureHashItem) CompareTo(other *PictureHashItem) int {
	if other == nil {
		return -1
	}
	if p.PictureID < other.PictureID {
		return -1
	}
	if p.PictureID > other.PictureID {
		return 1
	}
	return 0
}
