package main

import (
	"REST_golang/controller"
	"REST_golang/db"
	"REST_golang/repository"
	"REST_golang/router"
	"REST_golang/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)

	e.Logger.Fatal(e.Start(":8080"))
}
