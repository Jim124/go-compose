package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/go-compose-rest/database"
	"github.com/go-compose-rest/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	mysqlDb := os.Getenv("MYSQL_DB")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	database.Init(mysqlHost, mysqlDb, mysqlUser, mysqlPassword)
	port := os.Getenv("PORT")
	server := gin.Default()
	// server.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://localhost:5173"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))
	server.Use(cors.Default())
	routes.RegisterServer(server)
	server.Run(fmt.Sprintf(":%v", port))

}
