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
