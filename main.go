package main

import (
	"new_project/database"
	"new_project/handlers"
	"new_project/repository"
	"new_project/services"
)

func main() {
	//Init mongo
	if err := database.InitMongo(database.EnvMongoURI(), database.EnvMongoDBName()); err != nil {
		return
	}
	defer database.CloseMongoDB()
	////Router
	db := database.GetMongoDB()
	repo := repository.NewRepository(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	router := handler.InitRouter()
	router.Run(":8080")
}
