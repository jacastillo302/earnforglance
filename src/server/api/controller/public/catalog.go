package controller

import (
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/public"

	"github.com/gin-gonic/gin"
)

type CatalogController struct {
	CatalogUsecase domain.CatalogtUsecase
	Env            *bootstrap.Env
}

func (cc *CatalogController) GetProduct(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := c.ShouldBind(&ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: err.Error()})
		return

	}

	productResponse, err := cc.CatalogUsecase.GetProduct(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, productResponse)

}
