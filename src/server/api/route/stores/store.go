package route

import (
	"time"

	controller "earnforglance/server/api/controller/stores"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/stores"

	repository "earnforglance/server/repository/stores"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/stores"

	"github.com/gin-gonic/gin"
)

func StoreRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewStoreRepository(db, domain.CollectionStore)
	lc := &controller.StoreController{
		StoreUsecase: usecase.NewStoreUsecase(ur, timeout),
		Env:          env,
	}
	itemGroup := group.Group("/api/v1/stores")
	itemGroup.GET("/stores", lc.Fetch)
	itemGroup.GET("/store", lc.FetchByID)
	itemGroup.POST("/store", lc.Create)
	itemGroup.POST("/stores", lc.CreateMany)
	itemGroup.PUT("/store", lc.Update)
	itemGroup.DELETE("/store", lc.Delete)
}
