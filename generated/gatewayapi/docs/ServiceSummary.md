# ServiceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiInfo** | Pointer to **interface{}** | 仅当 &#x60;fields&#x3D;apiInfo&#x60; 时返回。 | [optional]
**LastProbe** | Pointer to **NullableString** | 最近一次探测时间。 | [optional]
**Name** | **string** | 服务名。 |
**Probe** | **string** | 探活路径。 |
**ProjectRoot** | Pointer to **NullableString** | 仅当 &#x60;fields&#x3D;projectRoot&#x60; 时返回。 | [optional]
**Protocol** | [**Protocol**](Protocol.md) | 协议。 |
**Revision** | **int64** | 单服务 revision。 |
**Status** | [**ServiceStatus**](ServiceStatus.md) | 状态。 |
**Version** | **string** | provider 版本。 |

## Methods

### NewServiceSummary

`func NewServiceSummary(name string, probe string, protocol Protocol, revision int64, status ServiceStatus, version string, ) *ServiceSummary`

NewServiceSummary instantiates a new ServiceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSummaryWithDefaults

`func NewServiceSummaryWithDefaults() *ServiceSummary`

NewServiceSummaryWithDefaults instantiates a new ServiceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiInfo

`func (o *ServiceSummary) GetApiInfo() interface{}`

GetApiInfo returns the ApiInfo field if non-nil, zero value otherwise.

### GetApiInfoOk

`func (o *ServiceSummary) GetApiInfoOk() (*interface{}, bool)`

GetApiInfoOk returns a tuple with the ApiInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiInfo

`func (o *ServiceSummary) SetApiInfo(v interface{})`

SetApiInfo sets ApiInfo field to given value.

### HasApiInfo

`func (o *ServiceSummary) HasApiInfo() bool`

HasApiInfo returns a boolean if a field has been set.

### SetApiInfoNil

`func (o *ServiceSummary) SetApiInfoNil(b bool)`

 SetApiInfoNil sets the value for ApiInfo to be an explicit nil

### UnsetApiInfo
`func (o *ServiceSummary) UnsetApiInfo()`

UnsetApiInfo ensures that no value is present for ApiInfo, not even an explicit nil
### GetLastProbe

`func (o *ServiceSummary) GetLastProbe() string`

GetLastProbe returns the LastProbe field if non-nil, zero value otherwise.

### GetLastProbeOk

`func (o *ServiceSummary) GetLastProbeOk() (*string, bool)`

GetLastProbeOk returns a tuple with the LastProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProbe

`func (o *ServiceSummary) SetLastProbe(v string)`

SetLastProbe sets LastProbe field to given value.

### HasLastProbe

`func (o *ServiceSummary) HasLastProbe() bool`

HasLastProbe returns a boolean if a field has been set.

### SetLastProbeNil

`func (o *ServiceSummary) SetLastProbeNil(b bool)`

 SetLastProbeNil sets the value for LastProbe to be an explicit nil

### UnsetLastProbe
`func (o *ServiceSummary) UnsetLastProbe()`

UnsetLastProbe ensures that no value is present for LastProbe, not even an explicit nil
### GetName

`func (o *ServiceSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceSummary) SetName(v string)`

SetName sets Name field to given value.


### GetProbe

`func (o *ServiceSummary) GetProbe() string`

GetProbe returns the Probe field if non-nil, zero value otherwise.

### GetProbeOk

`func (o *ServiceSummary) GetProbeOk() (*string, bool)`

GetProbeOk returns a tuple with the Probe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbe

`func (o *ServiceSummary) SetProbe(v string)`

SetProbe sets Probe field to given value.


### GetProjectRoot

`func (o *ServiceSummary) GetProjectRoot() string`

GetProjectRoot returns the ProjectRoot field if non-nil, zero value otherwise.

### GetProjectRootOk

`func (o *ServiceSummary) GetProjectRootOk() (*string, bool)`

GetProjectRootOk returns a tuple with the ProjectRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRoot

`func (o *ServiceSummary) SetProjectRoot(v string)`

SetProjectRoot sets ProjectRoot field to given value.

### HasProjectRoot

`func (o *ServiceSummary) HasProjectRoot() bool`

HasProjectRoot returns a boolean if a field has been set.

### SetProjectRootNil

`func (o *ServiceSummary) SetProjectRootNil(b bool)`

 SetProjectRootNil sets the value for ProjectRoot to be an explicit nil

### UnsetProjectRoot
`func (o *ServiceSummary) UnsetProjectRoot()`

UnsetProjectRoot ensures that no value is present for ProjectRoot, not even an explicit nil
### GetProtocol

`func (o *ServiceSummary) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ServiceSummary) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ServiceSummary) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetRevision

`func (o *ServiceSummary) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServiceSummary) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServiceSummary) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetStatus

`func (o *ServiceSummary) GetStatus() ServiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceSummary) GetStatusOk() (*ServiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceSummary) SetStatus(v ServiceStatus)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *ServiceSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceSummary) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
