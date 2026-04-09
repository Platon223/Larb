package getlogs

import (
	"encoding/json"
	"net/http"
	"slices"

	getJwtToken "github.com/Platon223/Larb/internal/domain/jwt"
)

type User struct {
	apiKey string
}

type Log struct {
	Id        string `json:"id"`
	ServiceId string `json:"service_id"`
	UserId    string `json:"user_id"`
	Message   string `json:"message"`
	Level     string `json:"level"`
	Time      string `json:"time"`
}

type Logs struct {
	ServiceId   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	LogList     []Log  `json:"logs"`
}

type Message struct {
	Message []Logs `json:"message"`
}

func ConfigUser(apiKey string) *User {
	user := &User{
		apiKey: apiKey,
	}

	return user
}

func (u *User) GetLogs(serviceId string) (Logs, error) {
	jwtToken, err := getJwtToken.GetJwt(u.apiKey)

	if err != nil {
		return Logs{}, err
	}

	req, err := http.NewRequest("POST", "https://logarbor.com/api/v1/logs/all_logs", nil)

	if err != nil {
		return Logs{}, err
	}

	req.Header.Set("Authorization", "Bearer "+jwtToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return Logs{}, err
	}

	defer resp.Body.Close()

	var logs Message
	json.NewDecoder(resp.Body).Decode(&logs)

	servicesLogs := Logs{}

	for _, log := range logs.Message {
		if log.ServiceId == serviceId {
			servicesLogs = log

			break
		}
	}

	slices.Reverse(servicesLogs.LogList)

	return servicesLogs, nil
}
