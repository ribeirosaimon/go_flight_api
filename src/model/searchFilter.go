package model

type SearchFilter struct {
	Page     uint8    `json:"page"`
	PageSize uint8    `json:"pageSize"`
	Filters  []Filter `json:"filters"`
}

type SearchFilterResult struct {
	Count  uint32        `json:"count"`
	Result []interface{} `json:"result"`
}

type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
