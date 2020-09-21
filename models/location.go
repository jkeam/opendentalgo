package models

// LocationBundle - Collection of Locations
type LocationBundle struct {
	ResourceType string     `json:"resourceType"`
	Type         string     `json:"type"`
	Total        int        `json:"total"`
	Entry        []Location `json:"entry"`
}

// Location - the patient information
type Location struct {
	FullURL  string           `json:"fullUrl"`
	Resource LocationResource `json:"resource"`
	Search   Search           `json:"search"`
}

// LocationResource - detail information on the patient
type LocationResource struct {
	ResourceType          string               `json:"resourceType"`
	ID                    string               `json:"id"`
	Identifiers           []Identifier         `json:"identifier"`
	Status                string               `json:"status"`
	Name                  string               `json:"name"`
	Mode                  string               `json:"mode"`
	Telecoms              []Telecom            `json:"telecom"`
	Address               Address              `json:"address"`
	PhysicalType          LocationPhysicalType `json:"physicalType"`
	ManagingOrganizations ManagingOrganization `json:"managingOrganization"`
}

// LocationPhysicalType - location physical type
type LocationPhysicalType struct {
	Coding []LocationCode `json:"coding"`
}

// LocationCode - location coding system
type LocationCode struct {
	System  string `json:"system"`
	Code    string `json:"code"`
	Display string `json:"display"`
}
