package bootstrap

import (
	"context"
	"fmt"
	"log"

	"earnforglance/server/service/data/mongo"
)

func NewMongoDatabase(env *Env) mongo.Client {
	//_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	mongodbURI := fmt.Sprintf("mongodb+srv://%s:%s@%s", dbUser, dbPass, dbHost)

	if env.AppEnv == "development" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

		if dbUser == "" || dbPass == "" {
			mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
		}
	} else {
		if dbPort != "" {
			mongodbURI = fmt.Sprintf("mongodb+srv://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)
		}
	}

	client, err := mongo.NewClient(mongodbURI)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
