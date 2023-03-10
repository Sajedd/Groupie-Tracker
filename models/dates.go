package model

type Dates struct {
	Id   int64    `json:"id"`
	Dates []string `json:"dates"`
}

var (
	DateStruct Dates
)

