package repository

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Account type struct
type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// AccountRepository is as an interface this repository
type AccountRepository interface {
	Close()
	PostAccount(ctx context.Context, a Account) error
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type postgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository is for wrapper postgres sql
func NewPostgresRepository(url string) (AccountRepository, error) {
	log.Println("repository.account_repository NewPostgresRepository begin")
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	log.Println("repository.account_repository Close begin")
	r.db.Close()
}

func (r *postgresRepository) Ping() error {
	log.Println("repository.account_repository Ping begin")
	return r.db.Ping()
}

func (r *postgresRepository) PostAccount(ctx context.Context, a Account) error {
	log.Println("repository.account_repository PostAccount begin")
	_, err := r.db.ExecContext(ctx, "INSERT INTO accounts(id, name) VALUES($1, $2)", a.ID, a.Name)
	return err
}

func (r *postgresRepository) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	log.Println("repository.account_repository GetAccountByID begin")
	row := r.db.QueryRowContext(ctx, "SELECT id, name FROM accounts WHERE id = $1", id)
	a := &Account{}
	if err := row.Scan(&a.ID, &a.Name); err != nil {
		return nil, err
	}
	return a, nil
}

func (r *postgresRepository) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	log.Println("repository.account_repository ListAccounts begin")
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, name FROM accounts ORDER BY id DESC OFFSET $1 LIMIT $2",
		skip,
		take,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []Account{}
	for rows.Next() {
		a := &Account{}
		if err = rows.Scan(&a.ID, &a.Name); err == nil {
			accounts = append(accounts, *a)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}
