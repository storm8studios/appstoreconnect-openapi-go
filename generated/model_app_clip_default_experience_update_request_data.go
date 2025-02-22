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

// checks if the AppClipDefaultExperienceUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceUpdateRequestData{}

// AppClipDefaultExperienceUpdateRequestData struct for AppClipDefaultExperienceUpdateRequestData
type AppClipDefaultExperienceUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppClipDefaultExperienceAttributes `json:"attributes,omitempty"`
	Relationships *AppClipDefaultExperienceUpdateRequestDataRelationships `json:"relationships,omitempty"`
}

// NewAppClipDefaultExperienceUpdateRequestData instantiates a new AppClipDefaultExperienceUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceUpdateRequestData(type_ string, id string) *AppClipDefaultExperienceUpdateRequestData {
	this := AppClipDefaultExperienceUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppClipDefaultExperienceUpdateRequestDataWithDefaults instantiates a new AppClipDefaultExperienceUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceUpdateRequestDataWithDefaults() *AppClipDefaultExperienceUpdateRequestData {
	this := AppClipDefaultExperienceUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipDefaultExperienceUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipDefaultExperienceUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppClipDefaultExperienceUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppClipDefaultExperienceUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceUpdateRequestData) GetAttributes() AppClipDefaultExperienceAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipDefaultExperienceAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceUpdateRequestData) GetAttributesOk() (*AppClipDefaultExperienceAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipDefaultExperienceAttributes and assigns it to the Attributes field.
func (o *AppClipDefaultExperienceUpdateRequestData) SetAttributes(v AppClipDefaultExperienceAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceUpdateRequestData) GetRelationships() AppClipDefaultExperienceUpdateRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppClipDefaultExperienceUpdateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceUpdateRequestData) GetRelationshipsOk() (*AppClipDefaultExperienceUpdateRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceUpdateRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppClipDefaultExperienceUpdateRequestDataRelationships and assigns it to the Relationships field.
func (o *AppClipDefaultExperienceUpdateRequestData) SetRelationships(v AppClipDefaultExperienceUpdateRequestDataRelationships) {
	o.Relationships = &v
}

func (o AppClipDefaultExperienceUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableAppClipDefaultExperienceUpdateRequestData struct {
	value *AppClipDefaultExperienceUpdateRequestData
	isSet bool
}

func (v NullableAppClipDefaultExperienceUpdateRequestData) Get() *AppClipDefaultExperienceUpdateRequestData {
	return v.value
}

func (v *NullableAppClipDefaultExperienceUpdateRequestData) Set(val *AppClipDefaultExperienceUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceUpdateRequestData(val *AppClipDefaultExperienceUpdateRequestData) *NullableAppClipDefaultExperienceUpdateRequestData {
	return &NullableAppClipDefaultExperienceUpdateRequestData{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


