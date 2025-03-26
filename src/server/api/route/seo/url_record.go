package route

import (
	"time"

	controller "earnforglance/server/api/controller/seo"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/seo"

	repository "earnforglance/server/repository/seo"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/seo"

	"github.com/gin-gonic/gin"
)

func UrlRecordRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUrlRecordRepository(db, domain.CollectionUrlRecord)
	lc := &controller.UrlRecordController{
		UrlRecordUsecase: usecase.NewUrlRecordUsecase(ur, timeout),
		Env:              env,
	}
	itemGroup := group.Group("/api/v1/seo")
	itemGroup.GET("/url_records", lc.Fetch)
	itemGroup.GET("/url_record", lc.FetchByID)
	itemGroup.POST("/url_record", lc.Create)
	itemGroup.POST("/url_records", lc.CreateMany)
	itemGroup.PUT("/url_record", lc.Update)
	itemGroup.DELETE("/url_record", lc.Delete)
}
