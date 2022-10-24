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
	celebrationSqlStruct := sqlbuilder.NewStruct(new(sqlCelebration)).For(sqlbuilder.PostgreSQL)

	sb := celebrationSqlStruct.SelectFrom(sqlCelebrationTable)
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
		var sqlCelebration sqlCelebration
		if err := rows.Scan(celebrationSqlStruct.Addr(&sqlCelebration)...); err != nil {
			return nil, errors.New("error scanning celebrations")
		}

		celebration, err := diade.NewCelebration(sqlCelebration.ID, sqlCelebration.Date.Format("2006-01-02"), sqlCelebration.Name, sqlCelebration.Status, sqlCelebration.Type)
		if err != nil {
			return nil, errors.New("error converting sql celebration to celebration")
		}

		celebrations = append(celebrations, celebration)
	}

	return celebrations, nil
}
