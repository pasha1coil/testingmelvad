package models

type Incr struct {
	Key   string `json:"key" validate:"required"`
	Value int    `json:"value" validate:"required"`
}

type Hmac struct {
	Text string `json:"text" validate:"required"`
	Key  string `json:"key" validate:"required"`
}

type Users struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required"`
}
