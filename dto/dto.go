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

type RosterSlotInsertionDTO struct {
	PlayerName      *string  `json:"player_name"`
	TeamID          *int     `json:"team_id"`
	Week            *int     `json:"week"`
	Year            *int     `json:"year"`
	Position        *string  `json:"position"`
	Benched         *bool    `json:"benched"`
	ProjectedPoints *float32 `json:"projected_points"`
	ActualPoints    *float32 `json:"actual_points"`
}

type RosterSlotResponseDTO struct {
	ID              *int     `json:"id"`
	PlayerID        *int     `json:"player_id"`
	TeamID          *int     `json:"team_id"`
	Week            *int     `json:"week"`
	Year            *int     `json:"year"`
	Position        *string  `json:"position"`
	Benched         *bool    `json:"benched"`
	ProjectedPoints *float32 `json:"projected_points"`
	ActualPoints    *float32 `json:"actual_points"`
}

type RosterSlotSelectionDTO struct {
	Team            *string  `json:"team"`
	Week            *int     `json:"week"`
	Year            *int     `json:"year"`
	Position        *string  `json:"position"`
	Benched         *bool    `json:"benched"`
	ProjectedPoints *float32 `json:"projected_points"`
	ActualPoints    *float32 `json:"actual_points"`
}

type PlayerSelectionDTO struct {
	ID         *int    `json:"id"`
	PlayerName *string `json:"player_name"`
}
