package route

import (
	"time"

	controller "earnforglance/server/api/controller/stores"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/stores"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/stores"
	usecase "earnforglance/server/usecase/stores"

	"github.com/gin-gonic/gin"
)

func StoreMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewStoreMappingRepository(db, domain.CollectionStoreMapping)
	lc := &controller.StoreMappingController{
		StoreMappingUsecase: usecase.NewStoreMappingUsecase(ur, timeout),
		Env:                 env,
	}
	itemGroup := group.Group("/api/v1/stores")
	itemGroup.GET("/store_mappings", lc.Fetch)
	itemGroup.GET("/store_mapping", lc.FetchByID)
	itemGroup.POST("/store_mapping", lc.Create)
	itemGroup.POST("/store_mappings", lc.CreateMany)
	itemGroup.PUT("/store_mapping", lc.Update)
	itemGroup.DELETE("/store_mapping", lc.Delete)
}
