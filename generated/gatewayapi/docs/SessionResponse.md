# SessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Caller** | [**CallerInfo**](CallerInfo.md) | 调用方进程信息。 |
**ExpiresAt** | **string** | 过期时间。 |
**Gateway** | [**GatewayInfo**](GatewayInfo.md) | 网关自身信息。 |
**Policy** | [**ProbePolicy**](ProbePolicy.md) | 探活策略。 |
**Token** | **string** | 写操作用的 bearer token。 |

## Methods

### NewSessionResponse

`func NewSessionResponse(caller CallerInfo, expiresAt string, gateway GatewayInfo, policy ProbePolicy, token string, ) *SessionResponse`

NewSessionResponse instantiates a new SessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionResponseWithDefaults

`func NewSessionResponseWithDefaults() *SessionResponse`

NewSessionResponseWithDefaults instantiates a new SessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaller

`func (o *SessionResponse) GetCaller() CallerInfo`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *SessionResponse) GetCallerOk() (*CallerInfo, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *SessionResponse) SetCaller(v CallerInfo)`

SetCaller sets Caller field to given value.


### GetExpiresAt

`func (o *SessionResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SessionResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SessionResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetGateway

`func (o *SessionResponse) GetGateway() GatewayInfo`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *SessionResponse) GetGatewayOk() (*GatewayInfo, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *SessionResponse) SetGateway(v GatewayInfo)`

SetGateway sets Gateway field to given value.


### GetPolicy

`func (o *SessionResponse) GetPolicy() ProbePolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *SessionResponse) GetPolicyOk() (*ProbePolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *SessionResponse) SetPolicy(v ProbePolicy)`

SetPolicy sets Policy field to given value.


### GetToken

`func (o *SessionResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SessionResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SessionResponse) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
