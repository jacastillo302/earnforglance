package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	controller "earnforglance/server/api/controller/affiliate"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/affiliate"
	"earnforglance/server/domain/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func getLoginToken(t *testing.T) string {

	// Create a mock request
	req, _ := http.Get("/api/login?email=test@gmail.com&password=test")
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Assert the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body to extract the token
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Return the token
	token, exists := response["token"]
	assert.True(t, exists, "Token not found in response")
	return token
}

func TestAffiliateController_Create(t *testing.T) {
	mockUsecase := new(mocks.AffiliateUsecase)
	controller := controller.AffiliateController{
		AffiliateUsecase: mockUsecase,
		Env:              &bootstrap.Env{},
	}

	mockUsecase.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

	body := "AdminComment=Test+Comment&FriendlyUrlName=test-url&Deleted=false&Active=true"
	req, _ := http.NewRequest(http.MethodPost, "/api/affiliate", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer test-token")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	controller.Create(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)
}

func TestAffiliateController_Update(t *testing.T) {
	mockUsecase := new(mocks.AffiliateUsecase)
	controller := controller.AffiliateController{
		AffiliateUsecase: mockUsecase,
		Env:              &bootstrap.Env{},
	}

	mockUsecase.On("Update", mock.Anything, mock.Anything).Return(nil).Once()

	body := "AdminComment=Updated+Comment&FriendlyUrlName=updated-url&Deleted=false&Active=true"
	req, _ := http.NewRequest(http.MethodPut, "/api/affiliate", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer test-token")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	controller.Update(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)
}

func TestAffiliateController_Delete(t *testing.T) {
	mockUsecase := new(mocks.AffiliateUsecase)
	controller := controller.AffiliateController{
		AffiliateUsecase: mockUsecase,
		Env:              &bootstrap.Env{},
	}

	mockUsecase.On("Delete", mock.Anything, "123").Return(nil).Once()

	req, _ := http.NewRequest(http.MethodDelete, "/api/affiliate?id=123", nil)
	req.Header.Set("Authorization", "Bearer test-token")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	controller.Delete(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)
}

func TestAffiliateController_FetchByID(t *testing.T) {
	mockUsecase := new(mocks.AffiliateUsecase)
	controller := controller.AffiliateController{
		AffiliateUsecase: mockUsecase,
		Env:              &bootstrap.Env{},
	}

	mockAffiliate := domain.Affiliate{
		AdminComment:    "Test Comment",
		FriendlyUrlName: "test-url",
		Deleted:         false,
		Active:          true,
	}

	mockUsecase.On("FetchByID", mock.Anything, "123").Return(&mockAffiliate, nil).Once()

	req, _ := http.NewRequest(http.MethodGet, "/api/affiliate?id=123", nil)
	req.Header.Set("Authorization", "Bearer test-token")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	controller.FetchByID(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)
}

func TestAffiliateController_Fetch(t *testing.T) {
	mockUsecase := new(mocks.AffiliateUsecase)
	controller := controller.AffiliateController{
		AffiliateUsecase: mockUsecase,
		Env:              &bootstrap.Env{},
	}

	mockAffiliates := []domain.Affiliate{
		{
			AdminComment:    "Test Comment 1",
			FriendlyUrlName: "test-url-1",
			Deleted:         false,
			Active:          true,
		},
		{
			AdminComment:    "Test Comment 2",
			FriendlyUrlName: "test-url-2",
			Deleted:         false,
			Active:          true,
		},
	}

	mockUsecase.On("Fetch", mock.Anything).Return(mockAffiliates, nil).Once()

	req, _ := http.NewRequest(http.MethodGet, "/api/affiliates", nil)
	req.Header.Set("Authorization", "Bearer test-token")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	controller.Fetch(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)
}
