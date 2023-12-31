package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/pasha1coil/testingmelvad/internal/handler"
	"github.com/pasha1coil/testingmelvad/internal/repository"
	"github.com/pasha1coil/testingmelvad/internal/repository/postgre"
	"github.com/pasha1coil/testingmelvad/internal/repository/redis"
	"github.com/pasha1coil/testingmelvad/internal/service"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.SetFormatter(new(log.JSONFormatter))
	// Загрузка переменных окружения из .env файла
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load .env file: %s", err.Error())
	}

	// Инициализация базы данных
	db, err := postgre.Migrate(&postgre.Config{
		Uname:      os.Getenv("DB_UNAME"),
		Pass:       os.Getenv("DB_PASS"),
		NameDB:     os.Getenv("DB_NAMEDB"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		SSL:        os.Getenv("DB_SSL"),
		DriverName: os.Getenv("DB_DriverName"),
	})
	if err != nil {
		log.Fatalf("failed to connect to the database: %s", err.Error())
	}
	defer db.Close()

	// Инициализация Redis
	redisClient, err := redis.InitRedis(&redis.RedisConfig{
		Addr: os.Getenv("REDIS_ADDR"),
		Port: os.Getenv("REDIS_PORT"),
		Pass: os.Getenv("REDIS_PASS"),
	})
	if err != nil {
		log.Fatalf("failed to connect to the Redis: %s", err.Error())
	}
	defer redisClient.Close()

	// Инициализация репозитория
	log.Infoln("Init repository...")
	rep := repository.NewTasksRepo(db, redisClient)
	// Инициализация сервиса
	log.Infoln("Init service...")
	srvc := service.NewTasksService(rep)

	// Инициализация GoFiber
	app := fiber.New(fiber.Config{
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	})
	app.Use(recover.New())

	// Инициализация API хэндлеров
	log.Infoln("Init handlers...")
	apiHandlers := handler.NewHandlers(srvc)
	app.Post("/redis/incr", apiHandlers.PostIncr)
	app.Post("/sign/hmacsha512", apiHandlers.PostHmac)
	app.Post("/postgres/users", apiHandlers.PostUsers)

	// Вывод роутов в консоль
	log.Infoln("POST /redis/incr")
	log.Infoln("POST /sign/hmacsha512")
	log.Infoln("POST /postgres/users")

	serverError := make(chan error, 1)

	// Starting server
	go func() {
		log.Infoln("Starting HTTP server on port:", os.Getenv("SRV_PORT"))
		if err := app.Listen(":" + os.Getenv("SRV_PORT")); err != nil {
			log.Errorf("Failed to start HTTP server: %s\n", err.Error())
			serverError <- err
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverError:
		log.Fatalf("Failed to start server, %v", err)
	case <-stop:
		log.Infoln("Shutting down the server...")
		if err := app.Shutdown(); err != nil {
			log.Infof("Graceful shutdown failed: %s\n", err.Error())
		} else {
			log.Infoln("Server stopped")
		}
	}
}
