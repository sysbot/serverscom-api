# \L2SegmentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateANewL2Segment**](L2SegmentApi.md#CreateANewL2Segment) | **Post** /v1/l2_segments | Create a new L2 segment
[**DeleteAnExistingL2Segment**](L2SegmentApi.md#DeleteAnExistingL2Segment) | **Delete** /v1/l2_segments/{l2_segment_id} | Delete an existing L2 segment
[**ListAllL2SegmentMembers**](L2SegmentApi.md#ListAllL2SegmentMembers) | **Get** /v1/l2_segments/{l2_segment_id}/members | List all L2 segment members
[**ListAllL2SegmentNetworks**](L2SegmentApi.md#ListAllL2SegmentNetworks) | **Get** /v1/l2_segments/{l2_segment_id}/networks | List all l2 segment networks
[**ListAllL2Segments**](L2SegmentApi.md#ListAllL2Segments) | **Get** /v1/l2_segments | List all L2 segments
[**ListAllLocationGroups**](L2SegmentApi.md#ListAllLocationGroups) | **Get** /v1/l2_segments/location_groups | List all Location groups
[**RetrieveAnExistingL2Segment**](L2SegmentApi.md#RetrieveAnExistingL2Segment) | **Get** /v1/l2_segments/{l2_segment_id} | Retrieve an existing L2 segment
[**UpdateAnExistingL2Segment**](L2SegmentApi.md#UpdateAnExistingL2Segment) | **Put** /v1/l2_segments/{l2_segment_id} | Update an existing L2 segment
[**UpdateAnExistingL2SegmentNetworks**](L2SegmentApi.md#UpdateAnExistingL2SegmentNetworks) | **Put** /v1/l2_segments/{l2_segment_id}/networks | Update an existing L2 segment networks



## CreateANewL2Segment

> V1L2SegmentsL2SegmentDetailed CreateANewL2Segment(ctx, optional)

Create a new L2 segment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateANewL2SegmentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateANewL2SegmentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject2** | [**optional.Interface of InlineObject2**](InlineObject2.md)|  | 

### Return type

[**V1L2SegmentsL2SegmentDetailed**](v1-l2_segments-_l2_segment_detailed.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnExistingL2Segment

> DeleteAnExistingL2Segment(ctx, l2SegmentId)

Delete an existing L2 segment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l2SegmentId** | **string**|  | 

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


## ListAllL2SegmentMembers

> []V1L2SegmentsL2Member ListAllL2SegmentMembers(ctx, l2SegmentId, optional)

List all L2 segment members

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l2SegmentId** | **string**|  | 
 **optional** | ***ListAllL2SegmentMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllL2SegmentMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]

### Return type

[**[]V1L2SegmentsL2Member**](v1-l2_segments-_l2_member.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllL2SegmentNetworks

> []V1L2SegmentsL2Network ListAllL2SegmentNetworks(ctx, l2SegmentId, optional)

List all l2 segment networks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l2SegmentId** | **string**|  | 
 **optional** | ***ListAllL2SegmentNetworksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllL2SegmentNetworksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]V1L2SegmentsL2Network**](v1-l2_segments-_l2_network.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllL2Segments

> []V1L2SegmentsL2Segment ListAllL2Segments(ctx, optional)

List all L2 segments

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAllL2SegmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllL2SegmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]V1L2SegmentsL2Segment**](v1-l2_segments-_l2_segment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllLocationGroups

> []V1L2SegmentsL2LocationGroup ListAllLocationGroups(ctx, optional)

List all Location groups

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAllLocationGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAllLocationGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupType** | **optional.String**|  group type | 
 **perPage** | **optional.Int32**|  per page | [default to 20]
 **page** | **optional.Int32**|  page | [default to 1]
 **sorting** | **optional.String**|  sorting | [default to id]
 **direction** | **optional.String**|  direction | [default to ASC]

### Return type

[**[]V1L2SegmentsL2LocationGroup**](v1-l2_segments-_l2_location_group.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAnExistingL2Segment

> V1L2SegmentsL2SegmentDetailed RetrieveAnExistingL2Segment(ctx, l2SegmentId)

Retrieve an existing L2 segment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l2SegmentId** | **string**|  | 

### Return type

[**V1L2SegmentsL2SegmentDetailed**](v1-l2_segments-_l2_segment_detailed.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnExistingL2Segment

> V1L2SegmentsL2SegmentDetailed UpdateAnExistingL2Segment(ctx, l2SegmentId, optional)

Update an existing L2 segment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l2SegmentId** | **string**|  | 
 **optional** | ***UpdateAnExistingL2SegmentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAnExistingL2SegmentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject3** | [**optional.Interface of InlineObject3**](InlineObject3.md)|  | 

### Return type

[**V1L2SegmentsL2SegmentDetailed**](v1-l2_segments-_l2_segment_detailed.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnExistingL2SegmentNetworks

> V1L2SegmentsL2SegmentDetailed UpdateAnExistingL2SegmentNetworks(ctx, l2SegmentId, optional)

Update an existing L2 segment networks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l2SegmentId** | **string**|  | 
 **optional** | ***UpdateAnExistingL2SegmentNetworksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAnExistingL2SegmentNetworksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject4** | [**optional.Interface of InlineObject4**](InlineObject4.md)|  | 

### Return type

[**V1L2SegmentsL2SegmentDetailed**](v1-l2_segments-_l2_segment_detailed.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

