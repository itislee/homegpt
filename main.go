package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("%s\n", "running")
	port, err := strconv.Atoi(os.Getenv("HOMEASSISTANT_PORT"))
	if err != nil {
		fmt.Printf("using default port 8123\n")
		port = 8123
	}
	ip := os.Getenv("HOMEASSISTANT_IP")
	if ip == "" {
		ip = "192.168.1.64"

	}
	token := os.Getenv("HOMEASSISTANT_TOKEN")
	keyid := os.Getenv("WENXIN_KEYID")
	keysecret := os.Getenv("WENXIN_KEYSECRET")

	fmt.Printf("ip:%s\nport:%d\ntoken:%s\nkeyid:%s\nkeysecret:%s\n",
		ip, port, token, keyid, keysecret)
	api := NewHomeAssistant(ip,
		port,
		token,
		NewOpenAIBot(os.Getenv("OPENAI_TOKEN")))

	api.Loop()
}
