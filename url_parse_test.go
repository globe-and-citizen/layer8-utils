package utilities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseURLPath(t *testing.T) {
	path, query := ParseURLPath("/init-tunnel?backend=https://wegotpoems.com")
	assert.Equal(t, "/init-tunnel", path)
	assert.Equal(t, "backend=https://wegotpoems.com", query)

	path, query = ParseURLPath("/poem?id=1")
	assert.Equal(t, "/poem", path)
	assert.Equal(t, "id=1", query)
}

func TestParseQueryParams(t *testing.T) {
	params := ParseQueryParams("key1=value1&key2=value2&key3=3")
	assert.Equal(t, map[string]string{"key1": "value1", "key2": "value2", "key3": "3"}, params)
}
