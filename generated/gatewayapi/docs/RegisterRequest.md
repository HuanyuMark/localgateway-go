# RegisterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiInfo** | Pointer to **interface{}** | provider 的 OpenAPI 文档，可后补。 | [optional]
**NotifyUrl** | Pointer to **NullableString** | 端口变更时的定向通知地址。 | [optional]
**Port** | **int32** | provider 监听端口。 |
**Probe** | Pointer to [**NullableProbeConfig**](ProbeConfig.md) | 探活配置；&#x60;probe.path&#x60; 必填。 | [optional]
**ProjectRoot** | Pointer to **NullableString** | provider 项目根目录。 | [optional]
**Protocol** | Pointer to [**Protocol**](Protocol.md) | 协议，缺省为 &#x60;http&#x60;。 | [optional]
**Version** | **string** | provider 版本。 |

## Methods

### NewRegisterRequest

`func NewRegisterRequest(port int32, version string, ) *RegisterRequest`

NewRegisterRequest instantiates a new RegisterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterRequestWithDefaults

`func NewRegisterRequestWithDefaults() *RegisterRequest`

NewRegisterRequestWithDefaults instantiates a new RegisterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiInfo

`func (o *RegisterRequest) GetApiInfo() interface{}`

GetApiInfo returns the ApiInfo field if non-nil, zero value otherwise.

### GetApiInfoOk

`func (o *RegisterRequest) GetApiInfoOk() (*interface{}, bool)`

GetApiInfoOk returns a tuple with the ApiInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiInfo

`func (o *RegisterRequest) SetApiInfo(v interface{})`

SetApiInfo sets ApiInfo field to given value.

### HasApiInfo

`func (o *RegisterRequest) HasApiInfo() bool`

HasApiInfo returns a boolean if a field has been set.

### SetApiInfoNil

`func (o *RegisterRequest) SetApiInfoNil(b bool)`

 SetApiInfoNil sets the value for ApiInfo to be an explicit nil

### UnsetApiInfo
`func (o *RegisterRequest) UnsetApiInfo()`

UnsetApiInfo ensures that no value is present for ApiInfo, not even an explicit nil
### GetNotifyUrl

`func (o *RegisterRequest) GetNotifyUrl() string`

GetNotifyUrl returns the NotifyUrl field if non-nil, zero value otherwise.

### GetNotifyUrlOk

`func (o *RegisterRequest) GetNotifyUrlOk() (*string, bool)`

GetNotifyUrlOk returns a tuple with the NotifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUrl

`func (o *RegisterRequest) SetNotifyUrl(v string)`

SetNotifyUrl sets NotifyUrl field to given value.

### HasNotifyUrl

`func (o *RegisterRequest) HasNotifyUrl() bool`

HasNotifyUrl returns a boolean if a field has been set.

### SetNotifyUrlNil

`func (o *RegisterRequest) SetNotifyUrlNil(b bool)`

 SetNotifyUrlNil sets the value for NotifyUrl to be an explicit nil

### UnsetNotifyUrl
`func (o *RegisterRequest) UnsetNotifyUrl()`

UnsetNotifyUrl ensures that no value is present for NotifyUrl, not even an explicit nil
### GetPort

`func (o *RegisterRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RegisterRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RegisterRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProbe

`func (o *RegisterRequest) GetProbe() ProbeConfig`

GetProbe returns the Probe field if non-nil, zero value otherwise.

### GetProbeOk

`func (o *RegisterRequest) GetProbeOk() (*ProbeConfig, bool)`

GetProbeOk returns a tuple with the Probe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbe

`func (o *RegisterRequest) SetProbe(v ProbeConfig)`

SetProbe sets Probe field to given value.

### HasProbe

`func (o *RegisterRequest) HasProbe() bool`

HasProbe returns a boolean if a field has been set.

### SetProbeNil

`func (o *RegisterRequest) SetProbeNil(b bool)`

 SetProbeNil sets the value for Probe to be an explicit nil

### UnsetProbe
`func (o *RegisterRequest) UnsetProbe()`

UnsetProbe ensures that no value is present for Probe, not even an explicit nil
### GetProjectRoot

`func (o *RegisterRequest) GetProjectRoot() string`

GetProjectRoot returns the ProjectRoot field if non-nil, zero value otherwise.

### GetProjectRootOk

`func (o *RegisterRequest) GetProjectRootOk() (*string, bool)`

GetProjectRootOk returns a tuple with the ProjectRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRoot

`func (o *RegisterRequest) SetProjectRoot(v string)`

SetProjectRoot sets ProjectRoot field to given value.

### HasProjectRoot

`func (o *RegisterRequest) HasProjectRoot() bool`

HasProjectRoot returns a boolean if a field has been set.

### SetProjectRootNil

`func (o *RegisterRequest) SetProjectRootNil(b bool)`

 SetProjectRootNil sets the value for ProjectRoot to be an explicit nil

### UnsetProjectRoot
`func (o *RegisterRequest) UnsetProjectRoot()`

UnsetProjectRoot ensures that no value is present for ProjectRoot, not even an explicit nil
### GetProtocol

`func (o *RegisterRequest) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RegisterRequest) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RegisterRequest) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *RegisterRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetVersion

`func (o *RegisterRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RegisterRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RegisterRequest) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
