package route

import (
	"time"

	controller "earnforglance/server/api/controller/catalog"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/catalog"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/catalog"
	usecase "earnforglance/server/usecase/catalog"

	"github.com/gin-gonic/gin"
)

func ProductAttributeCombinationRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductAttributeCombinationRepository(db, domain.CollectionProductAttributeCombination)
	lc := &controller.ProductAttributeCombinationController{
		ProductAttributeCombinationUsecase: usecase.NewProductAttributeCombinationUsecase(ur, timeout),
		Env:                                env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/product_attribute_combinations", lc.Fetch)
	itemGroup.GET("/product_attribute_combination", lc.FetchByID)
	itemGroup.POST("/product_attribute_combination", lc.Create)
	itemGroup.POST("/product_attribute_combinations", lc.CreateMany)
	itemGroup.PUT("/product_attribute_combination", lc.Update)
	itemGroup.DELETE("/product_attribute_combination", lc.Delete)
}
