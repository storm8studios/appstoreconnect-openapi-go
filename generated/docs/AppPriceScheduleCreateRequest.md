# AppPriceScheduleCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppPriceScheduleCreateRequestData**](AppPriceScheduleCreateRequestData.md) |  | 
**Included** | Pointer to [**[]AppPriceScheduleCreateRequestIncludedInner**](AppPriceScheduleCreateRequestIncludedInner.md) |  | [optional] 

## Methods

### NewAppPriceScheduleCreateRequest

`func NewAppPriceScheduleCreateRequest(data AppPriceScheduleCreateRequestData, ) *AppPriceScheduleCreateRequest`

NewAppPriceScheduleCreateRequest instantiates a new AppPriceScheduleCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPriceScheduleCreateRequestWithDefaults

`func NewAppPriceScheduleCreateRequestWithDefaults() *AppPriceScheduleCreateRequest`

NewAppPriceScheduleCreateRequestWithDefaults instantiates a new AppPriceScheduleCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppPriceScheduleCreateRequest) GetData() AppPriceScheduleCreateRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppPriceScheduleCreateRequest) GetDataOk() (*AppPriceScheduleCreateRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppPriceScheduleCreateRequest) SetData(v AppPriceScheduleCreateRequestData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppPriceScheduleCreateRequest) GetIncluded() []AppPriceScheduleCreateRequestIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppPriceScheduleCreateRequest) GetIncludedOk() (*[]AppPriceScheduleCreateRequestIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppPriceScheduleCreateRequest) SetIncluded(v []AppPriceScheduleCreateRequestIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppPriceScheduleCreateRequest) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


