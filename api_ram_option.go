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

// RamOptionApiService RamOptionApi service
type RamOptionApiService service

// ListAllAvailableRamOptionsForServerModelOpts Optional parameters for the method 'ListAllAvailableRamOptionsForServerModel'
type ListAllAvailableRamOptionsForServerModelOpts struct {
	PerPage   optional.Int32
	Page      optional.Int32
	Sorting   optional.String
	Direction optional.String
}

/*
ListAllAvailableRamOptionsForServerModel List all available ram options for server model
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param locationId
 * @param serverModelId
 * @param optional nil or *ListAllAvailableRamOptionsForServerModelOpts - Optional Parameters:
 * @param "PerPage" (optional.Int32) -   per page
 * @param "Page" (optional.Int32) -   page
 * @param "Sorting" (optional.String) -   sorting
 * @param "Direction" (optional.String) -   direction
@return []TheItemsSchema4
*/
func (a *RamOptionApiService) ListAllAvailableRamOptionsForServerModel(ctx _context.Context, locationId string, serverModelId string, localVarOptionals *ListAllAvailableRamOptionsForServerModelOpts) ([]TheItemsSchema4, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TheItemsSchema4
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/locations/{location_id}/order_options/server_models/{server_model_id}/ram"
	localVarPath = strings.Replace(localVarPath, "{"+"location_id"+"}", _neturl.QueryEscape(parameterToString(locationId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"server_model_id"+"}", _neturl.QueryEscape(parameterToString(serverModelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
			var v []TheItemsSchema4
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
