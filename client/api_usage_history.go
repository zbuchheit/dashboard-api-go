/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// UsageHistoryApiService UsageHistoryApi service
type UsageHistoryApiService service

type UsageHistoryApiGetNetworkApplianceUplinksUsageHistoryRequest struct {
	ctx context.Context
	ApiService *UsageHistoryApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	resolution *int32
}

// The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
func (r UsageHistoryApiGetNetworkApplianceUplinksUsageHistoryRequest) T0(t0 string) UsageHistoryApiGetNetworkApplianceUplinksUsageHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r UsageHistoryApiGetNetworkApplianceUplinksUsageHistoryRequest) T1(t1 string) UsageHistoryApiGetNetworkApplianceUplinksUsageHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 10 minutes.
func (r UsageHistoryApiGetNetworkApplianceUplinksUsageHistoryRequest) Timespan(timespan float32) UsageHistoryApiGetNetworkApplianceUplinksUsageHistoryRequest {
	r.timespan = &timespan
	return r
}

// The time resolution in seconds for returned data. The valid resolutions are: 60, 300, 600, 1800, 3600, 86400. The default is 60.
func (r UsageHistoryApiGetNetworkApplianceUplinksUsageHistoryRequest) Resolution(resolution int32) UsageHistoryApiGetNetworkApplianceUplinksUsageHistoryRequest {
	r.resolution = &resolution
	return r
}

func (r UsageHistoryApiGetNetworkApplianceUplinksUsageHistoryRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkApplianceUplinksUsageHistoryExecute(r)
}

/*
GetNetworkApplianceUplinksUsageHistory Get the sent and received bytes for each uplink of a network.

Get the sent and received bytes for each uplink of a network.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return UsageHistoryApiGetNetworkApplianceUplinksUsageHistoryRequest
*/
func (a *UsageHistoryApiService) GetNetworkApplianceUplinksUsageHistory(ctx context.Context, networkId string) UsageHistoryApiGetNetworkApplianceUplinksUsageHistoryRequest {
	return UsageHistoryApiGetNetworkApplianceUplinksUsageHistoryRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *UsageHistoryApiService) GetNetworkApplianceUplinksUsageHistoryExecute(r UsageHistoryApiGetNetworkApplianceUplinksUsageHistoryRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageHistoryApiService.GetNetworkApplianceUplinksUsageHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/uplinks/usageHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.resolution != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resolution", r.resolution, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UsageHistoryApiGetNetworkClientUsageHistoryRequest struct {
	ctx context.Context
	ApiService *UsageHistoryApiService
	networkId string
	clientId string
}

func (r UsageHistoryApiGetNetworkClientUsageHistoryRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkClientUsageHistoryExecute(r)
}

/*
GetNetworkClientUsageHistory Return the client's daily usage history

Return the client's daily usage history. Usage data is in kilobytes. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param clientId Client ID
 @return UsageHistoryApiGetNetworkClientUsageHistoryRequest
*/
func (a *UsageHistoryApiService) GetNetworkClientUsageHistory(ctx context.Context, networkId string, clientId string) UsageHistoryApiGetNetworkClientUsageHistoryRequest {
	return UsageHistoryApiGetNetworkClientUsageHistoryRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *UsageHistoryApiService) GetNetworkClientUsageHistoryExecute(r UsageHistoryApiGetNetworkClientUsageHistoryRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageHistoryApiService.GetNetworkClientUsageHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/clients/{clientId}/usageHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UsageHistoryApiGetNetworkWirelessUsageHistoryRequest struct {
	ctx context.Context
	ApiService *UsageHistoryApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	resolution *int32
	autoResolution *bool
	clientId *string
	deviceSerial *string
	apTag *string
	band *string
	ssid *int32
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r UsageHistoryApiGetNetworkWirelessUsageHistoryRequest) T0(t0 string) UsageHistoryApiGetNetworkWirelessUsageHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r UsageHistoryApiGetNetworkWirelessUsageHistoryRequest) T1(t1 string) UsageHistoryApiGetNetworkWirelessUsageHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
func (r UsageHistoryApiGetNetworkWirelessUsageHistoryRequest) Timespan(timespan float32) UsageHistoryApiGetNetworkWirelessUsageHistoryRequest {
	r.timespan = &timespan
	return r
}

// The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400.
func (r UsageHistoryApiGetNetworkWirelessUsageHistoryRequest) Resolution(resolution int32) UsageHistoryApiGetNetworkWirelessUsageHistoryRequest {
	r.resolution = &resolution
	return r
}

// Automatically select a data resolution based on the given timespan; this overrides the value specified by the &#39;resolution&#39; parameter. The default setting is false.
func (r UsageHistoryApiGetNetworkWirelessUsageHistoryRequest) AutoResolution(autoResolution bool) UsageHistoryApiGetNetworkWirelessUsageHistoryRequest {
	r.autoResolution = &autoResolution
	return r
}

// Filter results by network client to return per-device AP usage over time inner joined by the queried client&#39;s connection history.
func (r UsageHistoryApiGetNetworkWirelessUsageHistoryRequest) ClientId(clientId string) UsageHistoryApiGetNetworkWirelessUsageHistoryRequest {
	r.clientId = &clientId
	return r
}

// Filter results by device. Requires :band.
func (r UsageHistoryApiGetNetworkWirelessUsageHistoryRequest) DeviceSerial(deviceSerial string) UsageHistoryApiGetNetworkWirelessUsageHistoryRequest {
	r.deviceSerial = &deviceSerial
	return r
}

// Filter results by AP tag; either :clientId or :deviceSerial must be jointly specified.
func (r UsageHistoryApiGetNetworkWirelessUsageHistoryRequest) ApTag(apTag string) UsageHistoryApiGetNetworkWirelessUsageHistoryRequest {
	r.apTag = &apTag
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;).
func (r UsageHistoryApiGetNetworkWirelessUsageHistoryRequest) Band(band string) UsageHistoryApiGetNetworkWirelessUsageHistoryRequest {
	r.band = &band
	return r
}

// Filter results by SSID number.
func (r UsageHistoryApiGetNetworkWirelessUsageHistoryRequest) Ssid(ssid int32) UsageHistoryApiGetNetworkWirelessUsageHistoryRequest {
	r.ssid = &ssid
	return r
}

func (r UsageHistoryApiGetNetworkWirelessUsageHistoryRequest) Execute() ([]GetNetworkWirelessUsageHistory200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessUsageHistoryExecute(r)
}

/*
GetNetworkWirelessUsageHistory Return AP usage over time for a device or network client

Return AP usage over time for a device or network client

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return UsageHistoryApiGetNetworkWirelessUsageHistoryRequest
*/
func (a *UsageHistoryApiService) GetNetworkWirelessUsageHistory(ctx context.Context, networkId string) UsageHistoryApiGetNetworkWirelessUsageHistoryRequest {
	return UsageHistoryApiGetNetworkWirelessUsageHistoryRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkWirelessUsageHistory200ResponseInner
func (a *UsageHistoryApiService) GetNetworkWirelessUsageHistoryExecute(r UsageHistoryApiGetNetworkWirelessUsageHistoryRequest) ([]GetNetworkWirelessUsageHistory200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkWirelessUsageHistory200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageHistoryApiService.GetNetworkWirelessUsageHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/usageHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.resolution != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resolution", r.resolution, "")
	}
	if r.autoResolution != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "autoResolution", r.autoResolution, "")
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientId", r.clientId, "")
	}
	if r.deviceSerial != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deviceSerial", r.deviceSerial, "")
	}
	if r.apTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "apTag", r.apTag, "")
	}
	if r.band != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "band", r.band, "")
	}
	if r.ssid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ssid", r.ssid, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
