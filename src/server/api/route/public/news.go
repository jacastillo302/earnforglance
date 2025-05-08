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

func NewsItemRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewNewsItemRepository(db, domain.CollectionUser)
	tp := &controller.NewsItemController{
		NewsItemUsecase: usecase.NewnewsItemUsecase(ur, timeout),
		Env:             env,
	}
	group.POST("/news", tp.GetNewsItems)
}
