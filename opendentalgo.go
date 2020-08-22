package opendentalgo

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetAppointments - Get the appointments
func GetAppointments() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appKey := os.Getenv("APP_KEY")
	apiKey := os.Getenv("API_KEY")
	baseURL := os.Getenv("BASE_URL")
	api := NewAPI(baseURL, appKey, apiKey)
	resp, getAppointmentErr := api.GetAppointments()
	if getAppointmentErr != nil {
		log.Print("Error getting appointments")
		log.Print(getAppointmentErr)
		return ""
	}

	return resp
}
