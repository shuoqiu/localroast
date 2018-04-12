package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caalberts/localghost"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	port := "8888"
	server := NewServer(port, localghost.Schema{Path: "/"})
	assert.Equal(t, ":8888", server.Addr)
}

func TestNewMux(t *testing.T) {
	schemas := []localghost.Schema{
		localghost.Schema{
			Method:     "GET",
			Path:       "/",
			StatusCode: 200,
		},
		localghost.Schema{
			Method:     "POST",
			Path:       "/user",
			StatusCode: 201,
		},
	}

	mux := NewMux(schemas)
	server := httptest.NewServer(mux)
	defer server.Close()

	for _, schema := range schemas {

		var resp *http.Response
		var err error

		switch schema.Method {
		case http.MethodGet:
			resp, err = http.Get(server.URL + schema.Path)
		case http.MethodPost:
			resp, err = http.Post(server.URL+schema.Path, "", nil)
		}

		assert.Nil(t, err)
		assert.Equal(t, schema.StatusCode, resp.StatusCode)
	}
}
