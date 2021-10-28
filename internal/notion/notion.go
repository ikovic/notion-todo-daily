package notion

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/hashicorp/go-cleanhttp"
)

var client = cleanhttp.DefaultPooledClient()

const API_BASE_URL = "https://api.notion.com/v1"

var AUTH_TOKEN = os.Getenv("AUTH_TOKEN")

func SearchPages() {
	req, err := http.NewRequest("POST", API_BASE_URL+"/search", nil)

	fmt.Println(AUTH_TOKEN)

	if err != nil {
		// handle error
		fmt.Printf("Could not create API request %v\n", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+AUTH_TOKEN)
	req.Header.Set("Notion-Version", "2021-08-16")
	resp, err := client.Do(req)

	if err != nil {
		// handle error
		fmt.Printf("Could not reach API %v\n", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Could not read API response %v\n", err)
		return
	}

	fmt.Println(string(body))
}
