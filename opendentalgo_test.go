package opendentalgo

import (
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jkeam/opendentalgo/models"
	"github.com/stretchr/testify/assert"

	"github.com/jarcoal/httpmock"
)

var client *OpenDental

func TestMain(m *testing.M) {
	restyClient := resty.New()
	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()

	client = NewOpenDentalWithClient(restyClient)
	code := m.Run()

	httpmock.DeactivateAndReset()
	os.Exit(code)
}

func TestGetLocations(t *testing.T) {
	contents, readError := models.ReadFile("tests/locations.json")
	if readError != nil {
		t.Errorf("Unable to read fixture")
		return
	}

	httpmock.RegisterResponder("GET", "https://api.opendental.com/fhir/v2/location",
		httpmock.NewBytesResponder(200, contents))

	locations, err := client.GetLocations()
	if err != nil {
		t.Errorf("Failed to get locations")
		t.Error(err)
		return
	}

	assert.Equal(t, 9, len(locations.Entry), "Wrong number of locations")
}

func TestGetAppointments(t *testing.T) {
	contents, readError := models.ReadFile("tests/appointments.json")
	if readError != nil {
		t.Errorf("Unable to read fixture")
		return
	}

	httpmock.RegisterResponder("GET", "https://api.opendental.com/fhir/v2/appointment",
		httpmock.NewBytesResponder(200, contents))

	appointments, err := client.GetAppointments()
	if err != nil {
		t.Errorf("Failed to get appointments")
		t.Error(err)
		return
	}

	assert.Equal(t, 567, len(appointments.Entry), "Wrong number of appointments")
}

func TestGetPatients(t *testing.T) {
	contents, readError := models.ReadFile("tests/patients.json")
	if readError != nil {
		t.Errorf("Unable to read fixture")
		return
	}

	httpmock.RegisterResponder("GET", "https://api.opendental.com/fhir/v2/patient",
		httpmock.NewBytesResponder(200, contents))

	patients, err := client.GetPatients()
	if err != nil {
		t.Errorf("Failed to get patients")
		t.Error(err)
		return
	}

	assert.Equal(t, 259, len(patients.Entry), "Wrong number of patients")
}

func TestCreatePatient(t *testing.T) {
	contents, readError := models.ReadFile("tests/createPatient.json")
	if readError != nil {
		t.Errorf("Unable to read fixture")
		return
	}

	httpmock.RegisterResponder("POST", "https://api.opendental.com/fhir/v2/patient",
		httpmock.NewBytesResponder(200, contents))

	patient, err := client.CreatePatient("Boaty5", "McBoatface", "(202) 555-0001", "male", "1995-06-12")
	if err != nil {
		t.Error("Unable to create patient")
		t.Error(err)
	}

	assert.Equal(t, "335", patient.ID, "Patient ID does not match")
}

func TestFindPatient(t *testing.T) {
	contents, readError := models.ReadFile("tests/searchPatient.json")
	if readError != nil {
		t.Errorf("Unable to read fixture")
		return
	}

	httpmock.RegisterResponder("GET", "https://api.opendental.com/fhir/v2/patient",
		httpmock.NewBytesResponder(200, contents))

	patient, err := client.FindPatient("Boaty3", "McBoatface", "1996-09-19")
	if err != nil {
		t.Error("Unable to find patient")
		t.Error(err)
	}

	assert.Equal(t, "332", patient.ID, "Patient ID does not match")
}

func TestFindLocation(t *testing.T) {
	contents, readError := models.ReadFile("tests/searchLocation.json")
	if readError != nil {
		t.Errorf("Unable to read fixture")
		return
	}

	httpmock.RegisterResponder("GET", "https://api.opendental.com/fhir/v2/location",
		httpmock.NewBytesResponder(200, contents))

	location, err := client.FindLocation("Un")
	if err != nil {
		t.Error("Unable to find location")
		t.Error(err)
	}

	assert.Equal(t, "Un", location.Name, "Location name does not match")
}
