package models

// ManagingOrganization - managing organization of the patient
type ManagingOrganization struct {
	Reference string `json:"reference"`
	Display   string `json:"display"`
}
