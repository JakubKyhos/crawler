package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(baseURL string) (string, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	normUrl := strings.ToLower(url.Host) + strings.ToLower(strings.TrimRight(url.Path, "/"))
	return normUrl, nil
}
