package defaultcspheader_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/softwaremastermind/defaultcspheader"
)

type CSPTest struct {
	UpstreamCSPHeader string
	DefaultCSPHeader  string
	ExpectedCSPHeader string
}

func TestCSPHeaderMiddleware(t *testing.T) {
	tests := []CSPTest{
		{UpstreamCSPHeader: "", DefaultCSPHeader: "default_csp_value", ExpectedCSPHeader: "default_csp_value"},
		{UpstreamCSPHeader: "upstream_csp_value", DefaultCSPHeader: "default_csp_value", ExpectedCSPHeader: "upstream_csp_value"},
		{UpstreamCSPHeader: "upstream_csp_value", DefaultCSPHeader: "", ExpectedCSPHeader: "upstream_csp_value"},
	}
	for _, test := range tests {

		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Security-Policy", test.UpstreamCSPHeader)
			w.WriteHeader(200)
		})

		config := defaultcspheader.CreateConfig()
		config.DefaultCSPHeader = test.DefaultCSPHeader
		ctx := context.Background()
		plugin, err := defaultcspheader.New(ctx, next, config, "test123")

		if err != nil {
			t.Errorf("Failed to create plugin: %v", err)
		}

		recorder := httptest.NewRecorder()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
		if err != nil {
			t.Errorf("Error at creating request. Inner Error %s", err)
		}

		plugin.ServeHTTP(recorder, req)

		actualCSPHeader := recorder.Header().Get(defaultcspheader.CSPHeaderKey)
		if actualCSPHeader != test.ExpectedCSPHeader {
			t.Errorf("For default header '%s' and upstream CSP header '%s', expected '%s', got '%s'", test.DefaultCSPHeader, test.UpstreamCSPHeader, test.ExpectedCSPHeader, actualCSPHeader)

		}
	}
}
