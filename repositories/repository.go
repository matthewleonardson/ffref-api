package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/matthewleonardson/ffref-api/database"
	"github.com/matthewleonardson/ffref-api/domain"
)

func InsertTeam(team *domain.Team) (*domain.Team, error) {

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

	return team, nil

}

func SelectTeams() ([]domain.Team, error) {

	query := `SELECT id, team_name, manager, year_joined, year_left, created_at, updated_at FROM teams ORDER BY id`

	rows, err := database.DB.Db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var teams = []domain.Team{}

	for rows.Next() {
		var team domain.Team
		err = rows.Scan(&team.ID, &team.TeamName, &team.Manager, &team.YearJoined, &team.YearLeft, &team.CreatedAt, &team.UpdatedAt)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	return teams, nil

}

func SelectTeam(id *int) (*domain.Team, error) {

	query := `SELECT id, team_name, manager, year_joined, year_left, created_at, updated_at FROM teams WHERE id = @id`

	args := pgx.NamedArgs{
		"id": *id,
	}

	var team domain.Team

	err := database.DB.Db.QueryRow(context.Background(), query, args).Scan(&team.ID, &team.TeamName, &team.Manager, &team.YearJoined, &team.YearLeft, &team.CreatedAt, &team.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &team, nil
}

func SelectPlayerByName(name *string) (*domain.Player, error) {

	query := `SELECT id, player_name, created_at, updated_at FROM players WHERE UPPER(player_name) = UPPER(@name)`

	args := pgx.NamedArgs{
		"name": *name,
	}

	var player domain.Player

	err := database.DB.Db.QueryRow(context.Background(), query, args).Scan(&player.ID, &player.PlayerName, &player.CreatedAt, &player.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &player, nil

}

func InsertPlayer(player *domain.Player) (*domain.Player, error) {

	query := `INSERT INTO players (player_name, created_at) VALUES (@playerName, @createdAt) RETURNING id, created_at, updated_at`

	args := pgx.NamedArgs{
		"playerName": player.PlayerName,
		"createdAt":  time.Now(),
	}

	err := database.DB.Db.QueryRow(context.Background(), query, args).Scan(&player.ID, &player.CreatedAt, &player.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return player, nil

}

func InsertRosterSlot(rosterSlot *domain.RosterSlot) (*domain.RosterSlot, error) {

	query := `INSERT INTO roster_slot (player_id, team_id, week, year, position, benched, projected_points, actual_points, created_at) 
				VALUES (@playerID, @teamID, @week, @year, @position, @benched, @projectedPoints, @actualPoints, @createdAt) RETURNING id, created_at, updated_at`

	args := pgx.NamedArgs{
		"playerID":        rosterSlot.PlayerID,
		"teamID":          rosterSlot.TeamID,
		"week":            rosterSlot.Week,
		"year":            rosterSlot.Year,
		"position":        rosterSlot.Position,
		"benched":         rosterSlot.Benched,
		"projectedPoints": rosterSlot.ProjectedPoints,
		"actualPoints":    rosterSlot.ActualPoints,
		"createdAt":       time.Now(),
	}

	err := database.DB.Db.QueryRow(context.Background(), query, args).Scan(&rosterSlot.ID, &rosterSlot.CreatedAt, &rosterSlot.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return rosterSlot, nil

}
