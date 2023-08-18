package main

import (
	"context"
	"fmt"
	"log"
	"rest-api/config"
	"rest-api/constants"
	"rest-api/controllers"
	"rest-api/routes"
	"rest-api/services"

	//	"rest-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


var (
	mongoclient *mongo.Client
	ctx         context.Context
	server         *gin.Engine
)
func initRoutes(){
  routes.Default(server)
}

func initApp(mongoClient *mongo.Client){
	ctx = context.TODO()
	profileCollection := mongoClient.Database(constants.DatabaseName).Collection("profiles")
	profileService := services.NewProfileServiceInit(profileCollection, ctx)
	profileController := controllers.InitProfileController(profileService)
	routes.ProfileRoute(server,profileController)
}

func InitApp(mongoclient *mongo.Client){
	ctx:= context.TODO()
	transactionCollection := mongoclient.Database(constants.DatabaseName).Collection("transactions")
	transactionService := services.InitialiseTransactionService(transactionCollection,ctx)
	transactionController := controllers.InitTransController(transactionService)
	routes.TransactionRoute(server,transactionController)
	

}


func main(){
	server = gin.Default()
	mongoclient,err :=config.ConnectDataBase()
	defer   mongoclient.Disconnect(ctx)
	if err!=nil{
		panic(err)
	}
	// initRoutes()
	InitApp(mongoclient)
	fmt.Println("server running on port",constants.Port)
	log.Fatal(server.Run(constants.Port))
}

