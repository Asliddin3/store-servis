package repo

import (
	pb "github.com/Asliddin3/store-servis/genproto"
)

type StoreStorageI interface{
	Create(*pb.StoreRequest)(*pb.StoreResponse)
	GetStore(*pb.GetStoreInfoById)(*pb.StoreResponse)
}