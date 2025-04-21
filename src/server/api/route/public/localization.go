package route

import (
	"time"

	controller "earnforglance/server/api/controller/public"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/public"

	repository "earnforglance/server/repository/public"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/public"

	"github.com/gin-gonic/gin"
)

func LocalizationRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewLocalizationRepository(db, domain.CollectionUser)
	tp := &controller.LocalizationController{
		LocalizationUsecase: usecase.NewlocalizationUsecase(ur, timeout),
		Env:                 env,
	}
	group.GET("/localizations", tp.GetLocalizations)
}
