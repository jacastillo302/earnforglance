package route

import (
	"time"

	controller "earnforglance/server/api/controller/api"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/api"

	repository "earnforglance/server/repository/api"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/api"

	"github.com/gin-gonic/gin"
)

func ApiClientRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewApiClientRepository(db, domain.CollectionApiClient)
	lc := &controller.ApiClientController{
		ApiClientUsecase: usecase.NewApiClientUsecase(ur, timeout),
		Env:              env,
	}

	Group := group.Group("/api/v1/auth")

	group.GET("/apiclients", lc.Fetch)
	group.GET("/apiclient", lc.FetchByID)
	group.POST("/apiclient", lc.Create)
	Group.POST("/apiclients", lc.CreateMany)
	group.PUT("/apiclient", lc.Update)
	group.DELETE("/apiclient", lc.Delete)
}
