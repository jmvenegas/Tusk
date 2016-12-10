package model

type Safer interface {
	IsSafeToUse() bool
}

type Tester interface {
	GetTest() *Test
}

type Test struct {
	Pattern string `json:"pattern"`
	Date    string `json:"date"`
	Result  int    `json:result"`
	Table   string `json:"table"`
}
