

package repository

import (
	"context"
	"net/url"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/mittwald/goharbor-client/v3/apiv1/internal/api/client"
	integrationtest "github.com/mittwald/goharbor-client/v3/apiv1/testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/require"
)

var (
	u, _          = url.Parse(integrationtest.Host)
	swaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo      = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
)

func TestAPIRepositoryTagsGet(t *testing.T) {
	ctx := context.Background()


	c := NewClient(swaggerClient, authInfo)
	tagList, err := c.GetRepositoryTags(ctx, "copytrade/go-report-grpc-srv")

	require.NoError(t, err)
	require.NotNil(t, tagList)
	for _, tag := range tagList {
		t.Logf("%+v", tag)
	}
}

