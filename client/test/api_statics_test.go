/*
Meraki Dashboard API

Testing StaticsApiService

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

func Test_client_StaticsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test StaticsApiService CreateNetworkAppliancePrefixesDelegatedStatic", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.StaticsApi.CreateNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticsApiService DeleteNetworkAppliancePrefixesDelegatedStatic", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var staticDelegatedPrefixId string

        resp, httpRes, err := apiClient.StaticsApi.DeleteNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticsApiService GetNetworkAppliancePrefixesDelegatedStatic", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var staticDelegatedPrefixId string

        resp, httpRes, err := apiClient.StaticsApi.GetNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticsApiService GetNetworkAppliancePrefixesDelegatedStatics", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.StaticsApi.GetNetworkAppliancePrefixesDelegatedStatics(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test StaticsApiService UpdateNetworkAppliancePrefixesDelegatedStatic", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string
        var staticDelegatedPrefixId string

        resp, httpRes, err := apiClient.StaticsApi.UpdateNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
