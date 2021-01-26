package main

import (
	"context"

	pb "github.com/heru-wijaya/go-grpc-skeleton/model/pb"
	repo "github.com/heru-wijaya/go-grpc-skeleton/repository"
	"google.golang.org/grpc"
)

// Client type struct with service and connection
type Client struct {
	conn    *grpc.ClientConn
	service pb.AccountServiceClient
}

// NewClient for register new client
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewAccountServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close for closing the connection
func (c *Client) Close() {
	c.conn.Close()
}

// PostAccount for save new account
func (c *Client) PostAccount(ctx context.Context, name string) (*repo.Account, error) {
	r, err := c.service.PostAccount(
		ctx,
		&pb.PostAccountRequest{Name: name},
	)
	if err != nil {
		return nil, err
	}
	return &repo.Account{
		ID:   r.Account.Id,
		Name: r.Account.Name,
	}, nil
}

// GetAccount for get account by id
func (c *Client) GetAccount(ctx context.Context, id string) (*repo.Account, error) {
	r, err := c.service.GetAccount(
		ctx,
		&pb.GetAccountRequest{Id: id},
	)
	if err != nil {
		return nil, err
	}
	return &repo.Account{
		ID:   r.Account.Id,
		Name: r.Account.Name,
	}, nil
}

// GetAccounts for get list of account
func (c *Client) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]repo.Account, error) {
	r, err := c.service.GetAccounts(
		ctx,
		&pb.GetAccountsRequest{
			Skip: skip,
			Take: take,
		},
	)
	if err != nil {
		return nil, err
	}
	accounts := []repo.Account{}
	for _, a := range r.Accounts {
		accounts = append(accounts, repo.Account{
			ID:   a.Id,
			Name: a.Name,
		})
	}
	return accounts, nil
}
