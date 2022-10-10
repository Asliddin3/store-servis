package postgres

import (
	pb "github.com/Asliddin3/store-servis/genproto"
	"github.com/jmoiron/sqlx"
)

type storeRepo struct {
	db *sqlx.DB
}

func NewStoreRepo(db *sqlx.DB) *storeRepo {
	return &storeRepo{db: db}
}
func (r *storeRepo) Create(req *pb.StoreRequest) (*pb.StoreResponse, error) {
	storeResp := pb.StoreResponse{}
	err := r.db.QueryRow(`
	insert into stores(name) values($1) returning id ,name`, req.Name).Scan(
		&storeResp.Id, &storeResp.Name)
	if err != nil {
		return &pb.StoreResponse{}, err
	}
	addresses := []*pb.AddressResp{}
	for _, address := range req.Addresses {
		addressResp := pb.AddressResp{}
		err = r.db.QueryRow(`
		insert into addresses(district,street) values($1,$2)
		returning id,district,street`, address.District, address.Street).Scan(
			&addressResp.Id, &addressResp.District, &addressResp.Street,
		)
		addresses = append(addresses, &addressResp)
	}
	storeResp.Addresses = addresses
	return &storeResp, nil
}
func (r *storeRepo) GetStore(req *pb.GetStoreInfoById) (*pb.StoreResponse, error)
