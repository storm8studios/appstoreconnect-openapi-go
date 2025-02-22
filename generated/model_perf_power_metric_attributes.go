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

// checks if the PerfPowerMetricAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerfPowerMetricAttributes{}

// PerfPowerMetricAttributes struct for PerfPowerMetricAttributes
type PerfPowerMetricAttributes struct {
	Platform *string `json:"platform,omitempty"`
	MetricType *string `json:"metricType,omitempty"`
	DeviceType *string `json:"deviceType,omitempty"`
}

// NewPerfPowerMetricAttributes instantiates a new PerfPowerMetricAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerfPowerMetricAttributes() *PerfPowerMetricAttributes {
	this := PerfPowerMetricAttributes{}
	return &this
}

// NewPerfPowerMetricAttributesWithDefaults instantiates a new PerfPowerMetricAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerfPowerMetricAttributesWithDefaults() *PerfPowerMetricAttributes {
	this := PerfPowerMetricAttributes{}
	return &this
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *PerfPowerMetricAttributes) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfPowerMetricAttributes) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *PerfPowerMetricAttributes) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *PerfPowerMetricAttributes) SetPlatform(v string) {
	o.Platform = &v
}

// GetMetricType returns the MetricType field value if set, zero value otherwise.
func (o *PerfPowerMetricAttributes) GetMetricType() string {
	if o == nil || IsNil(o.MetricType) {
		var ret string
		return ret
	}
	return *o.MetricType
}

// GetMetricTypeOk returns a tuple with the MetricType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfPowerMetricAttributes) GetMetricTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MetricType) {
		return nil, false
	}
	return o.MetricType, true
}

// HasMetricType returns a boolean if a field has been set.
func (o *PerfPowerMetricAttributes) HasMetricType() bool {
	if o != nil && !IsNil(o.MetricType) {
		return true
	}

	return false
}

// SetMetricType gets a reference to the given string and assigns it to the MetricType field.
func (o *PerfPowerMetricAttributes) SetMetricType(v string) {
	o.MetricType = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *PerfPowerMetricAttributes) GetDeviceType() string {
	if o == nil || IsNil(o.DeviceType) {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfPowerMetricAttributes) GetDeviceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *PerfPowerMetricAttributes) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *PerfPowerMetricAttributes) SetDeviceType(v string) {
	o.DeviceType = &v
}

func (o PerfPowerMetricAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerfPowerMetricAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.MetricType) {
		toSerialize["metricType"] = o.MetricType
	}
	if !IsNil(o.DeviceType) {
		toSerialize["deviceType"] = o.DeviceType
	}
	return toSerialize, nil
}

type NullablePerfPowerMetricAttributes struct {
	value *PerfPowerMetricAttributes
	isSet bool
}

func (v NullablePerfPowerMetricAttributes) Get() *PerfPowerMetricAttributes {
	return v.value
}

func (v *NullablePerfPowerMetricAttributes) Set(val *PerfPowerMetricAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePerfPowerMetricAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePerfPowerMetricAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerfPowerMetricAttributes(val *PerfPowerMetricAttributes) *NullablePerfPowerMetricAttributes {
	return &NullablePerfPowerMetricAttributes{value: val, isSet: true}
}

func (v NullablePerfPowerMetricAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerfPowerMetricAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


