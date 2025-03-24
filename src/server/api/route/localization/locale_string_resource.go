package route

import (
	"time"

	controller "earnforglance/server/api/controller/localization"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/localization"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/localization"
	usecase "earnforglance/server/usecase/localization"

	"github.com/gin-gonic/gin"
)

func LocaleStringResourceRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewLocaleStringResourceRepository(db, domain.CollectionLocaleStringResource)
	lc := &controller.LocaleStringResourceController{
		LocaleStringResourceUsecase: usecase.NewLocaleStringResourceUsecase(ur, timeout),
		Env:                         env,
	}

	group.GET("/locale_string_resources", lc.Fetch)
	group.GET("/locale_string_resource", lc.FetchByID)
	group.POST("/locale_string_resource", lc.Create)
	group.POST("/locale_string_resources", lc.CreateMany)
	group.PUT("/locale_string_resource", lc.Update)
	group.DELETE("/locale_string_resource", lc.Delete)
}
