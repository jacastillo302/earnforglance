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

func LanguageRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewLanguageRepository(db, domain.CollectionLanguage)
	lc := &controller.LanguageController{
		LanguageUsecase: usecase.NewLanguageUsecase(ur, timeout),
		Env:             env,
	}
	itemGroup := group.Group("/api/v1/localization")
	itemGroup.GET("/languages", lc.Fetch)
	itemGroup.GET("/language", lc.FetchByID)
	itemGroup.POST("/language", lc.Create)
	itemGroup.POST("/languages", lc.CreateMany)
	itemGroup.PUT("/language", lc.Update)
	itemGroup.DELETE("/language", lc.Delete)
}
