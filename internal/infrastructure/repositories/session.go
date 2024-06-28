package repositories

import (
	"context"
	"onion/internal/domain/entities"
	"onion/internal/infrastructure/database"

	"github.com/jackc/pgx/v5"
)

type SessionPostgresRepository struct {
	db database.PostgresDB
}

func NewSessionPostgresRepository(db *database.PostgresDB) *SessionPostgresRepository {
	return &SessionPostgresRepository{
		db: *db,
	}
}

// Save saves a session to the database.
//
// It takes a context.Context and a pointer to an entities.Session as parameters.
// It returns an error.
func (r *SessionPostgresRepository) Save(ctx context.Context, session *entities.Session) error {
	query := `INSERT INTO "sessions" (identifier, user_id, valid_until) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, session.Identifier, session.User, session.ValidUntil)
	return err
}

// FindOne retrieves a session from the database based on its identifier.
//
// It takes a context.Context and a string identifier as parameters.
// It returns a pointer to entities.Session and an error.
func (r *SessionPostgresRepository) FindOne(ctx context.Context, identifier string) (*entities.Session, error) {
	query := `SELECT id, identifier, valid_until, agent, user, ip FROM "sessions" WHERE identifier = $1`
	row := r.db.QueryRow(ctx, query, identifier)
	return r.scan(row)
}

// FindMany retrieves multiple sessions from the PostgreSQL database based on the user ID.
//
// It takes a context.Context and a string user ID as parameters.
// It returns a slice of pointers to entities.Session and an error.
func (r *SessionPostgresRepository) FindMany(ctx context.Context, userId string) ([]*entities.Session, error) {
	query := `SELECT id, identifier, valid_until, agent, user, ip FROM "sessions" WHERE user = $1`
	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sessions := make([]*entities.Session, 0)
	for rows.Next() {
		session, err := r.scan(rows)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

// Delete deletes a session from the PostgreSQL database based on the provided identifier.
//
// Parameters:
// - ctx: The context.Context object for the request.
// - identifier: The identifier of the session to be deleted.
//
// Returns:
// - error: An error if the deletion operation fails, otherwise nil.
func (r *SessionPostgresRepository) Delete(ctx context.Context, identifier string) error {
	query := `DELETE FROM "sessions" WHERE identifier = $1`
	_, err := r.db.Exec(ctx, query, identifier)
	if err != nil {
		return err
	}
	return nil
}

func (r *SessionPostgresRepository) scan(row pgx.Row) (*entities.Session, error) {
	var session entities.Session
	err := row.Scan(&session.ID, &session.Identifier, &session.ValidUntil, &session.Agent, &session.User, &session.IP)
	if err != nil {
		return nil, err
	}
	return &session, nil
}
