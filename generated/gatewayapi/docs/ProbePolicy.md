# ProbePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpireAfterMs** | Pointer to **int64** | offline 后多久未恢复判 expired。 | [optional]
**OfflineAfterMisses** | Pointer to **int32** | 连续 miss 多少次判 offline。 | [optional]
**ProbeIntervalMs** | Pointer to **int32** | 探测周期（毫秒）。 | [optional]
**ProbeTimeoutMs** | Pointer to **int32** | 单次探测超时（毫秒）。 | [optional]

## Methods

### NewProbePolicy

`func NewProbePolicy() *ProbePolicy`

NewProbePolicy instantiates a new ProbePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProbePolicyWithDefaults

`func NewProbePolicyWithDefaults() *ProbePolicy`

NewProbePolicyWithDefaults instantiates a new ProbePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpireAfterMs

`func (o *ProbePolicy) GetExpireAfterMs() int64`

GetExpireAfterMs returns the ExpireAfterMs field if non-nil, zero value otherwise.

### GetExpireAfterMsOk

`func (o *ProbePolicy) GetExpireAfterMsOk() (*int64, bool)`

GetExpireAfterMsOk returns a tuple with the ExpireAfterMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfterMs

`func (o *ProbePolicy) SetExpireAfterMs(v int64)`

SetExpireAfterMs sets ExpireAfterMs field to given value.

### HasExpireAfterMs

`func (o *ProbePolicy) HasExpireAfterMs() bool`

HasExpireAfterMs returns a boolean if a field has been set.

### GetOfflineAfterMisses

`func (o *ProbePolicy) GetOfflineAfterMisses() int32`

GetOfflineAfterMisses returns the OfflineAfterMisses field if non-nil, zero value otherwise.

### GetOfflineAfterMissesOk

`func (o *ProbePolicy) GetOfflineAfterMissesOk() (*int32, bool)`

GetOfflineAfterMissesOk returns a tuple with the OfflineAfterMisses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineAfterMisses

`func (o *ProbePolicy) SetOfflineAfterMisses(v int32)`

SetOfflineAfterMisses sets OfflineAfterMisses field to given value.

### HasOfflineAfterMisses

`func (o *ProbePolicy) HasOfflineAfterMisses() bool`

HasOfflineAfterMisses returns a boolean if a field has been set.

### GetProbeIntervalMs

`func (o *ProbePolicy) GetProbeIntervalMs() int32`

GetProbeIntervalMs returns the ProbeIntervalMs field if non-nil, zero value otherwise.

### GetProbeIntervalMsOk

`func (o *ProbePolicy) GetProbeIntervalMsOk() (*int32, bool)`

GetProbeIntervalMsOk returns a tuple with the ProbeIntervalMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeIntervalMs

`func (o *ProbePolicy) SetProbeIntervalMs(v int32)`

SetProbeIntervalMs sets ProbeIntervalMs field to given value.

### HasProbeIntervalMs

`func (o *ProbePolicy) HasProbeIntervalMs() bool`

HasProbeIntervalMs returns a boolean if a field has been set.

### GetProbeTimeoutMs

`func (o *ProbePolicy) GetProbeTimeoutMs() int32`

GetProbeTimeoutMs returns the ProbeTimeoutMs field if non-nil, zero value otherwise.

### GetProbeTimeoutMsOk

`func (o *ProbePolicy) GetProbeTimeoutMsOk() (*int32, bool)`

GetProbeTimeoutMsOk returns a tuple with the ProbeTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeTimeoutMs

`func (o *ProbePolicy) SetProbeTimeoutMs(v int32)`

SetProbeTimeoutMs sets ProbeTimeoutMs field to given value.

### HasProbeTimeoutMs

`func (o *ProbePolicy) HasProbeTimeoutMs() bool`

HasProbeTimeoutMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
