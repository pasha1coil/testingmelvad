package models

type Incr struct {
	Key   string `json:"key" validate:"required"`
	Value int64  `json:"value" validate:"required"`
}

type Hash struct {
	Text string `json:"text" validate:"required"`
	Key  string `json:"key" validate:"required"`
}

type Users struct {
	Name string `json:"name" validate:"required"`
	Age  uint64 `json:"age" validate:"required"`
}
