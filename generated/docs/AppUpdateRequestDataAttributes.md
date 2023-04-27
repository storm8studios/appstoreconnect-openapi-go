# AppUpdateRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | Pointer to **string** |  | [optional] 
**PrimaryLocale** | Pointer to **string** |  | [optional] 
**SubscriptionStatusUrl** | Pointer to **string** |  | [optional] 
**SubscriptionStatusUrlVersion** | Pointer to [**SubscriptionStatusUrlVersion**](SubscriptionStatusUrlVersion.md) |  | [optional] 
**SubscriptionStatusUrlForSandbox** | Pointer to **string** |  | [optional] 
**SubscriptionStatusUrlVersionForSandbox** | Pointer to [**SubscriptionStatusUrlVersion**](SubscriptionStatusUrlVersion.md) |  | [optional] 
**AvailableInNewTerritories** | Pointer to **bool** |  | [optional] 
**ContentRightsDeclaration** | Pointer to **string** |  | [optional] 

## Methods

### NewAppUpdateRequestDataAttributes

`func NewAppUpdateRequestDataAttributes() *AppUpdateRequestDataAttributes`

NewAppUpdateRequestDataAttributes instantiates a new AppUpdateRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUpdateRequestDataAttributesWithDefaults

`func NewAppUpdateRequestDataAttributesWithDefaults() *AppUpdateRequestDataAttributes`

NewAppUpdateRequestDataAttributesWithDefaults instantiates a new AppUpdateRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *AppUpdateRequestDataAttributes) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *AppUpdateRequestDataAttributes) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *AppUpdateRequestDataAttributes) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *AppUpdateRequestDataAttributes) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetPrimaryLocale

`func (o *AppUpdateRequestDataAttributes) GetPrimaryLocale() string`

GetPrimaryLocale returns the PrimaryLocale field if non-nil, zero value otherwise.

### GetPrimaryLocaleOk

`func (o *AppUpdateRequestDataAttributes) GetPrimaryLocaleOk() (*string, bool)`

GetPrimaryLocaleOk returns a tuple with the PrimaryLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLocale

`func (o *AppUpdateRequestDataAttributes) SetPrimaryLocale(v string)`

SetPrimaryLocale sets PrimaryLocale field to given value.

### HasPrimaryLocale

`func (o *AppUpdateRequestDataAttributes) HasPrimaryLocale() bool`

HasPrimaryLocale returns a boolean if a field has been set.

### GetSubscriptionStatusUrl

`func (o *AppUpdateRequestDataAttributes) GetSubscriptionStatusUrl() string`

GetSubscriptionStatusUrl returns the SubscriptionStatusUrl field if non-nil, zero value otherwise.

### GetSubscriptionStatusUrlOk

`func (o *AppUpdateRequestDataAttributes) GetSubscriptionStatusUrlOk() (*string, bool)`

GetSubscriptionStatusUrlOk returns a tuple with the SubscriptionStatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatusUrl

`func (o *AppUpdateRequestDataAttributes) SetSubscriptionStatusUrl(v string)`

SetSubscriptionStatusUrl sets SubscriptionStatusUrl field to given value.

### HasSubscriptionStatusUrl

`func (o *AppUpdateRequestDataAttributes) HasSubscriptionStatusUrl() bool`

HasSubscriptionStatusUrl returns a boolean if a field has been set.

### GetSubscriptionStatusUrlVersion

`func (o *AppUpdateRequestDataAttributes) GetSubscriptionStatusUrlVersion() SubscriptionStatusUrlVersion`

GetSubscriptionStatusUrlVersion returns the SubscriptionStatusUrlVersion field if non-nil, zero value otherwise.

### GetSubscriptionStatusUrlVersionOk

`func (o *AppUpdateRequestDataAttributes) GetSubscriptionStatusUrlVersionOk() (*SubscriptionStatusUrlVersion, bool)`

GetSubscriptionStatusUrlVersionOk returns a tuple with the SubscriptionStatusUrlVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatusUrlVersion

`func (o *AppUpdateRequestDataAttributes) SetSubscriptionStatusUrlVersion(v SubscriptionStatusUrlVersion)`

SetSubscriptionStatusUrlVersion sets SubscriptionStatusUrlVersion field to given value.

### HasSubscriptionStatusUrlVersion

`func (o *AppUpdateRequestDataAttributes) HasSubscriptionStatusUrlVersion() bool`

HasSubscriptionStatusUrlVersion returns a boolean if a field has been set.

### GetSubscriptionStatusUrlForSandbox

`func (o *AppUpdateRequestDataAttributes) GetSubscriptionStatusUrlForSandbox() string`

GetSubscriptionStatusUrlForSandbox returns the SubscriptionStatusUrlForSandbox field if non-nil, zero value otherwise.

### GetSubscriptionStatusUrlForSandboxOk

`func (o *AppUpdateRequestDataAttributes) GetSubscriptionStatusUrlForSandboxOk() (*string, bool)`

GetSubscriptionStatusUrlForSandboxOk returns a tuple with the SubscriptionStatusUrlForSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatusUrlForSandbox

`func (o *AppUpdateRequestDataAttributes) SetSubscriptionStatusUrlForSandbox(v string)`

SetSubscriptionStatusUrlForSandbox sets SubscriptionStatusUrlForSandbox field to given value.

### HasSubscriptionStatusUrlForSandbox

`func (o *AppUpdateRequestDataAttributes) HasSubscriptionStatusUrlForSandbox() bool`

HasSubscriptionStatusUrlForSandbox returns a boolean if a field has been set.

### GetSubscriptionStatusUrlVersionForSandbox

`func (o *AppUpdateRequestDataAttributes) GetSubscriptionStatusUrlVersionForSandbox() SubscriptionStatusUrlVersion`

GetSubscriptionStatusUrlVersionForSandbox returns the SubscriptionStatusUrlVersionForSandbox field if non-nil, zero value otherwise.

### GetSubscriptionStatusUrlVersionForSandboxOk

`func (o *AppUpdateRequestDataAttributes) GetSubscriptionStatusUrlVersionForSandboxOk() (*SubscriptionStatusUrlVersion, bool)`

GetSubscriptionStatusUrlVersionForSandboxOk returns a tuple with the SubscriptionStatusUrlVersionForSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatusUrlVersionForSandbox

`func (o *AppUpdateRequestDataAttributes) SetSubscriptionStatusUrlVersionForSandbox(v SubscriptionStatusUrlVersion)`

SetSubscriptionStatusUrlVersionForSandbox sets SubscriptionStatusUrlVersionForSandbox field to given value.

### HasSubscriptionStatusUrlVersionForSandbox

`func (o *AppUpdateRequestDataAttributes) HasSubscriptionStatusUrlVersionForSandbox() bool`

HasSubscriptionStatusUrlVersionForSandbox returns a boolean if a field has been set.

### GetAvailableInNewTerritories

`func (o *AppUpdateRequestDataAttributes) GetAvailableInNewTerritories() bool`

GetAvailableInNewTerritories returns the AvailableInNewTerritories field if non-nil, zero value otherwise.

### GetAvailableInNewTerritoriesOk

`func (o *AppUpdateRequestDataAttributes) GetAvailableInNewTerritoriesOk() (*bool, bool)`

GetAvailableInNewTerritoriesOk returns a tuple with the AvailableInNewTerritories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableInNewTerritories

`func (o *AppUpdateRequestDataAttributes) SetAvailableInNewTerritories(v bool)`

SetAvailableInNewTerritories sets AvailableInNewTerritories field to given value.

### HasAvailableInNewTerritories

`func (o *AppUpdateRequestDataAttributes) HasAvailableInNewTerritories() bool`

HasAvailableInNewTerritories returns a boolean if a field has been set.

### GetContentRightsDeclaration

`func (o *AppUpdateRequestDataAttributes) GetContentRightsDeclaration() string`

GetContentRightsDeclaration returns the ContentRightsDeclaration field if non-nil, zero value otherwise.

### GetContentRightsDeclarationOk

`func (o *AppUpdateRequestDataAttributes) GetContentRightsDeclarationOk() (*string, bool)`

GetContentRightsDeclarationOk returns a tuple with the ContentRightsDeclaration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRightsDeclaration

`func (o *AppUpdateRequestDataAttributes) SetContentRightsDeclaration(v string)`

SetContentRightsDeclaration sets ContentRightsDeclaration field to given value.

### HasContentRightsDeclaration

`func (o *AppUpdateRequestDataAttributes) HasContentRightsDeclaration() bool`

HasContentRightsDeclaration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


