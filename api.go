package opendentalgo

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

// API - object to hit the endpoint
type API struct {
	client  *resty.Client
	appKey  string
	apiKey  string
	baseURL string
}

func (api *API) getResource(path string) (string, error) {
	url := fmt.Sprintf("%s/%s", api.baseURL, path)

	// resty.Response
	resp, err := api.client.R().EnableTrace().
		SetHeader("Authorization", fmt.Sprintf("ODFHIR %s/%s", api.appKey, api.apiKey)).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").Get(url)

	if err != nil {
		log.Printf("Error getting %s", path)
		log.Print(err)
		return "", err
	}

	return resp.String(), nil
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

// GetAppointments - gets all the appointments
func (api *API) GetAppointments() (string, error) {
	return api.getResource("appointment")
}

// GetPatients - gets all the patients
func (api *API) GetPatients() (string, error) {
	return api.getResource("patient")
}
