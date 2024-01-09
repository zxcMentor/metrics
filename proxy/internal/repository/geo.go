package repository

import (
	"database/sql"
	"metricsProm/proxy/internal/metrics"

	"sync"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
)

type GeoRepo interface {
	GetAddr(id int) (string, error)
	GetIDfrHist(request string) (int, error)
	GetAddrID(id int) (int, error)
	SaveSearchHist(request string) (int, error)
	SaveAddr(addr string) (int, error)
	SaveHistSearchAddr(searchHistID, addrID int) error
}

type PostgreGeoRepo struct {
	db     *sqlx.DB
	sqlBlb sq.StatementBuilderType
	sync.Mutex
}

func NewPostgreGeoRepo(db *sqlx.DB) PostgreGeoRepo {

	return PostgreGeoRepo{
		db:     db,
		sqlBlb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *PostgreGeoRepo) GetAddr(id int) (string, error) {
	query := `SELECT data FROM addresses WHERE id = $1`

	addr := ""
	err := r.db.QueryRow(query, id).Scan(addr)
	if err != nil {
		if err == sql.ErrNoRows {
			return addr, nil
		}

	}

	return addr, nil
}

func (r *PostgreGeoRepo) GetIDfrHist(request string) (int, error) {

	query := `SELECT id FROM search_history WHERE similarity(query, $1) >= 0.7`
	id := 0
	err := r.db.QueryRow(query, request).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return id, nil
		}

	}

	return id, err

}

func (r *PostgreGeoRepo) GetAddrID(id int) (int, error) {
	q := `SELECT address_id FROM history_search_address WHERE search_history_id = $1`

	addrId := 0
	err := r.db.QueryRow(q, id).Scan(&addrId)
	if err != nil {
		if err == sql.ErrNoRows {
			return addrId, nil
		}

	}

	return addrId, err

}

func (r *PostgreGeoRepo) SaveSearchHist(request string) (int, error) {
	query := `INSERT INTO search_history (query) VALUES ($1) RETURNING id`

	id := 0
	err := r.db.QueryRow(query, request).Scan(&id)

	return id, err
}

func (r *PostgreGeoRepo) SaveAddr(addr string) (int, error) {

	startTime := time.Now()
	query := `INSERT INTO address (data) VALUES ($1) RETURNING id`

	id := 0
	err := r.db.QueryRow(query, addr).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return id, nil
		}

	}

	elapsedTime := time.Since(startTime)
	metrics.NewTimeDB().With(prometheus.
		Labels{"method": "POST", "path": "/api/address/search"}).
		Observe(elapsedTime.Seconds())

	return id, err
}

func (r *PostgreGeoRepo) SaveHistSearchAddr(searchHistID, addrID int) error {
	query := `INSERT INTO history_search_address (search_history_id, address_id) VALUES ($1, $2)`

	_, err := r.db.Exec(query, searchHistID, addrID)
	if err != nil {
		return err
	}
	return nil
}
