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


// VerifiedDomainsAPIService VerifiedDomainsAPI service
type VerifiedDomainsAPIService service

type ApiCreateVerifiedDomainRequest struct {
	ctx context.Context
	ApiService *VerifiedDomainsAPIService
	body *VerifiedDomains2
}

func (r ApiCreateVerifiedDomainRequest) Body(body VerifiedDomains2) ApiCreateVerifiedDomainRequest {
	r.body = &body
	return r
}

func (r ApiCreateVerifiedDomainRequest) Execute() (*VerifiedDomains, *http.Response, error) {
	return r.ApiService.CreateVerifiedDomainExecute(r)
}

/*
CreateVerifiedDomain Add domain to account

Add a domain to the account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateVerifiedDomainRequest
*/
func (a *VerifiedDomainsAPIService) CreateVerifiedDomain(ctx context.Context) ApiCreateVerifiedDomainRequest {
	return ApiCreateVerifiedDomainRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VerifiedDomains
func (a *VerifiedDomainsAPIService) CreateVerifiedDomainExecute(r ApiCreateVerifiedDomainRequest) (*VerifiedDomains, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VerifiedDomains
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VerifiedDomainsAPIService.CreateVerifiedDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/verified-domains"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.body
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

type ApiDeleteVerifiedDomainRequest struct {
	ctx context.Context
	ApiService *VerifiedDomainsAPIService
	domainName string
}

func (r ApiDeleteVerifiedDomainRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteVerifiedDomainExecute(r)
}

/*
DeleteVerifiedDomain Delete domain

Delete a verified domain from the account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param domainName The domain name.
 @return ApiDeleteVerifiedDomainRequest
*/
func (a *VerifiedDomainsAPIService) DeleteVerifiedDomain(ctx context.Context, domainName string) ApiDeleteVerifiedDomainRequest {
	return ApiDeleteVerifiedDomainRequest{
		ApiService: a,
		ctx: ctx,
		domainName: domainName,
	}
}

// Execute executes the request
func (a *VerifiedDomainsAPIService) DeleteVerifiedDomainExecute(r ApiDeleteVerifiedDomainRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VerifiedDomainsAPIService.DeleteVerifiedDomain")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/verified-domains/{domain_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"domain_name"+"}", url.PathEscape(parameterValueToString(r.domainName, "domainName")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetVerifiedDomainRequest struct {
	ctx context.Context
	ApiService *VerifiedDomainsAPIService
	domainName string
}

func (r ApiGetVerifiedDomainRequest) Execute() (*VerifiedDomains, *http.Response, error) {
	return r.ApiService.GetVerifiedDomainExecute(r)
}

/*
GetVerifiedDomain Get domain info

Get the details for a single domain on the account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param domainName The domain name.
 @return ApiGetVerifiedDomainRequest
*/
func (a *VerifiedDomainsAPIService) GetVerifiedDomain(ctx context.Context, domainName string) ApiGetVerifiedDomainRequest {
	return ApiGetVerifiedDomainRequest{
		ApiService: a,
		ctx: ctx,
		domainName: domainName,
	}
}

// Execute executes the request
//  @return VerifiedDomains
func (a *VerifiedDomainsAPIService) GetVerifiedDomainExecute(r ApiGetVerifiedDomainRequest) (*VerifiedDomains, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VerifiedDomains
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VerifiedDomainsAPIService.GetVerifiedDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/verified-domains/{domain_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"domain_name"+"}", url.PathEscape(parameterValueToString(r.domainName, "domainName")), -1)

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

type ApiGetVerifiedDomainsRequest struct {
	ctx context.Context
	ApiService *VerifiedDomainsAPIService
}

func (r ApiGetVerifiedDomainsRequest) Execute() (*VerifiedDomains1, *http.Response, error) {
	return r.ApiService.GetVerifiedDomainsExecute(r)
}

/*
GetVerifiedDomains List sending domains

Get all of the sending domains on the account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetVerifiedDomainsRequest
*/
func (a *VerifiedDomainsAPIService) GetVerifiedDomains(ctx context.Context) ApiGetVerifiedDomainsRequest {
	return ApiGetVerifiedDomainsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VerifiedDomains1
func (a *VerifiedDomainsAPIService) GetVerifiedDomainsExecute(r ApiGetVerifiedDomainsRequest) (*VerifiedDomains1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VerifiedDomains1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VerifiedDomainsAPIService.GetVerifiedDomains")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/verified-domains"

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

type ApiVerifyDomainRequest struct {
	ctx context.Context
	ApiService *VerifiedDomainsAPIService
	domainName string
	body *VerifyADomainForSending
}

func (r ApiVerifyDomainRequest) Body(body VerifyADomainForSending) ApiVerifyDomainRequest {
	r.body = &body
	return r
}

func (r ApiVerifyDomainRequest) Execute() (*VerifiedDomains, *http.Response, error) {
	return r.ApiService.VerifyDomainExecute(r)
}

/*
VerifyDomain Verify domain

Verify a domain for sending.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param domainName The domain name.
 @return ApiVerifyDomainRequest
*/
func (a *VerifiedDomainsAPIService) VerifyDomain(ctx context.Context, domainName string) ApiVerifyDomainRequest {
	return ApiVerifyDomainRequest{
		ApiService: a,
		ctx: ctx,
		domainName: domainName,
	}
}

// Execute executes the request
//  @return VerifiedDomains
func (a *VerifiedDomainsAPIService) VerifyDomainExecute(r ApiVerifyDomainRequest) (*VerifiedDomains, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VerifiedDomains
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VerifiedDomainsAPIService.VerifyDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/verified-domains/{domain_name}/actions/verify"
	localVarPath = strings.Replace(localVarPath, "{"+"domain_name"+"}", url.PathEscape(parameterValueToString(r.domainName, "domainName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.body
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