package service

import (
	"context"

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

func (s *StoreService) Create(ctx context.Context, req *pb.StoreRequest) (*pb.StoreResponse, error) {
	store, err := s.storage.Store().Create(req)
	if err != nil {
		s.logger.Error("error while creating store", l.Any("error creating store", err))
		return &pb.StoreResponse{}, status.Error(codes.Internal, "something went wrong")
	}
	return store, nil
}

func (s *StoreService) GetStore(ctx context.Context, req *pb.GetStoreInfoById) (*pb.StoreResponse, error) {
	store, err := s.storage.Store().GetStore(req)
	if err != nil {
		s.logger.Error("error while geting store", l.Any("error geting stroe", err))
		return &pb.StoreResponse{}, status.Error(codes.Internal, "something went wrong")
	}
	return store, nil
}
