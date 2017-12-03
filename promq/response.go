package function

import (
	"encoding/json"
)

func responseJSON(resp *QueryRangeResponse) (string, error) {
	type valueEntry struct {
		Metric map[string]string `json:"metric"`
		Value  float64           `json:"value"`
	}
	type timeEntry struct {
		Time   int64         `json:"time"`
		Values []*valueEntry `json:"values"`
	}
	entryByTime := map[int64]*timeEntry{}

	for _, r := range resp.Data.Result {
		for _, v := range r.Values {
			t := v.Time()
			u := t.Unix()
			e, ok := entryByTime[u]
			if !ok {
				e = &timeEntry{
					Time:   u,
					Values: []*valueEntry{},
				}
				entryByTime[u] = e
			}

			val, err := v.Value()
			if err != nil {
				return "", err
			}
			e.Values = append(e.Values, &valueEntry{
				Metric: r.Metric,
				Value:  val,
			})
		}
	}

	s := make([]*timeEntry, len(entryByTime))
	i := 0
	for _, e := range entryByTime {
		s[i] = e
		i++
	}

	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
