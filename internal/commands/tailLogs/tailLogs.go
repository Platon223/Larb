package taillogs

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/zishang520/socket.io-client-go/socket"
)

type User struct {
	apiKey string
}

type NewLog struct {
	Message   string `json:"message"`
	UserId    string `json:"user_id"`
	ServiceId string `json:"service_id"`
}

func ConfigUser(apiKey string) *User {
	user := &User{
		apiKey: apiKey,
	}

	return user
}

func (u *User) TailLogs(serviceId string) error {

	manager := socket.NewManager("https://logarbor.com", nil)
	c := socket.NewSocket(manager, "/", nil)

	c.On("new-log", func(args ...any) {
		jsonBytes, _ := json.Marshal(args[0])

		var log NewLog
		if err := json.Unmarshal(jsonBytes, &log); err != nil {
			fmt.Println("Error occured while encoding the socket event.")

			return
		}

		if log.UserId == u.apiKey && log.ServiceId == serviceId {
			fmt.Println(log.Message)
		}
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Listening for logs... (press Ctrl+C to stop)")
	<-quit

	fmt.Println("\nStopping...")
	c.Close()

	return nil
}
