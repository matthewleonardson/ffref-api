package dto

import "github.com/matthewleonardson/ffref-api/domain"

func MapTeamInsertionDTO(dto *TeamInsertionDTO) *domain.Team {

	var team domain.Team

	team.TeamName = dto.TeamName
	team.Manager = dto.Manager
	team.YearJoined = dto.YearJoined
	team.YearLeft = dto.YearLeft

	return &team

}

func MapTeamSelectionDTO(team *domain.Team) *TeamSelectionDTO {

	var dto TeamSelectionDTO

	dto.ID = team.ID
	dto.TeamName = team.TeamName
	dto.Manager = team.Manager
	dto.YearJoined = team.YearJoined
	dto.YearLeft = team.YearLeft

	return &dto

}

func MapRosterSlotInsertionDTO(dto *RosterSlotInsertionDTO) *domain.RosterSlot {

	var rosterSlot domain.RosterSlot

	rosterSlot.TeamID = dto.TeamID
	rosterSlot.Week = dto.Week
	rosterSlot.Year = dto.Year
	rosterSlot.Position = dto.Position
	rosterSlot.Benched = dto.Benched
	rosterSlot.ProjectedPoints = dto.ProjectedPoints
	rosterSlot.ActualPoints = dto.ActualPoints

	return &rosterSlot

}

func MapRosterSlotSelectionDTO(rosterSlot *domain.RosterSlot) *RosterSlotSelectionDTO {

	var dto RosterSlotSelectionDTO

	dto.ID = rosterSlot.ID
	dto.PlayerID = rosterSlot.PlayerID
	dto.TeamID = rosterSlot.PlayerID
	dto.Week = rosterSlot.Week
	dto.Year = rosterSlot.Year
	dto.Position = rosterSlot.Position
	dto.Benched = rosterSlot.Benched
	dto.ProjectedPoints = rosterSlot.ProjectedPoints
	dto.ActualPoints = rosterSlot.ActualPoints

	return &dto

}

func MapPlayerSelectionDTO(player *domain.Player) *PlayerSelectionDTO {

	var dto PlayerSelectionDTO

	dto.ID = player.ID
	dto.PlayerName = player.PlayerName

	return &dto

}
