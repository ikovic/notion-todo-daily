package notion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func GetQueryDatabaseRequest(ctx context.Context, dbId string, query *PageQuery) *http.Request {
	authToken := ctx.Value(ntdCtx.ContextKey("AUTH_TOKEN")).(string)

	jsonBody, _ := json.Marshal(query)
	dbApiUrl := fmt.Sprintf("%s/databases/%s/query", API_BASE_URL, dbId)

	req, _ := http.NewRequest("POST", dbApiUrl, bytes.NewBuffer(jsonBody))
	prepareHeaders(req, authToken)

	return req
}
