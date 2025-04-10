package route

import (
	"time"

	"earnforglance/server/api/controller"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/security"
	"earnforglance/server/repository"
	"earnforglance/server/service/data/mongo"
	"earnforglance/server/usecase"

	"github.com/gin-gonic/gin"
)

func ProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
