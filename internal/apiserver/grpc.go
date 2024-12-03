package apiserver

import (
	"context"
	"log"
	"net"

	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/auth"
	pb "github.com/ividernvi/algohub-forum/internal/apiserver/service" // 替换为实际路径
	"github.com/ividernvi/algohub-forum/internal/apiserver/store"
	"github.com/ividernvi/algohub-forum/internal/apiserver/store/mysql"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/internal/model/options"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type grpcserver struct {
	pb.UnimplementedVerifyServer
}

var (
	db  store.Factory
	srv pb.Service
)

func (s *grpcserver) VerifyUser(ctx context.Context, in *pb.AccPasswd) (out *pb.VerifiedResult, err error) {
	logrus.Printf("VerifyUser called")
	var user *model.User
	if user, err = srv.Users().GetByAccount(ctx, in.Username, model.GetOptions{}); err != nil {
		logrus.Printf("grpc: get user failed: \033[1;31;40m%s\033[0m\n", err)
		return &pb.VerifiedResult{
			IsVerified: false,
		}, nil
	}

	if auth.Compare(user.Password, in.Password) != nil {
		logrus.Printf("verified failed: %s", err)
		return &pb.VerifiedResult{
			IsVerified: false,
		}, nil
	}

	return &pb.VerifiedResult{
		IsVerified: true,
	}, nil

}

func runGRPC() {
	go func() {
		db, _ = mysql.GetMySQLFactoryOr(*options.NewOptions().MySQLOptions)
		srv = pb.NewService(db)
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			logrus.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterVerifyServer(s, &grpcserver{})
		logrus.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}
