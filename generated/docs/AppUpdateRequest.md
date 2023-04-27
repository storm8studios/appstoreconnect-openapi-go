# AppUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppUpdateRequestData**](AppUpdateRequestData.md) |  | 
**Included** | Pointer to [**[]AppPriceInlineCreate**](AppPriceInlineCreate.md) |  | [optional] 

## Methods

### NewAppUpdateRequest

`func NewAppUpdateRequest(data AppUpdateRequestData, ) *AppUpdateRequest`

NewAppUpdateRequest instantiates a new AppUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUpdateRequestWithDefaults

`func NewAppUpdateRequestWithDefaults() *AppUpdateRequest`

NewAppUpdateRequestWithDefaults instantiates a new AppUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppUpdateRequest) GetData() AppUpdateRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppUpdateRequest) GetDataOk() (*AppUpdateRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppUpdateRequest) SetData(v AppUpdateRequestData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppUpdateRequest) GetIncluded() []AppPriceInlineCreate`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppUpdateRequest) GetIncludedOk() (*[]AppPriceInlineCreate, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppUpdateRequest) SetIncluded(v []AppPriceInlineCreate)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppUpdateRequest) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


