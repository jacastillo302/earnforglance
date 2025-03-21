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

func StoreRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewStoreRepository(db, domain.CollectionStore)
	lc := &controller.StoreController{
		StoreUsecase: usecase.NewStoreUsecase(ur, timeout),
		Env:          env,
	}

	group.GET("/stores", lc.Fetch)
	group.GET("/store", lc.FetchByID)
	group.POST("/store", lc.Create)
	group.PUT("/store", lc.Update)
	group.DELETE("/store", lc.Delete)
}
