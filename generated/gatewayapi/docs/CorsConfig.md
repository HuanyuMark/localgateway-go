# CorsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAll** | Pointer to **bool** | 一键放行所有来源、方法与请求头。 | [optional]
**AllowCredentials** | Pointer to **bool** | 是否允许携带凭据（Cookie / &#x60;Authorization&#x60;）。 | [optional]
**AllowHeaders** | Pointer to **[]string** | 允许的请求头；为空表示回显预检请求声明的请求头。 | [optional]
**AllowMethods** | Pointer to **[]string** | 允许的方法；为空表示回显预检请求声明的方法。 | [optional]
**AllowOrigins** | Pointer to **[]string** | 允许的来源，例如 &#x60;http://localhost:5173&#x60;；支持 &#x60;*&#x60; 与 &#x60;null&#x60;。 | [optional]
**ExposeHeaders** | Pointer to **[]string** | 暴露给浏览器的响应头。 | [optional]
**MaxAgeSeconds** | Pointer to **NullableInt64** | 预检结果缓存秒数；&#x60;None&#x60; 时不发送 &#x60;Access-Control-Max-Age&#x60;。 | [optional]

## Methods

### NewCorsConfig

`func NewCorsConfig() *CorsConfig`

NewCorsConfig instantiates a new CorsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorsConfigWithDefaults

`func NewCorsConfigWithDefaults() *CorsConfig`

NewCorsConfigWithDefaults instantiates a new CorsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAll

`func (o *CorsConfig) GetAllowAll() bool`

GetAllowAll returns the AllowAll field if non-nil, zero value otherwise.

### GetAllowAllOk

`func (o *CorsConfig) GetAllowAllOk() (*bool, bool)`

GetAllowAllOk returns a tuple with the AllowAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAll

`func (o *CorsConfig) SetAllowAll(v bool)`

SetAllowAll sets AllowAll field to given value.

### HasAllowAll

`func (o *CorsConfig) HasAllowAll() bool`

HasAllowAll returns a boolean if a field has been set.

### GetAllowCredentials

`func (o *CorsConfig) GetAllowCredentials() bool`

GetAllowCredentials returns the AllowCredentials field if non-nil, zero value otherwise.

### GetAllowCredentialsOk

`func (o *CorsConfig) GetAllowCredentialsOk() (*bool, bool)`

GetAllowCredentialsOk returns a tuple with the AllowCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCredentials

`func (o *CorsConfig) SetAllowCredentials(v bool)`

SetAllowCredentials sets AllowCredentials field to given value.

### HasAllowCredentials

`func (o *CorsConfig) HasAllowCredentials() bool`

HasAllowCredentials returns a boolean if a field has been set.

### GetAllowHeaders

`func (o *CorsConfig) GetAllowHeaders() []string`

GetAllowHeaders returns the AllowHeaders field if non-nil, zero value otherwise.

### GetAllowHeadersOk

`func (o *CorsConfig) GetAllowHeadersOk() (*[]string, bool)`

GetAllowHeadersOk returns a tuple with the AllowHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHeaders

`func (o *CorsConfig) SetAllowHeaders(v []string)`

SetAllowHeaders sets AllowHeaders field to given value.

### HasAllowHeaders

`func (o *CorsConfig) HasAllowHeaders() bool`

HasAllowHeaders returns a boolean if a field has been set.

### GetAllowMethods

`func (o *CorsConfig) GetAllowMethods() []string`

GetAllowMethods returns the AllowMethods field if non-nil, zero value otherwise.

### GetAllowMethodsOk

`func (o *CorsConfig) GetAllowMethodsOk() (*[]string, bool)`

GetAllowMethodsOk returns a tuple with the AllowMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMethods

`func (o *CorsConfig) SetAllowMethods(v []string)`

SetAllowMethods sets AllowMethods field to given value.

### HasAllowMethods

`func (o *CorsConfig) HasAllowMethods() bool`

HasAllowMethods returns a boolean if a field has been set.

### GetAllowOrigins

`func (o *CorsConfig) GetAllowOrigins() []string`

GetAllowOrigins returns the AllowOrigins field if non-nil, zero value otherwise.

### GetAllowOriginsOk

`func (o *CorsConfig) GetAllowOriginsOk() (*[]string, bool)`

GetAllowOriginsOk returns a tuple with the AllowOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOrigins

`func (o *CorsConfig) SetAllowOrigins(v []string)`

SetAllowOrigins sets AllowOrigins field to given value.

### HasAllowOrigins

`func (o *CorsConfig) HasAllowOrigins() bool`

HasAllowOrigins returns a boolean if a field has been set.

### GetExposeHeaders

`func (o *CorsConfig) GetExposeHeaders() []string`

GetExposeHeaders returns the ExposeHeaders field if non-nil, zero value otherwise.

### GetExposeHeadersOk

`func (o *CorsConfig) GetExposeHeadersOk() (*[]string, bool)`

GetExposeHeadersOk returns a tuple with the ExposeHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeHeaders

`func (o *CorsConfig) SetExposeHeaders(v []string)`

SetExposeHeaders sets ExposeHeaders field to given value.

### HasExposeHeaders

`func (o *CorsConfig) HasExposeHeaders() bool`

HasExposeHeaders returns a boolean if a field has been set.

### GetMaxAgeSeconds

`func (o *CorsConfig) GetMaxAgeSeconds() int64`

GetMaxAgeSeconds returns the MaxAgeSeconds field if non-nil, zero value otherwise.

### GetMaxAgeSecondsOk

`func (o *CorsConfig) GetMaxAgeSecondsOk() (*int64, bool)`

GetMaxAgeSecondsOk returns a tuple with the MaxAgeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAgeSeconds

`func (o *CorsConfig) SetMaxAgeSeconds(v int64)`

SetMaxAgeSeconds sets MaxAgeSeconds field to given value.

### HasMaxAgeSeconds

`func (o *CorsConfig) HasMaxAgeSeconds() bool`

HasMaxAgeSeconds returns a boolean if a field has been set.

### SetMaxAgeSecondsNil

`func (o *CorsConfig) SetMaxAgeSecondsNil(b bool)`

 SetMaxAgeSecondsNil sets the value for MaxAgeSeconds to be an explicit nil

### UnsetMaxAgeSeconds
`func (o *CorsConfig) UnsetMaxAgeSeconds()`

UnsetMaxAgeSeconds ensures that no value is present for MaxAgeSeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
