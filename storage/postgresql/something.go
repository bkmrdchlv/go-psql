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

type AnyData struct {
	SomethingInt           interface{} `json:"int"`
	SomethingString        interface{} `json:"string"`
	SomethingBool          interface{} `json:"bool"`
	SomethingFloat         interface{} `json:"float"`
	SomethingArrray        interface{} `json:"arrray"`
	SomethingMap           interface{} `json:"map"`
	SomethingStruct        interface{} `json:"struct"`
	SomethingPointer       interface{} `json:"pointer"`
	SomethingInterface     interface{} `json:"interface"`
	SomethingSlice         interface{} `json:"slice"`
	SomethingSliceOfStruct interface{} `json:"sliceOfStruct"`
	SomethingId            interface{} `json:"id"`
	SomethingTimestamp     interface{} `json:"timestamp"`
}

func (s *SomethingI) GetAllSomething(ctx context.Context, req *models.SomethingListRequest) (*models.SomethingList, error) {
	fmt.Println("GetAllSomething")
	fmt.Println(req)

	var somethingReadArray []models.SomethingRead

	query := `select  
	id,
	int,
	string, 
	map ,
	struct ,
	pointer,
	interface, 
	slice, 
	sliceofstruct , 
	timestamp 
	from something;`

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var something AnyData

		err := rows.Scan(
			&something.SomethingId,
			&something.SomethingInt,
			&something.SomethingString,
			&something.SomethingMap,
			&something.SomethingStruct,
			&something.SomethingPointer,
			&something.SomethingInterface,
			&something.SomethingSlice,
			&something.SomethingSliceOfStruct,
			&something.SomethingTimestamp,
		)

		if err != nil {
			return nil, err
		}
		fmt.Println(something)

		// somethingReadArray = append(somethingReadArray, something)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &models.SomethingList{
		Somethings: somethingReadArray,
	}, nil

}
func (s *SomethingI) GetByIDSomething(ctx context.Context, req *models.SomethingReadRequest) (*models.SomethingRead, error) {
	fmt.Println("GetByIDSomething")
	fmt.Println(req)
	return nil, nil
}
func (s *SomethingI) CreateSomething(ctx context.Context, req *models.SomethingCreate) (*models.SomethingCreateResponse, error) {
	fmt.Println("CreateSomething")
	fmt.Println(&req.SomethingArrray)
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
