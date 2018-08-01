package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReqIDGenerator(t *testing.T) {
	ts := httptest.NewServer(ReqIDGenerator(GetTestHandler()))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.NotNil(t, resp.Header.Get("X-Request-ID"))
}
