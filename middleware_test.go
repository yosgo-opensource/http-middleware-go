package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockMux struct{}

func ServeHTTP(rw http.ResponseWriter, r *http.Request) {}

func TestHTTPCORSHeaders(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest(http.MethodPost, "/_watermark", nil)
	w := httptest.NewRecorder()

	cors := HttpCORS(ServeHTTP)
	cors(w, req)

	res := w.Result()
	assert.Equal(res.Header.Get("Access-Control-Allow-Origin"), "*")
	assert.Equal(res.Header.Get("Access-Control-Allow-Headers"), "Content-Type")
}

func TestJSONResponseHeaders(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest(http.MethodPost, "/_watermark", nil)
	w := httptest.NewRecorder()

	cors := HttpJSONResponse(ServeHTTP)
	cors(w, req)

	res := w.Result()
	assert.Equal(res.Header.Get("Content-Type"), "application/json")
}
