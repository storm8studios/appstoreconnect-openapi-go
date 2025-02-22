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

// checks if the BetaTesterCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterCreateRequestDataRelationships{}

// BetaTesterCreateRequestDataRelationships struct for BetaTesterCreateRequestDataRelationships
type BetaTesterCreateRequestDataRelationships struct {
	BetaGroups *BetaTesterCreateRequestDataRelationshipsBetaGroups `json:"betaGroups,omitempty"`
	Builds *BetaGroupCreateRequestDataRelationshipsBuilds `json:"builds,omitempty"`
}

// NewBetaTesterCreateRequestDataRelationships instantiates a new BetaTesterCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterCreateRequestDataRelationships() *BetaTesterCreateRequestDataRelationships {
	this := BetaTesterCreateRequestDataRelationships{}
	return &this
}

// NewBetaTesterCreateRequestDataRelationshipsWithDefaults instantiates a new BetaTesterCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterCreateRequestDataRelationshipsWithDefaults() *BetaTesterCreateRequestDataRelationships {
	this := BetaTesterCreateRequestDataRelationships{}
	return &this
}

// GetBetaGroups returns the BetaGroups field value if set, zero value otherwise.
func (o *BetaTesterCreateRequestDataRelationships) GetBetaGroups() BetaTesterCreateRequestDataRelationshipsBetaGroups {
	if o == nil || IsNil(o.BetaGroups) {
		var ret BetaTesterCreateRequestDataRelationshipsBetaGroups
		return ret
	}
	return *o.BetaGroups
}

// GetBetaGroupsOk returns a tuple with the BetaGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterCreateRequestDataRelationships) GetBetaGroupsOk() (*BetaTesterCreateRequestDataRelationshipsBetaGroups, bool) {
	if o == nil || IsNil(o.BetaGroups) {
		return nil, false
	}
	return o.BetaGroups, true
}

// HasBetaGroups returns a boolean if a field has been set.
func (o *BetaTesterCreateRequestDataRelationships) HasBetaGroups() bool {
	if o != nil && !IsNil(o.BetaGroups) {
		return true
	}

	return false
}

// SetBetaGroups gets a reference to the given BetaTesterCreateRequestDataRelationshipsBetaGroups and assigns it to the BetaGroups field.
func (o *BetaTesterCreateRequestDataRelationships) SetBetaGroups(v BetaTesterCreateRequestDataRelationshipsBetaGroups) {
	o.BetaGroups = &v
}

// GetBuilds returns the Builds field value if set, zero value otherwise.
func (o *BetaTesterCreateRequestDataRelationships) GetBuilds() BetaGroupCreateRequestDataRelationshipsBuilds {
	if o == nil || IsNil(o.Builds) {
		var ret BetaGroupCreateRequestDataRelationshipsBuilds
		return ret
	}
	return *o.Builds
}

// GetBuildsOk returns a tuple with the Builds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterCreateRequestDataRelationships) GetBuildsOk() (*BetaGroupCreateRequestDataRelationshipsBuilds, bool) {
	if o == nil || IsNil(o.Builds) {
		return nil, false
	}
	return o.Builds, true
}

// HasBuilds returns a boolean if a field has been set.
func (o *BetaTesterCreateRequestDataRelationships) HasBuilds() bool {
	if o != nil && !IsNil(o.Builds) {
		return true
	}

	return false
}

// SetBuilds gets a reference to the given BetaGroupCreateRequestDataRelationshipsBuilds and assigns it to the Builds field.
func (o *BetaTesterCreateRequestDataRelationships) SetBuilds(v BetaGroupCreateRequestDataRelationshipsBuilds) {
	o.Builds = &v
}

func (o BetaTesterCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BetaGroups) {
		toSerialize["betaGroups"] = o.BetaGroups
	}
	if !IsNil(o.Builds) {
		toSerialize["builds"] = o.Builds
	}
	return toSerialize, nil
}

type NullableBetaTesterCreateRequestDataRelationships struct {
	value *BetaTesterCreateRequestDataRelationships
	isSet bool
}

func (v NullableBetaTesterCreateRequestDataRelationships) Get() *BetaTesterCreateRequestDataRelationships {
	return v.value
}

func (v *NullableBetaTesterCreateRequestDataRelationships) Set(val *BetaTesterCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterCreateRequestDataRelationships(val *BetaTesterCreateRequestDataRelationships) *NullableBetaTesterCreateRequestDataRelationships {
	return &NullableBetaTesterCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableBetaTesterCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


