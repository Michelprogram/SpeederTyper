package types

type Response struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}
