package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"project/internal/database"
	"project/internal/handlers"
	"project/internal/messagesService"
	"project/internal/userService"
	"project/internal/web/messages"
	"project/internal/web/users"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&messagesService.Message{})
	if err != nil {
		return
	}

	messagesRepo := messagesService.NewMessageRepository(database.DB)
	messageService := messagesService.NewMessageService(messagesRepo)
	MessagesHandler := handlers.NewHandler(messageService)

	userRepo := userService.NewUserRepository(database.DB)
	us2rService := userService.NewUserService(userRepo)
	UserHandlers := handlers.NewUserHandler(us2rService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := messages.NewStrictHandler(MessagesHandler, nil) // тут будет ошибка
	messages.RegisterHandlers(e, strictHandler)

	usersStrictHandler := users.NewStrictHandler(UserHandlers, nil) // This is where the users logic comes in
	users.RegisterHandlers(e, usersStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
