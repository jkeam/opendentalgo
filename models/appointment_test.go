package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializeAppointmentBundle(t *testing.T) {
	contents, err := ReadFile("../tests/appointments.json")
	if err != nil {
		t.Errorf("Error reading file")
		t.Error(err)
		return
	}

	data := &AppointmentBundle{}
	if err := json.Unmarshal(contents, data); err != nil {
		t.Errorf("Error unmarshalling appointments")
		t.Error(err)
		return
	}

	assert.Equal(t, 567, len(data.Entry), "Wrong number of appointments")
	assert.Equal(t, "searchset", data.Type, "Type does not match")
	assert.Equal(t, 567, data.Total, "Total does not match")

	appointment := data.Entry[0]
	assert.Equal(t, "https://api.opendental.com/fhir/v2/appointment/4", appointment.FullURL, "Wrong URL")

	appointmentDetail := appointment.Resource
	assert.Equal(t, "Appointment", appointmentDetail.ResourceType, "Resource is not matching")
	assert.Equal(t, "4", appointmentDetail.ID, "ID is not matching")
	assert.Equal(t, "fulfilled", appointmentDetail.Status, "Status is not matching")
	assert.Equal(t, 5, appointmentDetail.Priority, 5, "Priority is not matching")
	assert.Equal(t, "PerEx, Flo, DentAdj", appointmentDetail.Description, "Description is not matching")
	assert.Equal(t, "2017-03-17T08:00:00", appointmentDetail.Start, "Start is not matching")
	assert.Equal(t, "2017-03-17T08:40:00", appointmentDetail.End, "End is not matching")
	assert.Equal(t, 40, appointmentDetail.MinutesDuration, "Minutes duration is not matching")

	participants := appointmentDetail.Participants
	assert.Equal(t, 3, len(participants), "Participant count is not matching")

	participant := participants[0]
	assert.Equal(t, "required", participant.Required, "Required is not matching")
	assert.Equal(t, "needsaction", participant.Status, "Status is not matching")

	actor := participant.Actor
	assert.Equal(t, "patient/1", actor.Reference, "Actor reference is not matching")
	assert.Equal(t, "Hermione Granger", actor.Display, "Actor display is not matching")

	participantTypes := participant.Type
	assert.Equal(t, 1, len(participantTypes), "Participant type count is not matching")

	coding := participantTypes[0].Coding
	assert.Equal(t, 1, len(coding), "Coding type count is not matching")

	code := coding[0]
	assert.Equal(t, "http://hl7.org/fhir/participant-type", code.System, "Code system is not matching")
	assert.Equal(t, "PART", code.Code, "Code code is not matching")
	assert.Equal(t, "Participation", code.Display, "Code participation is not matching")

	meta := appointmentDetail.Meta
	assert.Equal(t, "2019-01-02T12:22:38", meta.LastUpdated, "Meta last updated is not matching")

	search := appointment.Search
	assert.Equal(t, "match", search.Mode, "Search mode last updated is not matching")
	assert.Equal(t, 1.0, search.Score, "Search score last updated is not matching")
}
