package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializePatientBundle(t *testing.T) {
	contents, err := ReadFile("../tests/patients.json")
	if err != nil {
		t.Errorf("Error reading file")
		return
	}

	data := &PatientBundle{}
	if err := json.Unmarshal(contents, data); err != nil {
		t.Errorf("Error unmarshalling patients")
		t.Error(err)
		return
	}

	assert.Equal(t, len(data.Entry), 259, "Wrong number of patients")
	assert.Equal(t, data.Total, 259, "Wrong number of patients")
	assert.Equal(t, data.Type, "searchset", "Type does not match")

	patient := data.Entry[0]
	assert.Equal(t, patient.FullURL, "https://api.opendental.com/fhir/v2/patient/1", "Full URL does not match")

	patientDetail := patient.Resource
	assert.Equal(t, patientDetail.ResourceType, "Patient", "Resource type does not match")
	assert.Equal(t, patientDetail.ID, "1", "ID does not match")
	assert.Equal(t, patientDetail.Active, true, "Active does not match")

	assert.Equal(t, len(patientDetail.Names), 1, "Names length does not match")
	name := patientDetail.Names[0]
	assert.Equal(t, name.Use, "usual", "Name use does not match")
	assert.Equal(t, name.Text, "Hermione Granger", "Name text does not match")
	assert.Equal(t, name.Family, "Granger", "Name family does not match")
	assert.Equal(t, name.Given, "Hermione", "Name given does not match")

	assert.Equal(t, len(patientDetail.Telecoms), 3, "Telecomes length does not match")
	email := patientDetail.Telecoms[0]
	assert.Equal(t, email.System, "email", "Telecom system does not match")
	assert.Equal(t, email.Value, "chris@opendental.com", "Telecom value does not match")
	assert.Equal(t, email.Use, "home", "Telecom use does not match")
	assert.Equal(t, email.Rank, 0, "Telecom rank does not match")

	assert.Equal(t, patientDetail.Gender, "female", "Gender does not match")
	assert.Equal(t, patientDetail.BirthDate, "1997-10-28T00:00:00", "BirthDate does not match")

	assert.Equal(t, len(patientDetail.Addresses), 1, "Address length does not match")
	address := patientDetail.Addresses[0]
	assert.Equal(t, address.Use, "home", "Address use does not match")
	assert.Equal(t, address.City, "Juniper", "Address city does not match")
	assert.Equal(t, address.District, "JI", "Address district does not match")
	assert.Equal(t, address.PostalCode, "24018", "Address postal code does not match")

	assert.Equal(t, len(address.Lines), 1, "Address lines length does not match")
	line := address.Lines[0]
	assert.Equal(t, line, "44 Mug Loop", "Address line does not match")

	assert.Equal(t, len(patientDetail.CareProviders), 1, "Care providers length does not match")
	careProvider := patientDetail.CareProviders[0]
	assert.Equal(t, careProvider.Reference, "practitioner/1", "Care provider reference does not match")
	assert.Equal(t, careProvider.Reference, "practitioner/1", "Care provider reference does not match")

	managingOrganization := patientDetail.ManagingOrganizations
	assert.Equal(t, managingOrganization.Reference, "organization/1", "Managing Org reference does not match")
	assert.Equal(t, managingOrganization.Display, "Hogwarts Hospital Wing", "Managing Org display does not match")

	meta := patientDetail.Meta
	assert.Equal(t, meta.LastUpdated, "2020-01-06T09:34:32", "Meta does not match")

	search := patient.Search
	assert.Equal(t, search.Mode, "match", "Search mode does not match")
	assert.Equal(t, search.Score, 1.0, "Search score does not match")
}
