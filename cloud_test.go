package ledgerclient_test

import (
	"context"
	ledgerclient "github.com/numary/numary-sdk-go"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCloudAuth(t *testing.T) {
	tok, err := ledgerclient.FetchToken(
		http.DefaultClient,
		ledgerclient.StagingAuthEndpoint,
		"YjgyY2ZjMDktNjZmYi.OWFhNWNjOGQtMTNlMS",
	)
	assert.NoError(t, err)

	client := newClient()
	_, _, err = client.TransactionsApi.ListTransactions(
		context.WithValue(context.Background(), ledgerclient.ContextAccessToken, tok),
		"tirelire-88550a7f",
	).Execute()
	assert.NoError(t, err)
}
