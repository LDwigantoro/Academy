package main

import (
	"latihan3/config"
	"latihan3/handler"
	"latihan3/repository"
	"latihan3/routes"
	"latihan3/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Database()
	config.AutoMigrate()

	app := fiber.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHanlder := handler.NewUserHandler(userUsecase)

	routes.Routes(app, userHanlder)

	app.Listen(":5555")
}
