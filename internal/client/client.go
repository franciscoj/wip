package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Get(ctx context.Context, url string, res interface{}) error {
	client := http.Client{}
	req, err := getReq(ctx, url)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("requesting: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}
	json.Unmarshal(body, res)

	return nil
}

func getReq(ctx context.Context, url string) (*http.Request, error) {
	token := os.Getenv("TODOIST_TOKEN")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Add("Authorization", "Bearer "+token)

	return req, nil
}
