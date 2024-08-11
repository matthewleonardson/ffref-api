package dto

type TeamInsertionDTO struct {
	TeamName   *string `json:"team_name"`
	Manager    *string `json:"manager"`
	YearJoined *int    `json:"year_joined"`
	YearLeft   *int    `json:"year_left"`
}

type TeamSelectionDTO struct {
	ID *int `json:"id"`
	TeamInsertionDTO
}
