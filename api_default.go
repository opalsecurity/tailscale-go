/*
Tailscale API

Tailscale API spec

API version: 1.0
Contact: hello@opal.dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opal

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiTailnetTailnetNameAclGetRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	tailnetName string
}

func (r ApiTailnetTailnetNameAclGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.TailnetTailnetNameAclGetExecute(r)
}

/*
TailnetTailnetNameAclGet Method for TailnetTailnetNameAclGet

Retrieves the ACL that is currently set for the given tailnet. Supply the tailnet of interest in the path. This endpoint can send back either the HuJSON of the ACL or a parsed JSON, depending on the Accept header.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tailnetName The name of the Tailnet
 @return ApiTailnetTailnetNameAclGetRequest
*/
func (a *DefaultApiService) TailnetTailnetNameAclGet(ctx context.Context, tailnetName string) ApiTailnetTailnetNameAclGetRequest {
	return ApiTailnetTailnetNameAclGetRequest{
		ApiService: a,
		ctx: ctx,
		tailnetName: tailnetName,
	}
}

// Execute executes the request
func (a *DefaultApiService) TailnetTailnetNameAclGetExecute(r ApiTailnetTailnetNameAclGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.TailnetTailnetNameAclGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tailnet/{tailnet_name}/acl"
	localVarPath = strings.Replace(localVarPath, "{"+"tailnet_name"+"}", url.PathEscape(parameterToString(r.tailnetName, "")), -1)

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
