package examples

type Live struct {
	LiveID      int64   `json:"live_id"`
	LiveName    *string `json:"live_name"`
	Status      *string `json:"status" default:"preparing"`
	LiveType    *int64  `json:"live_type" defalut:"live_type"`
	IsSpecial   *bool   `json:"is_special"`
	IsAvailable *bool   `json:"is_available" default:"true"`
}

type LiveAliased Live

type LiveIgnoreAliased = Live
