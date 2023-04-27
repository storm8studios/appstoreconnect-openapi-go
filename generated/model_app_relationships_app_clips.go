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

// checks if the AppRelationshipsAppClips type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationshipsAppClips{}

// AppRelationshipsAppClips struct for AppRelationshipsAppClips
type AppRelationshipsAppClips struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []AppClipAdvancedExperienceRelationshipsAppClipData `json:"data,omitempty"`
}

// NewAppRelationshipsAppClips instantiates a new AppRelationshipsAppClips object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsAppClips() *AppRelationshipsAppClips {
	this := AppRelationshipsAppClips{}
	return &this
}

// NewAppRelationshipsAppClipsWithDefaults instantiates a new AppRelationshipsAppClips object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsAppClipsWithDefaults() *AppRelationshipsAppClips {
	this := AppRelationshipsAppClips{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppRelationshipsAppClips) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsAppClips) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppRelationshipsAppClips) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *AppRelationshipsAppClips) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppRelationshipsAppClips) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsAppClips) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppRelationshipsAppClips) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppRelationshipsAppClips) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppRelationshipsAppClips) GetData() []AppClipAdvancedExperienceRelationshipsAppClipData {
	if o == nil || IsNil(o.Data) {
		var ret []AppClipAdvancedExperienceRelationshipsAppClipData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsAppClips) GetDataOk() ([]AppClipAdvancedExperienceRelationshipsAppClipData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppRelationshipsAppClips) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppClipAdvancedExperienceRelationshipsAppClipData and assigns it to the Data field.
func (o *AppRelationshipsAppClips) SetData(v []AppClipAdvancedExperienceRelationshipsAppClipData) {
	o.Data = v
}

func (o AppRelationshipsAppClips) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationshipsAppClips) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppRelationshipsAppClips struct {
	value *AppRelationshipsAppClips
	isSet bool
}

func (v NullableAppRelationshipsAppClips) Get() *AppRelationshipsAppClips {
	return v.value
}

func (v *NullableAppRelationshipsAppClips) Set(val *AppRelationshipsAppClips) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsAppClips) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsAppClips) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsAppClips(val *AppRelationshipsAppClips) *NullableAppRelationshipsAppClips {
	return &NullableAppRelationshipsAppClips{value: val, isSet: true}
}

func (v NullableAppRelationshipsAppClips) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsAppClips) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


