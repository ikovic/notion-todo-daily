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

const API_BASE_URL = "https://api.notion.com/v1"

type PageObject struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

type Response struct {
	Object  string `json:"object"`
	Results []PageObject
}

type PageQuery struct {
	PageSize int32 `json:"page_size"`
}

var NotionClient = cleanhttp.DefaultPooledClient()

func prepareHeaders(req *http.Request, authToken string) {
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Notion-Version", "2021-08-16")
	req.Header.Set("Content-Type", "application/json")
}

func QueryDatabase(ctx context.Context, dbId string, query *PageQuery) Response {
	authToken := ctx.Value(ntdCtx.ContextKey("AUTH_TOKEN")).(string)

	jsonBody, _ := json.Marshal(query)
	dbApiUrl := fmt.Sprintf("%s/databases/%s/query", API_BASE_URL, dbId)

	req, _ := http.NewRequest("POST", dbApiUrl, bytes.NewBuffer(jsonBody))
	prepareHeaders(req, authToken)

	resp, err := NotionClient.Do(req)
	if err != nil {
		log.Panicf("Could not reach API %v\n", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Panicf("Could not read API response %v\n", err)
	}

	response := GetResponseDTO(body)

	return response
}

func GetResponseDTO(body []byte) Response {
	var response Response
	json.Unmarshal(body, &response)
	return response
}
