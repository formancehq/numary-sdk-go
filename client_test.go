package ledgerclient_test

import (
	"context"
	"fmt"
	ledgerclient "github.com/numary/numary-sdk-go"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestClient(t *testing.T) {
	config := ledgerclient.NewConfiguration()
	config.Servers = ledgerclient.ServerConfigurations{{
		URL: "http://localhost:3068",
	}}
	client := ledgerclient.NewAPIClient(config)
	createTransactionRequest := client.TransactionsApi.CreateTransaction(context.Background(), "quickstart")
	createTransactionRequest = createTransactionRequest.TransactionData(ledgerclient.TransactionData{
		Postings: []ledgerclient.Posting{
			{
				Amount:      100,
				Asset:       "USD",
				Destination: "player1",
				Source:      "world",
			},
		},
		Reference: ledgerclient.PtrString(fmt.Sprint(rand.Int())),
	})
	_, _, err := createTransactionRequest.Execute()
	if err != nil {
		t.Fatal(err)
	}

	listTransactionsRequest := client.TransactionsApi.ListTransactions(context.Background(), "quickstart")
	listTransactionsResponse, _, err := listTransactionsRequest.Execute()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, 0, listTransactionsResponse.Cursor.Total)
	assert.NotEqual(t, 0, listTransactionsResponse.Cursor.PageSize)

	tx := listTransactionsResponse.Cursor.Data[0]
	getTransactionRequest := client.TransactionsApi.GetTransaction(context.Background(), "quickstart", tx.Txid)
	res, _, err := getTransactionRequest.Execute()
	assert.NoError(t, err)
	assert.Equal(t, tx.Txid, res.Data.Txid)

	_, err = client.TransactionsApi.
		AddMetadataOnTransaction(context.Background(), "quickstart", tx.Txid).
		RequestBody(map[string]interface{}{
			"foo": "bar",
		}).
		Execute()
	assert.NoError(t, err)

	_, err = client.TransactionsApi.RevertTransaction(context.Background(), "quickstart", tx.Txid).Execute()
	assert.NoError(t, err)

	_, _, err = client.MappingApi.
		UpdateMapping(context.Background(), "quickstart").
		Mapping(ledgerclient.Mapping{}).
		Execute()
	assert.NoError(t, err)

	m, _, err := client.MappingApi.GetMapping(context.Background(), "quickstart").Execute()
	assert.NoError(t, err)

	assert.Len(t, m.Data.Contracts, 0)

	_, _, err = client.ServerApi.GetInfo(context.Background()).Execute()
	assert.NoError(t, err)

	_, _, err = client.StatsApi.ReadStats(context.Background(), "quickstart").Execute()
	assert.NoError(t, err)

	listAccountsResponse, _, err := client.AccountsApi.ListAccounts(context.Background(), "quickstart").Execute()
	assert.NoError(t, err)

	firstAccount := listAccountsResponse.Cursor.Data[0]

	_, _, err = client.AccountsApi.GetAccount(context.Background(), "quickstart", firstAccount.Address).Execute()
	assert.NoError(t, err)

	_, err = client.AccountsApi.
		AddMetadataToAccount(context.Background(), "quickstart", firstAccount.Address).
		RequestBody(map[string]interface{}{
			"foo": "bar",
		}).
		Execute()
	assert.NoError(t, err)

	_, _, err = client.ScriptApi.RunScript(context.Background(), "quickstart").Script(ledgerclient.Script{
		Plain: `send [COIN 100] (
  source = @world
  destination = @centralbank
)
send [COIN 100] (
  source = @centralbank
  destination = @users:001
)`,
	}).Execute()
	assert.NoError(t, err)

}
