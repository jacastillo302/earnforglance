package route

import (
	"time"

	"earnforglance/server/api/controller"
	"earnforglance/server/bootstrap"
	"earnforglance/server/domain"
	"earnforglance/server/mongo"
	"earnforglance/server/repository"
	"earnforglance/server/usecase"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
