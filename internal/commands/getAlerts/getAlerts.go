package getalerts

import (
	"encoding/json"
	"errors"
	"net/http"

	getJwtToken "github.com/Platon223/Larb/internal/domain/jwt"
)

type User struct {
	apiKey string
}

type Alert struct {
	Id          string `json:"id"`
	Message     string `json:"message"`
	Level       string `json:"level"`
	Time        string `json:"time"`
	UserId      string `json:"user_id"`
	ServiceId   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	Viewed      bool   `json:"viewed"`
}

type Message struct {
	Message []Alert `json:"message"`
}

func ConfigUser(apiKey string) *User {
	user := &User{
		apiKey: apiKey,
	}

	return user
}

func (u *User) GetAlerts() ([]Alert, error) {
	jwtToken, err := getJwtToken.GetJwt(u.apiKey)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", "https://logarbor.com/api/v1/alerts/alerts", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+jwtToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Something went wrong while fetching your alerts. Please confirm that your api key is set correctly.")
	}

	defer resp.Body.Close()

	var alerts Message
	json.NewDecoder(resp.Body).Decode(&alerts)

	return alerts.Message, nil
}
