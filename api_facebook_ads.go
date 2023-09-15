/*
Mailchimp Marketing API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.55
Contact: apihelp@mailchimp.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// FacebookAdsAPIService FacebookAdsAPI service
type FacebookAdsAPIService service

type ApiGetAllFacebookAdsRequest struct {
	ctx context.Context
	ApiService *FacebookAdsAPIService
	fields *[]string
	excludeFields *[]string
	count *int32
	offset *int32
	sortField *string
	sortDir *string
}

// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
func (r ApiGetAllFacebookAdsRequest) Fields(fields []string) ApiGetAllFacebookAdsRequest {
	r.fields = &fields
	return r
}

// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
func (r ApiGetAllFacebookAdsRequest) ExcludeFields(excludeFields []string) ApiGetAllFacebookAdsRequest {
	r.excludeFields = &excludeFields
	return r
}

// The number of records to return. Default value is 10. Maximum value is 1000
func (r ApiGetAllFacebookAdsRequest) Count(count int32) ApiGetAllFacebookAdsRequest {
	r.count = &count
	return r
}

// Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0.
func (r ApiGetAllFacebookAdsRequest) Offset(offset int32) ApiGetAllFacebookAdsRequest {
	r.offset = &offset
	return r
}

// Returns files sorted by the specified field.
func (r ApiGetAllFacebookAdsRequest) SortField(sortField string) ApiGetAllFacebookAdsRequest {
	r.sortField = &sortField
	return r
}

// Determines the order direction for sorted results.
func (r ApiGetAllFacebookAdsRequest) SortDir(sortDir string) ApiGetAllFacebookAdsRequest {
	r.sortDir = &sortDir
	return r
}

func (r ApiGetAllFacebookAdsRequest) Execute() (*GetAllFacebookAds200Response, *http.Response, error) {
	return r.ApiService.GetAllFacebookAdsExecute(r)
}

/*
GetAllFacebookAds List facebook ads

Get list of Facebook ads.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllFacebookAdsRequest
*/
func (a *FacebookAdsAPIService) GetAllFacebookAds(ctx context.Context) ApiGetAllFacebookAdsRequest {
	return ApiGetAllFacebookAdsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetAllFacebookAds200Response
func (a *FacebookAdsAPIService) GetAllFacebookAdsExecute(r ApiGetAllFacebookAdsRequest) (*GetAllFacebookAds200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAllFacebookAds200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FacebookAdsAPIService.GetAllFacebookAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/facebook-ads"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "csv")
	}
	if r.excludeFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", r.excludeFields, "csv")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.sortField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_field", r.sortField, "")
	}
	if r.sortDir != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_dir", r.sortDir, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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
			var v ProblemDetailDocument
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetFacebookAdsIdRequest struct {
	ctx context.Context
	ApiService *FacebookAdsAPIService
	outreachId string
	fields *[]string
	excludeFields *[]string
}

// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
func (r ApiGetFacebookAdsIdRequest) Fields(fields []string) ApiGetFacebookAdsIdRequest {
	r.fields = &fields
	return r
}

// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
func (r ApiGetFacebookAdsIdRequest) ExcludeFields(excludeFields []string) ApiGetFacebookAdsIdRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ApiGetFacebookAdsIdRequest) Execute() (*GetFacebookAdsId200Response, *http.Response, error) {
	return r.ApiService.GetFacebookAdsIdExecute(r)
}

/*
GetFacebookAdsId Get facebook ad info

Get details of a Facebook ad.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param outreachId The outreach id.
 @return ApiGetFacebookAdsIdRequest
*/
func (a *FacebookAdsAPIService) GetFacebookAdsId(ctx context.Context, outreachId string) ApiGetFacebookAdsIdRequest {
	return ApiGetFacebookAdsIdRequest{
		ApiService: a,
		ctx: ctx,
		outreachId: outreachId,
	}
}

// Execute executes the request
//  @return GetFacebookAdsId200Response
func (a *FacebookAdsAPIService) GetFacebookAdsIdExecute(r ApiGetFacebookAdsIdRequest) (*GetFacebookAdsId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFacebookAdsId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FacebookAdsAPIService.GetFacebookAdsId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/facebook-ads/{outreach_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"outreach_id"+"}", url.PathEscape(parameterValueToString(r.outreachId, "outreachId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "csv")
	}
	if r.excludeFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", r.excludeFields, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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
			var v ProblemDetailDocument
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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