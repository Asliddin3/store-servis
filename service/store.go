package service

import (
	"context"
	"fmt"

	pb "github.com/Asliddin3/store-servis/genproto"
	l "github.com/Asliddin3/store-servis/pkg/logger"
	"github.com/Asliddin3/store-servis/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StoreService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewStoreService(db *sqlx.DB, log l.Logger) *StoreService {
	return &StoreService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *StoreService) Create(ctx context.Context,req *pb.StoreRequest) (*pb.StoreResponse,error){
	store,err:=s.storage.Product()
}


func (s *ProductService) Update(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	product, err := s.storage.Product().Update(req)
	if err != nil {
		s.logger.Error("error while updating product", l.Any("error updating product", err))
		return &pb.Product{}, status.Error(codes.Internal, "something went wrong")
	}
	return product, nil
}


