# ServiceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cors** | Pointer to [**NullableCorsConfig**](CorsConfig.md) | 跨域配置；未配置时不返回该字段。 | [optional]
**Endpoints** | **[]string** | apiInfo 中声明的路径，排序后返回。 |
**LastProbe** | Pointer to **NullableString** | 最近一次探测时间。 | [optional]
**Name** | **string** | 服务名。 |
**NotifyUrl** | Pointer to **NullableString** | 端口变更通知地址。 | [optional]
**Owner** | [**OwnerSummary**](OwnerSummary.md) | owner 摘要。 |
**Port** | **int32** | provider 端口。 |
**Probe** | [**ProbeConfig**](ProbeConfig.md) | 探活配置。 |
**ProjectRoot** | Pointer to **NullableString** | provider 项目根目录。 | [optional]
**Protocol** | [**Protocol**](Protocol.md) | 协议。 |
**Revision** | **int64** | 单服务 revision。 |
**Status** | [**ServiceStatus**](ServiceStatus.md) | 当前状态。 |
**UpdatedAt** | **string** | 最近更新时间。 |
**Version** | **string** | provider 版本。 |

## Methods

### NewServiceDetail

`func NewServiceDetail(endpoints []string, name string, owner OwnerSummary, port int32, probe ProbeConfig, protocol Protocol, revision int64, status ServiceStatus, updatedAt string, version string, ) *ServiceDetail`

NewServiceDetail instantiates a new ServiceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDetailWithDefaults

`func NewServiceDetailWithDefaults() *ServiceDetail`

NewServiceDetailWithDefaults instantiates a new ServiceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCors

`func (o *ServiceDetail) GetCors() CorsConfig`

GetCors returns the Cors field if non-nil, zero value otherwise.

### GetCorsOk

`func (o *ServiceDetail) GetCorsOk() (*CorsConfig, bool)`

GetCorsOk returns a tuple with the Cors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCors

`func (o *ServiceDetail) SetCors(v CorsConfig)`

SetCors sets Cors field to given value.

### HasCors

`func (o *ServiceDetail) HasCors() bool`

HasCors returns a boolean if a field has been set.

### SetCorsNil

`func (o *ServiceDetail) SetCorsNil(b bool)`

 SetCorsNil sets the value for Cors to be an explicit nil

### UnsetCors
`func (o *ServiceDetail) UnsetCors()`

UnsetCors ensures that no value is present for Cors, not even an explicit nil
### GetEndpoints

`func (o *ServiceDetail) GetEndpoints() []string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ServiceDetail) GetEndpointsOk() (*[]string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ServiceDetail) SetEndpoints(v []string)`

SetEndpoints sets Endpoints field to given value.


### GetLastProbe

`func (o *ServiceDetail) GetLastProbe() string`

GetLastProbe returns the LastProbe field if non-nil, zero value otherwise.

### GetLastProbeOk

`func (o *ServiceDetail) GetLastProbeOk() (*string, bool)`

GetLastProbeOk returns a tuple with the LastProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProbe

`func (o *ServiceDetail) SetLastProbe(v string)`

SetLastProbe sets LastProbe field to given value.

### HasLastProbe

`func (o *ServiceDetail) HasLastProbe() bool`

HasLastProbe returns a boolean if a field has been set.

### SetLastProbeNil

`func (o *ServiceDetail) SetLastProbeNil(b bool)`

 SetLastProbeNil sets the value for LastProbe to be an explicit nil

### UnsetLastProbe
`func (o *ServiceDetail) UnsetLastProbe()`

UnsetLastProbe ensures that no value is present for LastProbe, not even an explicit nil
### GetName

`func (o *ServiceDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceDetail) SetName(v string)`

SetName sets Name field to given value.


### GetNotifyUrl

`func (o *ServiceDetail) GetNotifyUrl() string`

GetNotifyUrl returns the NotifyUrl field if non-nil, zero value otherwise.

### GetNotifyUrlOk

`func (o *ServiceDetail) GetNotifyUrlOk() (*string, bool)`

GetNotifyUrlOk returns a tuple with the NotifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUrl

`func (o *ServiceDetail) SetNotifyUrl(v string)`

SetNotifyUrl sets NotifyUrl field to given value.

### HasNotifyUrl

`func (o *ServiceDetail) HasNotifyUrl() bool`

HasNotifyUrl returns a boolean if a field has been set.

### SetNotifyUrlNil

`func (o *ServiceDetail) SetNotifyUrlNil(b bool)`

 SetNotifyUrlNil sets the value for NotifyUrl to be an explicit nil

### UnsetNotifyUrl
`func (o *ServiceDetail) UnsetNotifyUrl()`

UnsetNotifyUrl ensures that no value is present for NotifyUrl, not even an explicit nil
### GetOwner

`func (o *ServiceDetail) GetOwner() OwnerSummary`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ServiceDetail) GetOwnerOk() (*OwnerSummary, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ServiceDetail) SetOwner(v OwnerSummary)`

SetOwner sets Owner field to given value.


### GetPort

`func (o *ServiceDetail) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ServiceDetail) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ServiceDetail) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProbe

`func (o *ServiceDetail) GetProbe() ProbeConfig`

GetProbe returns the Probe field if non-nil, zero value otherwise.

### GetProbeOk

`func (o *ServiceDetail) GetProbeOk() (*ProbeConfig, bool)`

GetProbeOk returns a tuple with the Probe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbe

`func (o *ServiceDetail) SetProbe(v ProbeConfig)`

SetProbe sets Probe field to given value.


### GetProjectRoot

`func (o *ServiceDetail) GetProjectRoot() string`

GetProjectRoot returns the ProjectRoot field if non-nil, zero value otherwise.

### GetProjectRootOk

`func (o *ServiceDetail) GetProjectRootOk() (*string, bool)`

GetProjectRootOk returns a tuple with the ProjectRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRoot

`func (o *ServiceDetail) SetProjectRoot(v string)`

SetProjectRoot sets ProjectRoot field to given value.

### HasProjectRoot

`func (o *ServiceDetail) HasProjectRoot() bool`

HasProjectRoot returns a boolean if a field has been set.

### SetProjectRootNil

`func (o *ServiceDetail) SetProjectRootNil(b bool)`

 SetProjectRootNil sets the value for ProjectRoot to be an explicit nil

### UnsetProjectRoot
`func (o *ServiceDetail) UnsetProjectRoot()`

UnsetProjectRoot ensures that no value is present for ProjectRoot, not even an explicit nil
### GetProtocol

`func (o *ServiceDetail) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ServiceDetail) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ServiceDetail) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetRevision

`func (o *ServiceDetail) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServiceDetail) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServiceDetail) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetStatus

`func (o *ServiceDetail) GetStatus() ServiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceDetail) GetStatusOk() (*ServiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceDetail) SetStatus(v ServiceStatus)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *ServiceDetail) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceDetail) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceDetail) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *ServiceDetail) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceDetail) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceDetail) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
