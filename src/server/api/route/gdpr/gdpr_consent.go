package route

import (
	"time"

	controller "earnforglance/server/api/controller/gdpr"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/gdpr"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/gdpr"
	usecase "earnforglance/server/usecase/gdpr"

	"github.com/gin-gonic/gin"
)

func GdprConsentRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewGdprConsentRepository(db, domain.CollectionGdprConsent)
	lc := &controller.GdprConsentController{
		GdprConsentUsecase: usecase.NewGdprConsentUsecase(ur, timeout),
		Env:                env,
	}

	group.GET("/gdpr_consents", lc.Fetch)
	group.GET("/gdpr_consent", lc.FetchByID)
	group.POST("/gdpr_consent", lc.Create)
	group.PUT("/gdpr_consent", lc.Update)
	group.DELETE("/gdpr_consent", lc.Delete)
}
