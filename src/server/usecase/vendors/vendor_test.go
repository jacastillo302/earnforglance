package usecase

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/vendors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestVendorUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.VendorRepository)
	timeout := time.Duration(10)
	usecase := NewVendorUsecase(mockRepo, timeout)

	vendorID := primitive.NewObjectID().Hex()

	updatedVendor := domain.Vendor{
		ID:                             primitive.NewObjectID(), // Existing ID of the record to update
		Name:                           "Updated Tech Supplies",
		Email:                          "support@updatedtechsupplies.com",
		Description:                    "Updated supplier of tech products.",
		PictureID:                      primitive.NewObjectID(),
		AddressID:                      primitive.NewObjectID(),
		AdminComment:                   "Updated vendor details.",
		Active:                         false,
		Deleted:                        true,
		DisplayOrder:                   2,
		MetaKeywords:                   "updated, tech, electronics",
		MetaDescription:                "Updated Tech Supplies - Electronics and more.",
		MetaTitle:                      "Updated Tech Supplies",
		PageSize:                       20,
		AllowCustomersToSelectPageSize: false,
		PageSizeOptions:                "20,40,60",
		PriceRangeFiltering:            false,
		PriceFrom:                      200.00,
		PriceTo:                        2000.00,
		ManuallyPriceRange:             true,
		PmCustomerID:                   new(primitive.ObjectID),
	}

	mockRepo.On("FetchByID", mock.Anything, vendorID).Return(updatedVendor, nil)

	result, err := usecase.FetchByID(context.Background(), vendorID)

	assert.NoError(t, err)
	assert.Equal(t, updatedVendor, result)
	mockRepo.AssertExpectations(t)
}

func TestVendorUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.VendorRepository)
	timeout := time.Duration(10)
	usecase := NewVendorUsecase(mockRepo, timeout)

	newVendor := &domain.Vendor{
		Name:                           "Tech Supplies",
		Email:                          "contact@techsupplies.com",
		Description:                    "Supplier of tech products and accessories.",
		PictureID:                      primitive.NewObjectID(),
		AddressID:                      primitive.NewObjectID(),
		AdminComment:                   "Preferred vendor for electronics.",
		Active:                         true,
		Deleted:                        false,
		DisplayOrder:                   1,
		MetaKeywords:                   "tech, electronics, gadgets",
		MetaDescription:                "Tech Supplies - Your source for electronics.",
		MetaTitle:                      "Tech Supplies",
		PageSize:                       10,
		AllowCustomersToSelectPageSize: true,
		PageSizeOptions:                "10,20,50",
		PriceRangeFiltering:            true,
		PriceFrom:                      100.00,
		PriceTo:                        1000.00,
		ManuallyPriceRange:             false,
		PmCustomerID:                   nil,
	}

	mockRepo.On("Create", mock.Anything, newVendor).Return(nil)

	err := usecase.Create(context.Background(), newVendor)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVendorUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.VendorRepository)
	timeout := time.Duration(10)
	usecase := NewVendorUsecase(mockRepo, timeout)

	updatedVendor := &domain.Vendor{
		ID:                             primitive.NewObjectID(), // Existing ID of the record to update
		Name:                           "Updated Tech Supplies",
		Email:                          "support@updatedtechsupplies.com",
		Description:                    "Updated supplier of tech products.",
		PictureID:                      primitive.NewObjectID(),
		AddressID:                      primitive.NewObjectID(),
		AdminComment:                   "Updated vendor details.",
		Active:                         false,
		Deleted:                        true,
		DisplayOrder:                   2,
		MetaKeywords:                   "updated, tech, electronics",
		MetaDescription:                "Updated Tech Supplies - Electronics and more.",
		MetaTitle:                      "Updated Tech Supplies",
		PageSize:                       20,
		AllowCustomersToSelectPageSize: false,
		PageSizeOptions:                "20,40,60",
		PriceRangeFiltering:            false,
		PriceFrom:                      200.00,
		PriceTo:                        2000.00,
		ManuallyPriceRange:             true,
		PmCustomerID:                   new(primitive.ObjectID),
	}
	*updatedVendor.PmCustomerID = primitive.NewObjectID()

	mockRepo.On("Update", mock.Anything, updatedVendor).Return(nil)

	err := usecase.Update(context.Background(), updatedVendor)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVendorUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.VendorRepository)
	timeout := time.Duration(10)
	usecase := NewVendorUsecase(mockRepo, timeout)

	vendorID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, vendorID).Return(nil)

	err := usecase.Delete(context.Background(), vendorID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVendorUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.VendorRepository)
	timeout := time.Duration(10)
	usecase := NewVendorUsecase(mockRepo, timeout)

	fetchedVendors := []domain.Vendor{
		{
			ID:                             primitive.NewObjectID(),
			Name:                           "Tech Supplies",
			Email:                          "contact@techsupplies.com",
			Description:                    "Supplier of tech products and accessories.",
			PictureID:                      primitive.NewObjectID(),
			AddressID:                      primitive.NewObjectID(),
			AdminComment:                   "Preferred vendor for electronics.",
			Active:                         true,
			Deleted:                        false,
			DisplayOrder:                   1,
			MetaKeywords:                   "tech, electronics, gadgets",
			MetaDescription:                "Tech Supplies - Your source for electronics.",
			MetaTitle:                      "Tech Supplies",
			PageSize:                       10,
			AllowCustomersToSelectPageSize: true,
			PageSizeOptions:                "10,20,50",
			PriceRangeFiltering:            true,
			PriceFrom:                      100.00,
			PriceTo:                        1000.00,
			ManuallyPriceRange:             false,
			PmCustomerID:                   nil,
		},
		{
			ID:                             primitive.NewObjectID(),
			Name:                           "Updated Tech Supplies",
			Email:                          "support@updatedtechsupplies.com",
			Description:                    "Updated supplier of tech products.",
			PictureID:                      primitive.NewObjectID(),
			AddressID:                      primitive.NewObjectID(),
			AdminComment:                   "Updated vendor details.",
			Active:                         false,
			Deleted:                        true,
			DisplayOrder:                   2,
			MetaKeywords:                   "updated, tech, electronics",
			MetaDescription:                "Updated Tech Supplies - Electronics and more.",
			MetaTitle:                      "Updated Tech Supplies",
			PageSize:                       20,
			AllowCustomersToSelectPageSize: false,
			PageSizeOptions:                "20,40,60",
			PriceRangeFiltering:            false,
			PriceFrom:                      200.00,
			PriceTo:                        2000.00,
			ManuallyPriceRange:             true,
			PmCustomerID:                   new(primitive.ObjectID),
		},
	}
	*fetchedVendors[1].PmCustomerID = primitive.NewObjectID()

	mockRepo.On("Fetch", mock.Anything).Return(fetchedVendors, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedVendors, result)
	mockRepo.AssertExpectations(t)
}
