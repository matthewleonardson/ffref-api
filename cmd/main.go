package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	handlers "github.com/matthewleonardson/ffref-api/controllers"
	"github.com/matthewleonardson/ffref-api/database"
)

func main() {

	database.ConnectDb()

	router := gin.Default()

	router.POST("/team", handlers.CreateTeam)
	router.GET("/team", handlers.ReadTeams)
	router.GET("/team/:id", handlers.ReadTeam)

	router.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT")))

	defer database.DB.Db.Close()

}
