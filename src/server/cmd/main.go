package main

import (
	"fmt"
	"time"

	route "earnforglance/server/api/route"
	"earnforglance/server/bootstrap"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Server Service Start")

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Use(static.Serve("/", static.LocalFile("../web/dist", false)))

	gin.Run(env.ServerAddress)

}
