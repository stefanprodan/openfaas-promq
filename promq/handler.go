package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	r, err := NewRequest(req)
	if err != nil {
		return fmt.Sprintf("Request parse error: %v", err)
	}

	start, end, step, err := r.GetQueryRange()
	if err != nil {
		return fmt.Sprintf("Query range parse error: %v", err)
	}

	c, err := NewClient(r.Server, "", "")
	if err != nil {
		return fmt.Sprintf("HTTP client error: %v", err)
	}

	response, err := c.QueryRange(r.Query, start, end, step)
	if err != nil {
		return fmt.Sprintf("Query error: %v", err)
	}

	json, err := responseJSON(response)
	if err != nil {
		return fmt.Sprintf("Query result parse error: %v", err)
	}

	return json
}
