package domain

// ManageInventoryMethod represents a method of inventory management
type ManageInventoryMethod int

const (
	// DontManageStock represents not tracking inventory for the product
	DontManageStock ManageInventoryMethod = 0

	// ManageStock represents tracking inventory for the product
	ManageStock ManageInventoryMethod = 1

	// ManageStockByAttributes represents tracking inventory for the product by product attributes
	ManageStockByAttributes ManageInventoryMethod = 2
)
