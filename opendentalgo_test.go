package opendentalgo

import (
	"log"
	"os"
	"testing"
)

var client *OpenDental

func TestMain(m *testing.M) {
	client = NewOpenDental()
	code := m.Run()
	os.Exit(code)
}
func TestGetLocations(t *testing.T) {
	locations, err := client.GetLocations()
	if err != nil {
		t.Errorf("Failed to get locations")
		t.Error(err)
		return
	}
	log.Print("Locations")
	log.Print(locations)
	log.Printf("Total %d", locations.Total)
}

func TestGetAppointments(t *testing.T) {
	appointments, err := client.GetAppointments()
	if err != nil {
		t.Errorf("Failed to get appointments")
		t.Error(err)
		return
	}
	log.Print("Appointments")
	log.Print(appointments)
}
func TestGetPatients(t *testing.T) {
	patients, err := client.GetPatients()
	if err != nil {
		t.Errorf("Failed to get patients")
		t.Error(err)
		return
	}
	log.Print("Patients")
	log.Print(patients)
	log.Printf("Total %d", patients.Total)
}

func TestCreatePatient(t *testing.T) {
	patient, err := client.CreatePatient("Boaty4", "McBoatface", "(202) 555-0001", "male", "1995-06-12")
	if err != nil {
		t.Error("Unable to create patient")
		t.Error(err)
	}

	log.Print("Create Patient")
	log.Print(patient)
}
func TestFindPatient(t *testing.T) {
	patient, err := client.FindPatient("Boaty3", "McBoatface", "1996-09-19")
	if err != nil {
		t.Error("Unable to find patient")
		t.Error(err)
	}

	log.Print("Find Patient")
	log.Print(patient)
}
func TestFindLocation(t *testing.T) {
	location, err := client.FindLocation("Un")
	if err != nil {
		t.Error("Unable to find location")
		t.Error(err)
	}

	log.Print("Find Location")
	log.Print(location)
}
