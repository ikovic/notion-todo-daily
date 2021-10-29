package notion

import (
	"context"
	"encoding/json"
	"io"
	"log"
)

const DATABASE_ID = "567c6989-c5d3-4a39-9e37-54c8636c11ba"

// Queries notion for the list of pages in the given database and returns the last created page
func GetLastPage(ctx context.Context) PageObject {
	// TODO we probably want to have a custom context to simplify this
	query := &PageQuery{PageSize: 1}
	req := GetQueryDatabaseRequest(ctx, DATABASE_ID, query)

	resp, err := NotionClient.Do(req)
	if err != nil {
		log.Panicf("Could not reach API %v\n", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Panicf("Could not read API response %v\n", err)
	}

	var response Response
	json.Unmarshal(body, &response)

	return response.Results[0]
}
