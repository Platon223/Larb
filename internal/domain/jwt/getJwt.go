package getJwtToken

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type JWT struct {
	Message     string `json:"message"`
	AccessToken string `json:"actk"`
}

func GetJwt(apiKey string) (string, error) {

	body, _ := json.Marshal(map[string]string{
		"user_id": apiKey,
	})

	req, err := http.NewRequest(
		"POST",
		"https://logarbor.com/auth/jwt",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return "", errors.New("Error occured while fetching the jwt")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != err {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("Something went wrong while getting your jwt key. Please make sure you api key is set correctly.")
	}

	defer resp.Body.Close()

	var jwt JWT
	json.NewDecoder(resp.Body).Decode(&jwt)

	return jwt.AccessToken, nil
}
