/*
Meraki Dashboard API

Testing SplashAuthorizationStatusApiService

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

func Test_client_SplashAuthorizationStatusApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SplashAuthorizationStatusApiService GetNetworkClientSplashAuthorizationStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var clientId string

        resp, httpRes, err := apiClient.SplashAuthorizationStatusApi.GetNetworkClientSplashAuthorizationStatus(context.Background(), networkId, clientId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SplashAuthorizationStatusApiService UpdateNetworkClientSplashAuthorizationStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var clientId string

        resp, httpRes, err := apiClient.SplashAuthorizationStatusApi.UpdateNetworkClientSplashAuthorizationStatus(context.Background(), networkId, clientId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
