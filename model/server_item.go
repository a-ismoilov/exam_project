package model

type Items struct {
	Words []map[string]int `json:"words"`
	Page  int              `json:"page"`
	Limit int              `json:"limit"`
}
