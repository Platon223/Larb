package getmetrics

import (
	"encoding/json"
	"errors"
	"net/http"
	"fmt"

	getJwtToken "github.com/Platon223/Larb/internal/domain/jwt"
)

type User struct {
	apiKey string
}

type Message struct {
	Message []LogCountMetric `json:"message"`
}

type LogSpeedMessage struct {
	Message []LogSpeedMetric `json:"message"`
}

type LogErrorMessage struct {
	Message []LogErrorMetric `json:"message"`
}

type LogCount struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type LogCountMetric struct {
	ServiceName string     `json:"service_name"`
	ServiceId   string     `json:"service_id"`
	LogCounts   []LogCount `json:"logs_metrics"`
}

type LogSpeedMetric struct {
	ServiceId   string  `json:"service_id"`
	ServiceName string  `json:"service_name"`
	Speed       float64 `json:"speed"`
}

type LogErrorMetric struct {
	ServiceName string  `json:"service_name"`
	ServiceId   string  `json:"service_id"`
	Rate        float64 `json:"rate"`
}

func ConfigUser(apiKey string) *User {
	u := &User{
		apiKey: apiKey,
	}

	return u
}

func (u *User) GetLogCountMetric() ([]LogCountMetric, error) {

	// Define regular get client
	regularGet := getJwtToken.RegularGet{}

	// Get the token from viper using regular client
	jwtToken, err := getJwtToken.WithType(regularGet, u.apiKey)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", "https://logarbor.com/api/v1/logs/metrics", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+jwtToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Something went wrong while fetching your log count metrics. Please confirm that your api key is set correctly.")
	}

	if resp.StatusCode == 401 {

		// Define the expired get client
		expiredGet := getJwtToken.ExpiredGet{}

		// Use the client to fetch the new token
		newJwtToken, err := getJwtToken.WithType(expiredGet, u.apiKey)

		if err != nil {
			return nil, fmt.Errorf("Something went wrong. Try again or check your internet connection.")
		}

		req, err = http.NewRequest("GET", "https://logarbor.com/api/v1/logs/metrics", nil)

		if err != nil {
			return nil, fmt.Errorf("Something went wrong. Try again or check your internet connection.")
		}

		req.Header.Set("Authorization", "Bearer "+newJwtToken)

		resp, err = http.DefaultClient.Do(req)

		if err != nil {
			return nil, fmt.Errorf("Status Code: %d. Something went wrong. Try again or check your internet connection.", resp.StatusCode)
		}	

		defer resp.Body.Close()
	
	}

	defer resp.Body.Close()

	var logcountMetrics Message
	json.NewDecoder(resp.Body).Decode(&logcountMetrics)

	return logcountMetrics.Message, nil

}

func (u *User) GetSpeedMetric() ([]LogSpeedMetric, error) {

	// Define a regular client
	regularClient := getJwtToken.RegularGet{}

	// Use the regular client to get the token from viper
	jwtToken, err := getJwtToken.WithType(regularClient, u.apiKey)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://logarbor.com/api/v1/logs/logs_speed_metric", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+jwtToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Something went wrong while fetching your log speed metrics. Please confirm that your api key is set correctly.")
	}

	if resp.StatusCode == 401 {

		// Define the expired get client
		expiredGet := getJwtToken.ExpiredGet{}

		// Use the client to fetch the new token
		newJwtToken, err := getJwtToken.WithType(expiredGet, u.apiKey)

		if err != nil {
			return nil, fmt.Errorf("Something went wrong. Try again or check your internet connection.")
		}

		req, err = http.NewRequest("GET", "https://logarbor.com/api/v1/logs/logs_speed_metric", nil)

		if err != nil {
			return nil, fmt.Errorf("Something went wrong. Try again or check your internet connection.")
		}

		req.Header.Set("Authorization", "Bearer "+newJwtToken)
		req.Header.Set("Content-Type", "application/json")

		resp, err = http.DefaultClient.Do(req)

		if err != nil {
			return nil, fmt.Errorf("Status Code: %d. Something went wrong. Try again or check your internet connection.", resp.StatusCode)
		}	

		defer resp.Body.Close()
	
	}

	defer resp.Body.Close()

	var logspeedMetrics LogSpeedMessage
	json.NewDecoder(resp.Body).Decode(&logspeedMetrics)

	return logspeedMetrics.Message, nil
}

func (u *User) GetErrorMetric() ([]LogErrorMetric, error) {

	// Define regular client
	regularClient := getJwtToken.RegularGet{}

	// Use the regular client to get token from viper
	jwtToken, err := getJwtToken.WithType(regularClient, u.apiKey)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://logarbor.com/api/v1/logs/error_rate_metric", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+jwtToken)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Something went wrong while fetching your log error rate metrics. Please confirm that your api key is set correctly.")
	}

	if resp.StatusCode == 401 {

		// Define the expired get client
		expiredGet := getJwtToken.ExpiredGet{}

		// Use the client to fetch the new token
		newJwtToken, err := getJwtToken.WithType(expiredGet, u.apiKey)

		if err != nil {
			return nil, fmt.Errorf("Something went wrong. Try again or check your internet connection.")
		}

		req, err = http.NewRequest("GET", "https://logarbor.com/api/v1/logs/error_rate_metric", nil)

		if err != nil {
			return nil, fmt.Errorf("Something went wrong. Try again or check your internet connection.")
		}

		req.Header.Set("Authorization", "Bearer "+newJwtToken)

		resp, err = http.DefaultClient.Do(req)

		if err != nil {
			return nil, fmt.Errorf("Status Code: %d. Something went wrong. Try again or check your internet connection.", resp.StatusCode)
		}	

		defer resp.Body.Close()
	
	}
	

	defer resp.Body.Close()

	var logerrorMessage LogErrorMessage
	json.NewDecoder(resp.Body).Decode(&logerrorMessage)

	return logerrorMessage.Message, nil
}
