package storage

import (
	"context"

	"github.com/go-psql/models"
)

type Storage interface {
	SomethingI() SomethingI

	CloseDB()
}

type SomethingI interface {
	GetAllSomething(ctx context.Context, req *models.SomethingListRequest) (*models.SomethingList, error)
	GetByIDSomething(ctx context.Context, req *models.SomethingReadRequest) (*models.SomethingRead, error)
	CreateSomething(ctx context.Context, req *models.SomethingCreate) (*models.SomethingCreateResponse, error)
	UpdateSomething(ctx context.Context, req *models.SomethingUpdate) (*models.SomethingUpdateResponse, error)
	PatchSomething(ctx context.Context, req *models.SomethingPatch) (*models.SomethingPatchResponse, error)
	DeleteSomething(ctx context.Context, req *models.SomethingDelete) (*models.SomethingDeleteResponse, error)
}
