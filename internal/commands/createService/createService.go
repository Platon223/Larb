package createservice

import (
	"bytes"
	"encoding/json"
	"net/http"
	"fmt"

	getJwtToken "github.com/Platon223/Larb/internal/domain/jwt"
)

type User struct {
	apiKey string
}

type CreateService struct {
	Name       string `json:"name"`
	AlertLevel string `json:"alert_level"`
}

type CreateServiceResponse struct {
	Message string `json:"message"`
}

func ConfigUser(apiKey string) *User {
	configedUser := &User{apiKey: apiKey}

	return configedUser
}

func (u *User) CreateService(name string, alertLevel string) (string, error) {

	// Define the regular get client
	regularGet := getJwtToken.RegularGet{}

	// Use the client to get the token from viper
	jwtToken, err := getJwtToken.WithType(regularGet, u.apiKey)

	if err != nil {
		return "", err
	}

	body := CreateService{
		Name:       name,
		AlertLevel: alertLevel,
	}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://logarbor.com/api/v1/services/create", bytes.NewBuffer(jsonBody))

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+jwtToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	if resp.StatusCode == 401 {
	
		// Define the expired get client
		expiredGet := getJwtToken.ExpiredGet{}

		// Use the client to fetch the new token
		newJwtToken, err := getJwtToken.WithType(expiredGet, u.apiKey)

		if err != nil {
			return "", fmt.Errorf("Something went wrong. Try again or check your internet connection.")
		}

		req, err = http.NewRequest("POST", "https://logarbor.com/api/v1/services/create", bytes.NewBuffer(jsonBody))

		if err != nil {
			return "", fmt.Errorf("Something went wrong. Try again or check your internet connection.")
		}

		req.Header.Set("Authorization", "Bearer "+newJwtToken)
		req.Header.Set("Content-Type", "application/json")

		resp, err = http.DefaultClient.Do(req)

		if err != nil {
			return "", fmt.Errorf("Status Code: %d. Something went wrong. Try again or check your internet connection.", resp.StatusCode)
		}	

		defer resp.Body.Close()

	}	

	defer resp.Body.Close()

	var message CreateServiceResponse
	json.NewDecoder(resp.Body).Decode(&message)

	return message.Message, nil
}
