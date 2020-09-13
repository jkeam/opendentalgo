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
	appointments := client.GetAppointments()
	log.Print("Appointments")
	log.Print(appointments)
}
func TestGetPatients(t *testing.T) {
	patients := client.GetPatients()
	log.Print("Patients")
	log.Print(patients)
}
