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

// checks if the ProfileRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileRelationships{}

// ProfileRelationships struct for ProfileRelationships
type ProfileRelationships struct {
	BundleId *CiProductRelationshipsBundleId `json:"bundleId,omitempty"`
	Devices *ProfileRelationshipsDevices `json:"devices,omitempty"`
	Certificates *ProfileRelationshipsCertificates `json:"certificates,omitempty"`
}

// NewProfileRelationships instantiates a new ProfileRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileRelationships() *ProfileRelationships {
	this := ProfileRelationships{}
	return &this
}

// NewProfileRelationshipsWithDefaults instantiates a new ProfileRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileRelationshipsWithDefaults() *ProfileRelationships {
	this := ProfileRelationships{}
	return &this
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *ProfileRelationships) GetBundleId() CiProductRelationshipsBundleId {
	if o == nil || IsNil(o.BundleId) {
		var ret CiProductRelationshipsBundleId
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileRelationships) GetBundleIdOk() (*CiProductRelationshipsBundleId, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *ProfileRelationships) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given CiProductRelationshipsBundleId and assigns it to the BundleId field.
func (o *ProfileRelationships) SetBundleId(v CiProductRelationshipsBundleId) {
	o.BundleId = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *ProfileRelationships) GetDevices() ProfileRelationshipsDevices {
	if o == nil || IsNil(o.Devices) {
		var ret ProfileRelationshipsDevices
		return ret
	}
	return *o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileRelationships) GetDevicesOk() (*ProfileRelationshipsDevices, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *ProfileRelationships) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given ProfileRelationshipsDevices and assigns it to the Devices field.
func (o *ProfileRelationships) SetDevices(v ProfileRelationshipsDevices) {
	o.Devices = &v
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *ProfileRelationships) GetCertificates() ProfileRelationshipsCertificates {
	if o == nil || IsNil(o.Certificates) {
		var ret ProfileRelationshipsCertificates
		return ret
	}
	return *o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileRelationships) GetCertificatesOk() (*ProfileRelationshipsCertificates, bool) {
	if o == nil || IsNil(o.Certificates) {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *ProfileRelationships) HasCertificates() bool {
	if o != nil && !IsNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given ProfileRelationshipsCertificates and assigns it to the Certificates field.
func (o *ProfileRelationships) SetCertificates(v ProfileRelationshipsCertificates) {
	o.Certificates = &v
}

func (o ProfileRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BundleId) {
		toSerialize["bundleId"] = o.BundleId
	}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !IsNil(o.Certificates) {
		toSerialize["certificates"] = o.Certificates
	}
	return toSerialize, nil
}

type NullableProfileRelationships struct {
	value *ProfileRelationships
	isSet bool
}

func (v NullableProfileRelationships) Get() *ProfileRelationships {
	return v.value
}

func (v *NullableProfileRelationships) Set(val *ProfileRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileRelationships(val *ProfileRelationships) *NullableProfileRelationships {
	return &NullableProfileRelationships{value: val, isSet: true}
}

func (v NullableProfileRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


