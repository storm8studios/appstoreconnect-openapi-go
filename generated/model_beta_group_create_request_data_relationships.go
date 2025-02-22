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

// checks if the BetaGroupCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaGroupCreateRequestDataRelationships{}

// BetaGroupCreateRequestDataRelationships struct for BetaGroupCreateRequestDataRelationships
type BetaGroupCreateRequestDataRelationships struct {
	App AppAvailabilityCreateRequestDataRelationshipsApp `json:"app"`
	Builds *BetaGroupCreateRequestDataRelationshipsBuilds `json:"builds,omitempty"`
	BetaTesters *BetaGroupCreateRequestDataRelationshipsBetaTesters `json:"betaTesters,omitempty"`
}

// NewBetaGroupCreateRequestDataRelationships instantiates a new BetaGroupCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupCreateRequestDataRelationships(app AppAvailabilityCreateRequestDataRelationshipsApp) *BetaGroupCreateRequestDataRelationships {
	this := BetaGroupCreateRequestDataRelationships{}
	this.App = app
	return &this
}

// NewBetaGroupCreateRequestDataRelationshipsWithDefaults instantiates a new BetaGroupCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupCreateRequestDataRelationshipsWithDefaults() *BetaGroupCreateRequestDataRelationships {
	this := BetaGroupCreateRequestDataRelationships{}
	return &this
}

// GetApp returns the App field value
func (o *BetaGroupCreateRequestDataRelationships) GetApp() AppAvailabilityCreateRequestDataRelationshipsApp {
	if o == nil {
		var ret AppAvailabilityCreateRequestDataRelationshipsApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestDataRelationships) GetAppOk() (*AppAvailabilityCreateRequestDataRelationshipsApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *BetaGroupCreateRequestDataRelationships) SetApp(v AppAvailabilityCreateRequestDataRelationshipsApp) {
	o.App = v
}

// GetBuilds returns the Builds field value if set, zero value otherwise.
func (o *BetaGroupCreateRequestDataRelationships) GetBuilds() BetaGroupCreateRequestDataRelationshipsBuilds {
	if o == nil || IsNil(o.Builds) {
		var ret BetaGroupCreateRequestDataRelationshipsBuilds
		return ret
	}
	return *o.Builds
}

// GetBuildsOk returns a tuple with the Builds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestDataRelationships) GetBuildsOk() (*BetaGroupCreateRequestDataRelationshipsBuilds, bool) {
	if o == nil || IsNil(o.Builds) {
		return nil, false
	}
	return o.Builds, true
}

// HasBuilds returns a boolean if a field has been set.
func (o *BetaGroupCreateRequestDataRelationships) HasBuilds() bool {
	if o != nil && !IsNil(o.Builds) {
		return true
	}

	return false
}

// SetBuilds gets a reference to the given BetaGroupCreateRequestDataRelationshipsBuilds and assigns it to the Builds field.
func (o *BetaGroupCreateRequestDataRelationships) SetBuilds(v BetaGroupCreateRequestDataRelationshipsBuilds) {
	o.Builds = &v
}

// GetBetaTesters returns the BetaTesters field value if set, zero value otherwise.
func (o *BetaGroupCreateRequestDataRelationships) GetBetaTesters() BetaGroupCreateRequestDataRelationshipsBetaTesters {
	if o == nil || IsNil(o.BetaTesters) {
		var ret BetaGroupCreateRequestDataRelationshipsBetaTesters
		return ret
	}
	return *o.BetaTesters
}

// GetBetaTestersOk returns a tuple with the BetaTesters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestDataRelationships) GetBetaTestersOk() (*BetaGroupCreateRequestDataRelationshipsBetaTesters, bool) {
	if o == nil || IsNil(o.BetaTesters) {
		return nil, false
	}
	return o.BetaTesters, true
}

// HasBetaTesters returns a boolean if a field has been set.
func (o *BetaGroupCreateRequestDataRelationships) HasBetaTesters() bool {
	if o != nil && !IsNil(o.BetaTesters) {
		return true
	}

	return false
}

// SetBetaTesters gets a reference to the given BetaGroupCreateRequestDataRelationshipsBetaTesters and assigns it to the BetaTesters field.
func (o *BetaGroupCreateRequestDataRelationships) SetBetaTesters(v BetaGroupCreateRequestDataRelationshipsBetaTesters) {
	o.BetaTesters = &v
}

func (o BetaGroupCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaGroupCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app"] = o.App
	if !IsNil(o.Builds) {
		toSerialize["builds"] = o.Builds
	}
	if !IsNil(o.BetaTesters) {
		toSerialize["betaTesters"] = o.BetaTesters
	}
	return toSerialize, nil
}

type NullableBetaGroupCreateRequestDataRelationships struct {
	value *BetaGroupCreateRequestDataRelationships
	isSet bool
}

func (v NullableBetaGroupCreateRequestDataRelationships) Get() *BetaGroupCreateRequestDataRelationships {
	return v.value
}

func (v *NullableBetaGroupCreateRequestDataRelationships) Set(val *BetaGroupCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupCreateRequestDataRelationships(val *BetaGroupCreateRequestDataRelationships) *NullableBetaGroupCreateRequestDataRelationships {
	return &NullableBetaGroupCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableBetaGroupCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


