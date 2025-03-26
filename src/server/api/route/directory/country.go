package route

import (
	"time"

	controller "earnforglance/server/api/controller/directory"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/directory"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/directory"
	usecase "earnforglance/server/usecase/directory"

	"github.com/gin-gonic/gin"
)

func CountryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCountryRepository(db, domain.CollectionCountry)
	lc := &controller.CountryController{
		CountryUsecase: usecase.NewCountryUsecase(ur, timeout),
		Env:            env,
	}

	itemGroup := group.Group("/api/v1/directory")
	itemGroup.GET("/countries", lc.Fetch)
	itemGroup.GET("/country", lc.FetchByID)
	itemGroup.POST("/country", lc.Create)
	itemGroup.POST("/countries", lc.CreateMany)
	itemGroup.PUT("/country", lc.Update)
	itemGroup.DELETE("/country", lc.Delete)
}
