# \RoutesApi

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRouteAsGPX**](RoutesApi.md#GetRouteAsGPX) | **Get** /routes/{id}/export_gpx | Export Route GPX
[**GetRouteAsTCX**](RoutesApi.md#GetRouteAsTCX) | **Get** /routes/{id}/export_tcx | Export Route TCX
[**GetRouteById**](RoutesApi.md#GetRouteById) | **Get** /routes/{id} | Get Route
[**GetRoutesByAthleteId**](RoutesApi.md#GetRoutesByAthleteId) | **Get** /athletes/{id}/routes | List Athlete Routes



## GetRouteAsGPX

> GetRouteAsGPX(ctx, id).Execute()

Export Route GPX



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
    r, err := apiClient.RoutesApi.GetRouteAsGPX(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutesApi.GetRouteAsGPX``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the route. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteAsGPXRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteAsTCX

> GetRouteAsTCX(ctx, id).Execute()

Export Route TCX



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
    r, err := apiClient.RoutesApi.GetRouteAsTCX(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutesApi.GetRouteAsTCX``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the route. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteAsTCXRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteById

> Route GetRouteById(ctx, id).Execute()

Get Route



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
    resp, r, err := apiClient.RoutesApi.GetRouteById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutesApi.GetRouteById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRouteById`: Route
    fmt.Fprintf(os.Stdout, "Response from `RoutesApi.GetRouteById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The identifier of the route. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Route**](Route.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutesByAthleteId

> []Route GetRoutesByAthleteId(ctx).Page(page).PerPage(perPage).Execute()

List Athlete Routes



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
    page := int32(56) // int32 | Page number. Defaults to 1. (optional)
    perPage := int32(56) // int32 | Number of items per page. Defaults to 30. (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutesApi.GetRoutesByAthleteId(context.Background()).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutesApi.GetRoutesByAthleteId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoutesByAthleteId`: []Route
    fmt.Fprintf(os.Stdout, "Response from `RoutesApi.GetRoutesByAthleteId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutesByAthleteIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Defaults to 1. | 
 **perPage** | **int32** | Number of items per page. Defaults to 30. | [default to 30]

### Return type

[**[]Route**](Route.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

