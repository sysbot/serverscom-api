# \KubernetesClusterApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KubernetesClusterNodes**](KubernetesClusterApi.md#KubernetesClusterNodes) | **Get** /v1/kubernetes_clusters/{kubernetes_cluster_id}/nodes | Kubernetes cluster nodes
[**KubernetesClusters**](KubernetesClusterApi.md#KubernetesClusters) | **Get** /v1/kubernetes_clusters | Kubernetes clusters
[**RetrieveAnExistingKubernetesCluster**](KubernetesClusterApi.md#RetrieveAnExistingKubernetesCluster) | **Get** /v1/kubernetes_clusters/{id} | Retrieve an existing kubernetes cluster
[**RetrieveAnExistingKubernetesClusterNode**](KubernetesClusterApi.md#RetrieveAnExistingKubernetesClusterNode) | **Get** /v1/kubernetes_clusters/{kubernetes_cluster_id}/nodes/{node_id} | Retrieve an existing kubernetes cluster node



## KubernetesClusterNodes

> []TheItemsSchema1 KubernetesClusterNodes(ctx, kubernetesClusterId, optional)
Kubernetes cluster nodes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kubernetesClusterId** | **string**|  | 
 **optional** | ***KubernetesClusterNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KubernetesClusterNodesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchPattern** | **optional.String**|  search pattern | 
 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]TheItemsSchema1**](The Items Schema_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KubernetesClusters

> []TheItemsSchema KubernetesClusters(ctx, optional)
Kubernetes clusters

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***KubernetesClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a KubernetesClustersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchPattern** | **optional.String**|  search pattern | 
 **locationId** | **optional.Int32**|  location | 
 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]TheItemsSchema**](The Items Schema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAnExistingKubernetesCluster

> InlineResponse2001 RetrieveAnExistingKubernetesCluster(ctx, id)
Retrieve an existing kubernetes cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAnExistingKubernetesClusterNode

> TheItemsSchema1 RetrieveAnExistingKubernetesClusterNode(ctx, kubernetesClusterId, nodeId)
Retrieve an existing kubernetes cluster node

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kubernetesClusterId** | **string**|  | 
**nodeId** | **string**|  | 

### Return type

[**TheItemsSchema1**](The Items Schema_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

