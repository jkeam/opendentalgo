package models

// PatientBundle - Collection of Patients
type PatientBundle struct {
	ResourceType string    `json:"resourceType"`
	Type         string    `json:"type"`
	Total        int       `json:"total"`
	Entry        []Patient `json:"entry"`
}

// Patient - the patient information
type Patient struct {
	FullURL  string          `json:"fullUrl"`
	Resource PatientResource `json:"resource"`
	Search   PatientSearch   `json:"search"`
}

// PatientResource - detail information on the patient
type PatientResource struct {
	ResourceType          string                              `json:"resourceType"`
	ID                    string                              `json:"id"`
	Active                bool                                `json:"active"`
	Names                 []PatientResourceName               `json:"name"`
	Gender                string                              `json:"gender"`
	BirthDate             string                              `json:"birthDate"`
	Telecoms              []PatientResourceTelecom            `json:"telecom"`
	Addresses             []PatientResourceAddress            `json:"address"`
	CareProviders         []PatientResourceCareProvider       `json:"careProvider"`
	ManagingOrganizations PatientResourceManagingOrganization `json:"managingOrganization"`
	Meta                  PatientResourceMeta                 `json:"meta"`
}

// PatientResourceName - name of the patient
type PatientResourceName struct {
	Use    string `json:"use"`
	Text   string `json:"text"`
	Family string `json:"family"`
	Given  string `json:"given"`
}

// PatientResourceTelecom - contact information for patient
type PatientResourceTelecom struct {
	System string `json:"system"`
	Value  string `json:"value"`
	Use    string `json:"use"`
	Rank   int    `json:"rank"`
}

// PatientResourceAddress - address of the patient
type PatientResourceAddress struct {
	Use        string   `json:"use"`
	Lines      []string `json:"line"`
	City       string   `json:"city"`
	District   string   `json:"district"`
	PostalCode string   `json:"postalCode"`
}

// PatientResourceCareProvider - care provider of the patient
type PatientResourceCareProvider struct {
	Reference string `json:"reference"`
	Display   string `json:"display"`
}

// PatientResourceManagingOrganization - managing organization of the patient
type PatientResourceManagingOrganization struct {
	Reference string `json:"reference"`
	Display   string `json:"display"`
}

// PatientResourceMeta - meta data related to the patient
type PatientResourceMeta struct {
	LastUpdated string `json:"lastUpdated"`
}

// PatientSearch - patient search
type PatientSearch struct {
	Mode  string  `json:"mode"`
	Score float64 `json:"score"`
}
