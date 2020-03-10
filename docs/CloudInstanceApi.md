# \CloudInstanceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveInstanceUpgrade**](CloudInstanceApi.md#ApproveInstanceUpgrade) | **Post** /v1/cloud_computing/instances/{instance_id}/approve_upgrade | Approve instance upgrade
[**CreateANewCloudInstance**](CloudInstanceApi.md#CreateANewCloudInstance) | **Post** /v1/cloud_computing/instances | Create a new cloud instance
[**CreatePtrForInstance**](CloudInstanceApi.md#CreatePtrForInstance) | **Post** /v1/cloud_computing/instances/{instance_id}/ptr_records | Create PTR for instance
[**DeleteInstance**](CloudInstanceApi.md#DeleteInstance) | **Delete** /v1/cloud_computing/instances/{instance_id} | Delete instance
[**DetetePtrForInstance**](CloudInstanceApi.md#DetetePtrForInstance) | **Delete** /v1/cloud_computing/instances/{instance_id}/ptr_records/{record_id} | Detete PTR for instance
[**ExitFromRescueState**](CloudInstanceApi.md#ExitFromRescueState) | **Post** /v1/cloud_computing/instances/{instance_id}/unrescue | Exit from rescue state
[**ListCloudInstances**](CloudInstanceApi.md#ListCloudInstances) | **Get** /v1/cloud_computing/instances | List cloud instances
[**MoveInstanceToRescueState**](CloudInstanceApi.md#MoveInstanceToRescueState) | **Post** /v1/cloud_computing/instances/{instance_id}/rescue | Move instance to rescue state
[**ReinstallInstanceWithImage**](CloudInstanceApi.md#ReinstallInstanceWithImage) | **Post** /v1/cloud_computing/instances/{instance_id}/reinstall | Reinstall instance with image
[**ReturnsInstancePtrRecords**](CloudInstanceApi.md#ReturnsInstancePtrRecords) | **Get** /v1/cloud_computing/instances/{instance_id}/ptr_records | Returns instance PTR records
[**RevertInstanceUpgrade**](CloudInstanceApi.md#RevertInstanceUpgrade) | **Post** /v1/cloud_computing/instances/{instance_id}/revert_upgrade | Revert instance upgrade
[**ShowCloudInstance**](CloudInstanceApi.md#ShowCloudInstance) | **Get** /v1/cloud_computing/instances/{instance_id} | Show cloud instance
[**SwitchPowerOff**](CloudInstanceApi.md#SwitchPowerOff) | **Post** /v1/cloud_computing/instances/{instance_id}/switch_off | Switch power off
[**SwitchPowerOn**](CloudInstanceApi.md#SwitchPowerOn) | **Post** /v1/cloud_computing/instances/{instance_id}/switch_on | Switch power on
[**UpdateCloudInstance**](CloudInstanceApi.md#UpdateCloudInstance) | **Put** /v1/cloud_computing/instances/{instance_id} | Update cloud instance
[**UpgradeInstance**](CloudInstanceApi.md#UpgradeInstance) | **Post** /v1/cloud_computing/instances/{instance_id}/upgrade | Upgrade instance



## ApproveInstanceUpgrade

> TheItemsSchema3 ApproveInstanceUpgrade(ctx, instanceId)

Approve instance upgrade

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string**|  | 

### Return type

[**TheItemsSchema3**](The_Items_Schema_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateANewCloudInstance

> TheItemsSchema3 CreateANewCloudInstance(ctx, regionId, name, flavorId, imageId, optional)

Create a new cloud instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **int32**|  region | 
**name** | **string**|  name | 
**flavorId** | **string**|  flavor | 
**imageId** | **string**|  image | 
 **optional** | ***CreateANewCloudInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateANewCloudInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **gpnEnabled** | **optional.Bool**|  gpn enabled | [default to false]
 **ipv6Enabled** | **optional.Bool**|  ipv6 enabled | [default to false]
 **sshKeyFingerprint** | **optional.String**|  ssh key fingerprint | 
 **backupCopies** | **optional.Int32**|  backup copies | 

### Return type

[**TheItemsSchema3**](The_Items_Schema_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePtrForInstance

> InlineResponse200 CreatePtrForInstance(ctx, data, ip, instanceId, optional)

Create PTR for instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**data** | **string**|  data | 
**ip** | **string**|  ip | 
**instanceId** | **string**|  | 
 **optional** | ***CreatePtrForInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePtrForInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ttl** | **optional.Int32**|  ttl | 
 **priority** | **optional.Int32**|  priority | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstance

> TheItemsSchema3 DeleteInstance(ctx, instanceId)

Delete instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string**|  | 

### Return type

[**TheItemsSchema3**](The_Items_Schema_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetetePtrForInstance

> DetetePtrForInstance(ctx, instanceId, recordId)

Detete PTR for instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string**|  | 
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


## ExitFromRescueState

> TheItemsSchema3 ExitFromRescueState(ctx, instanceId)

Exit from rescue state

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string**|  | 

### Return type

[**TheItemsSchema3**](The_Items_Schema_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudInstances

> []TheItemsSchema3 ListCloudInstances(ctx, optional)

List cloud instances

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCloudInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCloudInstancesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **optional.Int32**|  region | 
 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]TheItemsSchema3**](The_Items_Schema_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveInstanceToRescueState

> TheItemsSchema3 MoveInstanceToRescueState(ctx, instanceId, optional)

Move instance to rescue state

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string**|  | 
 **optional** | ***MoveInstanceToRescueStateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MoveInstanceToRescueStateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **imageId** | **optional.String**|  image | 

### Return type

[**TheItemsSchema3**](The_Items_Schema_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReinstallInstanceWithImage

> TheItemsSchema3 ReinstallInstanceWithImage(ctx, instanceId, optional)

Reinstall instance with image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string**|  | 
 **optional** | ***ReinstallInstanceWithImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReinstallInstanceWithImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **imageId** | **optional.String**|  image | 

### Return type

[**TheItemsSchema3**](The_Items_Schema_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnsInstancePtrRecords

> []InlineResponse200 ReturnsInstancePtrRecords(ctx, instanceId, optional)

Returns instance PTR records

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string**|  | 
 **optional** | ***ReturnsInstancePtrRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReturnsInstancePtrRecordsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]

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


## RevertInstanceUpgrade

> TheItemsSchema3 RevertInstanceUpgrade(ctx, instanceId)

Revert instance upgrade

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string**|  | 

### Return type

[**TheItemsSchema3**](The_Items_Schema_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowCloudInstance

> TheItemsSchema3 ShowCloudInstance(ctx, instanceId)

Show cloud instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string**|  | 

### Return type

[**TheItemsSchema3**](The_Items_Schema_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwitchPowerOff

> TheItemsSchema3 SwitchPowerOff(ctx, instanceId)

Switch power off

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string**|  | 

### Return type

[**TheItemsSchema3**](The_Items_Schema_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwitchPowerOn

> TheItemsSchema3 SwitchPowerOn(ctx, instanceId)

Switch power on

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string**|  | 

### Return type

[**TheItemsSchema3**](The_Items_Schema_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudInstance

> TheItemsSchema3 UpdateCloudInstance(ctx, instanceId, optional)

Update cloud instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string**|  | 
 **optional** | ***UpdateCloudInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCloudInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  name | 
 **backupCopies** | **optional.Int32**|  backup copies | 
 **ipv6Enabled** | **optional.Bool**|  ipv6 enabled | 
 **gpnEnabled** | **optional.Bool**|  gpn enabled | 

### Return type

[**TheItemsSchema3**](The_Items_Schema_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeInstance

> TheItemsSchema3 UpgradeInstance(ctx, instanceId, optional)

Upgrade instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string**|  | 
 **optional** | ***UpgradeInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpgradeInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flavorId** | **optional.String**|  flavor | 

### Return type

[**TheItemsSchema3**](The_Items_Schema_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

