//go:generate protoc --go_out=./model/pb --go_opt=module=github.com/heru-wijaya/go-grpc-skeleton/model/pb  ./model/account.proto
//go:generate protoc --go-grpc_out=./model/pb --go-grpc_opt=module=github.com/heru-wijaya/go-grpc-skeleton/model/pb ./model/account.proto
package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/heru-wijaya/go-grpc-skeleton/model/pb"
	service "github.com/heru-wijaya/go-grpc-skeleton/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	accountService service.AccountService
	pb.UnimplementedAccountServiceServer
}

// ListenGRPC for register server to grpc
func ListenGRPC(s service.AccountService, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterAccountServiceServer(serv, &grpcServer{accountService: s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	a, err := s.accountService.PostAccount(ctx, r.Name)
	if err != nil {
		return nil, err
	}
	return &pb.PostAccountResponse{Account: &pb.Account{
		Id:   a.ID,
		Name: a.Name,
	}}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	a, err := s.accountService.GetAccount(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountResponse{
		Account: &pb.Account{
			Id:   a.ID,
			Name: a.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	res, err := s.accountService.GetAccounts(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	accounts := []*pb.Account{}
	for _, p := range res {
		accounts = append(
			accounts,
			&pb.Account{
				Id:   p.ID,
				Name: p.Name,
			},
		)
	}
	return &pb.GetAccountsResponse{Accounts: accounts}, nil
}
