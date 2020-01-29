# \KubernetesBaremetalNodeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAllNetworksForAnExistingKubernetesBaremetalNode**](KubernetesBaremetalNodeApi.md#ListAllNetworksForAnExistingKubernetesBaremetalNode) | **Get** /v1/hosts/kubernetes_baremetal_nodes/{server_id}/networks | List all networks for an existing kubernetes baremetal node
[**ListAllPowerFeedsForAnExistingKubernetesBaremetalNode**](KubernetesBaremetalNodeApi.md#ListAllPowerFeedsForAnExistingKubernetesBaremetalNode) | **Get** /v1/hosts/kubernetes_baremetal_nodes/{server_id}/power_feeds | List all power feeds for an existing kubernetes baremetal node
[**RetrieveAnExistingKubernetesBaremetalNode**](KubernetesBaremetalNodeApi.md#RetrieveAnExistingKubernetesBaremetalNode) | **Get** /v1/hosts/kubernetes_baremetal_nodes/{server_id} | Retrieve an existing kubernetes baremetal node



## ListAllNetworksForAnExistingKubernetesBaremetalNode

> []Network ListAllNetworksForAnExistingKubernetesBaremetalNode(ctx, serverId, optional)

List all networks for an existing kubernetes baremetal node

To list all of the networks for specific kubernetes baremetal node please send `GET /v1/hosts/kubernetes_baremetal_nodes/{server_id}/networks`  The response contains list of networks which classified by `interface_type`, `distribution_method`  - `interface_type` points on which interface this network configured. - `distribution_method` describes how network will be distributed.   In case of the `gateway`, will be reserved first 4 IP addresses for *enterprise location*,   and 1 IP address for *standard location* from the network.   In the case of the `route`, the network will be available fully for customer use. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string**|  | 
 **optional** | ***ListAllNetworksForAnExistingKubernetesBaremetalNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllNetworksForAnExistingKubernetesBaremetalNodeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchPattern** | **optional.String**|  search pattern | 
 **family** | **optional.String**|  family | 
 **interfaceType** | **optional.String**|  interface type | 
 **distributionMethod** | **optional.String**|  distribution method | 
 **additional** | **optional.Bool**|  additional | 
 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]Network**](network.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllPowerFeedsForAnExistingKubernetesBaremetalNode

> []PowerFeed ListAllPowerFeedsForAnExistingKubernetesBaremetalNode(ctx, serverId)

List all power feeds for an existing kubernetes baremetal node

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string**|  | 

### Return type

[**[]PowerFeed**](Power_feed.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAnExistingKubernetesBaremetalNode

> TheKubernetesBaremetalNodeEntitySchema RetrieveAnExistingKubernetesBaremetalNode(ctx, serverId)

Retrieve an existing kubernetes baremetal node

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string**|  | 

### Return type

[**TheKubernetesBaremetalNodeEntitySchema**](The_Kubernetes_Baremetal_Node_Entity_Schema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

