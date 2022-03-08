package postgresql

import (
	"github.com/lib/pq"

	"github.com/benchttp/server/benchttp"
)

func (c ComputedStatsService) ListMetadataByUserID(userID int) ([]benchttp.Metadata, error) {
	metadataList := []benchttp.Metadata{}

	stmt, err := c.db.Prepare(`SELECT tag, finished_at FROM metadata WHERE user_id = $1 ORDER BY finished_at DESC`)
	if err != nil {
		return []benchttp.Metadata{}, ErrPreparingStmt
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return []benchttp.Metadata{}, ErrExecutingPreparedStmt
	}
	defer rows.Close()

	for rows.Next() {
		metadata := benchttp.Metadata{}
		err = rows.Scan(
			&metadata.Tag,
			&metadata.FinishedAt,
		)
		if err != nil {
			return nil, err
		}
		metadataList = append(metadataList, metadata)
	}

	return metadataList, nil
}

func (c ComputedStatsService) FindComputedStatsByID(metadataID int) (benchttp.ComputedStats, error) {
	computedStats := benchttp.ComputedStats{}

	stmt := `SELECT
			m.tag,
			m.finished_at,
			c.code_1xx,
			c.code_2xx,
			c.code_3xx,
			c.code_4xx,
			c.code_5xx,
			t.min,
			t.max,
			t.mean,
			t.median,
			t.variance,
			t.deciles
		FROM public.metadata AS m
		INNER JOIN public.codestats AS c ON c.metadata_id = m.id
		INNER JOIN public.timestats AS t ON t.metadata_id = m.id
		WHERE m.id = $1
		ORDER BY m.finished_at DESC`

	row := c.db.QueryRow(stmt, metadataID)
	err := row.Scan(
		&computedStats.Metadata.Tag,
		&computedStats.Metadata.FinishedAt,
		&computedStats.Codestats.Code1xx,
		&computedStats.Codestats.Code2xx,
		&computedStats.Codestats.Code3xx,
		&computedStats.Codestats.Code4xx,
		&computedStats.Codestats.Code5xx,
		&computedStats.Timestats.Min,
		&computedStats.Timestats.Max,
		&computedStats.Timestats.Mean,
		&computedStats.Timestats.Median,
		&computedStats.Timestats.Variance,
		(*pq.Float64Array)(&computedStats.Timestats.Deciles),
	)
	if err != nil {
		return computedStats, ErrScanningRows
	}

	return computedStats, nil
}
