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
	api := NewHomeAssistant(os.Getenv("HOMEASSISTANT_IP"),
	port,
	os.Getenv("HOMEASSISTANT_TOKEN"), 
	NewWenXinBot(os.Getenv("WENXIN_KEYID"), os.Getenv("WENXIN_KEYSECRET")))

	api.Loop()
}
