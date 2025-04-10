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

func SignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
