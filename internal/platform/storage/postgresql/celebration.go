package postgresql

import "time"

const (
	sqlCelebrationTable string = "celebrations"
)

type sqlCelebration struct {
	ID     string    `db:"id"`
	Date   time.Time `db:"date"`
	Name   string    `db:"name"`
	Status string    `db:"status"`
	Type   string    `db:"type"`
}
