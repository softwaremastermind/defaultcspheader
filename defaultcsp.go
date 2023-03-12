package defaultcspheader

import (
	"context"
	"net/http"
	"os"
)

const CSPHeaderKey = "Content-Security-Policy"

type Config struct {
	DefaultCSPHeader string `json:"defaultCspHeader,omitempty"`
}

func CreateConfig() *Config {
	return &Config{
		DefaultCSPHeader: "",
	}
}

type DefaultCSPHeaderPlugin struct {
	next   http.Handler
	name   string
	config Config
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	os.Stdout.WriteString("\n\nStarted DefaultCSPHeader plugin\n\n")
	return &DefaultCSPHeaderPlugin{
		next:   next,
		config: *config,
		name:   name,
	}, nil
}

// ServeHTTP implements http.Handler
func (d *DefaultCSPHeaderPlugin) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	os.Stdout.WriteString("\nDefaultCSPHeader plugin: call upstream\n\n")

	mrw := MyResponseWriter{
		w, d.config.DefaultCSPHeader,
	}
	d.next.ServeHTTP(&mrw, r)

}

type MyResponseWriter struct {
	http.ResponseWriter
	DefaultCSPEntry string
}

func (m *MyResponseWriter) WriteHeader(statusCode int) {
	actualCSP := m.ResponseWriter.Header().Get(CSPHeaderKey)

	if actualCSP == "" {
		m.ResponseWriter.Header().Del(CSPHeaderKey)
		m.ResponseWriter.Header().Add(CSPHeaderKey, m.DefaultCSPEntry)
	}
	m.ResponseWriter.WriteHeader(statusCode)
}
