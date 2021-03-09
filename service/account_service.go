package service

import (
	"context"
	"log"

	repo "github.com/heru-wijaya/go-grpc-skeleton/repository"
	"github.com/segmentio/ksuid"
)

// AccountService is as an interface this service
type AccountService interface {
	PostAccount(ctx context.Context, name string) (*repo.Account, error)
	GetAccount(ctx context.Context, id string) (*repo.Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]repo.Account, error)
}

type accountService struct {
	repository repo.AccountRepository
}

// NewService is for wrapper this service
func NewService(r repo.AccountRepository) AccountService {
	return &accountService{r}
}

func (s *accountService) PostAccount(ctx context.Context, name string) (*repo.Account, error) {
	log.Println("service.account_service postAccount begin")
	a := &repo.Account{
		Name: name,
		ID:   ksuid.New().String(),
	}
	if err := s.repository.PutAccount(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *accountService) GetAccount(ctx context.Context, id string) (*repo.Account, error) {
	log.Println("service.account_service GetAccount begin")
	return s.repository.GetAccountByID(ctx, id)
}

func (s *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]repo.Account, error) {
	log.Println("service.account_service GetAccounts begin")
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	return s.repository.ListAccounts(ctx, skip, take)
}
