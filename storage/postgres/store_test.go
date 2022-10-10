package postgres

import (
	"testing"

	"github.com/Asliddin3/store-servis/config"
	pb "github.com/Asliddin3/store-servis/genproto"
	"github.com/Asliddin3/store-servis/pkg/db"
	"github.com/Asliddin3/store-servis/storage/repo"
	"github.com/stretchr/testify/suite"
)

type ProductSuiteTest struct {
	suite.Suite
	CleanUpFunc func()
	Repository repo.StoreStorageI
}

func (suite *ProductSuiteTest) SetupSuite() {
	pgPool, cleanupfunc := db.ConnectToDbForSuite(config.Load())

	suite.Repository = NewStoreRepo(pgPool)
	suite.CleanUpFunc = cleanupfunc
}

func (s *ProductSuiteTest) TestProductCrud(){
	storeCreate:=pb.StoreRequest{
		Name: "Malika bozor",
		Addresses: []*pb.Address{
			{
				District: "Yakasaroy",
				Street: "Shota Rustaveli",
			},{
				District: "Mirzo Ulugbek",
				Street: "Farhadskiy",
			},
		},
	}
	storeResp,err:=s.Repository.Create(&storeCreate)
	s.Nil(err)
	s.NotNil(storeResp)
	storeRequest,err:=s.Repository.GetStore(&pb.GetStoreInfoById{
		Id: storeResp.Id,
	})
	s.Nil(err)
	s.NotNil(storeRequest)
}

func (suite *ProductSuiteTest) TearDownSuite() {
	suite.CleanUpFunc()
}


func TestProductRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProductSuiteTest))
}
