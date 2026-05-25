package osv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const osvURL = "https://api.osv.dev/v1/query"

type QueryReq struct {
	Package struct {
		Name      string `json:"name"`
		Ecosystem string `json:"ecosystem"`
	} `json:"package"`
	Version string `json:"version"`
}

type Vuln struct {
	ID       string `json:"id"`
	Summary  string `json:"summary,omitempty"`
	Details  string `json:"details,omitempty"`
	Refs     []any  `json:"references,omitempty"`
	Severity string `json:"severity,omitempty"`
}

type Client struct {
	httpClient *http.Client
	cache      sync.Map
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 15 * time.Second},
	}
}

func (c *Client) cacheKey(ecosys, name, version string) string {
	return ecosys + "|" + name + "|" + version
}

func (c *Client) Query(ecosys, name, version string) ([]Vuln, error) {
	key := c.cacheKey(ecosys, name, version)
	if v, ok := c.cache.Load(key); ok {
		return v.([]Vuln), nil
	}

	req := QueryReq{}
	req.Package.Name = name
	req.Package.Ecosystem = ecosys
	req.Version = version
	bs, _ := json.Marshal(req)

	var lastErr error
	for i := 0; i < 3; i++ {
		resp, err := c.httpClient.Post(osvURL, "application/json", bytes.NewReader(bs))
		if err != nil {
			lastErr = err
			time.Sleep(time.Duration(200*(i+1)) * time.Millisecond)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 429 {
			time.Sleep(time.Duration(500*(i+1)) * time.Millisecond)
			continue
		}

		if resp.StatusCode != 200 {
			lastErr = fmt.Errorf("osv: status %d", resp.StatusCode)
			time.Sleep(time.Duration(200*(i+1)) * time.Millisecond)
			continue
		}

		var out struct {
			Vulns []struct {
				ID       string `json:"id"`
				Summary  string `json:"summary"`
				Details  string `json:"details"`
				Severity interface{} `json:"severity"`
				Refs     []any  `json:"references"`
			} `json:"vulns"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
			lastErr = err
			time.Sleep(time.Duration(200*(i+1)) * time.Millisecond)
			continue
		}

		var res []Vuln
		for _, v := range out.Vulns {
			sev := "UNKNOWN"
			switch vv := v.Severity.(type) {
			case string:
				sev = vv
			case []interface{}:
				if len(vv) > 0 {
					if s, ok := vv[0].(string); ok {
						sev = s
					}
				}
			}
			res = append(res, Vuln{
				ID: v.ID, Summary: v.Summary, Details: v.Details,
				Refs: v.Refs, Severity: sev,
			})
		}
		c.cache.Store(key, res)
		return res, nil
	}
	return nil, lastErr
}