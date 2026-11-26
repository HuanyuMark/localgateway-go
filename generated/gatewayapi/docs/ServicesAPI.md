# \ServicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceCors**](ServicesAPI.md#DeleteServiceCors) | **Delete** /gw/services/{serviceName}/cors | &#x60;DELETE /gw/services/{serviceName}/cors&#x60;
[**GetService**](ServicesAPI.md#GetService) | **Get** /gw/services/{serviceName} | &#x60;GET /gw/services/{serviceName}&#x60;
[**GetServiceApiInfo**](ServicesAPI.md#GetServiceApiInfo) | **Get** /gw/services/{serviceName}/api-info | &#x60;GET /gw/services/{serviceName}/api-info&#x60;
[**ListServices**](ServicesAPI.md#ListServices) | **Get** /gw/services | &#x60;GET /gw/services&#x60;
[**PatchService**](ServicesAPI.md#PatchService) | **Patch** /gw/services/{serviceName} | &#x60;PATCH /gw/services/{serviceName}&#x60;
[**PutServiceCors**](ServicesAPI.md#PutServiceCors) | **Put** /gw/services/{serviceName}/cors | &#x60;PUT /gw/services/{serviceName}/cors&#x60;
[**RegisterService**](ServicesAPI.md#RegisterService) | **Put** /gw/services/{serviceName} | &#x60;PUT /gw/services/{serviceName}&#x60;
[**UnregisterService**](ServicesAPI.md#UnregisterService) | **Delete** /gw/services/{serviceName} | &#x60;DELETE /gw/services/{serviceName}&#x60;



## DeleteServiceCors

> DeleteServiceCors(ctx, serviceName).Execute()

`DELETE /gw/services/{serviceName}/cors`



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/gatewayapi"
)

func main() {
	serviceName := "serviceName_example" // string | 服务名

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServicesAPI.DeleteServiceCors(context.Background(), serviceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.DeleteServiceCors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | 服务名 |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceCorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetService

> ServiceDetail GetService(ctx, serviceName).Execute()

`GET /gw/services/{serviceName}`

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/gatewayapi"
)

func main() {
	serviceName := "serviceName_example" // string | 服务名

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.GetService(context.Background(), serviceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.GetService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetService`: ServiceDetail
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | 服务名 |

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDetail**](ServiceDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceApiInfo

> interface{} GetServiceApiInfo(ctx, serviceName).Execute()

`GET /gw/services/{serviceName}/api-info`

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/gatewayapi"
)

func main() {
	serviceName := "serviceName_example" // string | 服务名

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.GetServiceApiInfo(context.Background(), serviceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.GetServiceApiInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceApiInfo`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.GetServiceApiInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | 服务名 |

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceApiInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.oai.openapi+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServices

> ServiceListResponse ListServices(ctx).NameContains(nameContains).Page(page).Limit(limit).Fields(fields).Execute()

`GET /gw/services`

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/gatewayapi"
)

func main() {
	nameContains := "nameContains_example" // string | 服务名子串，大小写不敏感 (optional)
	page := int32(56) // int32 | 页码，从 1 开始 (optional)
	limit := int32(56) // int32 | 每页数量，默认 20，上限 200 (optional)
	fields := "fields_example" // string | 扩展字段：apiInfo,projectRoot (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ListServices(context.Background()).NameContains(nameContains).Page(page).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ListServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServices`: ServiceListResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ListServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameContains** | **string** | 服务名子串，大小写不敏感 |
 **page** | **int32** | 页码，从 1 开始 |
 **limit** | **int32** | 每页数量，默认 20，上限 200 |
 **fields** | **string** | 扩展字段：apiInfo,projectRoot |

### Return type

[**ServiceListResponse**](ServiceListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchService

> RegistrationResponse PatchService(ctx, serviceName).PatchRequest(patchRequest).Execute()

`PATCH /gw/services/{serviceName}`

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/gatewayapi"
)

func main() {
	serviceName := "serviceName_example" // string | 服务名
	patchRequest := *openapiclient.NewPatchRequest() // PatchRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.PatchService(context.Background(), serviceName).PatchRequest(patchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.PatchService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchService`: RegistrationResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.PatchService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | 服务名 |

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchRequest** | [**PatchRequest**](PatchRequest.md) |  |

### Return type

[**RegistrationResponse**](RegistrationResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceCors

> CorsResponse PutServiceCors(ctx, serviceName).CorsConfig(corsConfig).Execute()

`PUT /gw/services/{serviceName}/cors`



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/gatewayapi"
)

func main() {
	serviceName := "serviceName_example" // string | 服务名
	corsConfig := *openapiclient.NewCorsConfig() // CorsConfig |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.PutServiceCors(context.Background(), serviceName).CorsConfig(corsConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.PutServiceCors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceCors`: CorsResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.PutServiceCors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | 服务名 |

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceCorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **corsConfig** | [**CorsConfig**](CorsConfig.md) |  |

### Return type

[**CorsResponse**](CorsResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterService

> RegistrationResponse RegisterService(ctx, serviceName).RegisterRequest(registerRequest).Execute()

`PUT /gw/services/{serviceName}`

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/gatewayapi"
)

func main() {
	serviceName := "serviceName_example" // string | 服务名
	registerRequest := *openapiclient.NewRegisterRequest(int32(123), "Version_example") // RegisterRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.RegisterService(context.Background(), serviceName).RegisterRequest(registerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.RegisterService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterService`: RegistrationResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.RegisterService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | 服务名 |

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registerRequest** | [**RegisterRequest**](RegisterRequest.md) |  |

### Return type

[**RegistrationResponse**](RegistrationResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterService

> UnregisterService(ctx, serviceName).Execute()

`DELETE /gw/services/{serviceName}`

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/gatewayapi"
)

func main() {
	serviceName := "serviceName_example" // string | 服务名

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServicesAPI.UnregisterService(context.Background(), serviceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.UnregisterService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | 服务名 |

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
