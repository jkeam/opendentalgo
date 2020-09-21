package opendentalgo

import (
	"log"
	"os"

	"github.com/jkeam/opendentalgo/models"
	"github.com/joho/godotenv"
)

// OpenDental - main developer interface
type OpenDental struct {
	Endpoint *API
}

// NewOpenDental - Constructor
func NewOpenDental() *OpenDental {
	client := &OpenDental{}
	return client.init()
}

// init - Get the client ready
func (openDental *OpenDental) init() *OpenDental {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appKey := os.Getenv("APP_KEY")
	apiKey := os.Getenv("API_KEY")
	baseURL := os.Getenv("BASE_URL")
	openDental.Endpoint = NewAPI(baseURL, appKey, apiKey)
	return openDental
}

// GetAppointments - Get the appointments
func (openDental *OpenDental) GetAppointments() (*models.AppointmentBundle, error) {
	return openDental.Endpoint.GetAppointments()
}

// GetPatients - Get the patients
func (openDental *OpenDental) GetPatients() (*models.PatientBundle, error) {
	return openDental.Endpoint.GetPatients()
}

// CreatePatient - Create patient
func (openDental *OpenDental) CreatePatient(
	firstName string,
	lastName string,
	cellPhone string,
	gender string,
	birthDate string,
) (*models.PatientResource, error) {
	return openDental.Endpoint.CreatePatient(firstName, lastName, cellPhone, gender, birthDate)
}

// FindPatient - Finds the patient
func (openDental *OpenDental) FindPatient(
	firstName string,
	lastName string,
	birthDate string,
) (*models.PatientResource, error) {
	return openDental.Endpoint.FindPatient(firstName, lastName, birthDate)
}
