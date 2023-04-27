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

// checks if the AppAvailability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAvailability{}

// AppAvailability struct for AppAvailability
type AppAvailability struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppAvailabilityAttributes `json:"attributes,omitempty"`
	Relationships *AppAvailabilityRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewAppAvailability instantiates a new AppAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAvailability(type_ string, id string, links ResourceLinks) *AppAvailability {
	this := AppAvailability{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewAppAvailabilityWithDefaults instantiates a new AppAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAvailabilityWithDefaults() *AppAvailability {
	this := AppAvailability{}
	return &this
}

// GetType returns the Type field value
func (o *AppAvailability) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppAvailability) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppAvailability) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppAvailability) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppAvailability) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppAvailability) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppAvailability) GetAttributes() AppAvailabilityAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppAvailabilityAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAvailability) GetAttributesOk() (*AppAvailabilityAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppAvailability) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppAvailabilityAttributes and assigns it to the Attributes field.
func (o *AppAvailability) SetAttributes(v AppAvailabilityAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppAvailability) GetRelationships() AppAvailabilityRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppAvailabilityRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAvailability) GetRelationshipsOk() (*AppAvailabilityRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppAvailability) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppAvailabilityRelationships and assigns it to the Relationships field.
func (o *AppAvailability) SetRelationships(v AppAvailabilityRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *AppAvailability) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppAvailability) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppAvailability) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o AppAvailability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAvailability) ToMap() (map[string]interface{}, error) {
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

type NullableAppAvailability struct {
	value *AppAvailability
	isSet bool
}

func (v NullableAppAvailability) Get() *AppAvailability {
	return v.value
}

func (v *NullableAppAvailability) Set(val *AppAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAvailability(val *AppAvailability) *NullableAppAvailability {
	return &NullableAppAvailability{value: val, isSet: true}
}

func (v NullableAppAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


