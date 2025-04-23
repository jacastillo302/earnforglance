package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/public"

	"github.com/gin-gonic/gin"
)

type GdprConsentController struct {
	GdprConsentUsecase domain.GdprConsentUsecase
	Env                *bootstrap.Env
}

func (cc *GdprConsentController) GetGdprConsents(c *gin.Context) {
	var request domain.GdprConsentRequest

	filter, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Failed to read request body"})
		return
	}

	err = json.Unmarshal(filter, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Invalid request body"})
		return
	}

	productResponse, err := cc.GdprConsentUsecase.GetGdprConsents(c, request)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, productResponse)
}
