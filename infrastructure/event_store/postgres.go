package event_store

import (
	"database/sql"
	"github.com/avgalaida/library/domain"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db,
	}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()
}

func (r *PostgresRepository) InsertAggregate(a domain.BasedAggregate) {
	r.db.Exec("INSERT INTO aggregates(id,meta,created_at) VALUES($1,$2,$3)",
		a.ID,
		a.Meta,
		a.CreatedAt,
	)
}

func (r *PostgresRepository) GetAggregate(id string) domain.BasedAggregate {
	aggregateRows, _ := r.db.Query("SELECT * FROM aggregates WHERE id = $1", id)
	defer aggregateRows.Close()

	var aggregate domain.BasedAggregate
	for aggregateRows.Next() {
		aggregateRows.Scan(&aggregate.ID, &aggregate.Meta, &aggregate.CreatedAt)
	}
	return aggregate
}

func (r *PostgresRepository) UpdateAggregateRevision(id string) {
	r.db.Exec("UPDATE aggregates SET meta=meta+1 WHERE id = $1", id)
}

func (r *PostgresRepository) GetAggregateVersion(id string, version string) (domain.BasedAggregate, []domain.BasedEvent) {
	var aggregate domain.BasedAggregate
	aggregateRows, _ := r.db.Query("SELECT * FROM aggregates WHERE id = $1", id)
	defer aggregateRows.Close()
	for aggregateRows.Next() {
		aggregateRows.Scan(&aggregate.ID, &aggregate.Meta, &aggregate.CreatedAt)
	}

	var events []domain.BasedEvent
	eventRows, _ := r.db.Query("SELECT * FROM events WHERE aggregate_id = $1 and revision <= $2", aggregate.ID, version)
	for eventRows.Next() {
		event := domain.BasedEvent{}
		eventRows.Scan(
			&event.ID,
			&event.AggregateID,
			&event.CreatedAt,
			&event.UserID,
			&event.Revision,
			&event.Data,
			&event.Type,
		)
		events = append(events, event)
	}
	return aggregate, events
}

func (r *PostgresRepository) InsertEvent(e domain.BasedEvent) {
	r.db.Exec("INSERT INTO events(id, aggregate_id, created_at, user_id, revision, delta_data, event_type) VALUES($1,$2,$3,$4,$5,$6,$7)",
		e.ID,
		e.AggregateID,
		e.CreatedAt,
		e.UserID,
		e.Revision,
		e.Data,
		e.Type,
	)
}

func (r *PostgresRepository) GetAll() map[domain.BasedAggregate][]domain.BasedEvent {
	aggregateRows, _ := r.db.Query("SELECT * FROM aggregates")
	defer aggregateRows.Close()

	var aggregates []domain.BasedAggregate
	for aggregateRows.Next() {
		aggregate := domain.BasedAggregate{}
		aggregateRows.Scan(&aggregate.ID, &aggregate.Meta, &aggregate.CreatedAt)
		aggregates = append(aggregates, aggregate)
	}

	aemap := make(map[domain.BasedAggregate][]domain.BasedEvent)

	for _, aggregate := range aggregates {
		var events []domain.BasedEvent
		eventRows, _ := r.db.Query("SELECT * FROM events WHERE aggregate_id = $1", aggregate.ID)
		for eventRows.Next() {
			event := domain.BasedEvent{}
			eventRows.Scan(
				&event.ID,
				&event.AggregateID,
				&event.CreatedAt,
				&event.UserID,
				&event.Revision,
				&event.Data,
				&event.Type,
			)
			events = append(events, event)
		}
		aemap[aggregate] = events
	}

	return aemap
}
