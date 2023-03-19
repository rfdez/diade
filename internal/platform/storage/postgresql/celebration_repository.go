package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/huandu/go-sqlbuilder"
	diade "github.com/rfdez/diade/internal"
	"github.com/rfdez/diade/kit/errors"
)

type celebrationRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewCelebrationRepository creates a new celebration repository.
func NewCelebrationRepository(db *sql.DB, dbTimeout time.Duration) diade.CelebrationRepository {
	return &celebrationRepository{db: db, dbTimeout: dbTimeout}
}

// SearchByDate implements the CelebrationRepository interface.
func (r *celebrationRepository) SearchByDate(ctx context.Context, date diade.CelebrationDate) ([]diade.Celebration, error) {
	celebrationSQLStruct := sqlbuilder.NewStruct(new(sqlCelebration)).For(sqlbuilder.PostgreSQL)

	sb := celebrationSQLStruct.SelectFrom(sqlCelebrationTable)
	sb.Where(sb.E("date", date.String()))

	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, errors.New("error querying celebrations")
	}
	defer rows.Close()

	var celebrations []diade.Celebration
	for rows.Next() {
		var sqlC sqlCelebration
		if err := rows.Scan(celebrationSQLStruct.Addr(&sqlC)...); err != nil {
			return nil, errors.New("error scanning celebrations")
		}

		celebration, err := diade.NewCelebration(sqlC.ID, sqlC.Date.Format("2006-01-02"), sqlC.Name, sqlC.Status, sqlC.Type)
		if err != nil {
			return nil, errors.New("error converting sql celebration to celebration")
		}

		celebrations = append(celebrations, celebration)
	}

	return celebrations, nil
}
