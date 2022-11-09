/*
Ledger API

Testing AccountsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ledgerclient

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    client "github.com/numary/numary-sdk-go"
)

func Test_ledgerclient_AccountsApiService(t *testing.T) {

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)

    t.Run("Test AccountsApiService AddMetadataToAccount", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ledger string
        var address string

        resp, httpRes, err := apiClient.AccountsApi.AddMetadataToAccount(context.Background(), ledger, address).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AccountsApiService CountAccounts", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ledger string

        resp, httpRes, err := apiClient.AccountsApi.CountAccounts(context.Background(), ledger).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AccountsApiService GetAccount", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ledger string
        var address string

        resp, httpRes, err := apiClient.AccountsApi.GetAccount(context.Background(), ledger, address).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AccountsApiService ListAccounts", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ledger string

        resp, httpRes, err := apiClient.AccountsApi.ListAccounts(context.Background(), ledger).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
