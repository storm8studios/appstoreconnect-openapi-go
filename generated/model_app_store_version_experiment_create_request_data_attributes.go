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

// checks if the AppStoreVersionExperimentCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentCreateRequestDataAttributes{}

// AppStoreVersionExperimentCreateRequestDataAttributes struct for AppStoreVersionExperimentCreateRequestDataAttributes
type AppStoreVersionExperimentCreateRequestDataAttributes struct {
	Name string `json:"name"`
	TrafficProportion int32 `json:"trafficProportion"`
}

// NewAppStoreVersionExperimentCreateRequestDataAttributes instantiates a new AppStoreVersionExperimentCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentCreateRequestDataAttributes(name string, trafficProportion int32) *AppStoreVersionExperimentCreateRequestDataAttributes {
	this := AppStoreVersionExperimentCreateRequestDataAttributes{}
	this.Name = name
	this.TrafficProportion = trafficProportion
	return &this
}

// NewAppStoreVersionExperimentCreateRequestDataAttributesWithDefaults instantiates a new AppStoreVersionExperimentCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentCreateRequestDataAttributesWithDefaults() *AppStoreVersionExperimentCreateRequestDataAttributes {
	this := AppStoreVersionExperimentCreateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *AppStoreVersionExperimentCreateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppStoreVersionExperimentCreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetTrafficProportion returns the TrafficProportion field value
func (o *AppStoreVersionExperimentCreateRequestDataAttributes) GetTrafficProportion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TrafficProportion
}

// GetTrafficProportionOk returns a tuple with the TrafficProportion field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentCreateRequestDataAttributes) GetTrafficProportionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrafficProportion, true
}

// SetTrafficProportion sets field value
func (o *AppStoreVersionExperimentCreateRequestDataAttributes) SetTrafficProportion(v int32) {
	o.TrafficProportion = v
}

func (o AppStoreVersionExperimentCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["trafficProportion"] = o.TrafficProportion
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentCreateRequestDataAttributes struct {
	value *AppStoreVersionExperimentCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppStoreVersionExperimentCreateRequestDataAttributes) Get() *AppStoreVersionExperimentCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppStoreVersionExperimentCreateRequestDataAttributes) Set(val *AppStoreVersionExperimentCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentCreateRequestDataAttributes(val *AppStoreVersionExperimentCreateRequestDataAttributes) *NullableAppStoreVersionExperimentCreateRequestDataAttributes {
	return &NullableAppStoreVersionExperimentCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


