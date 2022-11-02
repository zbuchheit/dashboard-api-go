/*
Meraki Dashboard API

Testing ConnectivityEventsApiService

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

func Test_client_ConnectivityEventsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ConnectivityEventsApiService GetNetworkWirelessClientConnectivityEvents", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var clientId string

        resp, httpRes, err := apiClient.ConnectivityEventsApi.GetNetworkWirelessClientConnectivityEvents(context.Background(), networkId, clientId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
