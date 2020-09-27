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

	assert.Equal(t, 259, len(data.Entry), "Wrong number of patients")
	assert.Equal(t, 259, data.Total, "Wrong number of patients")
	assert.Equal(t, "searchset", data.Type, "Type does not match")

	patient := data.Entry[0]
	assert.Equal(t, "https://api.opendental.com/fhir/v2/patient/1", patient.FullURL, "Full URL does not match")

	patientDetail := patient.Resource
	assert.Equal(t, patientDetail.ResourceType, "Patient", patientDetail.ResourceType, "Resource type does not match")
	assert.Equal(t, "1", patientDetail.ID, "ID does not match")
	assert.Equal(t, true, patientDetail.Active, "Active does not match")

	assert.Equal(t, 1, len(patientDetail.Names), "Names length does not match")
	name := patientDetail.Names[0]
	assert.Equal(t, "usual", name.Use, "Name use does not match")
	assert.Equal(t, "Hermione Granger", name.Text, "Name text does not match")
	assert.Equal(t, "Granger", name.Family, "Name family does not match")
	assert.Equal(t, "Hermione", name.Given, "Name given does not match")

	assert.Equal(t, len(patientDetail.Telecoms), 3, "Telecomes length does not match")
	email := patientDetail.Telecoms[0]
	assert.Equal(t, "email", email.System, "Telecom system does not match")
	assert.Equal(t, "chris@opendental.com", email.Value, "Telecom value does not match")
	assert.Equal(t, "home", email.Use, "Telecom use does not match")
	assert.Equal(t, 0, email.Rank, "Telecom rank does not match")

	assert.Equal(t, "female", patientDetail.Gender, "Gender does not match")
	assert.Equal(t, "1997-10-28T00:00:00", patientDetail.BirthDate, "BirthDate does not match")

	assert.Equal(t, 1, len(patientDetail.Addresses), "Address length does not match")
	address := patientDetail.Addresses[0]
	assert.Equal(t, "home", address.Use, "Address use does not match")
	assert.Equal(t, "Juniper", address.City, "Address city does not match")
	assert.Equal(t, "JI", address.District, "Address district does not match")
	assert.Equal(t, "24018", address.PostalCode, "Address postal code does not match")

	assert.Equal(t, 1, len(address.Lines), "Address lines length does not match")
	line := address.Lines[0]
	assert.Equal(t, "44 Mug Loop", line, "Address line does not match")

	assert.Equal(t, 1, len(patientDetail.CareProviders), "Care providers length does not match")
	careProvider := patientDetail.CareProviders[0]
	assert.Equal(t, "practitioner/1", careProvider.Reference, "Care provider reference does not match")
	assert.Equal(t, "Madame Pomprey, DMD", careProvider.Display, "Care provider display does not match")

	managingOrganization := patientDetail.ManagingOrganizations
	assert.Equal(t, "organization/1", managingOrganization.Reference, "Managing Org reference does not match")
	assert.Equal(t, "Hogwarts Hospital Wing", managingOrganization.Display, "Managing Org display does not match")

	meta := patientDetail.Meta
	assert.Equal(t, "2020-01-06T09:34:32", meta.LastUpdated, "Meta does not match")

	search := patient.Search
	assert.Equal(t, "match", search.Mode, "Search mode does not match")
	assert.Equal(t, 1.0, search.Score, "Search score does not match")
}
