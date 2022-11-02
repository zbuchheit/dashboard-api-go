/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// AccessPoliciesApiService AccessPoliciesApi service
type AccessPoliciesApiService service

type AccessPoliciesApiCreateNetworkSwitchAccessPolicyRequest struct {
	ctx context.Context
	ApiService *AccessPoliciesApiService
	networkId string
	createNetworkSwitchAccessPolicy *CreateNetworkSwitchAccessPolicyRequest
}

func (r AccessPoliciesApiCreateNetworkSwitchAccessPolicyRequest) CreateNetworkSwitchAccessPolicy(createNetworkSwitchAccessPolicy CreateNetworkSwitchAccessPolicyRequest) AccessPoliciesApiCreateNetworkSwitchAccessPolicyRequest {
	r.createNetworkSwitchAccessPolicy = &createNetworkSwitchAccessPolicy
	return r
}

func (r AccessPoliciesApiCreateNetworkSwitchAccessPolicyRequest) Execute() (*GetNetworkSwitchAccessPolicies200ResponseInner, *http.Response, error) {
	return r.ApiService.CreateNetworkSwitchAccessPolicyExecute(r)
}

/*
CreateNetworkSwitchAccessPolicy Create an access policy for a switch network

Create an access policy for a switch network. If you would like to enable Meraki Authentication, set radiusServers to empty array.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return AccessPoliciesApiCreateNetworkSwitchAccessPolicyRequest
*/
func (a *AccessPoliciesApiService) CreateNetworkSwitchAccessPolicy(ctx context.Context, networkId string) AccessPoliciesApiCreateNetworkSwitchAccessPolicyRequest {
	return AccessPoliciesApiCreateNetworkSwitchAccessPolicyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return GetNetworkSwitchAccessPolicies200ResponseInner
func (a *AccessPoliciesApiService) CreateNetworkSwitchAccessPolicyExecute(r AccessPoliciesApiCreateNetworkSwitchAccessPolicyRequest) (*GetNetworkSwitchAccessPolicies200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkSwitchAccessPolicies200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessPoliciesApiService.CreateNetworkSwitchAccessPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/accessPolicies"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkSwitchAccessPolicy == nil {
		return localVarReturnValue, nil, reportError("createNetworkSwitchAccessPolicy is required and must be specified")
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
	localVarPostBody = r.createNetworkSwitchAccessPolicy
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type AccessPoliciesApiDeleteNetworkSwitchAccessPolicyRequest struct {
	ctx context.Context
	ApiService *AccessPoliciesApiService
	networkId string
	accessPolicyNumber string
}

func (r AccessPoliciesApiDeleteNetworkSwitchAccessPolicyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkSwitchAccessPolicyExecute(r)
}

/*
DeleteNetworkSwitchAccessPolicy Delete an access policy for a switch network

Delete an access policy for a switch network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param accessPolicyNumber
 @return AccessPoliciesApiDeleteNetworkSwitchAccessPolicyRequest
*/
func (a *AccessPoliciesApiService) DeleteNetworkSwitchAccessPolicy(ctx context.Context, networkId string, accessPolicyNumber string) AccessPoliciesApiDeleteNetworkSwitchAccessPolicyRequest {
	return AccessPoliciesApiDeleteNetworkSwitchAccessPolicyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		accessPolicyNumber: accessPolicyNumber,
	}
}

// Execute executes the request
func (a *AccessPoliciesApiService) DeleteNetworkSwitchAccessPolicyExecute(r AccessPoliciesApiDeleteNetworkSwitchAccessPolicyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessPoliciesApiService.DeleteNetworkSwitchAccessPolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accessPolicyNumber"+"}", url.PathEscape(parameterToString(r.accessPolicyNumber, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type AccessPoliciesApiGetNetworkSwitchAccessPoliciesRequest struct {
	ctx context.Context
	ApiService *AccessPoliciesApiService
	networkId string
}

func (r AccessPoliciesApiGetNetworkSwitchAccessPoliciesRequest) Execute() ([]GetNetworkSwitchAccessPolicies200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchAccessPoliciesExecute(r)
}

/*
GetNetworkSwitchAccessPolicies List the access policies for a switch network

List the access policies for a switch network. Only returns access policies with 'my RADIUS server' as authentication method

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return AccessPoliciesApiGetNetworkSwitchAccessPoliciesRequest
*/
func (a *AccessPoliciesApiService) GetNetworkSwitchAccessPolicies(ctx context.Context, networkId string) AccessPoliciesApiGetNetworkSwitchAccessPoliciesRequest {
	return AccessPoliciesApiGetNetworkSwitchAccessPoliciesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkSwitchAccessPolicies200ResponseInner
func (a *AccessPoliciesApiService) GetNetworkSwitchAccessPoliciesExecute(r AccessPoliciesApiGetNetworkSwitchAccessPoliciesRequest) ([]GetNetworkSwitchAccessPolicies200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkSwitchAccessPolicies200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessPoliciesApiService.GetNetworkSwitchAccessPolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/accessPolicies"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type AccessPoliciesApiGetNetworkSwitchAccessPolicyRequest struct {
	ctx context.Context
	ApiService *AccessPoliciesApiService
	networkId string
	accessPolicyNumber string
}

func (r AccessPoliciesApiGetNetworkSwitchAccessPolicyRequest) Execute() (*GetNetworkSwitchAccessPolicies200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchAccessPolicyExecute(r)
}

/*
GetNetworkSwitchAccessPolicy Return a specific access policy for a switch network

Return a specific access policy for a switch network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param accessPolicyNumber
 @return AccessPoliciesApiGetNetworkSwitchAccessPolicyRequest
*/
func (a *AccessPoliciesApiService) GetNetworkSwitchAccessPolicy(ctx context.Context, networkId string, accessPolicyNumber string) AccessPoliciesApiGetNetworkSwitchAccessPolicyRequest {
	return AccessPoliciesApiGetNetworkSwitchAccessPolicyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		accessPolicyNumber: accessPolicyNumber,
	}
}

// Execute executes the request
//  @return GetNetworkSwitchAccessPolicies200ResponseInner
func (a *AccessPoliciesApiService) GetNetworkSwitchAccessPolicyExecute(r AccessPoliciesApiGetNetworkSwitchAccessPolicyRequest) (*GetNetworkSwitchAccessPolicies200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkSwitchAccessPolicies200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessPoliciesApiService.GetNetworkSwitchAccessPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accessPolicyNumber"+"}", url.PathEscape(parameterToString(r.accessPolicyNumber, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type AccessPoliciesApiUpdateNetworkSwitchAccessPolicyRequest struct {
	ctx context.Context
	ApiService *AccessPoliciesApiService
	networkId string
	accessPolicyNumber string
	updateNetworkSwitchAccessPolicy *UpdateNetworkSwitchAccessPolicyRequest
}

func (r AccessPoliciesApiUpdateNetworkSwitchAccessPolicyRequest) UpdateNetworkSwitchAccessPolicy(updateNetworkSwitchAccessPolicy UpdateNetworkSwitchAccessPolicyRequest) AccessPoliciesApiUpdateNetworkSwitchAccessPolicyRequest {
	r.updateNetworkSwitchAccessPolicy = &updateNetworkSwitchAccessPolicy
	return r
}

func (r AccessPoliciesApiUpdateNetworkSwitchAccessPolicyRequest) Execute() (*GetNetworkSwitchAccessPolicies200ResponseInner, *http.Response, error) {
	return r.ApiService.UpdateNetworkSwitchAccessPolicyExecute(r)
}

/*
UpdateNetworkSwitchAccessPolicy Update an access policy for a switch network

Update an access policy for a switch network. If you would like to enable Meraki Authentication, set radiusServers to empty array.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param accessPolicyNumber
 @return AccessPoliciesApiUpdateNetworkSwitchAccessPolicyRequest
*/
func (a *AccessPoliciesApiService) UpdateNetworkSwitchAccessPolicy(ctx context.Context, networkId string, accessPolicyNumber string) AccessPoliciesApiUpdateNetworkSwitchAccessPolicyRequest {
	return AccessPoliciesApiUpdateNetworkSwitchAccessPolicyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		accessPolicyNumber: accessPolicyNumber,
	}
}

// Execute executes the request
//  @return GetNetworkSwitchAccessPolicies200ResponseInner
func (a *AccessPoliciesApiService) UpdateNetworkSwitchAccessPolicyExecute(r AccessPoliciesApiUpdateNetworkSwitchAccessPolicyRequest) (*GetNetworkSwitchAccessPolicies200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkSwitchAccessPolicies200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessPoliciesApiService.UpdateNetworkSwitchAccessPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accessPolicyNumber"+"}", url.PathEscape(parameterToString(r.accessPolicyNumber, "")), -1)

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
	localVarPostBody = r.updateNetworkSwitchAccessPolicy
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
