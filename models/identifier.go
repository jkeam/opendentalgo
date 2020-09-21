package models

// Identifier - identifier object
type Identifier struct {
	Use   string         `json:"use"`
	Type  IdentifierType `json:"type"`
	Value string         `json:"value"`
}

// IdentifierType - identifier description text
type IdentifierType struct {
	Text string `json:"text"`
}
