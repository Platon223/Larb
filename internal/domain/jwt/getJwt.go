package getJwtToken

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type JWT struct {
	Message     string `json:"message"`
	AccessToken string `json:"actk"`
}

type WithTypeProvider interface {
	ReturnType() string
}

type RegularGet struct {
}

type ExpiredGet struct {
}

func (r RegularGet) ReturnType() string {

	return "regular"
}

func (e ExpiredGet) ReturnType() string {

	return "expired"
}

func WithType(w WithTypeProvider, apiKey string) (string, error) {
	token, err := GetJwt(apiKey, w.ReturnType())
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetJwt(apiKey string, getType string) (string, error) {

	switch getType {
	case "regular":
		token := viper.GetString("jwt")
		return token, nil
	case "expired":
		body, _ := json.Marshal(map[string]string{
			"user_id": apiKey,
		})

		// Timeout if the api hangs

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(
			ctx,
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

		if err != nil {
			if ctx.Err() == context.DeadlineExceeded {
				return "", errors.New("Request timed out. Please check your connection or try again.")
			}
			return "", err
		}

		if resp.StatusCode != 200 {
			return "", errors.New("Something went wrong while getting your jwt key. Please make sure you api key is set correctly.")
		}

		defer resp.Body.Close()

		var jwt JWT
		json.NewDecoder(resp.Body).Decode(&jwt)

		// Save to viper
		viper.Set("jwt", jwt.AccessToken)

		err = viper.WriteConfig()

		if err != nil {
			viper.SafeWriteConfig()
		}

		return jwt.AccessToken, nil
	default:
		return "", errors.New("Invalid jwt type get provided")

	}

}
