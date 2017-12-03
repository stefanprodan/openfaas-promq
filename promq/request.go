package function

import (
	"encoding/json"
	"errors"
	"os"
)

type Request struct {
	Format string
	Server string
	Query  string
	Start  string
	End    string
	Step   string
}

func NewRequest(data []byte) (*Request, error) {
	r := &Request{}
	if err := json.Unmarshal(data, r); err != nil {
		return nil, err
	}

	promURL := os.Getenv("PROMETHEUS_URL")
	if r.Server == "" && len(promURL) > 0 {
		r.Server = promURL
	}

	if r.Server == "" {
		return nil, errors.New("no Prometheus server specified")
	}

	if r.Query == "" {
		return nil, errors.New("no query specified")
	}

	if r.Start == "" {
		r.Start = "1 hour ago"
	}

	if r.End == "" {
		r.End = "now"
	}

	if r.Step == "" {
		r.Step = "1m"
	}

	return r, nil
}
