package notion

import (
	"context"
)

const DATABASE_ID = "567c6989-c5d3-4a39-9e37-54c8636c11ba"

// Queries notion for the list of pages in the given database and returns the last created page
func GetLastPage(ctx context.Context) PageObject {
	query := &PageQuery{PageSize: 1}
	response := QueryDatabase(ctx, DATABASE_ID, query)

	return response.Results[0]
}
