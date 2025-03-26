package route

import (
	"time"

	controller "earnforglance/server/api/controller/gdpr"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/gdpr"

	repository "earnforglance/server/repository/gdpr"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/gdpr"

	"github.com/gin-gonic/gin"
)

func GdprConsentRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewGdprConsentRepository(db, domain.CollectionGdprConsent)
	lc := &controller.GdprConsentController{
		GdprConsentUsecase: usecase.NewGdprConsentUsecase(ur, timeout),
		Env:                env,
	}
	itemGroup := group.Group("/api/v1/gdpr")
	itemGroup.GET("/gdpr_consents", lc.Fetch)
	itemGroup.GET("/gdpr_consent", lc.FetchByID)
	itemGroup.POST("/gdpr_consent", lc.Create)
	itemGroup.POST("/gdpr_consents", lc.CreateMany)
	itemGroup.PUT("/gdpr_consent", lc.Update)
	itemGroup.DELETE("/gdpr_consent", lc.Delete)
}
