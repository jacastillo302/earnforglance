package route

import (
	"time"

	controller "earnforglance/server/api/controller/localization"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/localization"

	repository "earnforglance/server/repository/localization"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/localization"

	"github.com/gin-gonic/gin"
)

func LocaleStringResourceRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewLocaleStringResourceRepository(db, domain.CollectionLocaleStringResource)
	lc := &controller.LocaleStringResourceController{
		LocaleStringResourceUsecase: usecase.NewLocaleStringResourceUsecase(ur, timeout),
		Env:                         env,
	}
	itemGroup := group.Group("/api/v1/localization")
	itemGroup.GET("/locale_string_resources", lc.Fetch)
	itemGroup.GET("/locale_string_resource", lc.FetchByID)
	itemGroup.POST("/locale_string_resource", lc.Create)
	itemGroup.POST("/locale_string_resources", lc.CreateMany)
	itemGroup.PUT("/locale_string_resource", lc.Update)
	itemGroup.DELETE("/locale_string_resource", lc.Delete)
}
