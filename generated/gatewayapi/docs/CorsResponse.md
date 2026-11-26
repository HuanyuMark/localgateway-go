# CorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cors** | Pointer to [**NullableCorsConfig**](CorsConfig.md) | 当前跨域配置；&#x60;DELETE&#x60; 后为 &#x60;None&#x60;。 | [optional]
**Name** | **string** | 服务名。 |
**Revision** | **int64** | 单服务 revision。 |

## Methods

### NewCorsResponse

`func NewCorsResponse(name string, revision int64, ) *CorsResponse`

NewCorsResponse instantiates a new CorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorsResponseWithDefaults

`func NewCorsResponseWithDefaults() *CorsResponse`

NewCorsResponseWithDefaults instantiates a new CorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCors

`func (o *CorsResponse) GetCors() CorsConfig`

GetCors returns the Cors field if non-nil, zero value otherwise.

### GetCorsOk

`func (o *CorsResponse) GetCorsOk() (*CorsConfig, bool)`

GetCorsOk returns a tuple with the Cors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCors

`func (o *CorsResponse) SetCors(v CorsConfig)`

SetCors sets Cors field to given value.

### HasCors

`func (o *CorsResponse) HasCors() bool`

HasCors returns a boolean if a field has been set.

### SetCorsNil

`func (o *CorsResponse) SetCorsNil(b bool)`

 SetCorsNil sets the value for Cors to be an explicit nil

### UnsetCors
`func (o *CorsResponse) UnsetCors()`

UnsetCors ensures that no value is present for Cors, not even an explicit nil
### GetName

`func (o *CorsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CorsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CorsResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRevision

`func (o *CorsResponse) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *CorsResponse) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *CorsResponse) SetRevision(v int64)`

SetRevision sets Revision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
