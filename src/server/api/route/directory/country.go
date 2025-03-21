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

	group.GET("/countries", lc.Fetch)
	group.GET("/country", lc.FetchByID)
	group.POST("/country", lc.Create)
	group.PUT("/country", lc.Update)
	group.DELETE("/country", lc.Delete)
}
