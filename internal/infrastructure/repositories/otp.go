package repositories

import (
	"context"
	"onion/internal/domain/entities"
	"onion/internal/infrastructure/database"
)

type OtpPostgresRepository struct {
	db database.PostgresDB
}

func NewOtpPostgresRepository(db *database.PostgresDB) *OtpPostgresRepository {
	return &OtpPostgresRepository{
		db: *db,
	}
}

// Save saves an OTP entity to the PostgreSQL database.
//
// Parameters:
// - ctx: The context.Context used to control the execution of the function.
// - otp: A pointer to an entities.OTP representing the OTP entity to be saved.
//
// Returns:
// - error: An error if the operation fails, otherwise nil.
func (r *OtpPostgresRepository) Save(ctx context.Context, otp *entities.OTP) error {
	query := `INSERT INTO "otps" (code, user_id, valid_until) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, otp.Code, otp.User, otp.ValidUntil)
	return err
}

// Delete deletes an OTP entity from the PostgreSQL database based on the provided code and user.
//
// Parameters:
// - ctx: The context.Context used to control the execution of the function.
// - code: The code of the OTP entity to be deleted.
// - user: The user associated with the OTP entity.
//
// Returns:
// - error: An error if the deletion operation fails, otherwise nil.
func (r *OtpPostgresRepository) Delete(ctx context.Context, code string, user string) error {
	query := `DELETE FROM "otps" WHERE code = $1 AND email = $2`
	_, err := r.db.Exec(ctx, query, code, user)
	if err != nil {
		return err
	}
	return nil
}
