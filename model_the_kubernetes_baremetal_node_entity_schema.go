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
	"time"
)

// TheKubernetesBaremetalNodeEntitySchema struct for TheKubernetesBaremetalNodeEntitySchema
type TheKubernetesBaremetalNodeEntitySchema struct {
	Id                          string     `json:"id,omitempty"`
	KubernetesClusterId         string     `json:"kubernetes_cluster_id,omitempty"`
	KubernetesClusterNodeId     string     `json:"kubernetes_cluster_node_id,omitempty"`
	KubernetesClusterNodeNumber int32      `json:"kubernetes_cluster_node_number,omitempty"`
	Title                       *string    `json:"title,omitempty"`
	Type                        string     `json:"type,omitempty"`
	Status                      string     `json:"status,omitempty"`
	Configuration               string     `json:"configuration,omitempty"`
	LocationId                  int32      `json:"location_id,omitempty"`
	PrivateIpv4Address          *string    `json:"private_ipv4_address,omitempty"`
	PublicIpv4Address           *string    `json:"public_ipv4_address,omitempty"`
	LeaseStartAt                *string    `json:"lease_start_at,omitempty"`
	ScheduledReleaseAt          *time.Time `json:"scheduled_release_at,omitempty"`
	CreatedAt                   time.Time  `json:"created_at,omitempty"`
	UpdatedAt                   time.Time  `json:"updated_at,omitempty"`
}
