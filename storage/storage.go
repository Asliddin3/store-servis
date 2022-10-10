package storage

import (
	"github.com/Asliddin3/store-servis/storage/postgres"
	"github.com/Asliddin3/store-servis/storage/repo"
	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Store() repo.StoreStorageI
}

type storagePg struct {
	db          *sqlx.DB
	productRepo repo.StoreStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		storeRepo: postgres.NewStoreRepo(db),
	}
}
func (s storagePg) Product() repo.StoreStorageI {
	return s.productRepo
}
