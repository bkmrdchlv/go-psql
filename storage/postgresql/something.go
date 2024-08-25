package postgresql

import (
	"context"
	"fmt"

	"github.com/go-psql/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SomethingI struct {
	db *pgxpool.Pool
}

func NewSomethingI(db *pgxpool.Pool) *SomethingI {
	return &SomethingI{
		db: db,
	}
}

func (s *SomethingI) GetAllSomething(ctx context.Context, req *models.SomethingListRequest) (*models.SomethingList, error) {
	fmt.Println("GetAllSomething")
	fmt.Println(req)
	return nil, nil
}
func (s *SomethingI) GetByIDSomething(ctx context.Context, req *models.SomethingReadRequest) (*models.SomethingRead, error) {
	fmt.Println("GetByIDSomething")
	fmt.Println(req)
	return nil, nil
}
func (s *SomethingI) CreateSomething(ctx context.Context, req *models.SomethingCreate) (*models.SomethingCreateResponse, error) {
	fmt.Println("CreateSomething")
	fmt.Println(&req.SomethingArray)
	return nil, nil
}
func (s *SomethingI) UpdateSomething(ctx context.Context, req *models.SomethingUpdate) (*models.SomethingUpdateResponse, error) {
	fmt.Println("UpdateSomething")
	fmt.Println(*req)
	return nil, nil
}
func (s *SomethingI) PatchSomething(ctx context.Context, req *models.SomethingPatch) (*models.SomethingPatchResponse, error) {
	fmt.Println("PatchSomething")
	fmt.Println(req)
	return nil, nil
}
func (s *SomethingI) DeleteSomething(ctx context.Context, req *models.SomethingDelete) (*models.SomethingDeleteResponse, error) {
	fmt.Println("DeleteSomething")
	fmt.Println(req)
	return nil, nil
}
