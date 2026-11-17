# ShutdownRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPort** | Pointer to **NullableInt32** | 重启后使用的新端口；&#x60;None&#x60; 表示不换端口。 | [optional]

## Methods

### NewShutdownRequest

`func NewShutdownRequest() *ShutdownRequest`

NewShutdownRequest instantiates a new ShutdownRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShutdownRequestWithDefaults

`func NewShutdownRequestWithDefaults() *ShutdownRequest`

NewShutdownRequestWithDefaults instantiates a new ShutdownRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPort

`func (o *ShutdownRequest) GetNextPort() int32`

GetNextPort returns the NextPort field if non-nil, zero value otherwise.

### GetNextPortOk

`func (o *ShutdownRequest) GetNextPortOk() (*int32, bool)`

GetNextPortOk returns a tuple with the NextPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPort

`func (o *ShutdownRequest) SetNextPort(v int32)`

SetNextPort sets NextPort field to given value.

### HasNextPort

`func (o *ShutdownRequest) HasNextPort() bool`

HasNextPort returns a boolean if a field has been set.

### SetNextPortNil

`func (o *ShutdownRequest) SetNextPortNil(b bool)`

 SetNextPortNil sets the value for NextPort to be an explicit nil

### UnsetNextPort
`func (o *ShutdownRequest) UnsetNextPort()`

UnsetNextPort ensures that no value is present for NextPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
