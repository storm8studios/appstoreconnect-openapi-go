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

// checks if the CiMacOsVersionAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiMacOsVersionAttributes{}

// CiMacOsVersionAttributes struct for CiMacOsVersionAttributes
type CiMacOsVersionAttributes struct {
	Version *string `json:"version,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewCiMacOsVersionAttributes instantiates a new CiMacOsVersionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiMacOsVersionAttributes() *CiMacOsVersionAttributes {
	this := CiMacOsVersionAttributes{}
	return &this
}

// NewCiMacOsVersionAttributesWithDefaults instantiates a new CiMacOsVersionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiMacOsVersionAttributesWithDefaults() *CiMacOsVersionAttributes {
	this := CiMacOsVersionAttributes{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CiMacOsVersionAttributes) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiMacOsVersionAttributes) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CiMacOsVersionAttributes) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CiMacOsVersionAttributes) SetVersion(v string) {
	o.Version = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CiMacOsVersionAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiMacOsVersionAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CiMacOsVersionAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CiMacOsVersionAttributes) SetName(v string) {
	o.Name = &v
}

func (o CiMacOsVersionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiMacOsVersionAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCiMacOsVersionAttributes struct {
	value *CiMacOsVersionAttributes
	isSet bool
}

func (v NullableCiMacOsVersionAttributes) Get() *CiMacOsVersionAttributes {
	return v.value
}

func (v *NullableCiMacOsVersionAttributes) Set(val *CiMacOsVersionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCiMacOsVersionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCiMacOsVersionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiMacOsVersionAttributes(val *CiMacOsVersionAttributes) *NullableCiMacOsVersionAttributes {
	return &NullableCiMacOsVersionAttributes{value: val, isSet: true}
}

func (v NullableCiMacOsVersionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiMacOsVersionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


