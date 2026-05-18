package initCommd

import (
	"fmt"
	"net/http"

	"github.com/Platon223/Larb/internal/domain/jwt"
	"github.com/spf13/viper"
)

func Init(apiKey string) string {

	jwtToken, err := getJwtToken.GetJwt(apiKey, "expired")

	if err != nil {
		return "Please check your internet connection."
	}

	req, err := http.NewRequest("POST", "https://logarbor.com/api/v1/home/credentials/username", nil)

	if err != nil {
		return "Please check your internet connection."
	}

	req.Header.Set("Authorization", "Bearer "+jwtToken)

	resp, err := http.DefaultClient.Do(req)

	if resp.StatusCode != 200 {
		return fmt.Sprintf("Status Code: %d. Please provide a valid api key or check your internet connection.", resp.StatusCode)
	}

	defer resp.Body.Close()

	viper.Set("apiKey", apiKey)

	err = viper.WriteConfig()

	if err != nil {
		viper.SafeWriteConfig()
	}

	return "Success"
}
