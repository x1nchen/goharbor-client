package repository

import (
	"context"

	"github.com/mittwald/goharbor-client/v3/apiv1/internal/api/client"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v3/apiv1/internal/api/client/products"
	model "github.com/mittwald/goharbor-client/v3/apiv1/model"
)

// RESTClient is a subclient for handling repository related actions.
type RESTClient struct {
	// The swagger client
	Client *client.Harbor

	// AuthInfo contain auth information, which are provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(cl *client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Client:   cl,
		AuthInfo: authInfo,
	}
}

type Client interface {
	GetRepositoryTags(ctx context.Context, repoName string) ([]*model.DetailedTag, error)
}

// GetRepositoryTags
// ref https://github.com/goharbor/harbor/blob/v1.10.0/api/harbor/swagger.yaml#L1004-L1045
func (c *RESTClient) GetRepositoryTags(ctx context.Context, repoName string) ([]*model.DetailedTag, error) {
	result, err := c.Client.Products.GetRepositoryTags(&products.GetRepositoryTagsParams{
		RepoName: &repoName,
		Context:  ctx,
	}, c.AuthInfo)

	if err != nil {
		return nil, err
	}

	return result.GetPayload(), nil
}
