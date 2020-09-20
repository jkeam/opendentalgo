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

	assert.Equal(t, len(data.Entry), 567, "Wrong number of appointments")
	assert.Equal(t, data.Type, "searchset", "Type does not match")
	assert.Equal(t, data.Total, 567, "Total does not match")

	appointment := data.Entry[0]
	assert.Equal(t, appointment.FullURL, "https://api.opendental.com/fhir/v2/appointment/4", "Wrong URL")

	appointmentDetail := appointment.Resource
	assert.Equal(t, appointmentDetail.ResourceType, "Appointment", "Resource is not matching")
	assert.Equal(t, appointmentDetail.ID, "4", "ID is not matching")
	assert.Equal(t, appointmentDetail.Status, "fulfilled", "Status is not matching")
	assert.Equal(t, appointmentDetail.Priority, 5, "Priority is not matching")
	assert.Equal(t, appointmentDetail.Description, "PerEx, Flo, DentAdj", "Description is not matching")
	assert.Equal(t, appointmentDetail.Start, "2017-03-17T08:00:00", "Start is not matching")
	assert.Equal(t, appointmentDetail.End, "2017-03-17T08:40:00", "End is not matching")
	assert.Equal(t, appointmentDetail.MinutesDuration, 40, "Minutes duration is not matching")

	participants := appointmentDetail.Participants
	assert.Equal(t, len(participants), 3, "Participant count is not matching")

	participant := participants[0]
	assert.Equal(t, participant.Required, "required", "Required is not matching")
	assert.Equal(t, participant.Status, "needsaction", "Status is not matching")

	actor := participant.Actor
	assert.Equal(t, actor.Reference, "patient/1", "Actor reference is not matching")
	assert.Equal(t, actor.Display, "Hermione Granger", "Actor display is not matching")

	participantTypes := participant.Type
	assert.Equal(t, len(participantTypes), 1, "Participant type count is not matching")

	coding := participantTypes[0].Coding
	assert.Equal(t, len(coding), 1, "Coding type count is not matching")

	code := coding[0]
	assert.Equal(t, code.System, "http://hl7.org/fhir/participant-type", "Code system is not matching")
	assert.Equal(t, code.Code, "PART", "Code code is not matching")
	assert.Equal(t, code.Display, "Participation", "Code participation is not matching")

	meta := appointmentDetail.Meta
	assert.Equal(t, meta.LastUpdated, "2019-01-02T12:22:38", "Meta last updated is not matching")

	search := appointment.Search
	assert.Equal(t, search.Mode, "match", "Search mode last updated is not matching")
	assert.Equal(t, search.Score, 1.0, "Search score last updated is not matching")
}
