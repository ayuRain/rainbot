package chat

import "fmt"

func GenMiniMaxHeader(ApiKey string) *map[string]string {
	return &map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", ApiKey),
		"Content-Type":  "application/json",
	}
}
