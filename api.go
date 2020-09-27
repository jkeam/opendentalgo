package opendentalgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/jkeam/opendentalgo/models"
)

// API - object to hit the endpoint
type API struct {
	client  *resty.Client
	appKey  string
	apiKey  string
	baseURL string
}

func (api *API) createClient() *resty.Request {
	return api.client.R().EnableTrace().
		SetHeader("Authorization", fmt.Sprintf("ODFHIR %s/%s", api.appKey, api.apiKey)).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json")
}

func (api *API) getResource(path string, queryParams map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", api.baseURL, path)

	var request *resty.Request
	if queryParams == nil {
		request = api.createClient()
	} else {
		request = api.createClient().SetQueryParams(queryParams)
	}

	resp, err := request.Get(url)
	if err != nil {
		log.Printf("Error getting %s", path)
		log.Print(err)
		return nil, err
	}

	return resp.Body(), nil
}

func (api *API) createResource(path string, body []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", api.baseURL, path)

	// resty.Response
	resp, err := api.createClient().SetBody(body).Post(url)
	if err != nil {
		log.Printf("Error posting %s", path)
		log.Print(err)
		return nil, err
	}

	return resp.Body(), nil
}

// NewAPI - create the api object
func NewAPI(baseURL string, appKey string, apiKey string, restyClient *resty.Client) *API {
	client := restyClient
	if client == nil {
		client = resty.New()
	}

	return &API{
		client:  client,
		appKey:  appKey,
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// GetLocations - Get collection of locations
func (api *API) GetLocations() (*models.LocationBundle, error) {
	contents, getErr := api.getResource("location", nil)
	if getErr != nil {
		log.Print("Error making api call for locations")
		return nil, getErr
	}

	data := &models.LocationBundle{}
	if err := json.Unmarshal(contents, data); err != nil {
		log.Print("Error unmarshalling location")
		return nil, err
	}
	return data, nil
}

// GetAppointments - Get collection of appointments
func (api *API) GetAppointments() (*models.AppointmentBundle, error) {
	contents, getErr := api.getResource("appointment", nil)
	if getErr != nil {
		log.Print("Error making api call for appointments")
		return nil, getErr
	}

	data := &models.AppointmentBundle{}
	if err := json.Unmarshal(contents, data); err != nil {
		log.Print("Error unmarshalling appointments")
		return nil, err
	}
	return data, nil
}

// GetPatients - Get collection of patients
func (api *API) GetPatients() (*models.PatientBundle, error) {
	contents, getErr := api.getResource("patient", nil)
	if getErr != nil {
		log.Print("Error making api call for patients")
		return nil, getErr
	}

	data := &models.PatientBundle{}
	if err := json.Unmarshal(contents, data); err != nil {
		log.Print("Error unmarshalling patients")
		return nil, err
	}
	return data, nil
}

// CreatePatient - Create patient
func (api *API) CreatePatient(
	firstName string,
	lastName string,
	cellPhone string,
	gender string,
	birthDate string,
) (*models.PatientResource, error) {
	telecoms := make([]models.Telecom, 1)
	telecoms[0] = models.Telecom{
		System: "phone",
		Value:  cellPhone,
		Use:    "mobile",
	}
	names := make([]models.PatientResourceName, 1)
	names[0] = models.PatientResourceName{
		Use:    "usual",
		Family: lastName,
		Given:  firstName,
	}

	patient := &models.PatientResource{
		Active:    true,
		Names:     names,
		Telecoms:  telecoms,
		Gender:    gender,
		BirthDate: birthDate,
	}
	body, err := json.Marshal(patient)
	if err != nil {
		log.Print("Unable to marshall input")
		return nil, err
	}

	contents, postErr := api.createResource("patient", body)
	if postErr != nil {
		log.Print("Error making api call for create patient")
		return nil, postErr
	}

	data := &models.PatientResource{}
	if err := json.Unmarshal(contents, data); err != nil {
		log.Print("Error unmarshalling patient")
		return nil, err
	}

	if len(data.Issues) > 0 {
		issue := data.Issues[0]
		if issue.Severity == "error" {
			return nil, errors.New(issue.Details.Text)
		}
	}
	return data, nil
}

// FindPatient - finds the single patient based on query parameters
// will return nil if nothing or more than 1 found
func (api *API) FindPatient(firstName string, lastName string, birthDate string) (*models.PatientResource, error) {
	queryParams := make(map[string]string)
	queryParams["family"] = lastName
	queryParams["given"] = firstName
	queryParams["birthDate"] = birthDate
	contents, getErr := api.getResource("patient", queryParams)
	if getErr != nil {
		log.Print("Error making api call to find patients")
		return nil, getErr
	}

	data := &models.PatientBundle{}
	if err := json.Unmarshal(contents, data); err != nil {
		log.Print("Error unmarshalling patients")
		return nil, err
	}

	if len(data.Entry) != 1 {
		return nil, nil
	}

	return &data.Entry[0].Resource, nil
}

// FindLocation - finds the single location based on query parameters
// will return nil if nothing or more than 1 found
func (api *API) FindLocation(name string) (*models.LocationResource, error) {
	queryParams := make(map[string]string)
	queryParams["name"] = name
	contents, getErr := api.getResource("location", queryParams)
	if getErr != nil {
		log.Print("Error making api call to find location")
		return nil, getErr
	}

	data := &models.LocationBundle{}
	if err := json.Unmarshal(contents, data); err != nil {
		log.Print("Error unmarshalling locations")
		return nil, err
	}

	if len(data.Entry) != 1 {
		return nil, nil
	}

	return &data.Entry[0].Resource, nil
}
