// Package main является точкой входа.
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ArtemST2006/Avito_internship/backend/internal/delivery/http/handler"
	"github.com/ArtemST2006/Avito_internship/backend/internal/repository"
	"github.com/ArtemST2006/Avito_internship/backend/internal/server"
	"github.com/ArtemST2006/Avito_internship/backend/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	db, err := repository.InitBD() // инициализация бд
	if err != nil {
		logrus.Fatalf("fatal initializetion db {main.go}, %s", err.Error())
	}

	// сервис разделё на 3 слоя
	repos := repository.NewRepository(db)   // слой репозитория для работы с бд
	services := service.NewService(repos)   // слой сервисов для работы с бизнес логикой
	handler := handler.NewHandler(services) // слой хэндлеров для отловки запросов

	srv := new(server.Server)
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.Run("8000", handler.InitRoutes()); err != nil {
			logrus.Fatalf("error in init http server{main.go }: %s", err.Error())
		}
	}()

	logrus.Print("Server up")
	<-done
	logrus.Print("Server down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error with shutting down %s", err.Error())
	}

	closer, err := db.DB()
	if err2 := closer.Close(); err != nil || err2 != nil {
		logrus.Fatalf("error with clossing db %s", err.Error())
	}
}
