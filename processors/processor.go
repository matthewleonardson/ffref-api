package processors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matthewleonardson/ffref-api/dto"
	"github.com/matthewleonardson/ffref-api/repositories"
)

func ProcessCreateTeam(c *gin.Context, teamDTO *dto.TeamInsertionDTO) {

	team := dto.MapTeamInsertionDTO(teamDTO)
	response, err := repositories.InsertTeam(team)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest})
	} else {
		c.IndentedJSON(http.StatusCreated, response)
	}

}

func ProcessReadTeams(c *gin.Context) {

	response, err := repositories.SelectTeams()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError})
	} else {
		c.IndentedJSON(http.StatusOK, response)
	}

}

func ProcessReadTeam(c *gin.Context, id *int) {

	response, err := repositories.SelectTeam(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusBadRequest})
	} else {
		c.IndentedJSON(http.StatusOK, response)
	}

}
