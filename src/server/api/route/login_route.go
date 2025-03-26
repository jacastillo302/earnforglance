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

func LoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
