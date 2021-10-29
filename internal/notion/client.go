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

type ToDO struct {
	Text []struct {
		Type string `json:"type"`
		Text struct {
			Content string      `json:"content"`
			Link    interface{} `json:"link"`
		} `json:"text"`
		Annotations struct {
			Bold          bool   `json:"bold"`
			Italic        bool   `json:"italic"`
			Strikethrough bool   `json:"strikethrough"`
			Underline     bool   `json:"underline"`
			Code          bool   `json:"code"`
			Color         string `json:"color"`
		} `json:"annotations"`
		PlainText string      `json:"plain_text"`
		Href      interface{} `json:"href"`
	} `json:"text"`
	Checked bool `json:"checked"`
}

type PageObject struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// https://developers.notion.com/reference/block
type BlockObject struct {
	Object      string `json:"object"`
	Id          string `json:"id"`
	HasChildren bool   `json:"has_children"`
	Type        string `json:"type"`
	ToDo        ToDO   `json:"to_do"`
}

type PageResponseDTO struct {
	Object  string `json:"object"`
	Results []PageObject
}

type BlockResponseDTO struct {
	Object  string `json:"object"`
	Results []BlockObject
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

func QueryDatabase(ctx context.Context, dbId string, query *PageQuery) PageResponseDTO {
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

	response := getPageResponseDTO(body)

	return response
}

func RetrieveBlockChildren(ctx context.Context, blockId string) BlockResponseDTO {
	authToken := ctx.Value(ntdCtx.ContextKey("AUTH_TOKEN")).(string)
	retrieveBlockApiUrl := fmt.Sprintf("%s/blocks/%s/children", API_BASE_URL, blockId)

	req, _ := http.NewRequest("GET", retrieveBlockApiUrl, nil)
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

	response := getBlockResponseDTO(body)

	return response
}

func getPageResponseDTO(body []byte) PageResponseDTO {
	var response PageResponseDTO
	json.Unmarshal(body, &response)
	return response
}

func getBlockResponseDTO(body []byte) BlockResponseDTO {
	var response BlockResponseDTO
	json.Unmarshal(body, &response)
	return response
}
