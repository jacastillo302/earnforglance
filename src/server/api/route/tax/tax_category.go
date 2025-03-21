package route

import (
	"time"

	controller "earnforglance/server/api/controller/tax"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/tax"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/tax"
	usecase "earnforglance/server/usecase/tax"

	"github.com/gin-gonic/gin"
)

func TaxCategoryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewTaxCategoryRepository(db, domain.CollectionTaxCategory)
	lc := &controller.TaxCategoryController{
		TaxCategoryUsecase: usecase.NewTaxCategoryUsecase(ur, timeout),
		Env:                env,
	}

	group.GET("/taxcategorys", lc.Fetch)
	group.GET("/taxcategory", lc.FetchByID)
	group.POST("/taxcategory", lc.Create)
	group.PUT("/taxcategory", lc.Update)
	group.DELETE("/taxcategory", lc.Delete)
}
