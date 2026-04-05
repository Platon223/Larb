package allservices

import (
	"encoding/json"
	"errors"
	"net/http"

	getJwtToken "github.com/Platon223/Larb/internal/domain/jwt"
)

type User struct {
	apiKey string
}

type Service struct {
	Id           string       `json:"id"`
	Name         string       `json:"name"`
	AlertLevel   string       `json:"alert_level"`
	UserId       string       `json:"user_id"`
	LogRetention LogRetention `json:"log_retention"`
	Health       string       `json:"health"`
	LogCount     int          `json:"total_logs"`
}

type Message struct {
	Message []Service `json:"message"`
}

type LogRetention struct {
	Date string `json:"$date"`
}

// Function that creates a User type which will allow to then use the GetAllServices method

func ConfigUser(apiKey string) *User {
	configedUser := &User{apiKey: apiKey}

	return configedUser
}

// Method that fetches user's services

func (u *User) GetAllServices() ([]Service, error) {

	jwtToken, err := getJwtToken.GetJwt(u.apiKey)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://logarbor.com/api/v1/services/all_services", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+jwtToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Something went wrong while fetching your services. Please confirm that your api key is set correctly.")
	}

	defer resp.Body.Close()

	var services Message
	json.NewDecoder(resp.Body).Decode(&services)

	return services.Message, nil
}
