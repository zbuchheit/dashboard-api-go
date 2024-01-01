/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
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


// LiveToolsApiService LiveToolsApi service
type LiveToolsApiService service

type LiveToolsApiBlinkDeviceLedsRequest struct {
	ctx context.Context
	ApiService *LiveToolsApiService
	serial string
	blinkDeviceLedsRequest *BlinkDeviceLedsRequest
}

func (r LiveToolsApiBlinkDeviceLedsRequest) BlinkDeviceLedsRequest(blinkDeviceLedsRequest BlinkDeviceLedsRequest) LiveToolsApiBlinkDeviceLedsRequest {
	r.blinkDeviceLedsRequest = &blinkDeviceLedsRequest
	return r
}

func (r LiveToolsApiBlinkDeviceLedsRequest) Execute() (*BlinkDeviceLeds202Response, *http.Response, error) {
	return r.ApiService.BlinkDeviceLedsExecute(r)
}

/*
BlinkDeviceLeds Blink the LEDs on a device

Blink the LEDs on a device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return LiveToolsApiBlinkDeviceLedsRequest
*/
func (a *LiveToolsApiService) BlinkDeviceLeds(ctx context.Context, serial string) LiveToolsApiBlinkDeviceLedsRequest {
	return LiveToolsApiBlinkDeviceLedsRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return BlinkDeviceLeds202Response
func (a *LiveToolsApiService) BlinkDeviceLedsExecute(r LiveToolsApiBlinkDeviceLedsRequest) (*BlinkDeviceLeds202Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BlinkDeviceLeds202Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LiveToolsApiService.BlinkDeviceLeds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/blinkLeds"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.blinkDeviceLedsRequest
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

type LiveToolsApiCreateDeviceLiveToolsPingRequest struct {
	ctx context.Context
	ApiService *LiveToolsApiService
	serial string
	createDeviceLiveToolsPingRequest *CreateDeviceLiveToolsPingRequest
}

func (r LiveToolsApiCreateDeviceLiveToolsPingRequest) CreateDeviceLiveToolsPingRequest(createDeviceLiveToolsPingRequest CreateDeviceLiveToolsPingRequest) LiveToolsApiCreateDeviceLiveToolsPingRequest {
	r.createDeviceLiveToolsPingRequest = &createDeviceLiveToolsPingRequest
	return r
}

func (r LiveToolsApiCreateDeviceLiveToolsPingRequest) Execute() (*CreateDeviceLiveToolsPing201Response, *http.Response, error) {
	return r.ApiService.CreateDeviceLiveToolsPingExecute(r)
}

/*
CreateDeviceLiveToolsPing Enqueue a job to ping a target host from the device

Enqueue a job to ping a target host from the device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return LiveToolsApiCreateDeviceLiveToolsPingRequest
*/
func (a *LiveToolsApiService) CreateDeviceLiveToolsPing(ctx context.Context, serial string) LiveToolsApiCreateDeviceLiveToolsPingRequest {
	return LiveToolsApiCreateDeviceLiveToolsPingRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return CreateDeviceLiveToolsPing201Response
func (a *LiveToolsApiService) CreateDeviceLiveToolsPingExecute(r LiveToolsApiCreateDeviceLiveToolsPingRequest) (*CreateDeviceLiveToolsPing201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateDeviceLiveToolsPing201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LiveToolsApiService.CreateDeviceLiveToolsPing")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/liveTools/ping"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createDeviceLiveToolsPingRequest == nil {
		return localVarReturnValue, nil, reportError("createDeviceLiveToolsPingRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.createDeviceLiveToolsPingRequest
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

type LiveToolsApiCreateDeviceLiveToolsPingDeviceRequest struct {
	ctx context.Context
	ApiService *LiveToolsApiService
	serial string
	createDeviceLiveToolsPingDeviceRequest *CreateDeviceLiveToolsPingDeviceRequest
}

func (r LiveToolsApiCreateDeviceLiveToolsPingDeviceRequest) CreateDeviceLiveToolsPingDeviceRequest(createDeviceLiveToolsPingDeviceRequest CreateDeviceLiveToolsPingDeviceRequest) LiveToolsApiCreateDeviceLiveToolsPingDeviceRequest {
	r.createDeviceLiveToolsPingDeviceRequest = &createDeviceLiveToolsPingDeviceRequest
	return r
}

func (r LiveToolsApiCreateDeviceLiveToolsPingDeviceRequest) Execute() (*CreateDeviceLiveToolsPing201Response, *http.Response, error) {
	return r.ApiService.CreateDeviceLiveToolsPingDeviceExecute(r)
}

/*
CreateDeviceLiveToolsPingDevice Enqueue a job to check connectivity status to the device

Enqueue a job to check connectivity status to the device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return LiveToolsApiCreateDeviceLiveToolsPingDeviceRequest
*/
func (a *LiveToolsApiService) CreateDeviceLiveToolsPingDevice(ctx context.Context, serial string) LiveToolsApiCreateDeviceLiveToolsPingDeviceRequest {
	return LiveToolsApiCreateDeviceLiveToolsPingDeviceRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return CreateDeviceLiveToolsPing201Response
func (a *LiveToolsApiService) CreateDeviceLiveToolsPingDeviceExecute(r LiveToolsApiCreateDeviceLiveToolsPingDeviceRequest) (*CreateDeviceLiveToolsPing201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateDeviceLiveToolsPing201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LiveToolsApiService.CreateDeviceLiveToolsPingDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/liveTools/pingDevice"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.createDeviceLiveToolsPingDeviceRequest
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

type LiveToolsApiCycleDeviceSwitchPortsRequest struct {
	ctx context.Context
	ApiService *LiveToolsApiService
	serial string
	cycleDeviceSwitchPortsRequest *CycleDeviceSwitchPortsRequest
}

func (r LiveToolsApiCycleDeviceSwitchPortsRequest) CycleDeviceSwitchPortsRequest(cycleDeviceSwitchPortsRequest CycleDeviceSwitchPortsRequest) LiveToolsApiCycleDeviceSwitchPortsRequest {
	r.cycleDeviceSwitchPortsRequest = &cycleDeviceSwitchPortsRequest
	return r
}

func (r LiveToolsApiCycleDeviceSwitchPortsRequest) Execute() (*CycleDeviceSwitchPorts200Response, *http.Response, error) {
	return r.ApiService.CycleDeviceSwitchPortsExecute(r)
}

/*
CycleDeviceSwitchPorts Cycle a set of switch ports

Cycle a set of switch ports

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return LiveToolsApiCycleDeviceSwitchPortsRequest
*/
func (a *LiveToolsApiService) CycleDeviceSwitchPorts(ctx context.Context, serial string) LiveToolsApiCycleDeviceSwitchPortsRequest {
	return LiveToolsApiCycleDeviceSwitchPortsRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return CycleDeviceSwitchPorts200Response
func (a *LiveToolsApiService) CycleDeviceSwitchPortsExecute(r LiveToolsApiCycleDeviceSwitchPortsRequest) (*CycleDeviceSwitchPorts200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CycleDeviceSwitchPorts200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LiveToolsApiService.CycleDeviceSwitchPorts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/switch/ports/cycle"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cycleDeviceSwitchPortsRequest == nil {
		return localVarReturnValue, nil, reportError("cycleDeviceSwitchPortsRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.cycleDeviceSwitchPortsRequest
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

type LiveToolsApiGetDeviceLiveToolsPingRequest struct {
	ctx context.Context
	ApiService *LiveToolsApiService
	serial string
	id string
}

func (r LiveToolsApiGetDeviceLiveToolsPingRequest) Execute() (*DevicesSerialLiveToolsPingPostRequestMessage, *http.Response, error) {
	return r.ApiService.GetDeviceLiveToolsPingExecute(r)
}

/*
GetDeviceLiveToolsPing Return a ping job

Return a ping job. Latency unit in response is in milliseconds. Size is in bytes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @param id ID
 @return LiveToolsApiGetDeviceLiveToolsPingRequest
*/
func (a *LiveToolsApiService) GetDeviceLiveToolsPing(ctx context.Context, serial string, id string) LiveToolsApiGetDeviceLiveToolsPingRequest {
	return LiveToolsApiGetDeviceLiveToolsPingRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
		id: id,
	}
}

// Execute executes the request
//  @return DevicesSerialLiveToolsPingPostRequestMessage
func (a *LiveToolsApiService) GetDeviceLiveToolsPingExecute(r LiveToolsApiGetDeviceLiveToolsPingRequest) (*DevicesSerialLiveToolsPingPostRequestMessage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DevicesSerialLiveToolsPingPostRequestMessage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LiveToolsApiService.GetDeviceLiveToolsPing")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/liveTools/ping/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type LiveToolsApiGetDeviceLiveToolsPingDeviceRequest struct {
	ctx context.Context
	ApiService *LiveToolsApiService
	serial string
	id string
}

func (r LiveToolsApiGetDeviceLiveToolsPingDeviceRequest) Execute() (*GetDeviceLiveToolsPingDevice200Response, *http.Response, error) {
	return r.ApiService.GetDeviceLiveToolsPingDeviceExecute(r)
}

/*
GetDeviceLiveToolsPingDevice Return a ping device job

Return a ping device job. Latency unit in response is in milliseconds. Size is in bytes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @param id ID
 @return LiveToolsApiGetDeviceLiveToolsPingDeviceRequest
*/
func (a *LiveToolsApiService) GetDeviceLiveToolsPingDevice(ctx context.Context, serial string, id string) LiveToolsApiGetDeviceLiveToolsPingDeviceRequest {
	return LiveToolsApiGetDeviceLiveToolsPingDeviceRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
		id: id,
	}
}

// Execute executes the request
//  @return GetDeviceLiveToolsPingDevice200Response
func (a *LiveToolsApiService) GetDeviceLiveToolsPingDeviceExecute(r LiveToolsApiGetDeviceLiveToolsPingDeviceRequest) (*GetDeviceLiveToolsPingDevice200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetDeviceLiveToolsPingDevice200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LiveToolsApiService.GetDeviceLiveToolsPingDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/liveTools/pingDevice/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type LiveToolsApiRebootDeviceRequest struct {
	ctx context.Context
	ApiService *LiveToolsApiService
	serial string
}

func (r LiveToolsApiRebootDeviceRequest) Execute() (*RebootDevice202Response, *http.Response, error) {
	return r.ApiService.RebootDeviceExecute(r)
}

/*
RebootDevice Reboot a device

Reboot a device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return LiveToolsApiRebootDeviceRequest
*/
func (a *LiveToolsApiService) RebootDevice(ctx context.Context, serial string) LiveToolsApiRebootDeviceRequest {
	return LiveToolsApiRebootDeviceRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return RebootDevice202Response
func (a *LiveToolsApiService) RebootDeviceExecute(r LiveToolsApiRebootDeviceRequest) (*RebootDevice202Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RebootDevice202Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LiveToolsApiService.RebootDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/reboot"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterValueToString(r.serial, "serial")), -1)

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
