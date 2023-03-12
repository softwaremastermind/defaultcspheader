package defaultcspheader_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestIntegration(t *testing.T) {
	identifier := tc.StackIdentifier("traefik_test")
	compose, err := tc.NewDockerComposeWith(tc.WithStackFiles("./docker-compose.integration.yml"), identifier)
	assert.NoError(t, err, "NewDockerComposeAPIWith()")

	t.Cleanup(func() {
		assert.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "compose.Down()")
	})

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	err = compose.
		WaitForService("mockserver", wait.NewLogStrategy("started on port: 1080")).Up(ctx, tc.Wait(true))
	assert.NoError(t, err, "compose.Up()")
	fmt.Print("Docker Compose based test startegy is up\n")
	sc, err := compose.ServiceContainer(ctx, "traefik")
	assert.NoError(t, err, "DockerContainer()")
	endpoint, err := sc.PortEndpoint(ctx, "8000", "http")
	fmt.Printf("endpoint for tests is : %v\n", endpoint)
	assert.NoError(t, err, "PortEndpoint()")
	endpointTest := checkcspheader(endpoint)
	endpointTest.TestSubpath(ctx, t, "path1", "test-middleware")
	endpointTest.TestSubpath(ctx, t, "path2", "test-upstream")

}

type checkcspheader string

func (c checkcspheader) TestSubpath(ctx context.Context, t *testing.T, subpath string, expectedCSPHeader string) {
	url, err := url.JoinPath(string(c), subpath)
	assert.NoError(t, err, "Cannot Compose path")
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	assert.NoError(t, err, "Create New Request")
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "cannot execute http request")

	responseCSP := resp.Header.Get("Content-Security-Policy")
	assert.Equal(t, expectedCSPHeader, responseCSP, "url: %s", url)
}
