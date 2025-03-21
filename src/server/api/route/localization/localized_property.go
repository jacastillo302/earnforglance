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

func LocalizedPropertyRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewLocalizedPropertyRepository(db, domain.CollectionLocalizedProperty)
	lc := &controller.LocalizedPropertyController{
		LocalizedPropertyUsecase: usecase.NewLocalizedPropertyUsecase(ur, timeout),
		Env:                      env,
	}

	group.GET("/localized_properties", lc.Fetch)
	group.GET("/localized_property", lc.FetchByID)
	group.POST("/localized_property", lc.Create)
	group.PUT("/localized_property", lc.Update)
	group.DELETE("/localized_property", lc.Delete)
}
