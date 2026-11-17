# GatewayInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | **string** | 管理面地址，例如 &#x60;http://127.0.0.1:47821&#x60;。 |
**Pid** | **int32** | 网关进程 ID。 |
**Version** | **string** | 网关版本。 |

## Methods

### NewGatewayInfo

`func NewGatewayInfo(baseUrl string, pid int32, version string, ) *GatewayInfo`

NewGatewayInfo instantiates a new GatewayInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayInfoWithDefaults

`func NewGatewayInfoWithDefaults() *GatewayInfo`

NewGatewayInfoWithDefaults instantiates a new GatewayInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *GatewayInfo) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *GatewayInfo) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *GatewayInfo) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetPid

`func (o *GatewayInfo) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *GatewayInfo) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *GatewayInfo) SetPid(v int32)`

SetPid sets Pid field to given value.


### GetVersion

`func (o *GatewayInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GatewayInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GatewayInfo) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
