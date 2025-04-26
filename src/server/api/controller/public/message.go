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

func (cc *NewsLetterController) NewsLetterInactivate(c *gin.Context) {
	guid := c.Query("guid")
	if guid == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Missing guid parameter"})
		return
	}

	productResponse, err := cc.NewsLetterUsecase.NewsLetterInactivate(c, guid)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, productResponse)
}

func (cc *NewsLetterController) NewsLetterUnSubscribe(c *gin.Context) {
	var request domain.NewsLetterRequest

	news, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Failed to read request body"})
		return
	}

	err = json.Unmarshal(news, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Invalid request body"})
		return
	}

	productResponse, err := cc.NewsLetterUsecase.NewsLetterUnSubscribe(c, request)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, productResponse)
}

func (cc *NewsLetterController) NewsLetterActivation(c *gin.Context) {
	guid := c.Query("guid")
	if guid == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Missing guid parameter"})
		return
	}

	productResponse, err := cc.NewsLetterUsecase.NewsLetterActivation(c, guid)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, productResponse)
}

func (cc *NewsLetterController) NewsLetterSubscription(c *gin.Context) {
	var request domain.NewsLetterRequest

	news, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Failed to read request body"})
		return
	}

	err = json.Unmarshal(news, &request)
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
