# Go API client for client

# Authentication We use bearerAuth with JSON Web Tokens for Authentication. You must send this token in the Authorization header when making requests to protected resources: ``` Authorization: Bearer <JWTtoken> ```  You can find more information [here](https://swagger.io/docs/specification/authentication/bearer-authentication/)   # Pagination We use [github style pagination](https://developer.github.com/v3/#pagination) with [WebLinking](https://tools.ietf.org/html/rfc5988). Maximum 100 results per page.  # Rate limits 2000 requests per account per hour. You can see your limits in the following response headers * `X-RateLimit-Limit`  for a total limit * `X-RateLimit-Remaining` for remaining limit * `X-RateLimit-Reset` for timestamp when the oldest request will expire 

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./client"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DedicatedServerApi* | [**AbortReleaseForAnExistingDedicatedServer**](docs/DedicatedServerApi.md#abortreleaseforanexistingdedicatedserver) | **Post** /v1/hosts/dedicated_servers/{server_id}/abort_release | Abort release for an existing dedicated server
*DedicatedServerApi* | [**CreateANewDedicatedServer**](docs/DedicatedServerApi.md#createanewdedicatedserver) | **Post** /v1/hosts/dedicated_servers | Create a new dedicated server
*DedicatedServerApi* | [**CreatePtrRecordForServerNetworks**](docs/DedicatedServerApi.md#createptrrecordforservernetworks) | **Post** /v1/hosts/dedicated_servers/{server_id}/ptr_records | Create PTR record for server networks
*DedicatedServerApi* | [**DeleteAnExistingPtrRecord**](docs/DedicatedServerApi.md#deleteanexistingptrrecord) | **Delete** /v1/hosts/dedicated_servers/{server_id}/ptr_records/{record_id} | Delete an existing PTR record
*DedicatedServerApi* | [**ListAllNetworksForAnExistingDedicatedServer**](docs/DedicatedServerApi.md#listallnetworksforanexistingdedicatedserver) | **Get** /v1/hosts/dedicated_servers/{server_id}/networks | List all networks for an existing dedicated server
*DedicatedServerApi* | [**ListAllPowerFeedsForAnExistingDedicatedServer**](docs/DedicatedServerApi.md#listallpowerfeedsforanexistingdedicatedserver) | **Get** /v1/hosts/dedicated_servers/{server_id}/power_feeds | List all power feeds for an existing dedicated server
*DedicatedServerApi* | [**ListAllPtrRecordsForServerNetworks**](docs/DedicatedServerApi.md#listallptrrecordsforservernetworks) | **Get** /v1/hosts/dedicated_servers/{server_id}/ptr_records | List all PTR records for server networks
*DedicatedServerApi* | [**RetrieveAnExistingDedicatedServer**](docs/DedicatedServerApi.md#retrieveanexistingdedicatedserver) | **Get** /v1/hosts/dedicated_servers/{server_id} | Retrieve an existing dedicated server
*DedicatedServerApi* | [**ScheduleReleaseForAnExistingDedicatedServer**](docs/DedicatedServerApi.md#schedulereleaseforanexistingdedicatedserver) | **Post** /v1/hosts/dedicated_servers/{server_id}/schedule_release | Schedule release for an existing dedicated server
*HostsApi* | [**ListAllHosts**](docs/HostsApi.md#listallhosts) | **Get** /v1/hosts | List all hosts
*KubernetesBaremetalNodeApi* | [**ListAllNetworksForAnExistingKubernetesBaremetalNode**](docs/KubernetesBaremetalNodeApi.md#listallnetworksforanexistingkubernetesbaremetalnode) | **Get** /v1/hosts/kubernetes_baremetal_nodes/{server_id}/networks | List all networks for an existing kubernetes baremetal node
*KubernetesBaremetalNodeApi* | [**ListAllPowerFeedsForAnExistingKubernetesBaremetalNode**](docs/KubernetesBaremetalNodeApi.md#listallpowerfeedsforanexistingkubernetesbaremetalnode) | **Get** /v1/hosts/kubernetes_baremetal_nodes/{server_id}/power_feeds | List all power feeds for an existing kubernetes baremetal node
*KubernetesBaremetalNodeApi* | [**RetrieveAnExistingKubernetesBaremetalNode**](docs/KubernetesBaremetalNodeApi.md#retrieveanexistingkubernetesbaremetalnode) | **Get** /v1/hosts/kubernetes_baremetal_nodes/{node_id} | Retrieve an existing kubernetes baremetal node
*KubernetesClusterApi* | [**KubernetesClusterNodes**](docs/KubernetesClusterApi.md#kubernetesclusternodes) | **Get** /v1/kubernetes_clusters/{kubernetes_cluster_id}/nodes | Kubernetes cluster nodes
*KubernetesClusterApi* | [**KubernetesClusters**](docs/KubernetesClusterApi.md#kubernetesclusters) | **Get** /v1/kubernetes_clusters | Kubernetes clusters
*KubernetesClusterApi* | [**RetrieveAnExistingKubernetesCluster**](docs/KubernetesClusterApi.md#retrieveanexistingkubernetescluster) | **Get** /v1/kubernetes_clusters/{id} | Retrieve an existing kubernetes cluster
*KubernetesClusterApi* | [**RetrieveAnExistingKubernetesClusterNode**](docs/KubernetesClusterApi.md#retrieveanexistingkubernetesclusternode) | **Get** /v1/kubernetes_clusters/{kubernetes_cluster_id}/nodes/{node_id} | Retrieve an existing kubernetes cluster node
*L2SegmentApi* | [**CreateANewL2Segment**](docs/L2SegmentApi.md#createanewl2segment) | **Post** /v1/l2_segments | Create a new L2 segment
*L2SegmentApi* | [**DeleteAnExistingL2Segment**](docs/L2SegmentApi.md#deleteanexistingl2segment) | **Delete** /v1/l2_segments/{l2_segment_id} | Delete an existing L2 segment
*L2SegmentApi* | [**ListAllL2SegmentMembers**](docs/L2SegmentApi.md#listalll2segmentmembers) | **Get** /v1/l2_segments/{l2_segment_id}/members | List all L2 segment members
*L2SegmentApi* | [**ListAllL2SegmentNetworks**](docs/L2SegmentApi.md#listalll2segmentnetworks) | **Get** /v1/l2_segments/{l2_segment_id}/networks | List all l2 segment networks
*L2SegmentApi* | [**ListAllL2Segments**](docs/L2SegmentApi.md#listalll2segments) | **Get** /v1/l2_segments | List all L2 segments
*L2SegmentApi* | [**RetrieveAnExistingL2Segment**](docs/L2SegmentApi.md#retrieveanexistingl2segment) | **Get** /v1/l2_segments/{l2_segment_id} | Retrieve an existing L2 segment
*L2SegmentApi* | [**UpdateAnExistingL2Segment**](docs/L2SegmentApi.md#updateanexistingl2segment) | **Put** /v1/l2_segments/{l2_segment_id} | Update an existing L2 segment
*L2SegmentApi* | [**UpdateAnExistingL2SegmentNetworks**](docs/L2SegmentApi.md#updateanexistingl2segmentnetworks) | **Put** /v1/l2_segments/{l2_segment_id}/networks | Update an existing L2 segment networks
*LoadBalancerApi* | [**CreateANewL4LoadBalancers**](docs/LoadBalancerApi.md#createanewl4loadbalancers) | **Post** /v1/load_balancers/l4 | Create a new L4 load balancers
*LoadBalancerApi* | [**CreateANewL7LoadBalancers**](docs/LoadBalancerApi.md#createanewl7loadbalancers) | **Post** /v1/load_balancers/l7 | Create a new L7 load balancers
*LoadBalancerApi* | [**DeleteAnExisitingL4LoadBalancer**](docs/LoadBalancerApi.md#deleteanexisitingl4loadbalancer) | **Delete** /v1/load_balancers/l4/{id} | Delete an exisiting L4 load balancer
*LoadBalancerApi* | [**DeleteAnExistingL7LoadBalancer**](docs/LoadBalancerApi.md#deleteanexistingl7loadbalancer) | **Delete** /v1/load_balancers/l7/{id} | Delete an existing L7 load balancer
*LoadBalancerApi* | [**LoadBalancers**](docs/LoadBalancerApi.md#loadbalancers) | **Get** /v1/load_balancers | Load balancers
*LoadBalancerApi* | [**RetrieveAnExistingL4LoadBalancer**](docs/LoadBalancerApi.md#retrieveanexistingl4loadbalancer) | **Get** /v1/load_balancers/l4/{id} | Retrieve an existing L4 load balancer
*LoadBalancerApi* | [**RetrieveAnExistingL7LoadBalancer**](docs/LoadBalancerApi.md#retrieveanexistingl7loadbalancer) | **Get** /v1/load_balancers/l7/{id} | Retrieve an existing L7 load balancer
*LoadBalancerApi* | [**UpdateAnExistingL4LoadBalancer**](docs/LoadBalancerApi.md#updateanexistingl4loadbalancer) | **Put** /v1/load_balancers/l4/{id} | Update an existing L4 load balancer
*LoadBalancerApi* | [**UpdateAnExistingL7LoadBalancer**](docs/LoadBalancerApi.md#updateanexistingl7loadbalancer) | **Put** /v1/load_balancers/l7/{id} | Update an existing L7 load balancer
*LocationApi* | [**Locations**](docs/LocationApi.md#locations) | **Get** /v1/locations | Locations
*LocationApi* | [**RetrieveAnExisitingLocation**](docs/LocationApi.md#retrieveanexisitinglocation) | **Get** /v1/locations/{location_id} | Retrieve an exisiting location
*SSLCertificateApi* | [**CreateANewCustomSslCertificate**](docs/SSLCertificateApi.md#createanewcustomsslcertificate) | **Post** /v1/ssl_certificates/custom | Create a new custom SSL certificate
*SSLCertificateApi* | [**RetrieveAnExistingCustomSslCertificate**](docs/SSLCertificateApi.md#retrieveanexistingcustomsslcertificate) | **Get** /v1/ssl_certificates/custom/{id} | Retrieve an existing custom ssl certificate
*SSLCertificateApi* | [**SSLCertificates**](docs/SSLCertificateApi.md#sslcertificates) | **Get** /v1/ssl_certificates | SSL Certificates


## Documentation For Models

 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineObject2](docs/InlineObject2.md)
 - [InlineObject3](docs/InlineObject3.md)
 - [InlineObject4](docs/InlineObject4.md)
 - [InlineObject5](docs/InlineObject5.md)
 - [InlineObject6](docs/InlineObject6.md)
 - [InlineObject7](docs/InlineObject7.md)
 - [InlineObject8](docs/InlineObject8.md)
 - [InlineObject9](docs/InlineObject9.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse201](docs/InlineResponse201.md)
 - [InlineResponse2011](docs/InlineResponse2011.md)
 - [InlineResponse422](docs/InlineResponse422.md)
 - [Network](docs/Network.md)
 - [PowerFeed](docs/PowerFeed.md)
 - [TheDedicatedServerEntitySchema](docs/TheDedicatedServerEntitySchema.md)
 - [TheItemsSchema](docs/TheItemsSchema.md)
 - [TheItemsSchema1](docs/TheItemsSchema1.md)
 - [TheItemsSchema2](docs/TheItemsSchema2.md)
 - [TheItemsSchema3](docs/TheItemsSchema3.md)
 - [TheItemsSchema4](docs/TheItemsSchema4.md)
 - [TheKubernetesBaremetalNodeEntitySchema](docs/TheKubernetesBaremetalNodeEntitySchema.md)
 - [TheRootSchema](docs/TheRootSchema.md)
 - [TheRootSchema1](docs/TheRootSchema1.md)
 - [V1HostsDedicatedServersDrives](docs/V1HostsDedicatedServersDrives.md)
 - [V1HostsDedicatedServersDrivesLayout](docs/V1HostsDedicatedServersDrivesLayout.md)
 - [V1HostsDedicatedServersDrivesPartitions](docs/V1HostsDedicatedServersDrivesPartitions.md)
 - [V1HostsDedicatedServersDrivesSlots](docs/V1HostsDedicatedServersDrivesSlots.md)
 - [V1HostsDedicatedServersEntity](docs/V1HostsDedicatedServersEntity.md)
 - [V1HostsDedicatedServersHosts](docs/V1HostsDedicatedServersHosts.md)
 - [V1HostsDedicatedServersShortEntity](docs/V1HostsDedicatedServersShortEntity.md)
 - [V1HostsDedicatedServersUplinkModels](docs/V1HostsDedicatedServersUplinkModels.md)
 - [V1HostsDedicatedServersUplinkModelsPrivate](docs/V1HostsDedicatedServersUplinkModelsPrivate.md)
 - [V1HostsDedicatedServersUplinkModelsPublic](docs/V1HostsDedicatedServersUplinkModelsPublic.md)
 - [V1HostsKubernetesBaremetalNodeEntity](docs/V1HostsKubernetesBaremetalNodeEntity.md)
 - [V1L2SegmentsL2Member](docs/V1L2SegmentsL2Member.md)
 - [V1L2SegmentsL2Network](docs/V1L2SegmentsL2Network.md)
 - [V1L2SegmentsL2Segment](docs/V1L2SegmentsL2Segment.md)
 - [V1L2SegmentsL2SegmentDetailed](docs/V1L2SegmentsL2SegmentDetailed.md)
 - [V1L2SegmentsL2SegmentIdNetworksCreate](docs/V1L2SegmentsL2SegmentIdNetworksCreate.md)
 - [V1L2SegmentsMembers](docs/V1L2SegmentsMembers.md)
 - [V1LoadBalancersL4Rules](docs/V1LoadBalancersL4Rules.md)
 - [V1LoadBalancersL4Upstreams](docs/V1LoadBalancersL4Upstreams.md)
 - [V1LoadBalancersL7IdRules](docs/V1LoadBalancersL7IdRules.md)
 - [V1LoadBalancersL7Rules](docs/V1LoadBalancersL7Rules.md)


## Documentation For Authorization



## bearerAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Author



