package ledgerclient_test

import (
	"context"
	"fmt"
	ledgerclient "github.com/numary/numary-sdk-go"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func newClient() *ledgerclient.APIClient {
	config := ledgerclient.NewConfiguration()
	config.Servers = ledgerclient.ServerConfigurations{{
		URL: "http://localhost:3068",
	}}
	return ledgerclient.NewAPIClient(config)
}

func TestClient(t *testing.T) {
	client := newClient()
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

	assert.Len(t, m.Data.Get().Contracts, 0)

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

func TestCreateTransaction(t *testing.T) {
	type testCase struct {
		name          string
		data          []ledgerclient.TransactionData
		expectedError ledgerclient.ErrorCode
	}

	cases := []testCase{
		{
			name: "nominal",
			data: []ledgerclient.TransactionData{
				{
					Postings: []ledgerclient.Posting{
						{
							Amount:      100,
							Asset:       "USD",
							Destination: "player1",
							Source:      "world",
						},
					},
				},
			},
		},
		{
			name: "ref-conflict",
			data: []ledgerclient.TransactionData{
				{
					Postings: []ledgerclient.Posting{
						{
							Amount:      100,
							Asset:       "USD",
							Destination: "player1",
							Source:      "world",
						},
					},
					Reference: ledgerclient.PtrString("ref"),
				},
				{
					Postings: []ledgerclient.Posting{
						{
							Amount:      100,
							Asset:       "USD",
							Destination: "player1",
							Source:      "world",
						},
					},
					Reference: ledgerclient.PtrString("ref"),
				},
			},
			expectedError: ledgerclient.CONFLICT,
		},
	}

	client := newClient()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ledger := uuid.New()
			for i := 0; i < len(c.data)-1; i++ {
				_, _, err := client.TransactionsApi.
					CreateTransaction(context.Background(), ledger).
					TransactionData(c.data[i]).
					Execute()
				assert.NoError(t, err)
			}

			_, _, err := client.TransactionsApi.
				CreateTransaction(context.Background(), ledger).
				TransactionData(c.data[len(c.data)-1]).
				Execute()
			if c.expectedError != "" {
				assert.Error(t, err)
				assert.IsType(t, ledgerclient.GenericOpenAPIError{}, err)
				assert.Equal(t, c.expectedError, err.(ledgerclient.GenericOpenAPIError).Model().(ledgerclient.ErrorResponse).ErrorCode)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestTransactionBatch(t *testing.T) {

	type testCase struct {
		name           string
		batch          []ledgerclient.TransactionData
		expectedErrors []ledgerclient.ErrorCode
	}

	cases := []testCase{
		{
			name: "nominal",
			batch: []ledgerclient.TransactionData{
				{
					Postings: []ledgerclient.Posting{
						{
							Amount:      100,
							Asset:       "USD",
							Destination: "player1",
							Source:      "world",
						},
					},
				},
				{
					Postings: []ledgerclient.Posting{
						{
							Amount:      100,
							Asset:       "USD",
							Destination: "player2",
							Source:      "world",
						},
					},
				},
			},
		},
		{
			name: "conflict",
			batch: []ledgerclient.TransactionData{
				{
					Postings: []ledgerclient.Posting{
						{
							Amount:      100,
							Asset:       "USD",
							Destination: "player1",
							Source:      "world",
						},
					},
					Reference: ledgerclient.PtrString("ref"),
				},
				{
					Postings: []ledgerclient.Posting{
						{
							Amount:      100,
							Asset:       "USD",
							Destination: "player2",
							Source:      "world",
						},
					},
					Reference: ledgerclient.PtrString("ref"),
				},
			},
			expectedErrors: []ledgerclient.ErrorCode{"", ledgerclient.CONFLICT},
		},
	}

	client := newClient()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, _, err := client.TransactionsApi.
				CreateTransactions(context.Background(), uuid.New()).
				Transactions(ledgerclient.Transactions{
					Transactions: c.batch,
				}).
				Execute()

			if len(c.expectedErrors) > 0 {
				assert.IsType(t, ledgerclient.GenericOpenAPIError{}, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
