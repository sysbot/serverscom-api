# \CloudSnapshotsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstanceSnapshots**](CloudSnapshotsApi.md#CreateInstanceSnapshots) | **Post** /v1/cloud_computing/regions/{region_id}/snapshots | Create instance snapshots
[**DeleteSnapshot**](CloudSnapshotsApi.md#DeleteSnapshot) | **Delete** /v1/cloud_computing/regions/{region_id}/snapshots/{snapshot_id} | Delete snapshot
[**ListCloudSnapshots**](CloudSnapshotsApi.md#ListCloudSnapshots) | **Get** /v1/cloud_computing/regions/{region_id}/snapshots | List cloud snapshots



## CreateInstanceSnapshots

> TheItemsSchema5 CreateInstanceSnapshots(ctx, name, instanceId, regionId)

Create instance snapshots

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  name | 
**instanceId** | **string**|  instance | 
**regionId** | **string**|  | 

### Return type

[**TheItemsSchema5**](The_Items_Schema_5.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSnapshot

> DeleteSnapshot(ctx, regionId, snapshotId, optional)

Delete snapshot

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **string**|  | 
**snapshotId** | **string**|  | 
 **optional** | ***DeleteSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSnapshotOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceId** | **optional.String**|  instance | 
 **isBackup** | **optional.Bool**|  is backup | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudSnapshots

> []TheItemsSchema5 ListCloudSnapshots(ctx, regionId, optional)

List cloud snapshots

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **string**|  | 
 **optional** | ***ListCloudSnapshotsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCloudSnapshotsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceId** | **optional.String**|  instance | 
 **isBackup** | **optional.Bool**|  is backup | 
 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]

### Return type

[**[]TheItemsSchema5**](The_Items_Schema_5.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

