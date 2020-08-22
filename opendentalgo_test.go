package opendentalgo

import (
	"log"
	"testing"
)

func TestAbs(t *testing.T) {
	appointments := GetAppointments()
	log.Print(appointments)
}
