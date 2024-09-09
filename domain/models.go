package domain

import "time"

type Team struct {
	ID         *int
	TeamName   *string
	Manager    *string
	YearJoined *int
	YearLeft   *int
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

type Player struct {
	ID         *int
	PlayerName *string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

type RosterSlot struct {
	ID              *int
	PlayerID        *int
	TeamID          *int
	Week            *int
	Year            *int
	Position        *string
	Benched         bool
	ProjectedPoints *float32
	ActualPoints    *float32
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}
