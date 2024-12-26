package branch

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"

	pb "github.com/ttydes/liby/services/branch/protogen"
)

type Branch struct {
	Id       int32
	Name     string
	Address  string
	City     string
	Postcode string
	Country  string
}

func MkBranch(ctx context.Context, db *pgxpool.Pool, req *pb.MkBranchReq) (*pb.MkBranchResp, error) {
	query := `
		INSERT INTO branches (name, address, city, postcode, country)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := db.Exec(ctx, query, req.Name, req.Address, req.City, req.Postcode, req.Country)
	if err != nil {
		return nil, fmt.Errorf("Failed to create branch: %w", err)
	}

	return &pb.MkBranchResp{Success: true, Message: "Branch created successfully"}, nil
}

func GetBranch(ctx context.Context, db *pgxpool.Pool, req *pb.GetBranchReq) (*pb.GetBranchResp, error) {
	query := `
		SELECT id, name, address, city, postcode, country
		FROM branches
		WHERE id = $1;
	`
	resp := &pb.GetBranchResp{Branch: &pb.Branch{}}

	if err := db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Branch.Id,
		&resp.Branch.Name,
		&resp.Branch.Address,
		&resp.Branch.City,
		&resp.Branch.Postcode,
		&resp.Branch.Country,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Branch with Id %d not found", req.Id)
		}
		return nil, fmt.Errorf("Failed to get branch: %w", err)
	}

	return resp, nil
}

func UpdateBranch(ctx context.Context, db *pgxpool.Pool, req *pb.UpdateBranchReq) (*pb.UpdateBranchResp, error) {
	query := `
		 UPDATE branches
			SET 
				name = CASE WHEN $2 <> '' THEN $2 ELSE name END,
				address = CASE WHEN $3 <> '' THEN $3 ELSE address END,
				city = CASE WHEN $4 <> '' THEN $4 ELSE city END,
				postcode = CASE WHEN $5 <> '' THEN $5 ELSE postcode END,
				country = CASE WHEN $6 <> '' THEN $6 ELSE country END
			WHERE id = $1;
	`
	_, err := db.Exec(ctx, query, req.Id, req.Name, req.Address, req.City, req.Postcode, req.Country)
	if err != nil {
		return nil, fmt.Errorf("Failed to update branch: %w", err)
	}

	return &pb.UpdateBranchResp{Success: true, Message: "Branch updated successfully"}, nil
}

func DeleteBranch(ctx context.Context, db *pgxpool.Pool, req *pb.DeleteBranchReq) (*pb.DeleteBranchResp, error) {
	query := `
		DELETE FROM branches
		WHERE id = $1;
	`
	_, err := db.Exec(ctx, query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("Failed to delete branch: %w", err)
	}

	return &pb.DeleteBranchResp{Success: true, Message: "Branch deleted successfully"}, nil
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
