package initCommd

import (
	"github.com/spf13/viper"
)

func Init(apiKey string) string {

	viper.Set("apiKey", apiKey)

	err := viper.WriteConfig()

	if err != nil {
		viper.SafeWriteConfig()
	}

	return "Success"
}
