/*
 * Servers.com OpenAPI Specification
 *
 * # Authentication We use bearerAuth with JSON Web Tokens for Authentication. You must send this token in the Authorization header when making requests to protected resources: ``` Authorization: Bearer <JWTtoken> ```  You can find more information [here](https://swagger.io/docs/specification/authentication/bearer-authentication/)   # Pagination We use [github style pagination](https://developer.github.com/v3/#pagination) with [WebLinking](https://tools.ietf.org/html/rfc5988). Maximum 100 results per page.  # Rate limits 2000 requests per account per hour. You can see your limits in the following response headers * `X-RateLimit-Limit`  for a total limit * `X-RateLimit-Remaining` for remaining limit * `X-RateLimit-Reset` for timestamp when the oldest request will expire
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// V1L2SegmentsL2LocationGroup struct for V1L2SegmentsL2LocationGroup
type V1L2SegmentsL2LocationGroup struct {
	Id          int32   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	GroupType   string  `json:"group_type,omitempty"`
	LocationIds []int32 `json:"location_ids,omitempty"`
}