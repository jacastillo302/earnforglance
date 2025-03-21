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

	group.GET("/product_attribute_combinations", lc.Fetch)
	group.GET("/product_attribute_combination", lc.FetchByID)
	group.POST("/product_attribute_combination", lc.Create)
	group.PUT("/product_attribute_combination", lc.Update)
	group.DELETE("/product_attribute_combination", lc.Delete)
}
