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

	group.GET("/store_mappings", lc.Fetch)
	group.GET("/store_mapping", lc.FetchByID)
	group.POST("/store_mapping", lc.Create)
	group.PUT("/store_mapping", lc.Update)
	group.DELETE("/store_mapping", lc.Delete)
}
