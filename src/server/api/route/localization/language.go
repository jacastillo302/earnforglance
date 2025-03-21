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

func LanguageRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewLanguageRepository(db, domain.CollectionLanguage)
	lc := &controller.LanguageController{
		LanguageUsecase: usecase.NewLanguageUsecase(ur, timeout),
		Env:             env,
	}

	group.GET("/languages", lc.Fetch)
	group.GET("/language", lc.FetchByID)
	group.POST("/language", lc.Create)
	group.PUT("/language", lc.Update)
	group.DELETE("/language", lc.Delete)
}
