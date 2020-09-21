package models

// Search - patient search
type Search struct {
	Mode  string  `json:"mode"`
	Score float64 `json:"score"`
}
