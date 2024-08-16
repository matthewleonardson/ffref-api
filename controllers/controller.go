package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matthewleonardson/ffref-api/dto"
	"github.com/matthewleonardson/ffref-api/processors"
)

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

func CreateRosterSlot(c *gin.Context) {

	var rosterSlotDTO dto.RosterSlotInsertionDTO
	c.BindJSON(&rosterSlotDTO)
	processors.ProcessCreateRosterSlot(c, &rosterSlotDTO)

}

func ReadPlayer(c *gin.Context) {

	name := strings.ReplaceAll(c.Param("name"), "-", " ")
	processors.ProcessReadPlayer(c, &name)

}
