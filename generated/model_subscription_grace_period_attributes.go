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

// checks if the SubscriptionGracePeriodAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGracePeriodAttributes{}

// SubscriptionGracePeriodAttributes struct for SubscriptionGracePeriodAttributes
type SubscriptionGracePeriodAttributes struct {
	OptIn *bool `json:"optIn,omitempty"`
	SandboxOptIn *bool `json:"sandboxOptIn,omitempty"`
	Duration *SubscriptionGracePeriodDuration `json:"duration,omitempty"`
	RenewalType *string `json:"renewalType,omitempty"`
}

// NewSubscriptionGracePeriodAttributes instantiates a new SubscriptionGracePeriodAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGracePeriodAttributes() *SubscriptionGracePeriodAttributes {
	this := SubscriptionGracePeriodAttributes{}
	return &this
}

// NewSubscriptionGracePeriodAttributesWithDefaults instantiates a new SubscriptionGracePeriodAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGracePeriodAttributesWithDefaults() *SubscriptionGracePeriodAttributes {
	this := SubscriptionGracePeriodAttributes{}
	return &this
}

// GetOptIn returns the OptIn field value if set, zero value otherwise.
func (o *SubscriptionGracePeriodAttributes) GetOptIn() bool {
	if o == nil || IsNil(o.OptIn) {
		var ret bool
		return ret
	}
	return *o.OptIn
}

// GetOptInOk returns a tuple with the OptIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionGracePeriodAttributes) GetOptInOk() (*bool, bool) {
	if o == nil || IsNil(o.OptIn) {
		return nil, false
	}
	return o.OptIn, true
}

// HasOptIn returns a boolean if a field has been set.
func (o *SubscriptionGracePeriodAttributes) HasOptIn() bool {
	if o != nil && !IsNil(o.OptIn) {
		return true
	}

	return false
}

// SetOptIn gets a reference to the given bool and assigns it to the OptIn field.
func (o *SubscriptionGracePeriodAttributes) SetOptIn(v bool) {
	o.OptIn = &v
}

// GetSandboxOptIn returns the SandboxOptIn field value if set, zero value otherwise.
func (o *SubscriptionGracePeriodAttributes) GetSandboxOptIn() bool {
	if o == nil || IsNil(o.SandboxOptIn) {
		var ret bool
		return ret
	}
	return *o.SandboxOptIn
}

// GetSandboxOptInOk returns a tuple with the SandboxOptIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionGracePeriodAttributes) GetSandboxOptInOk() (*bool, bool) {
	if o == nil || IsNil(o.SandboxOptIn) {
		return nil, false
	}
	return o.SandboxOptIn, true
}

// HasSandboxOptIn returns a boolean if a field has been set.
func (o *SubscriptionGracePeriodAttributes) HasSandboxOptIn() bool {
	if o != nil && !IsNil(o.SandboxOptIn) {
		return true
	}

	return false
}

// SetSandboxOptIn gets a reference to the given bool and assigns it to the SandboxOptIn field.
func (o *SubscriptionGracePeriodAttributes) SetSandboxOptIn(v bool) {
	o.SandboxOptIn = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SubscriptionGracePeriodAttributes) GetDuration() SubscriptionGracePeriodDuration {
	if o == nil || IsNil(o.Duration) {
		var ret SubscriptionGracePeriodDuration
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionGracePeriodAttributes) GetDurationOk() (*SubscriptionGracePeriodDuration, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SubscriptionGracePeriodAttributes) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given SubscriptionGracePeriodDuration and assigns it to the Duration field.
func (o *SubscriptionGracePeriodAttributes) SetDuration(v SubscriptionGracePeriodDuration) {
	o.Duration = &v
}

// GetRenewalType returns the RenewalType field value if set, zero value otherwise.
func (o *SubscriptionGracePeriodAttributes) GetRenewalType() string {
	if o == nil || IsNil(o.RenewalType) {
		var ret string
		return ret
	}
	return *o.RenewalType
}

// GetRenewalTypeOk returns a tuple with the RenewalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionGracePeriodAttributes) GetRenewalTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RenewalType) {
		return nil, false
	}
	return o.RenewalType, true
}

// HasRenewalType returns a boolean if a field has been set.
func (o *SubscriptionGracePeriodAttributes) HasRenewalType() bool {
	if o != nil && !IsNil(o.RenewalType) {
		return true
	}

	return false
}

// SetRenewalType gets a reference to the given string and assigns it to the RenewalType field.
func (o *SubscriptionGracePeriodAttributes) SetRenewalType(v string) {
	o.RenewalType = &v
}

func (o SubscriptionGracePeriodAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGracePeriodAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptIn) {
		toSerialize["optIn"] = o.OptIn
	}
	if !IsNil(o.SandboxOptIn) {
		toSerialize["sandboxOptIn"] = o.SandboxOptIn
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.RenewalType) {
		toSerialize["renewalType"] = o.RenewalType
	}
	return toSerialize, nil
}

type NullableSubscriptionGracePeriodAttributes struct {
	value *SubscriptionGracePeriodAttributes
	isSet bool
}

func (v NullableSubscriptionGracePeriodAttributes) Get() *SubscriptionGracePeriodAttributes {
	return v.value
}

func (v *NullableSubscriptionGracePeriodAttributes) Set(val *SubscriptionGracePeriodAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGracePeriodAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGracePeriodAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGracePeriodAttributes(val *SubscriptionGracePeriodAttributes) *NullableSubscriptionGracePeriodAttributes {
	return &NullableSubscriptionGracePeriodAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionGracePeriodAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGracePeriodAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


