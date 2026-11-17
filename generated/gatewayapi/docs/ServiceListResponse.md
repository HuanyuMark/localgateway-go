# ServiceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ServiceSummary**](ServiceSummary.md) | 当前页。 |
**Limit** | **int32** | 每页数量。 |
**Page** | **int32** | 页码，从 1 开始。 |
**Revision** | **int64** | 全局 revision。 |
**Total** | **int64** | 过滤后的总数。 |

## Methods

### NewServiceListResponse

`func NewServiceListResponse(items []ServiceSummary, limit int32, page int32, revision int64, total int64, ) *ServiceListResponse`

NewServiceListResponse instantiates a new ServiceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceListResponseWithDefaults

`func NewServiceListResponseWithDefaults() *ServiceListResponse`

NewServiceListResponseWithDefaults instantiates a new ServiceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ServiceListResponse) GetItems() []ServiceSummary`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ServiceListResponse) GetItemsOk() (*[]ServiceSummary, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ServiceListResponse) SetItems(v []ServiceSummary)`

SetItems sets Items field to given value.


### GetLimit

`func (o *ServiceListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ServiceListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ServiceListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetPage

`func (o *ServiceListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ServiceListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ServiceListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetRevision

`func (o *ServiceListResponse) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServiceListResponse) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServiceListResponse) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetTotal

`func (o *ServiceListResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ServiceListResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ServiceListResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
