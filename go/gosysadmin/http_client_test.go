package gosysadmin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequest(t *testing.T) {
	t.Parallel()
	customHeader := map[string]string{
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0",
		"accept-language": "en-US,en;q=0.6",
	}
	respBody := GetRequest("https://google.com", customHeader)
	assert.Contains(t, respBody, "ag")
}
