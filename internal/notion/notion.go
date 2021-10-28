package notion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hashicorp/go-cleanhttp"
	ntdCtx "github.com/ikovic/notion-todo-daily/internal/ctx"
)

var client = cleanhttp.DefaultPooledClient()

const API_BASE_URL = "https://api.notion.com/v1"
const DATABASE_ID = "567c6989-c5d3-4a39-9e37-54c8636c11ba"

type PageQuery struct {
	PageSize int32 `json:"page_size"`
}

func SearchPages(ctx context.Context) {
	// TODO we probably want to have a custom context to simplify this
	authToken := ctx.Value(ntdCtx.ContextKey("AUTH_TOKEN")).(string)
	query := &PageQuery{PageSize: 1}
	jsonBody, _ := json.Marshal(query)
	dbApiUrl := fmt.Sprintf("%s/databases/%s/query", API_BASE_URL, DATABASE_ID)

	req, err := http.NewRequest("POST", dbApiUrl, bytes.NewBuffer(jsonBody))

	if err != nil {
		log.Panicf("Could not create API request %v\n", err)
	}

	// TODO needs to be reused across all the requests
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Notion-Version", "2021-08-16")
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Panicf("Could not reach API %v\n", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Panicf("Could not read API response %v\n", err)
	}

	fmt.Println(string(body))
}
