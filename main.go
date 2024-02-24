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
		fmt.Printf("can't parse HOMEASSISTANT_PORT\n")
		return
	}
	ip := os.Getenv("HOMEASSISTANT_IP")
	token := os.Getenv("HOMEASSISTANT_TOKEN")
	keyid := os.Getenv("WENXIN_KEYID")
	keysecret := os.Getenv("WENXIN_KEYSECRET")
	fmt.Printf("ip:%s\nport:%d\ntoken:%s\nkeyid:%s\nkeysecret:%s\n",
		ip, port, token, keyid, keysecret)
	api := NewHomeAssistant(ip,
		port,
		token,
		NewWenXinBot(keyid, keysecret))

	api.Loop()
}
