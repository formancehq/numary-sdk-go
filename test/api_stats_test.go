/*
Ledger API

Testing StatsApiService

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

func Test_ledgerclient_StatsApiService(t *testing.T) {

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)

    t.Run("Test StatsApiService ReadStats", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ledger string

        resp, httpRes, err := apiClient.StatsApi.ReadStats(context.Background(), ledger).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
