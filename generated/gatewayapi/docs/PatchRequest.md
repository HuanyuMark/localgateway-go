# PatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiInfo** | Pointer to **interface{}** | 新的 apiInfo。 | [optional]
**Probe** | Pointer to [**NullableProbeConfig**](ProbeConfig.md) | 新的探活配置。 | [optional]
**ProjectRoot** | Pointer to **NullableString** | 新的项目根目录。 | [optional]

## Methods

### NewPatchRequest

`func NewPatchRequest() *PatchRequest`

NewPatchRequest instantiates a new PatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchRequestWithDefaults

`func NewPatchRequestWithDefaults() *PatchRequest`

NewPatchRequestWithDefaults instantiates a new PatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiInfo

`func (o *PatchRequest) GetApiInfo() interface{}`

GetApiInfo returns the ApiInfo field if non-nil, zero value otherwise.

### GetApiInfoOk

`func (o *PatchRequest) GetApiInfoOk() (*interface{}, bool)`

GetApiInfoOk returns a tuple with the ApiInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiInfo

`func (o *PatchRequest) SetApiInfo(v interface{})`

SetApiInfo sets ApiInfo field to given value.

### HasApiInfo

`func (o *PatchRequest) HasApiInfo() bool`

HasApiInfo returns a boolean if a field has been set.

### SetApiInfoNil

`func (o *PatchRequest) SetApiInfoNil(b bool)`

 SetApiInfoNil sets the value for ApiInfo to be an explicit nil

### UnsetApiInfo
`func (o *PatchRequest) UnsetApiInfo()`

UnsetApiInfo ensures that no value is present for ApiInfo, not even an explicit nil
### GetProbe

`func (o *PatchRequest) GetProbe() ProbeConfig`

GetProbe returns the Probe field if non-nil, zero value otherwise.

### GetProbeOk

`func (o *PatchRequest) GetProbeOk() (*ProbeConfig, bool)`

GetProbeOk returns a tuple with the Probe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbe

`func (o *PatchRequest) SetProbe(v ProbeConfig)`

SetProbe sets Probe field to given value.

### HasProbe

`func (o *PatchRequest) HasProbe() bool`

HasProbe returns a boolean if a field has been set.

### SetProbeNil

`func (o *PatchRequest) SetProbeNil(b bool)`

 SetProbeNil sets the value for Probe to be an explicit nil

### UnsetProbe
`func (o *PatchRequest) UnsetProbe()`

UnsetProbe ensures that no value is present for Probe, not even an explicit nil
### GetProjectRoot

`func (o *PatchRequest) GetProjectRoot() string`

GetProjectRoot returns the ProjectRoot field if non-nil, zero value otherwise.

### GetProjectRootOk

`func (o *PatchRequest) GetProjectRootOk() (*string, bool)`

GetProjectRootOk returns a tuple with the ProjectRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRoot

`func (o *PatchRequest) SetProjectRoot(v string)`

SetProjectRoot sets ProjectRoot field to given value.

### HasProjectRoot

`func (o *PatchRequest) HasProjectRoot() bool`

HasProjectRoot returns a boolean if a field has been set.

### SetProjectRootNil

`func (o *PatchRequest) SetProjectRootNil(b bool)`

 SetProjectRootNil sets the value for ProjectRoot to be an explicit nil

### UnsetProjectRoot
`func (o *PatchRequest) UnsetProjectRoot()`

UnsetProjectRoot ensures that no value is present for ProjectRoot, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
