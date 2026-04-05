package createservice

import (
	"bytes"
	"encoding/json"
	"net/http"

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

	jwtToken, err := getJwtToken.GetJwt(u.apiKey)

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
		return "", nil
	}

	req.Header.Set("Authorization", "Bearer "+jwtToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var message CreateServiceResponse
	json.NewDecoder(resp.Body).Decode(&message)

	return message.Message, nil
}
