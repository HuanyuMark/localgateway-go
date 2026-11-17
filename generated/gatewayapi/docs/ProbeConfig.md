# ProbeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | 探活路径，必须以 &#x60;/&#x60; 开头，例如 &#x60;/health&#x60;。 |
**TimeoutMs** | Pointer to **int32** | 单次探活超时（毫秒）。 | [optional]

## Methods

### NewProbeConfig

`func NewProbeConfig(path string, ) *ProbeConfig`

NewProbeConfig instantiates a new ProbeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProbeConfigWithDefaults

`func NewProbeConfigWithDefaults() *ProbeConfig`

NewProbeConfigWithDefaults instantiates a new ProbeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ProbeConfig) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ProbeConfig) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ProbeConfig) SetPath(v string)`

SetPath sets Path field to given value.


### GetTimeoutMs

`func (o *ProbeConfig) GetTimeoutMs() int32`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *ProbeConfig) GetTimeoutMsOk() (*int32, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *ProbeConfig) SetTimeoutMs(v int32)`

SetTimeoutMs sets TimeoutMs field to given value.

### HasTimeoutMs

`func (o *ProbeConfig) HasTimeoutMs() bool`

HasTimeoutMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
