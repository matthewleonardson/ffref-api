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
