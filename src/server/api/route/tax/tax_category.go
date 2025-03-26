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
	itemGroup.GET("/taxcategories", lc.Fetch)
	itemGroup.GET("/taxcategory", lc.FetchByID)
	itemGroup.POST("/taxcategory", lc.Create)
	itemGroup.POST("/taxcategories", lc.CreateMany)
	itemGroup.PUT("/taxcategory", lc.Update)
	itemGroup.DELETE("/taxcategory", lc.Delete)
}
