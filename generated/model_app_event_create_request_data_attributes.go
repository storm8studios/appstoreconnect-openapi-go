/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreopenapi

import (
	"encoding/json"
)

// checks if the AppEventCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventCreateRequestDataAttributes{}

// AppEventCreateRequestDataAttributes struct for AppEventCreateRequestDataAttributes
type AppEventCreateRequestDataAttributes struct {
	ReferenceName string `json:"referenceName"`
	Badge *string `json:"badge,omitempty"`
	DeepLink *string `json:"deepLink,omitempty"`
	PurchaseRequirement *string `json:"purchaseRequirement,omitempty"`
	PrimaryLocale *string `json:"primaryLocale,omitempty"`
	Priority *string `json:"priority,omitempty"`
	Purpose *string `json:"purpose,omitempty"`
	TerritorySchedules []AppEventAttributesTerritorySchedulesInner `json:"territorySchedules,omitempty"`
}

// NewAppEventCreateRequestDataAttributes instantiates a new AppEventCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventCreateRequestDataAttributes(referenceName string) *AppEventCreateRequestDataAttributes {
	this := AppEventCreateRequestDataAttributes{}
	this.ReferenceName = referenceName
	return &this
}

// NewAppEventCreateRequestDataAttributesWithDefaults instantiates a new AppEventCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventCreateRequestDataAttributesWithDefaults() *AppEventCreateRequestDataAttributes {
	this := AppEventCreateRequestDataAttributes{}
	return &this
}

// GetReferenceName returns the ReferenceName field value
func (o *AppEventCreateRequestDataAttributes) GetReferenceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value
// and a boolean to check if the value has been set.
func (o *AppEventCreateRequestDataAttributes) GetReferenceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceName, true
}

// SetReferenceName sets field value
func (o *AppEventCreateRequestDataAttributes) SetReferenceName(v string) {
	o.ReferenceName = v
}

// GetBadge returns the Badge field value if set, zero value otherwise.
func (o *AppEventCreateRequestDataAttributes) GetBadge() string {
	if o == nil || IsNil(o.Badge) {
		var ret string
		return ret
	}
	return *o.Badge
}

// GetBadgeOk returns a tuple with the Badge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventCreateRequestDataAttributes) GetBadgeOk() (*string, bool) {
	if o == nil || IsNil(o.Badge) {
		return nil, false
	}
	return o.Badge, true
}

// HasBadge returns a boolean if a field has been set.
func (o *AppEventCreateRequestDataAttributes) HasBadge() bool {
	if o != nil && !IsNil(o.Badge) {
		return true
	}

	return false
}

// SetBadge gets a reference to the given string and assigns it to the Badge field.
func (o *AppEventCreateRequestDataAttributes) SetBadge(v string) {
	o.Badge = &v
}

// GetDeepLink returns the DeepLink field value if set, zero value otherwise.
func (o *AppEventCreateRequestDataAttributes) GetDeepLink() string {
	if o == nil || IsNil(o.DeepLink) {
		var ret string
		return ret
	}
	return *o.DeepLink
}

// GetDeepLinkOk returns a tuple with the DeepLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventCreateRequestDataAttributes) GetDeepLinkOk() (*string, bool) {
	if o == nil || IsNil(o.DeepLink) {
		return nil, false
	}
	return o.DeepLink, true
}

// HasDeepLink returns a boolean if a field has been set.
func (o *AppEventCreateRequestDataAttributes) HasDeepLink() bool {
	if o != nil && !IsNil(o.DeepLink) {
		return true
	}

	return false
}

// SetDeepLink gets a reference to the given string and assigns it to the DeepLink field.
func (o *AppEventCreateRequestDataAttributes) SetDeepLink(v string) {
	o.DeepLink = &v
}

// GetPurchaseRequirement returns the PurchaseRequirement field value if set, zero value otherwise.
func (o *AppEventCreateRequestDataAttributes) GetPurchaseRequirement() string {
	if o == nil || IsNil(o.PurchaseRequirement) {
		var ret string
		return ret
	}
	return *o.PurchaseRequirement
}

// GetPurchaseRequirementOk returns a tuple with the PurchaseRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventCreateRequestDataAttributes) GetPurchaseRequirementOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseRequirement) {
		return nil, false
	}
	return o.PurchaseRequirement, true
}

// HasPurchaseRequirement returns a boolean if a field has been set.
func (o *AppEventCreateRequestDataAttributes) HasPurchaseRequirement() bool {
	if o != nil && !IsNil(o.PurchaseRequirement) {
		return true
	}

	return false
}

// SetPurchaseRequirement gets a reference to the given string and assigns it to the PurchaseRequirement field.
func (o *AppEventCreateRequestDataAttributes) SetPurchaseRequirement(v string) {
	o.PurchaseRequirement = &v
}

// GetPrimaryLocale returns the PrimaryLocale field value if set, zero value otherwise.
func (o *AppEventCreateRequestDataAttributes) GetPrimaryLocale() string {
	if o == nil || IsNil(o.PrimaryLocale) {
		var ret string
		return ret
	}
	return *o.PrimaryLocale
}

// GetPrimaryLocaleOk returns a tuple with the PrimaryLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventCreateRequestDataAttributes) GetPrimaryLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryLocale) {
		return nil, false
	}
	return o.PrimaryLocale, true
}

// HasPrimaryLocale returns a boolean if a field has been set.
func (o *AppEventCreateRequestDataAttributes) HasPrimaryLocale() bool {
	if o != nil && !IsNil(o.PrimaryLocale) {
		return true
	}

	return false
}

// SetPrimaryLocale gets a reference to the given string and assigns it to the PrimaryLocale field.
func (o *AppEventCreateRequestDataAttributes) SetPrimaryLocale(v string) {
	o.PrimaryLocale = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *AppEventCreateRequestDataAttributes) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventCreateRequestDataAttributes) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *AppEventCreateRequestDataAttributes) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *AppEventCreateRequestDataAttributes) SetPriority(v string) {
	o.Priority = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *AppEventCreateRequestDataAttributes) GetPurpose() string {
	if o == nil || IsNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventCreateRequestDataAttributes) GetPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *AppEventCreateRequestDataAttributes) HasPurpose() bool {
	if o != nil && !IsNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *AppEventCreateRequestDataAttributes) SetPurpose(v string) {
	o.Purpose = &v
}

// GetTerritorySchedules returns the TerritorySchedules field value if set, zero value otherwise.
func (o *AppEventCreateRequestDataAttributes) GetTerritorySchedules() []AppEventAttributesTerritorySchedulesInner {
	if o == nil || IsNil(o.TerritorySchedules) {
		var ret []AppEventAttributesTerritorySchedulesInner
		return ret
	}
	return o.TerritorySchedules
}

// GetTerritorySchedulesOk returns a tuple with the TerritorySchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventCreateRequestDataAttributes) GetTerritorySchedulesOk() ([]AppEventAttributesTerritorySchedulesInner, bool) {
	if o == nil || IsNil(o.TerritorySchedules) {
		return nil, false
	}
	return o.TerritorySchedules, true
}

// HasTerritorySchedules returns a boolean if a field has been set.
func (o *AppEventCreateRequestDataAttributes) HasTerritorySchedules() bool {
	if o != nil && !IsNil(o.TerritorySchedules) {
		return true
	}

	return false
}

// SetTerritorySchedules gets a reference to the given []AppEventAttributesTerritorySchedulesInner and assigns it to the TerritorySchedules field.
func (o *AppEventCreateRequestDataAttributes) SetTerritorySchedules(v []AppEventAttributesTerritorySchedulesInner) {
	o.TerritorySchedules = v
}

func (o AppEventCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["referenceName"] = o.ReferenceName
	if !IsNil(o.Badge) {
		toSerialize["badge"] = o.Badge
	}
	if !IsNil(o.DeepLink) {
		toSerialize["deepLink"] = o.DeepLink
	}
	if !IsNil(o.PurchaseRequirement) {
		toSerialize["purchaseRequirement"] = o.PurchaseRequirement
	}
	if !IsNil(o.PrimaryLocale) {
		toSerialize["primaryLocale"] = o.PrimaryLocale
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	if !IsNil(o.TerritorySchedules) {
		toSerialize["territorySchedules"] = o.TerritorySchedules
	}
	return toSerialize, nil
}

type NullableAppEventCreateRequestDataAttributes struct {
	value *AppEventCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppEventCreateRequestDataAttributes) Get() *AppEventCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppEventCreateRequestDataAttributes) Set(val *AppEventCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventCreateRequestDataAttributes(val *AppEventCreateRequestDataAttributes) *NullableAppEventCreateRequestDataAttributes {
	return &NullableAppEventCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppEventCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


