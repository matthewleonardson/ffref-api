package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/matthewleonardson/ffref-api/database"
	"github.com/matthewleonardson/ffref-api/domain"
	"github.com/matthewleonardson/ffref-api/dto"
)

func InsertTeam(team *domain.Team) (*dto.TeamSelectionDTO, error) {

	query := `INSERT INTO teams (team_name, manager, year_joined, year_left, created_at) VALUES (@teamName, @manager, @yearJoined, @yearLeft, @createdAt) RETURNING id, created_at, updated_at`

	args := pgx.NamedArgs{
		"teamName":   team.TeamName,
		"manager":    team.Manager,
		"yearJoined": team.YearJoined,
		"yearLeft":   team.YearLeft,
		"createdAt":  time.Now(),
	}

	err := database.DB.Db.QueryRow(context.Background(), query, args).Scan(&team.ID, &team.CreatedAt, &team.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return dto.MapTeamSelectionDTO(team), nil

}

func SelectTeams() ([]dto.TeamSelectionDTO, error) {

	query := `SELECT id, team_name, manager, year_joined, year_left FROM teams ORDER BY id`

	rows, err := database.DB.Db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var dtoSlice = []dto.TeamSelectionDTO{}

	for rows.Next() {
		var dto dto.TeamSelectionDTO
		err = rows.Scan(&dto.ID, &dto.TeamName, &dto.Manager, &dto.YearJoined, &dto.YearLeft)
		if err != nil {
			return nil, err
		}
		dtoSlice = append(dtoSlice, dto)
	}

	return dtoSlice, nil

}

func SelectTeam(id *int) (*dto.TeamSelectionDTO, error) {

	query := `SELECT id, team_name, manager, year_joined, year_left FROM teams WHERE id = @id`

	args := pgx.NamedArgs{
		"id": *id,
	}

	var dto dto.TeamSelectionDTO

	err := database.DB.Db.QueryRow(context.Background(), query, args).Scan(&dto.ID, &dto.TeamName, &dto.Manager, &dto.YearJoined, &dto.YearLeft)

	if err != nil {
		return nil, err
	}

	return &dto, nil
}
