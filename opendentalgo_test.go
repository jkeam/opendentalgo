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
}
