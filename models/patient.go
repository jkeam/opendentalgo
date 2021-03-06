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
	Search   Search          `json:"search"`
}

// PatientResource - detail information on the patient
type PatientResource struct {
	ResourceType          string                        `json:"resourceType"`
	Issues                []PatientResourceIssue        `json:"issue,omitempty"`
	ID                    string                        `json:"id"`
	Identifiers           []Identifier                  `json:"identifier"`
	Active                bool                          `json:"active"`
	Names                 []PatientResourceName         `json:"name"`
	Gender                string                        `json:"gender"`
	BirthDate             string                        `json:"birthDate"`
	Telecoms              []Telecom                     `json:"telecom"`
	Addresses             []Address                     `json:"address"`
	CareProviders         []PatientResourceCareProvider `json:"careProvider"`
	ManagingOrganizations ManagingOrganization          `json:"managingOrganization"`
	Meta                  PatientResourceMeta           `json:"meta"`
}

// PatientResourceIssue - issue created as a result of the API call
type PatientResourceIssue struct {
	Severity string                     `json:"severity"`
	Code     string                     `json:"code"`
	Details  PatientResourceIssueDetail `json:"details"`
}

// PatientResourceIssueDetail - issue detail
type PatientResourceIssueDetail struct {
	Text string `json:"text"`
}

// PatientResourceName - name of the patient
type PatientResourceName struct {
	Use    string `json:"use"`
	Text   string `json:"text"`
	Family string `json:"family"`
	Given  string `json:"given"`
}

// PatientResourceCareProvider - care provider of the patient
type PatientResourceCareProvider struct {
	Reference string `json:"reference"`
	Display   string `json:"display"`
}

// PatientResourceMeta - meta data related to the patient
type PatientResourceMeta struct {
	LastUpdated string `json:"lastUpdated"`
}
