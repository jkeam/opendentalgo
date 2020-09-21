package models

// AppointmentBundle - Bundle of appointments
type AppointmentBundle struct {
	ResourceType string        `json:"resourceType"`
	Type         string        `json:"type"`
	Total        int           `json:"total"`
	Entry        []Appointment `json:"entry"`
}

// Appointment - appointment information
type Appointment struct {
	FullURL  string              `json:"fullUrl"`
	Resource AppointmentResource `json:"resource"`
	Search   AppointmentSearch   `json:"search"`
}

// AppointmentResource - detail information about the appointment
type AppointmentResource struct {
	ResourceType    string                   `json:"resourceType"`
	ID              string                   `json:"id"`
	Identifiers     []Identifier             `json:"identifier"`
	Status          string                   `json:"status"`
	Priority        int                      `json:"priority"`
	Description     string                   `json:"description"`
	Start           string                   `json:"start"`
	End             string                   `json:"end"`
	MinutesDuration int                      `json:"minutesDuration"`
	Participants    []AppointmentParticipant `json:"participant"`
	Meta            AppointmentMeta          `json:"meta"`
}

// AppointmentParticipant - appointment participant information
type AppointmentParticipant struct {
	Type     []AppointmentParticipantType `json:"type"`
	Actor    AppointmentParticipantActor  `json:"actor"`
	Required string                       `json:"required"`
	Status   string                       `json:"status"`
}

// AppointmentParticipantType - type of participant for the user
type AppointmentParticipantType struct {
	Coding []AppointmentParticipantTypeCoding `json:"coding"`
}

// AppointmentParticipantTypeCoding - participant type code
type AppointmentParticipantTypeCoding struct {
	System  string `json:"system"`
	Code    string `json:"code"`
	Display string `json:"display"`
}

// AppointmentParticipantActor - the person who is the appointment participant
type AppointmentParticipantActor struct {
	Reference string `json:"reference"`
	Display   string `json:"display"`
}

// AppointmentMeta - meta data for the appointment
type AppointmentMeta struct {
	LastUpdated string `json:"lastUpdated"`
}

// AppointmentSearch - appointment search details
type AppointmentSearch struct {
	Mode  string  `json:"mode"`
	Score float64 `json:"score"`
}
