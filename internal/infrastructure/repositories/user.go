package repositories

import (
	"context"
	"onion/internal/domain/entities"
	"onion/internal/infrastructure/database"

	"github.com/jackc/pgx/v5"
)

type UserPostgresRepository struct {
	db database.PostgresDB
}

func NewUserPostgresRepository(db *database.PostgresDB) *UserPostgresRepository {
	return &UserPostgresRepository{
		db: *db,
	}
}

// Save saves a user to the PostgreSQL database.
//
// It takes a context.Context and a *entities.User as parameters.
// It returns a *entities.User and an error.
func (r *UserPostgresRepository) Save(ctx context.Context, user *entities.User) (*entities.User, error) {
	query := `INSERT INTO "users" (username, email, blocked, avatar, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, $6)
    ON CONFLICT (email)
    DO UPDATE SET
        updated_at = EXCLUDED.updated_at
	RETURNING 
		id, username, email, blocked, avatar, created_at, updated_at`
	row := r.db.QueryRow(ctx, query, user.Username, user.Email, user.Blocked, user.Avatar, user.CreatedAt, user.UpdatedAt)
	return r.scan(row)
}

// FindById retrieves a user from the PostgreSQL database based on their email.
//
// It takes a context.Context and a string email as parameters.
// It returns a pointer to entities.User and an error.
func (r *UserPostgresRepository) FindById(ctx context.Context, email string) (*entities.User, error) {
	query := `SELECT id, username, email, blocked, avatar, created_at, updated_at FROM "users" WHERE email = $1`
	row := r.db.QueryRow(ctx, query, email)
	return r.scan(row)
}

// FindMany retrieves multiple users from the PostgreSQL database based on their IDs.
//
// It takes a context.Context and a slice of string IDs as parameters.
// It returns a slice of pointers to entities.User and an error.
func (r *UserPostgresRepository) FindMany(ctx context.Context, ids []string) ([]*entities.User, error) {
	query := `
	SELECT id, username, email, blocked, avatar, created_at, updated_at
	FROM "users"
	WHERE id = ANY($1)
	`
	rows, err := r.db.Query(ctx, query, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Blocked, &user.Avatar, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// FindWithValidOTP retrieves a user from the PostgreSQL database along with a valid OTP.
//
// It takes a context.Context and a string email and code as parameters.
// It returns a pointer to entities.User and an error.
func (r *UserPostgresRepository) FindWithValidOTP(ctx context.Context, email, code string) (*entities.User, error) {
	query := `SELECT id, username, email, blocked, avatar, created_at, updated_at FROM "users" WHERE (SELECT COUNT(*) FROM "otps" WHERE code = $1 AND email = $2 AND used = false AND valid_until > NOW()) > 0`
	row := r.db.QueryRow(ctx, query, email, code)
	return r.scan(row)
}

func (r *UserPostgresRepository) scan(row pgx.Row) (*entities.User, error) {
	var user entities.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Blocked, &user.Avatar, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
