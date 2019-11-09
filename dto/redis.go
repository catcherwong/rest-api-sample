package dto

type SetStringRequest struct {
	Key        string `json:"key"`
	Value      string `json:"val"`
	Expiration int64  `json:"exp"`
}
