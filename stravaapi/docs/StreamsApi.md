# \StreamsApi

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActivityStreams**](StreamsApi.md#GetActivityStreams) | **Get** /activities/{id}/streams | Get Activity Streams
[**GetRouteStreams**](StreamsApi.md#GetRouteStreams) | **Get** /routes/{id}/streams | Get Route Streams
[**GetSegmentEffortStreams**](StreamsApi.md#GetSegmentEffortStreams) | **Get** /segment_efforts/{id}/streams | Get Segment Effort Streams
[**GetSegmentStreams**](StreamsApi.md#GetSegmentStreams) | **Get** /segments/{id}/streams | Get Segment Streams



## GetActivityStreams

> StreamSet GetActivityStreams(ctx, id).Keys(keys).KeyByType(keyByType).Execute()

Get Activity Streams



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stevenpelley/snowflake-strava/stravaapi"
)

func main() {
    id := int64(789) // int64 | The identifier of the activity.
    keys := []string{"Keys_example"} // []string | Desired stream types.
    keyByType := true // bool | Must be true. (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StreamsApi.GetActivityStreams(context.Background(), id).Keys(keys).KeyByType(keyByType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetActivityStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActivityStreams`: StreamSet
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.GetActivityStreams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the activity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keys** | **[]string** | Desired stream types. | 
 **keyByType** | **bool** | Must be true. | [default to true]

### Return type

[**StreamSet**](StreamSet.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteStreams

> StreamSet GetRouteStreams(ctx, id).Execute()

Get Route Streams



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stevenpelley/snowflake-strava/stravaapi"
)

func main() {
    id := int64(789) // int64 | The identifier of the route.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StreamsApi.GetRouteStreams(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetRouteStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRouteStreams`: StreamSet
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.GetRouteStreams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the route. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StreamSet**](StreamSet.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegmentEffortStreams

> StreamSet GetSegmentEffortStreams(ctx, id).Keys(keys).KeyByType(keyByType).Execute()

Get Segment Effort Streams



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stevenpelley/snowflake-strava/stravaapi"
)

func main() {
    id := int64(789) // int64 | The identifier of the segment effort.
    keys := []string{"Keys_example"} // []string | The types of streams to return.
    keyByType := true // bool | Must be true. (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StreamsApi.GetSegmentEffortStreams(context.Background(), id).Keys(keys).KeyByType(keyByType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetSegmentEffortStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegmentEffortStreams`: StreamSet
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.GetSegmentEffortStreams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the segment effort. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentEffortStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keys** | **[]string** | The types of streams to return. | 
 **keyByType** | **bool** | Must be true. | [default to true]

### Return type

[**StreamSet**](StreamSet.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegmentStreams

> StreamSet GetSegmentStreams(ctx, id).Keys(keys).KeyByType(keyByType).Execute()

Get Segment Streams



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stevenpelley/snowflake-strava/stravaapi"
)

func main() {
    id := int64(789) // int64 | The identifier of the segment.
    keys := []string{"Keys_example"} // []string | The types of streams to return.
    keyByType := true // bool | Must be true. (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StreamsApi.GetSegmentStreams(context.Background(), id).Keys(keys).KeyByType(keyByType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamsApi.GetSegmentStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegmentStreams`: StreamSet
    fmt.Fprintf(os.Stdout, "Response from `StreamsApi.GetSegmentStreams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keys** | **[]string** | The types of streams to return. | 
 **keyByType** | **bool** | Must be true. | [default to true]

### Return type

[**StreamSet**](StreamSet.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

