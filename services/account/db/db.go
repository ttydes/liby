package account

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"

	pb "github.com/ttydes/liby/services/account/protogen"
)

type Account struct {
	Id        int32
	Fname     string
	Dob       int32
	CreatedAt time.Time
}

func MkAccount(ctx context.Context, db *pgxpool.Pool, req *pb.MkAccountReq) (*pb.MkAccountResp, error) {
	query := `
		INSERT INTO accounts (id, fname, dob)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE
		SET fname = EXCLUDED.fname, dob = EXCLUDED.dob;
	`
	_, err := db.Exec(ctx, query, req.Id, req.Fname, req.Dob)
	if err != nil {
		return nil, fmt.Errorf("Failed to create user: %w", err)
	}

	return &pb.MkAccountResp{Id: req.Id}, nil
}

func UpdateAccount(ctx context.Context, db *pgxpool.Pool, req *pb.UpdateAccountReq) (*pb.UpdateAccountResp, error) {
	query := `
		UPDATE accounts
		SET fname = $2, dob = $3
		WHERE id = $1;
	`
	_, err := db.Exec(ctx, query, req.Id, req.Fname, req.Dob)
	if err != nil {
		return nil, fmt.Errorf("Failed to update user: %w", err)
	}

	return &pb.UpdateAccountResp{}, nil
}

func DeleteAccount(ctx context.Context, db *pgxpool.Pool, req *pb.DeleteAccountReq) (*pb.DeleteAccountResp, error) {
	query := `
		DELETE FROM accounts
		WHERE id = $1;
	`
	_, err := db.Exec(ctx, query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("Failed to delete user: %w", err)
	}

	return &pb.DeleteAccountResp{}, nil
}

func GetAccount(ctx context.Context, db *pgxpool.Pool, req *pb.GetAccountReq) (*pb.GetAccountResp, error) {
	query := `
		SELECT id, fname, dob
		FROM accounts
		WHERE id = $1;
	`
	var resp pb.GetAccountResp
	err := db.QueryRow(ctx, query, req.Id).Scan(&resp.Id, &resp.Fname, &resp.Dob)
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
	url := fmt.Sprintf("postgres://%s:%s@%s/%s", user, pass, host, name)

	db, err := pgxpool.Connect(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	return db, nil
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
