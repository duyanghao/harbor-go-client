package swagger

import (
	"context"
	"crypto/tls"
	"github.com/antihax/optional"
	"net/http"
	"testing"
)

func TestNewAPIClient(t *testing.T) {
	cfg := &Configuration{
		BasePath:      "https://your.harbor.domain.name/api",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
	client := NewAPIClient(cfg)
	auth := context.WithValue(context.Background(), ContextBasicAuth, BasicAuth{
		UserName: "admin",
		Password: "Harbor12345",
	})
	projs, _, err := client.ProductsApi.ProjectsGet(auth, &ProjectsGetOpts{
		Public: optional.NewBool(true),
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, proj := range projs {
		t.Logf(proj.Name)
	}
}
