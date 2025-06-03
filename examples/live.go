package examples

type Live struct {
	LiveID      int64       `json:"live_id"`
	LiveName    *string     `json:"live_name"`
	Status      *StatusType `json:"status"`
	LiveType    *int64      `json:"live_type"`
	IsSpecial   *bool       `json:"is_special"`
	IsAvailable *bool       `json:"is_available"`
	Rate        *float64    `json:"rate"`
	DefaultRate *float64    `json:"default_rate"`
}

type LiveAliased Live

type LiveIgnoreAliased = Live

type StatusType int64
