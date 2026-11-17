# RegistrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastProbe** | Pointer to **NullableString** | 最近一次探测时间。 | [optional]
**Name** | **string** | 服务名。 |
**Revision** | **int64** | 单服务 revision。 |
**Status** | [**ServiceStatus**](ServiceStatus.md) | 注册后的状态；注册时立即探测一次。 |

## Methods

### NewRegistrationResponse

`func NewRegistrationResponse(name string, revision int64, status ServiceStatus, ) *RegistrationResponse`

NewRegistrationResponse instantiates a new RegistrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationResponseWithDefaults

`func NewRegistrationResponseWithDefaults() *RegistrationResponse`

NewRegistrationResponseWithDefaults instantiates a new RegistrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastProbe

`func (o *RegistrationResponse) GetLastProbe() string`

GetLastProbe returns the LastProbe field if non-nil, zero value otherwise.

### GetLastProbeOk

`func (o *RegistrationResponse) GetLastProbeOk() (*string, bool)`

GetLastProbeOk returns a tuple with the LastProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProbe

`func (o *RegistrationResponse) SetLastProbe(v string)`

SetLastProbe sets LastProbe field to given value.

### HasLastProbe

`func (o *RegistrationResponse) HasLastProbe() bool`

HasLastProbe returns a boolean if a field has been set.

### SetLastProbeNil

`func (o *RegistrationResponse) SetLastProbeNil(b bool)`

 SetLastProbeNil sets the value for LastProbe to be an explicit nil

### UnsetLastProbe
`func (o *RegistrationResponse) UnsetLastProbe()`

UnsetLastProbe ensures that no value is present for LastProbe, not even an explicit nil
### GetName

`func (o *RegistrationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistrationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistrationResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRevision

`func (o *RegistrationResponse) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *RegistrationResponse) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *RegistrationResponse) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetStatus

`func (o *RegistrationResponse) GetStatus() ServiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegistrationResponse) GetStatusOk() (*ServiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegistrationResponse) SetStatus(v ServiceStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
