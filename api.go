package opendentalgo

import (
	"encoding/json"
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

func (api *API) getResource(path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", api.baseURL, path)

	// resty.Response
	resp, err := api.client.R().EnableTrace().
		SetHeader("Authorization", fmt.Sprintf("ODFHIR %s/%s", api.appKey, api.apiKey)).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").Get(url)

	if err != nil {
		log.Printf("Error getting %s", path)
		log.Print(err)
		return nil, err
	}

	return resp.Body(), nil
}

// NewAPI - create the api object
func NewAPI(baseURL string, appKey string, apiKey string) *API {
	return &API{
		client:  resty.New(),
		appKey:  appKey,
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// GetAppointments - Get collection of appointments
func (api *API) GetAppointments() (*models.AppointmentBundle, error) {
	contents, getErr := api.getResource("appointment")
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
	contents, getErr := api.getResource("patient")
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
