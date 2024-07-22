package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/KseniiaTD/time-tracker/internal/database"
	"github.com/KseniiaTD/time-tracker/internal/logger"
	"github.com/KseniiaTD/time-tracker/internal/router"
	"github.com/KseniiaTD/time-tracker/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	_ "github.com/KseniiaTD/time-tracker/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2" // http-swagger middleware
)

// @title Swagger API
// @version 1.0
// @description Swagger API for project Time tracker
func main() {
	fmt.Println(runtime.GOROOT())
	err := godotenv.Load("../../config.env")
	if err != nil {
		log.Fatal(err)
	}
	debug := os.Getenv("DEBUG")
	debugValue, err := strconv.ParseBool(debug)
	if err != nil {
		log.Fatal(err)
	}

	if debugValue {
		logger.Logger().SetLevel(logrus.DebugLevel)
	} else {
		logger.Logger().SetLevel(logrus.InfoLevel)
	}

	db, err := database.Connect()
	if err != nil {
		log.Panic(err)
	}
	logger.Logger().Debug("Database started")
	defer db.Disconnect()

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx := context.Background()
	//defer cancel()

	srv := service.New(db, ctx)
	router := router.New(srv)
	port := os.Getenv("SERVICE_PORT")

	router.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL(
			"http://localhost:" + port + "/swagger/doc.json")))
	logger.Logger().Info("Service started")

	log.Panic(http.ListenAndServe(":"+port, router))
}
