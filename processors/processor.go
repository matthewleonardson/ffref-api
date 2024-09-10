package processors

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/matthewleonardson/ffref-api/domain"
	"github.com/matthewleonardson/ffref-api/dto"
	"github.com/matthewleonardson/ffref-api/repositories"
)

func ProcessCreateTeam(c *gin.Context, teamDTO *dto.TeamInsertionDTO) {

	team := dto.MapTeamInsertionDTO(teamDTO)
	response, err := repositories.InsertTeam(team)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest})
	} else {
		c.IndentedJSON(http.StatusCreated, dto.MapTeamSelectionDTO(response))
	}

}

func ProcessReadTeams(c *gin.Context) {

	response, err := repositories.SelectTeams()

	var dtos = []dto.TeamSelectionDTO{}
	for _, x := range response {
		dtos = append(dtos, *dto.MapTeamSelectionDTO(&x))
	}

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError})
	} else {
		c.IndentedJSON(http.StatusOK, dtos)
	}

}

func ProcessReadTeam(c *gin.Context, id *int) {

	response, err := repositories.SelectTeam(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError})
	} else {
		c.IndentedJSON(http.StatusOK, dto.MapTeamSelectionDTO(response))
	}

}

func ProcessCreateRosterSlot(c *gin.Context, rosterSlotDTO *dto.RosterSlotInsertionDTO) {

	rosterSlot := dto.MapRosterSlotInsertionDTO(rosterSlotDTO)

	var playerID *int

	response, err := repositories.SelectPlayerByName(rosterSlotDTO.PlayerName)

	if errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows) {
		var player domain.Player
		player.PlayerName = rosterSlotDTO.PlayerName
		updatedPlayer, err := repositories.InsertPlayer(&player)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError})
			return
		}

		playerID = updatedPlayer.ID

	} else {
		playerID = response.ID
	}

	rosterSlot.PlayerID = playerID

	updatedRosterSlot, err := repositories.InsertRosterSlot(rosterSlot)

	if err != nil {
		c.IndentedJSON(410, gin.H{"status_code": http.StatusBadRequest})
	} else {
		c.IndentedJSON(http.StatusOK, dto.MapRosterSlotResponseDTO(updatedRosterSlot))
	}

}

func ProcessReadPlayer(c *gin.Context, name *string) {

	response, err := repositories.SelectPlayerByName(name)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest})
	} else {
		c.IndentedJSON(http.StatusOK, dto.MapPlayerSelectionDTO(response))
	}

}

func ProcessReadPlayerHistory(c *gin.Context, name *string) {

	player, err := repositories.SelectPlayerByName(name)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest})
	}

	response, err := repositories.SelectPlayerHistoryByID(player.ID)

	var dtos = []dto.RosterSlotSelectionDTO{}
	for _, x := range response {

		dto := dto.MapRosterSlotSelectionDTO(&x)
		team, err := repositories.SelectTeam(x.TeamID)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest})
		}

		dto.Team = team.TeamName

		dtos = append(dtos, *dto)
	}

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"player_name": player.PlayerName, "history": dtos})
	}

}

func ProcessReadPlayerHistoryByYear(c *gin.Context, name *string, year *int) {

	player, err := repositories.SelectPlayerByName(name)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest})
	}

	response, err := repositories.SelectPlayerHistoryByIDAndYear(player.ID, year)

	var dtos = []dto.RosterSlotSelectionDTO{}
	for _, x := range response {

		dto := dto.MapRosterSlotSelectionDTO(&x)
		team, err := repositories.SelectTeam(x.TeamID)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest})
		}

		dto.Team = team.TeamName

		dtos = append(dtos, *dto)
	}

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"player_name": player.PlayerName, "history": dtos})
	}
}
