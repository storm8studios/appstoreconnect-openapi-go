# CiBuildRunCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | Pointer to [**CiBuildRunCreateRequestDataAttributes**](CiBuildRunCreateRequestDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**CiBuildRunCreateRequestDataRelationships**](CiBuildRunCreateRequestDataRelationships.md) |  | [optional] 

## Methods

### NewCiBuildRunCreateRequestData

`func NewCiBuildRunCreateRequestData(type_ string, ) *CiBuildRunCreateRequestData`

NewCiBuildRunCreateRequestData instantiates a new CiBuildRunCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiBuildRunCreateRequestDataWithDefaults

`func NewCiBuildRunCreateRequestDataWithDefaults() *CiBuildRunCreateRequestData`

NewCiBuildRunCreateRequestDataWithDefaults instantiates a new CiBuildRunCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CiBuildRunCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CiBuildRunCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CiBuildRunCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *CiBuildRunCreateRequestData) GetAttributes() CiBuildRunCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CiBuildRunCreateRequestData) GetAttributesOk() (*CiBuildRunCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CiBuildRunCreateRequestData) SetAttributes(v CiBuildRunCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CiBuildRunCreateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *CiBuildRunCreateRequestData) GetRelationships() CiBuildRunCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CiBuildRunCreateRequestData) GetRelationshipsOk() (*CiBuildRunCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CiBuildRunCreateRequestData) SetRelationships(v CiBuildRunCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CiBuildRunCreateRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


