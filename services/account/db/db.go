package account

// package github.com/ttydes/liby/services/account/protogen

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	pb "github.com/ttydes/liby/services/account/protogen"
)

type Account struct {
	Id        int32
	Fname     string
	Dob       int32
	CreatedAt time.Time
}

const (
	insertUserQuery = `
		INSERT INTO accounts (id, fname, dob)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE
		SET fname = EXCLUDED.fname, dob = EXCLUDED.dob;
	`

	updateUserQuery = `
		UPDATE accounts
		SET fname = $2, dob = $3
		WHERE id = $1;
	`

	deleteUserQuery = `
		DELETE FROM accounts
		WHERE id = $1;
	`

	getUserQuery = `
		SELECT id, fname, dob
		FROM accounts
		WHERE id = $1;
	`
)

func MkAccount(ctx context.Context, db *sqlx.DB, req *pb.MkAccountReq) (*pb.MkAccountResp, error) {
	switch req.Operation {
	case pb.Operation_CREATE:
		_, err := db.ExecContext(ctx, insertUserQuery, req.Id, req.Fname, req.Dob)
		if err != nil {
			return nil, fmt.Errorf("Failed to create user: %w", err)
		}
	case pb.Operation_UPDATE:
		_, err := db.ExecContext(ctx, updateUserQuery, req.Id, req.Fname, req.Dob)
		if err != nil {
			return nil, fmt.Errorf("Failed to update user: %w", err)
		}
	default:
		return nil, errors.New("Invalid operation")
	}

	return &pb.MkAccountResp{Id: req.Id}, nil
}

func UpdateAccount(ctx context.Context, db *sqlx.DB, req *pb.UpdateAccountReq) (*pb.UpdateAccountResp, error) {
	_, err := db.ExecContext(ctx, updateUserQuery, req.Id, req.Fname, req.Dob)
	if err != nil {
		return nil, fmt.Errorf("Failed to update user: %w", err)
	}

	return &pb.UpdateAccountResp{}, nil
}

func DeleteAccount(ctx context.Context, db *sqlx.DB, req *pb.DeleteAccountReq) (*pb.DeleteAccountResp, error) {
	_, err := db.ExecContext(ctx, deleteUserQuery, req.Id)
	if err != nil {
		return nil, fmt.Errorf("Failed to delete user: %w", err)
	}

	return &pb.DeleteAccountResp{}, nil
}

func GetAccount(ctx context.Context, db *sqlx.DB, req *pb.GetAccountReq) (*pb.GetAccountResp, error) {
	var resp pb.GetAccountResp
	err := db.GetContext(ctx, &resp, getUserQuery, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with ID %s not found", req.Id)
		}
		return nil, fmt.Errorf("Failed to get user: %w", err)
	}

	return &resp, nil
}

func MkClient(ctx context.Context) (*pgxpool.Pool, error) {
	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")

	if err := ValidateEnvs(name, host, user, pass); err != nil {
		return nil, err
	}
	// url := "postgres://tydes:pass@localhost:5432/lms_accounts"
	url := fmt.Sprintf("postgres://%s:%s@%s/%s", user, pass, host, name)
	client, err := pgxpool.Connect(ctx, url)

	if err != nil {
		return nil, fmt.Errorf("Failed to connect to %s", err)
	}

	return client, nil
}

func ValidateEnvs(envs ...string) error {
	var missing []string
	for _, x := range envs {
		if x == "" {
			missing = append(missing, x)
		}
	}
	if len(missing) > 0 {
		fmt.Fprintf(os.Stdout, "Missing required environment variables: %v\n", missing)
		return errors.New("missing required environment variables")
	}
	return nil
}
