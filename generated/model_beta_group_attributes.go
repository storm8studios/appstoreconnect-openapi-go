/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreopenapi

import (
	"encoding/json"
	"time"
)

// checks if the BetaGroupAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaGroupAttributes{}

// BetaGroupAttributes struct for BetaGroupAttributes
type BetaGroupAttributes struct {
	Name *string `json:"name,omitempty"`
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	IsInternalGroup *bool `json:"isInternalGroup,omitempty"`
	HasAccessToAllBuilds *bool `json:"hasAccessToAllBuilds,omitempty"`
	PublicLinkEnabled *bool `json:"publicLinkEnabled,omitempty"`
	PublicLinkId *string `json:"publicLinkId,omitempty"`
	PublicLinkLimitEnabled *bool `json:"publicLinkLimitEnabled,omitempty"`
	PublicLinkLimit *int32 `json:"publicLinkLimit,omitempty"`
	PublicLink *string `json:"publicLink,omitempty"`
	FeedbackEnabled *bool `json:"feedbackEnabled,omitempty"`
	IosBuildsAvailableForAppleSiliconMac *bool `json:"iosBuildsAvailableForAppleSiliconMac,omitempty"`
}

// NewBetaGroupAttributes instantiates a new BetaGroupAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupAttributes() *BetaGroupAttributes {
	this := BetaGroupAttributes{}
	return &this
}

// NewBetaGroupAttributesWithDefaults instantiates a new BetaGroupAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupAttributesWithDefaults() *BetaGroupAttributes {
	this := BetaGroupAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BetaGroupAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BetaGroupAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BetaGroupAttributes) SetName(v string) {
	o.Name = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *BetaGroupAttributes) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupAttributes) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *BetaGroupAttributes) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *BetaGroupAttributes) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetIsInternalGroup returns the IsInternalGroup field value if set, zero value otherwise.
func (o *BetaGroupAttributes) GetIsInternalGroup() bool {
	if o == nil || IsNil(o.IsInternalGroup) {
		var ret bool
		return ret
	}
	return *o.IsInternalGroup
}

// GetIsInternalGroupOk returns a tuple with the IsInternalGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupAttributes) GetIsInternalGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInternalGroup) {
		return nil, false
	}
	return o.IsInternalGroup, true
}

// HasIsInternalGroup returns a boolean if a field has been set.
func (o *BetaGroupAttributes) HasIsInternalGroup() bool {
	if o != nil && !IsNil(o.IsInternalGroup) {
		return true
	}

	return false
}

// SetIsInternalGroup gets a reference to the given bool and assigns it to the IsInternalGroup field.
func (o *BetaGroupAttributes) SetIsInternalGroup(v bool) {
	o.IsInternalGroup = &v
}

// GetHasAccessToAllBuilds returns the HasAccessToAllBuilds field value if set, zero value otherwise.
func (o *BetaGroupAttributes) GetHasAccessToAllBuilds() bool {
	if o == nil || IsNil(o.HasAccessToAllBuilds) {
		var ret bool
		return ret
	}
	return *o.HasAccessToAllBuilds
}

// GetHasAccessToAllBuildsOk returns a tuple with the HasAccessToAllBuilds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupAttributes) GetHasAccessToAllBuildsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAccessToAllBuilds) {
		return nil, false
	}
	return o.HasAccessToAllBuilds, true
}

// HasHasAccessToAllBuilds returns a boolean if a field has been set.
func (o *BetaGroupAttributes) HasHasAccessToAllBuilds() bool {
	if o != nil && !IsNil(o.HasAccessToAllBuilds) {
		return true
	}

	return false
}

// SetHasAccessToAllBuilds gets a reference to the given bool and assigns it to the HasAccessToAllBuilds field.
func (o *BetaGroupAttributes) SetHasAccessToAllBuilds(v bool) {
	o.HasAccessToAllBuilds = &v
}

// GetPublicLinkEnabled returns the PublicLinkEnabled field value if set, zero value otherwise.
func (o *BetaGroupAttributes) GetPublicLinkEnabled() bool {
	if o == nil || IsNil(o.PublicLinkEnabled) {
		var ret bool
		return ret
	}
	return *o.PublicLinkEnabled
}

// GetPublicLinkEnabledOk returns a tuple with the PublicLinkEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupAttributes) GetPublicLinkEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicLinkEnabled) {
		return nil, false
	}
	return o.PublicLinkEnabled, true
}

// HasPublicLinkEnabled returns a boolean if a field has been set.
func (o *BetaGroupAttributes) HasPublicLinkEnabled() bool {
	if o != nil && !IsNil(o.PublicLinkEnabled) {
		return true
	}

	return false
}

// SetPublicLinkEnabled gets a reference to the given bool and assigns it to the PublicLinkEnabled field.
func (o *BetaGroupAttributes) SetPublicLinkEnabled(v bool) {
	o.PublicLinkEnabled = &v
}

// GetPublicLinkId returns the PublicLinkId field value if set, zero value otherwise.
func (o *BetaGroupAttributes) GetPublicLinkId() string {
	if o == nil || IsNil(o.PublicLinkId) {
		var ret string
		return ret
	}
	return *o.PublicLinkId
}

// GetPublicLinkIdOk returns a tuple with the PublicLinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupAttributes) GetPublicLinkIdOk() (*string, bool) {
	if o == nil || IsNil(o.PublicLinkId) {
		return nil, false
	}
	return o.PublicLinkId, true
}

// HasPublicLinkId returns a boolean if a field has been set.
func (o *BetaGroupAttributes) HasPublicLinkId() bool {
	if o != nil && !IsNil(o.PublicLinkId) {
		return true
	}

	return false
}

// SetPublicLinkId gets a reference to the given string and assigns it to the PublicLinkId field.
func (o *BetaGroupAttributes) SetPublicLinkId(v string) {
	o.PublicLinkId = &v
}

// GetPublicLinkLimitEnabled returns the PublicLinkLimitEnabled field value if set, zero value otherwise.
func (o *BetaGroupAttributes) GetPublicLinkLimitEnabled() bool {
	if o == nil || IsNil(o.PublicLinkLimitEnabled) {
		var ret bool
		return ret
	}
	return *o.PublicLinkLimitEnabled
}

// GetPublicLinkLimitEnabledOk returns a tuple with the PublicLinkLimitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupAttributes) GetPublicLinkLimitEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicLinkLimitEnabled) {
		return nil, false
	}
	return o.PublicLinkLimitEnabled, true
}

// HasPublicLinkLimitEnabled returns a boolean if a field has been set.
func (o *BetaGroupAttributes) HasPublicLinkLimitEnabled() bool {
	if o != nil && !IsNil(o.PublicLinkLimitEnabled) {
		return true
	}

	return false
}

// SetPublicLinkLimitEnabled gets a reference to the given bool and assigns it to the PublicLinkLimitEnabled field.
func (o *BetaGroupAttributes) SetPublicLinkLimitEnabled(v bool) {
	o.PublicLinkLimitEnabled = &v
}

// GetPublicLinkLimit returns the PublicLinkLimit field value if set, zero value otherwise.
func (o *BetaGroupAttributes) GetPublicLinkLimit() int32 {
	if o == nil || IsNil(o.PublicLinkLimit) {
		var ret int32
		return ret
	}
	return *o.PublicLinkLimit
}

// GetPublicLinkLimitOk returns a tuple with the PublicLinkLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupAttributes) GetPublicLinkLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.PublicLinkLimit) {
		return nil, false
	}
	return o.PublicLinkLimit, true
}

// HasPublicLinkLimit returns a boolean if a field has been set.
func (o *BetaGroupAttributes) HasPublicLinkLimit() bool {
	if o != nil && !IsNil(o.PublicLinkLimit) {
		return true
	}

	return false
}

// SetPublicLinkLimit gets a reference to the given int32 and assigns it to the PublicLinkLimit field.
func (o *BetaGroupAttributes) SetPublicLinkLimit(v int32) {
	o.PublicLinkLimit = &v
}

// GetPublicLink returns the PublicLink field value if set, zero value otherwise.
func (o *BetaGroupAttributes) GetPublicLink() string {
	if o == nil || IsNil(o.PublicLink) {
		var ret string
		return ret
	}
	return *o.PublicLink
}

// GetPublicLinkOk returns a tuple with the PublicLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupAttributes) GetPublicLinkOk() (*string, bool) {
	if o == nil || IsNil(o.PublicLink) {
		return nil, false
	}
	return o.PublicLink, true
}

// HasPublicLink returns a boolean if a field has been set.
func (o *BetaGroupAttributes) HasPublicLink() bool {
	if o != nil && !IsNil(o.PublicLink) {
		return true
	}

	return false
}

// SetPublicLink gets a reference to the given string and assigns it to the PublicLink field.
func (o *BetaGroupAttributes) SetPublicLink(v string) {
	o.PublicLink = &v
}

// GetFeedbackEnabled returns the FeedbackEnabled field value if set, zero value otherwise.
func (o *BetaGroupAttributes) GetFeedbackEnabled() bool {
	if o == nil || IsNil(o.FeedbackEnabled) {
		var ret bool
		return ret
	}
	return *o.FeedbackEnabled
}

// GetFeedbackEnabledOk returns a tuple with the FeedbackEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupAttributes) GetFeedbackEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FeedbackEnabled) {
		return nil, false
	}
	return o.FeedbackEnabled, true
}

// HasFeedbackEnabled returns a boolean if a field has been set.
func (o *BetaGroupAttributes) HasFeedbackEnabled() bool {
	if o != nil && !IsNil(o.FeedbackEnabled) {
		return true
	}

	return false
}

// SetFeedbackEnabled gets a reference to the given bool and assigns it to the FeedbackEnabled field.
func (o *BetaGroupAttributes) SetFeedbackEnabled(v bool) {
	o.FeedbackEnabled = &v
}

// GetIosBuildsAvailableForAppleSiliconMac returns the IosBuildsAvailableForAppleSiliconMac field value if set, zero value otherwise.
func (o *BetaGroupAttributes) GetIosBuildsAvailableForAppleSiliconMac() bool {
	if o == nil || IsNil(o.IosBuildsAvailableForAppleSiliconMac) {
		var ret bool
		return ret
	}
	return *o.IosBuildsAvailableForAppleSiliconMac
}

// GetIosBuildsAvailableForAppleSiliconMacOk returns a tuple with the IosBuildsAvailableForAppleSiliconMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupAttributes) GetIosBuildsAvailableForAppleSiliconMacOk() (*bool, bool) {
	if o == nil || IsNil(o.IosBuildsAvailableForAppleSiliconMac) {
		return nil, false
	}
	return o.IosBuildsAvailableForAppleSiliconMac, true
}

// HasIosBuildsAvailableForAppleSiliconMac returns a boolean if a field has been set.
func (o *BetaGroupAttributes) HasIosBuildsAvailableForAppleSiliconMac() bool {
	if o != nil && !IsNil(o.IosBuildsAvailableForAppleSiliconMac) {
		return true
	}

	return false
}

// SetIosBuildsAvailableForAppleSiliconMac gets a reference to the given bool and assigns it to the IosBuildsAvailableForAppleSiliconMac field.
func (o *BetaGroupAttributes) SetIosBuildsAvailableForAppleSiliconMac(v bool) {
	o.IosBuildsAvailableForAppleSiliconMac = &v
}

func (o BetaGroupAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaGroupAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.IsInternalGroup) {
		toSerialize["isInternalGroup"] = o.IsInternalGroup
	}
	if !IsNil(o.HasAccessToAllBuilds) {
		toSerialize["hasAccessToAllBuilds"] = o.HasAccessToAllBuilds
	}
	if !IsNil(o.PublicLinkEnabled) {
		toSerialize["publicLinkEnabled"] = o.PublicLinkEnabled
	}
	if !IsNil(o.PublicLinkId) {
		toSerialize["publicLinkId"] = o.PublicLinkId
	}
	if !IsNil(o.PublicLinkLimitEnabled) {
		toSerialize["publicLinkLimitEnabled"] = o.PublicLinkLimitEnabled
	}
	if !IsNil(o.PublicLinkLimit) {
		toSerialize["publicLinkLimit"] = o.PublicLinkLimit
	}
	if !IsNil(o.PublicLink) {
		toSerialize["publicLink"] = o.PublicLink
	}
	if !IsNil(o.FeedbackEnabled) {
		toSerialize["feedbackEnabled"] = o.FeedbackEnabled
	}
	if !IsNil(o.IosBuildsAvailableForAppleSiliconMac) {
		toSerialize["iosBuildsAvailableForAppleSiliconMac"] = o.IosBuildsAvailableForAppleSiliconMac
	}
	return toSerialize, nil
}

type NullableBetaGroupAttributes struct {
	value *BetaGroupAttributes
	isSet bool
}

func (v NullableBetaGroupAttributes) Get() *BetaGroupAttributes {
	return v.value
}

func (v *NullableBetaGroupAttributes) Set(val *BetaGroupAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupAttributes(val *BetaGroupAttributes) *NullableBetaGroupAttributes {
	return &NullableBetaGroupAttributes{value: val, isSet: true}
}

func (v NullableBetaGroupAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


