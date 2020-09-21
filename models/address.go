package models

// Address - address of the patient
type Address struct {
	Use        string   `json:"use"`
	Lines      []string `json:"line"`
	City       string   `json:"city"`
	District   string   `json:"district"`
	PostalCode string   `json:"postalCode"`
}
