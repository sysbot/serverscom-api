/*
 * Servers.com OpenAPI Specification
 *
 * # Authentication We use bearerAuth with JSON Web Tokens for Authentication. You must send this token in the Authorization header when making requests to protected resources: ``` Authorization: Bearer <JWTtoken> ```  You can find more information [here](https://swagger.io/docs/specification/authentication/bearer-authentication/)   # Pagination We use [github style pagination](https://developer.github.com/v3/#pagination) with [WebLinking](https://tools.ietf.org/html/rfc5988). Maximum 100 results per page.  # Rate limits 2000 requests per account per hour. You can see your limits in the following response headers * `X-RateLimit-Limit`  for a total limit * `X-RateLimit-Remaining` for remaining limit * `X-RateLimit-Reset` for timestamp when the oldest request will expire
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	_context "context"
	"github.com/antihax/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// KubernetesBaremetalNodeApiService KubernetesBaremetalNodeApi service
type KubernetesBaremetalNodeApiService service

// ListAllNetworksForAnExistingKubernetesBaremetalNodeOpts Optional parameters for the method 'ListAllNetworksForAnExistingKubernetesBaremetalNode'
type ListAllNetworksForAnExistingKubernetesBaremetalNodeOpts struct {
	SearchPattern      optional.String
	Family             optional.String
	InterfaceType      optional.String
	DistributionMethod optional.String
	Additional         optional.Bool
	PerPage            optional.Int32
	Page               optional.Int32
	Sorting            optional.String
	Direction          optional.String
}

/*
ListAllNetworksForAnExistingKubernetesBaremetalNode List all networks for an existing kubernetes baremetal node
To list all of the networks for specific kubernetes baremetal node please send &#x60;GET /v1/hosts/kubernetes_baremetal_nodes/{server_id}/networks&#x60;  The response contains list of networks which classified by &#x60;interface_type&#x60;, &#x60;distribution_method&#x60;  - &#x60;interface_type&#x60; points on which interface this network configured. - &#x60;distribution_method&#x60; describes how network will be distributed.   In case of the &#x60;gateway&#x60;, will be reserved first 4 IP addresses for *enterprise location*,   and 1 IP address for *standard location* from the network.   In the case of the &#x60;route&#x60;, the network will be available fully for customer use.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serverId
 * @param optional nil or *ListAllNetworksForAnExistingKubernetesBaremetalNodeOpts - Optional Parameters:
 * @param "SearchPattern" (optional.String) -   search pattern
 * @param "Family" (optional.String) -   family
 * @param "InterfaceType" (optional.String) -   interface type
 * @param "DistributionMethod" (optional.String) -   distribution method
 * @param "Additional" (optional.Bool) -   additional
 * @param "PerPage" (optional.Int32) -   per page
 * @param "Page" (optional.Int32) -   page
 * @param "Sorting" (optional.String) -   sorting
 * @param "Direction" (optional.String) -   direction
@return []Network
*/
func (a *KubernetesBaremetalNodeApiService) ListAllNetworksForAnExistingKubernetesBaremetalNode(ctx _context.Context, serverId string, localVarOptionals *ListAllNetworksForAnExistingKubernetesBaremetalNodeOpts) ([]Network, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Network
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/hosts/kubernetes_baremetal_nodes/{server_id}/networks"
	localVarPath = strings.Replace(localVarPath, "{"+"server_id"+"}", _neturl.QueryEscape(parameterToString(serverId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.SearchPattern.IsSet() {
		localVarQueryParams.Add("search_pattern", parameterToString(localVarOptionals.SearchPattern.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Family.IsSet() {
		localVarQueryParams.Add("family", parameterToString(localVarOptionals.Family.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.InterfaceType.IsSet() {
		localVarQueryParams.Add("interface_type", parameterToString(localVarOptionals.InterfaceType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DistributionMethod.IsSet() {
		localVarQueryParams.Add("distribution_method", parameterToString(localVarOptionals.DistributionMethod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Additional.IsSet() {
		localVarQueryParams.Add("additional", parameterToString(localVarOptionals.Additional.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PerPage.IsSet() {
		localVarQueryParams.Add("per_page", parameterToString(localVarOptionals.PerPage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sorting.IsSet() {
		localVarQueryParams.Add("sorting", parameterToString(localVarOptionals.Sorting.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Direction.IsSet() {
		localVarQueryParams.Add("direction", parameterToString(localVarOptionals.Direction.Value(), ""))
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v []Network
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 412 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v TheRootSchema
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ListAllPowerFeedsForAnExistingKubernetesBaremetalNode List all power feeds for an existing kubernetes baremetal node
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serverId
@return []PowerFeed
*/
func (a *KubernetesBaremetalNodeApiService) ListAllPowerFeedsForAnExistingKubernetesBaremetalNode(ctx _context.Context, serverId string) ([]PowerFeed, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []PowerFeed
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/hosts/kubernetes_baremetal_nodes/{server_id}/power_feeds"
	localVarPath = strings.Replace(localVarPath, "{"+"server_id"+"}", _neturl.QueryEscape(parameterToString(serverId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v []PowerFeed
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 412 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
RetrieveAnExistingKubernetesBaremetalNode Retrieve an existing kubernetes baremetal node
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serverId
@return TheKubernetesBaremetalNodeEntitySchema
*/
func (a *KubernetesBaremetalNodeApiService) RetrieveAnExistingKubernetesBaremetalNode(ctx _context.Context, serverId string) (TheKubernetesBaremetalNodeEntitySchema, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TheKubernetesBaremetalNodeEntitySchema
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/hosts/kubernetes_baremetal_nodes/{server_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"server_id"+"}", _neturl.QueryEscape(parameterToString(serverId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v TheKubernetesBaremetalNodeEntitySchema
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
