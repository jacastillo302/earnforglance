package route

import (
	"time"

	controller "earnforglance/server/api/controller/install"
	"earnforglance/server/bootstrap"

	repository "earnforglance/server/repository/install"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/install"

	"github.com/gin-gonic/gin"
)

func InstallRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewInstallRepository(db)
	lc := &controller.InstallController{
		InstallUsecase: usecase.NewInstallUsecase(ur, timeout),
		Env:            env,
	}

	group.GET("/ping_database", lc.PingDatabase)
	group.GET("/install_stores", lc.InstallStores)

}
