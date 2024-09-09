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

	baseEndpoint := os.Getenv("BASE_ENDPOINT")

	router.POST(baseEndpoint + "/team", handlers.CreateTeam)
	router.GET(baseEndpoint + "/team", handlers.ReadTeams)
	router.POST(baseEndpoint + "/rosterslot", handlers.CreateRosterSlot)
	router.GET(baseEndpoint + "/team/:id", handlers.ReadTeam)
	router.GET(baseEndpoint + "/player/:name", handlers.ReadPlayer)

	router.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT")))

	defer database.DB.Db.Close()

}
