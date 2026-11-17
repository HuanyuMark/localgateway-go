# OwnerSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImagePath** | **string** | 进程映像完整路径。 |
**Pid** | **int32** | 进程 ID（仅供展示，可能被复用）。 |
**UserName** | Pointer to **NullableString** | 用户名，取不到时为空。 | [optional]
**UserSid** | **string** | Windows 用户 SID。 |

## Methods

### NewOwnerSummary

`func NewOwnerSummary(imagePath string, pid int32, userSid string, ) *OwnerSummary`

NewOwnerSummary instantiates a new OwnerSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerSummaryWithDefaults

`func NewOwnerSummaryWithDefaults() *OwnerSummary`

NewOwnerSummaryWithDefaults instantiates a new OwnerSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImagePath

`func (o *OwnerSummary) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *OwnerSummary) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *OwnerSummary) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.


### GetPid

`func (o *OwnerSummary) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *OwnerSummary) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *OwnerSummary) SetPid(v int32)`

SetPid sets Pid field to given value.


### GetUserName

`func (o *OwnerSummary) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *OwnerSummary) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *OwnerSummary) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *OwnerSummary) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *OwnerSummary) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *OwnerSummary) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetUserSid

`func (o *OwnerSummary) GetUserSid() string`

GetUserSid returns the UserSid field if non-nil, zero value otherwise.

### GetUserSidOk

`func (o *OwnerSummary) GetUserSidOk() (*string, bool)`

GetUserSidOk returns a tuple with the UserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSid

`func (o *OwnerSummary) SetUserSid(v string)`

SetUserSid sets UserSid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
