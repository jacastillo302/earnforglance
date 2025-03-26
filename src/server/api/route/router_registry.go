package route

import (
	"time"

	"earnforglance/server/bootstrap"
	"earnforglance/server/service/data/mongo"

	"github.com/gin-gonic/gin"
)

// RouterRegistry maintains a centralized registry of router functions
var RouterRegistry = make(map[string]map[string]RouterFunc)

// RegisterRouter adds a router function to the registry
func RegisterRouter(module, name string, routerFunc RouterFunc) {
	if _, exists := RouterRegistry[module]; !exists {
		RouterRegistry[module] = make(map[string]RouterFunc)
	}
	RouterRegistry[module][name] = routerFunc
}

// SetupFromRegistry initializes all routers from the registry
func SetupFromRegistry(env *bootstrap.Env, timeout time.Duration, db mongo.Database, router *gin.RouterGroup) {
	for _, moduleRouters := range RouterRegistry {
		for _, routerFunc := range moduleRouters {
			routerFunc(env, timeout, db, router)
		}
	}
}
