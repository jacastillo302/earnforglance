package route

import (
	"time"

	controller "earnforglance/server/api/controller/tax"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/tax"

	repository "earnforglance/server/repository/tax"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/tax"

	"github.com/gin-gonic/gin"
)

func TaxCategoryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewTaxCategoryRepository(db, domain.CollectionTaxCategory)
	lc := &controller.TaxCategoryController{
		TaxCategoryUsecase: usecase.NewTaxCategoryUsecase(ur, timeout),
		Env:                env,
	}
	itemGroup := group.Group("/api/v1/tax")
	itemGroup.GET("/tax_categories", lc.Fetch)
	itemGroup.GET("/tax_category", lc.FetchByID)
	itemGroup.POST("/tax_category", lc.Create)
	itemGroup.POST("/tax_categories", lc.CreateMany)
	itemGroup.PUT("/tax_category", lc.Update)
	itemGroup.DELETE("/tax_category", lc.Delete)
}
