package models

// Something is a struct that represents a something
type Something struct {
	SomethingInt    int               `json:"int"`
	SomethingString string            `json:"string"`
	SomethingBool   bool              `json:"bool"`
	SomethingFloat  float64           `json:"float"`
	SomethingArray  []string          `json:"array"`
	SomethingMap    map[string]string `json:"map"`
	SomethingStruct struct {
		StructInt    int    `json:"int"`
		StructString string `json:"string"`
	} `json:"struct"`
	SomethingPointer       *string     `json:"pointer"`
	SomethingInterface     interface{} `json:"interface"`
	SomethingSlice         []string    `json:"slice"`
	SomethingSliceOfStruct []struct {
		StructInt    int    `json:"int"`
		StructString string `json:"string"`
	} `json:"sliceOfStruct"`

	SomethingId        int `json:"id"`
	SomethingTimestamp int `json:"timestamp"`
}

// SomethingCreate is a struct that represents a something create
type SomethingCreate struct {
	SomethingInt    int               `json:"int"`
	SomethingString string            `json:"string"`
	SomethingBool   bool              `json:"bool"`
	SomethingFloat  float64           `json:"float"`
	SomethingArray  []string          `json:"array"`
	SomethingMap    map[string]string `json:"map"`
	SomethingStruct struct {
		StructInt    int    `json:"int"`
		StructString string `json:"string"`
	} `json:"struct"`
	SomethingPointer       *string     `json:"pointer"`
	SomethingInterface     interface{} `json:"interface"`
	SomethingSlice         []string    `json:"slice"`
	SomethingSliceOfStruct []struct {
		StructInt    int    `json:"int"`
		StructString string `json:"string"`
	} `json:"sliceOfStruct"`

	SomethingId        int `json:"id"`
	SomethingTimestamp int `json:"timestamp"`
}

type SomethingCreateResponse struct {
	SomethingId int `json:"id"`

	SomethingTimestamp int `json:"timestamp"`
}

// SomethingRead is a struct that represents a something read
type SomethingRead struct {
	SomethingInt    int               `json:"int"`
	SomethingString string            `json:"string"`
	SomethingBool   bool              `json:"bool"`
	SomethingFloat  float64           `json:"float"`
	SomethingArray  []string          `json:"array"`
	SomethingMap    map[string]string `json:"map"`
	SomethingStruct struct {
		StructInt    int    `json:"int"`
		StructString string `json:"string"`
	} `json:"struct"`
	SomethingPointer       *string     `json:"pointer"`
	SomethingInterface     interface{} `json:"interface"`
	SomethingSlice         []string    `json:"slice"`
	SomethingSliceOfStruct []struct {
		StructInt    int    `json:"int"`
		StructString string `json:"string"`
	} `json:"sliceOfStruct"`

	SomethingId        int `json:"id"`
	SomethingTimestamp int `json:"timestamp"`
}

type SomethingReadRequest struct {
	SomethingId int `json:"id"`
}

// SomethingList is a struct that represents a something list
type SomethingList struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`

	Somethings []SomethingRead `json:"somethings"`
}

// SomethingListRequest is a struct that represents a something list request
type SomethingListRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// SomethingUpdate is a struct that represents a something update
type SomethingUpdate struct {
	SomethingInt    *int               `json:"int"`
	SomethingString *string            `json:"string"`
	SomethingBool   *bool              `json:"bool"`
	SomethingFloat  *float64           `json:"float"`
	SomethingArray  *[]string          `json:"array"`
	SomethingMap    *map[string]string `json:"map"`
	SomethingStruct *struct {
		StructInt    *int    `json:"int"`
		StructString *string `json:"string"`
	} `json:"struct"`
	SomethingPointer       *string     `json:"pointer"`
	SomethingInterface     interface{} `json:"interface"`
	SomethingSlice         *[]string   `json:"slice"`
	SomethingSliceOfStruct *[]struct {
		StructInt    *int    `json:"int"`
		StructString *string `json:"string"`
	} `json:"sliceOfStruct"`

	SomethingId        *int `json:"id"`
	SomethingTimestamp *int `json:"timestamp"`
}

type SomethingUpdateResponse struct {
	SomethingId int `json:"id"`

	SomethingTimestamp int `json:"timestamp"`
}

// SomethingPatch is a struct that represents a something patch
type SomethingPatch struct {
	SomethingKey   string      `json:"key"`
	SomethingValue interface{} `json:"value"`
	SomethingId    int         `json:"id"`

	SomethingTimestamp int `json:"timestamp"`
}

// SomethingPatchResponse is a struct that represents a something patch response
type SomethingPatchResponse struct {
	SomethingId int `json:"id"`

	SomethingTimestamp int `json:"timestamp"`
}

// SomethingDelete is a struct that represents a something delete
type SomethingDelete struct {
	SomethingId int `json:"id"`
}

// SomethingDeleteResponse is a struct that represents a something delete response
type SomethingDeleteResponse struct {
	SomethingId int `json:"id"`
}
