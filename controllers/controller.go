package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matthewleonardson/ffref-api/dto"
	"github.com/matthewleonardson/ffref-api/processors"
)

// func CreatePlayer(c *gin.Context) {
// 	var newPlayer models.Player
// 	c.BindJSON(&newPlayer)
// 	database.DB.Db.Creatse(&newPlayer)
// 	c.IndentedJSON(http.StatusCreated, newPlayer)

// }

// func CreateTeam(c *gin.Context) {
// 	var newTeam models.Team
// 	c.BindJSON(&newTeam)
// 	database.DB.Db.Create(&newTeam)
// 	c.IndentedJSON(http.StatusCreated, newTeam)

// }

// func CreateRosterSlot(c *gin.Context) {
// 	var newRosterSlot models.RosterSlot
// 	c.BindJSON(&newRosterSlot)
// 	database.DB.Db.Create(&newRosterSlot)
// 	c.IndentedJSON(http.StatusCreated, newRosterSlot)

// }

func CreateTeam(c *gin.Context) {

	var teamDTO dto.TeamInsertionDTO
	c.BindJSON(&teamDTO)
	processors.ProcessCreateTeam(c, &teamDTO)

}

func ReadTeams(c *gin.Context) {

	processors.ProcessReadTeams(c)

}

func ReadTeam(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error_code": http.StatusBadRequest})
		return
	}

	processors.ProcessReadTeam(c, &id)

}
