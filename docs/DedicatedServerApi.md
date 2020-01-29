# \DedicatedServerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortReleaseForAnExistingDedicatedServer**](DedicatedServerApi.md#AbortReleaseForAnExistingDedicatedServer) | **Post** /v1/hosts/dedicated_servers/{server_id}/abort_release | Abort release for an existing dedicated server
[**CreateANewDedicatedServer**](DedicatedServerApi.md#CreateANewDedicatedServer) | **Post** /v1/hosts/dedicated_servers | Create a new dedicated server
[**CreatePtrRecordForServerNetworks**](DedicatedServerApi.md#CreatePtrRecordForServerNetworks) | **Post** /v1/hosts/dedicated_servers/{server_id}/ptr_records | Create PTR record for server networks
[**DeleteAnExistingPtrRecord**](DedicatedServerApi.md#DeleteAnExistingPtrRecord) | **Delete** /v1/hosts/dedicated_servers/{server_id}/ptr_records/{record_id} | Delete an existing PTR record
[**ListAllConnectionsForAnExistingDedicatedServer**](DedicatedServerApi.md#ListAllConnectionsForAnExistingDedicatedServer) | **Get** /v1/hosts/dedicated_servers/{lease_id}/connections | List all connections for an existing dedicated server
[**ListAllNetworksForAnExistingDedicatedServer**](DedicatedServerApi.md#ListAllNetworksForAnExistingDedicatedServer) | **Get** /v1/hosts/dedicated_servers/{server_id}/networks | List all networks for an existing dedicated server
[**ListAllPowerFeedsForAnExistingDedicatedServer**](DedicatedServerApi.md#ListAllPowerFeedsForAnExistingDedicatedServer) | **Get** /v1/hosts/dedicated_servers/{server_id}/power_feeds | List all power feeds for an existing dedicated server
[**ListAllPtrRecordsForServerNetworks**](DedicatedServerApi.md#ListAllPtrRecordsForServerNetworks) | **Get** /v1/hosts/dedicated_servers/{server_id}/ptr_records | List all PTR records for server networks
[**RetrieveAnExistingDedicatedServer**](DedicatedServerApi.md#RetrieveAnExistingDedicatedServer) | **Get** /v1/hosts/dedicated_servers/{server_id} | Retrieve an existing dedicated server
[**ScheduleReleaseForAnExistingDedicatedServer**](DedicatedServerApi.md#ScheduleReleaseForAnExistingDedicatedServer) | **Post** /v1/hosts/dedicated_servers/{server_id}/schedule_release | Schedule release for an existing dedicated server



## AbortReleaseForAnExistingDedicatedServer

> V1HostsDedicatedServersEntity AbortReleaseForAnExistingDedicatedServer(ctx, serverId)

Abort release for an existing dedicated server

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string**|  | 

### Return type

[**V1HostsDedicatedServersEntity**](v1-hosts-dedicated_servers-_entity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateANewDedicatedServer

> []V1HostsDedicatedServersEntity CreateANewDedicatedServer(ctx, optional)

Create a new dedicated server

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateANewDedicatedServerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateANewDedicatedServerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**optional.Interface of InlineObject**](InlineObject.md)|  | 

### Return type

[**[]V1HostsDedicatedServersEntity**](v1-hosts-dedicated_servers-_entity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePtrRecordForServerNetworks

> InlineResponse200 CreatePtrRecordForServerNetworks(ctx, serverId, optional)

Create PTR record for server networks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string**|  | 
 **optional** | ***CreatePtrRecordForServerNetworksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePtrRecordForServerNetworksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject1** | [**optional.Interface of InlineObject1**](InlineObject1.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnExistingPtrRecord

> DeleteAnExistingPtrRecord(ctx, serverId, recordId)

Delete an existing PTR record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string**|  | 
**recordId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllConnectionsForAnExistingDedicatedServer

> []Connection ListAllConnectionsForAnExistingDedicatedServer(ctx, leaseId, optional)

List all connections for an existing dedicated server

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**leaseId** | **string**|  | 
 **optional** | ***ListAllConnectionsForAnExistingDedicatedServerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllConnectionsForAnExistingDedicatedServerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to port]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]Connection**](connection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllNetworksForAnExistingDedicatedServer

> []Network ListAllNetworksForAnExistingDedicatedServer(ctx, serverId, optional)

List all networks for an existing dedicated server

To list all of the networks for specific dedicated server please send `GET /v1/hosts/dedicated_servers/{server_id}/networks`  The response contains list of networks which classified by `interface_type`, `distribution_method`  - `interface_type` points on which interface this network configured. - `distribution_method` describes how network will be distributed.   In case of the `gateway`, will be reserved first 4 IP addresses for *enterprise location*,   and 1 IP address for *standard location* from the network.   In the case of the `route`, the network will be available fully for customer use. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string**|  | 
 **optional** | ***ListAllNetworksForAnExistingDedicatedServerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllNetworksForAnExistingDedicatedServerOpts struct


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


## ListAllPowerFeedsForAnExistingDedicatedServer

> []PowerFeed ListAllPowerFeedsForAnExistingDedicatedServer(ctx, serverId)

List all power feeds for an existing dedicated server

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


## ListAllPtrRecordsForServerNetworks

> []InlineResponse200 ListAllPtrRecordsForServerNetworks(ctx, serverId, optional)

List all PTR records for server networks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string**|  | 
 **optional** | ***ListAllPtrRecordsForServerNetworksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllPtrRecordsForServerNetworksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]InlineResponse200**](inline_response_200.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAnExistingDedicatedServer

> V1HostsDedicatedServersEntity RetrieveAnExistingDedicatedServer(ctx, serverId)

Retrieve an existing dedicated server

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string**|  | 

### Return type

[**V1HostsDedicatedServersEntity**](v1-hosts-dedicated_servers-_entity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleReleaseForAnExistingDedicatedServer

> V1HostsDedicatedServersEntity ScheduleReleaseForAnExistingDedicatedServer(ctx, serverId)

Schedule release for an existing dedicated server

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string**|  | 

### Return type

[**V1HostsDedicatedServersEntity**](v1-hosts-dedicated_servers-_entity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

