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

// checks if the AppClipHeaderImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipHeaderImage{}

// AppClipHeaderImage struct for AppClipHeaderImage
type AppClipHeaderImage struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppClipAdvancedExperienceImageAttributes `json:"attributes,omitempty"`
	Relationships *AppClipHeaderImageRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewAppClipHeaderImage instantiates a new AppClipHeaderImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipHeaderImage(type_ string, id string, links ResourceLinks) *AppClipHeaderImage {
	this := AppClipHeaderImage{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewAppClipHeaderImageWithDefaults instantiates a new AppClipHeaderImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipHeaderImageWithDefaults() *AppClipHeaderImage {
	this := AppClipHeaderImage{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipHeaderImage) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipHeaderImage) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipHeaderImage) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppClipHeaderImage) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppClipHeaderImage) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppClipHeaderImage) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppClipHeaderImage) GetAttributes() AppClipAdvancedExperienceImageAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipAdvancedExperienceImageAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipHeaderImage) GetAttributesOk() (*AppClipAdvancedExperienceImageAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppClipHeaderImage) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipAdvancedExperienceImageAttributes and assigns it to the Attributes field.
func (o *AppClipHeaderImage) SetAttributes(v AppClipAdvancedExperienceImageAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppClipHeaderImage) GetRelationships() AppClipHeaderImageRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppClipHeaderImageRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipHeaderImage) GetRelationshipsOk() (*AppClipHeaderImageRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppClipHeaderImage) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppClipHeaderImageRelationships and assigns it to the Relationships field.
func (o *AppClipHeaderImage) SetRelationships(v AppClipHeaderImageRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *AppClipHeaderImage) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppClipHeaderImage) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppClipHeaderImage) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o AppClipHeaderImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipHeaderImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppClipHeaderImage struct {
	value *AppClipHeaderImage
	isSet bool
}

func (v NullableAppClipHeaderImage) Get() *AppClipHeaderImage {
	return v.value
}

func (v *NullableAppClipHeaderImage) Set(val *AppClipHeaderImage) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipHeaderImage) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipHeaderImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipHeaderImage(val *AppClipHeaderImage) *NullableAppClipHeaderImage {
	return &NullableAppClipHeaderImage{value: val, isSet: true}
}

func (v NullableAppClipHeaderImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipHeaderImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


