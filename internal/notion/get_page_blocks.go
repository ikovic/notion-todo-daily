package notion

import (
	"context"
)

const PAGE_ID = "d35d836d-fa4c-4929-b47e-0020ac9eb350"

// Queries notion for the list of pages in the given database and returns the last created page
func GetPageBlocks(ctx context.Context) []BlockObject {
	response := RetrieveBlockChildren(ctx, PAGE_ID)

	return response.Results
}
