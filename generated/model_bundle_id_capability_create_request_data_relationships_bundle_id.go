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

// checks if the BundleIdCapabilityCreateRequestDataRelationshipsBundleId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleIdCapabilityCreateRequestDataRelationshipsBundleId{}

// BundleIdCapabilityCreateRequestDataRelationshipsBundleId struct for BundleIdCapabilityCreateRequestDataRelationshipsBundleId
type BundleIdCapabilityCreateRequestDataRelationshipsBundleId struct {
	Data BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData `json:"data"`
}

// NewBundleIdCapabilityCreateRequestDataRelationshipsBundleId instantiates a new BundleIdCapabilityCreateRequestDataRelationshipsBundleId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdCapabilityCreateRequestDataRelationshipsBundleId(data BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) *BundleIdCapabilityCreateRequestDataRelationshipsBundleId {
	this := BundleIdCapabilityCreateRequestDataRelationshipsBundleId{}
	this.Data = data
	return &this
}

// NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdWithDefaults instantiates a new BundleIdCapabilityCreateRequestDataRelationshipsBundleId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdWithDefaults() *BundleIdCapabilityCreateRequestDataRelationshipsBundleId {
	this := BundleIdCapabilityCreateRequestDataRelationshipsBundleId{}
	return &this
}

// GetData returns the Data field value
func (o *BundleIdCapabilityCreateRequestDataRelationshipsBundleId) GetData() BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData {
	if o == nil {
		var ret BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BundleIdCapabilityCreateRequestDataRelationshipsBundleId) GetDataOk() (*BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BundleIdCapabilityCreateRequestDataRelationshipsBundleId) SetData(v BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) {
	o.Data = v
}

func (o BundleIdCapabilityCreateRequestDataRelationshipsBundleId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleIdCapabilityCreateRequestDataRelationshipsBundleId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleId struct {
	value *BundleIdCapabilityCreateRequestDataRelationshipsBundleId
	isSet bool
}

func (v NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleId) Get() *BundleIdCapabilityCreateRequestDataRelationshipsBundleId {
	return v.value
}

func (v *NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleId) Set(val *BundleIdCapabilityCreateRequestDataRelationshipsBundleId) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleId) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdCapabilityCreateRequestDataRelationshipsBundleId(val *BundleIdCapabilityCreateRequestDataRelationshipsBundleId) *NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleId {
	return &NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleId{value: val, isSet: true}
}

func (v NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


