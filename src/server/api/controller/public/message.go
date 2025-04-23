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

type NewsLetterController struct {
	NewsLetterUsecase domain.NewsLetterUsecase
	Env               *bootstrap.Env
}

func (cc *NewsLetterController) NewsLetterSubscription(c *gin.Context) {
	var request domain.NewsLetterRequest

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

	productResponse, err := cc.NewsLetterUsecase.NewsLetterSubscription(c, request, c.ClientIP())
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, productResponse)
}
