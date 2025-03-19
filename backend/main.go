package main

import (
	"fmt"
	"os"

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
	routes.RegisterServer(server)
	server.Run(fmt.Sprintf(":%v", port))

}
