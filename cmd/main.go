package main

import (
	"fmt"
	"log"

	"github.com/yahyaammar-dev/pacebe/cmd/api"
	"github.com/yahyaammar-dev/pacebe/configs"
	"github.com/yahyaammar-dev/pacebe/db"
	"github.com/yahyaammar-dev/pacebe/services/event"
	"github.com/yahyaammar-dev/pacebe/services/logger"
	"github.com/yahyaammar-dev/pacebe/types"
	"gorm.io/gorm"
)

// @title Base APIs
// @version 1.0
// @description APIs for base project
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
func main() {
	dbInstance, err := db.NewSQLiteStorage()
	if err != nil {
		log.Fatal(err)
	}

	initStorage(dbInstance)

	// Migrations
	dbInstance.AutoMigrate(&types.User{}, &types.Role{}, &types.Product{})

	// Seeders
	// seeder := db.NewSeeder(dbInstance)
	// seeder.CreateProducts()

	// logger
	if err := logger.Init(); err != nil {
		log.Fatal(err)
	}

	// events and listeners
	event.NewListener()

	// sockets
	// socketServer := socket.NewConnection()
	// go func() {
	// 	if err := socketServer.Run(); err != nil {
	// 		log.Fatalf("Socket server failed: %v", err)
	// 	}
	// }()

	// redis
	// redix.InitRedis()

	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), dbInstance)
	if err := server.Run(); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}
func initStorage(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
