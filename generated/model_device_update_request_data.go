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

// checks if the DeviceUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceUpdateRequestData{}

// DeviceUpdateRequestData struct for DeviceUpdateRequestData
type DeviceUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *DeviceUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewDeviceUpdateRequestData instantiates a new DeviceUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceUpdateRequestData(type_ string, id string) *DeviceUpdateRequestData {
	this := DeviceUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewDeviceUpdateRequestDataWithDefaults instantiates a new DeviceUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceUpdateRequestDataWithDefaults() *DeviceUpdateRequestData {
	this := DeviceUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *DeviceUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeviceUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeviceUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *DeviceUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DeviceUpdateRequestData) GetAttributes() DeviceUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret DeviceUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUpdateRequestData) GetAttributesOk() (*DeviceUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DeviceUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given DeviceUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *DeviceUpdateRequestData) SetAttributes(v DeviceUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o DeviceUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableDeviceUpdateRequestData struct {
	value *DeviceUpdateRequestData
	isSet bool
}

func (v NullableDeviceUpdateRequestData) Get() *DeviceUpdateRequestData {
	return v.value
}

func (v *NullableDeviceUpdateRequestData) Set(val *DeviceUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceUpdateRequestData(val *DeviceUpdateRequestData) *NullableDeviceUpdateRequestData {
	return &NullableDeviceUpdateRequestData{value: val, isSet: true}
}

func (v NullableDeviceUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


