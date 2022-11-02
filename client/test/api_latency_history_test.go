/*
Meraki Dashboard API

Testing LatencyHistoryApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_client_LatencyHistoryApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test LatencyHistoryApiService GetNetworkWirelessClientLatencyHistory", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var clientId string

        resp, httpRes, err := apiClient.LatencyHistoryApi.GetNetworkWirelessClientLatencyHistory(context.Background(), networkId, clientId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test LatencyHistoryApiService GetNetworkWirelessLatencyHistory", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.LatencyHistoryApi.GetNetworkWirelessLatencyHistory(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
