# CallerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImagePath** | **string** | 进程映像完整路径。 |
**Pid** | **int32** | 调用方进程 ID。 |
**UserName** | Pointer to **NullableString** | 用户名，取不到时为空。 | [optional]
**UserSid** | **string** | 调用方 Windows 用户 SID。 |

## Methods

### NewCallerInfo

`func NewCallerInfo(imagePath string, pid int32, userSid string, ) *CallerInfo`

NewCallerInfo instantiates a new CallerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallerInfoWithDefaults

`func NewCallerInfoWithDefaults() *CallerInfo`

NewCallerInfoWithDefaults instantiates a new CallerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImagePath

`func (o *CallerInfo) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *CallerInfo) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *CallerInfo) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.


### GetPid

`func (o *CallerInfo) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *CallerInfo) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *CallerInfo) SetPid(v int32)`

SetPid sets Pid field to given value.


### GetUserName

`func (o *CallerInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *CallerInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *CallerInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *CallerInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *CallerInfo) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *CallerInfo) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetUserSid

`func (o *CallerInfo) GetUserSid() string`

GetUserSid returns the UserSid field if non-nil, zero value otherwise.

### GetUserSidOk

`func (o *CallerInfo) GetUserSidOk() (*string, bool)`

GetUserSidOk returns a tuple with the UserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSid

`func (o *CallerInfo) SetUserSid(v string)`

SetUserSid sets UserSid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
