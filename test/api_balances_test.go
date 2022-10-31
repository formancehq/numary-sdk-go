/*
Ledger API

Testing BalancesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ledgerclient

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    client "./openapi"
)

func Test_ledgerclient_BalancesApiService(t *testing.T) {

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)

    t.Run("Test BalancesApiService GetBalances", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ledger string

        resp, httpRes, err := apiClient.BalancesApi.GetBalances(context.Background(), ledger).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BalancesApiService GetBalancesAggregated", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ledger string

        resp, httpRes, err := apiClient.BalancesApi.GetBalancesAggregated(context.Background(), ledger).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
