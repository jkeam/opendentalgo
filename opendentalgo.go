package opendentalgo

import (
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/jkeam/opendentalgo/models"
	"github.com/joho/godotenv"
)

// OpenDental - main developer interface
type OpenDental struct {
	Endpoint *API
}

// NewOpenDental - Default Constructor
func NewOpenDental() *OpenDental {
	client := &OpenDental{}
	return client.init(nil)
}

// NewOpenDentalWithClient - Constructor, optionally passing in resty client
func NewOpenDentalWithClient(restyClient *resty.Client) *OpenDental {
	client := &OpenDental{}
	return client.init(restyClient)
}

// init - Get the client ready
func (openDental *OpenDental) init(restyClient *resty.Client) *OpenDental {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appKey := os.Getenv("APP_KEY")
	apiKey := os.Getenv("API_KEY")
	baseURL := os.Getenv("BASE_URL")
	openDental.Endpoint = NewAPI(baseURL, appKey, apiKey, restyClient)
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

// GetLocations - Get collection of locations
func (openDental *OpenDental) GetLocations() (*models.LocationBundle, error) {
	return openDental.Endpoint.GetLocations()
}

// FindLocation - Find the location
func (openDental *OpenDental) FindLocation(name string) (*models.LocationResource, error) {
	return openDental.Endpoint.FindLocation(name)
}
