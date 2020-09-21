package models

// Telecom - contact information for patient
type Telecom struct {
	System string `json:"system"`
	Value  string `json:"value"`
	Use    string `json:"use"`
	Rank   int    `json:"rank"`
}
