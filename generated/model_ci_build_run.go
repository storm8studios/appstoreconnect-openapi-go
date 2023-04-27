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

// checks if the CiBuildRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBuildRun{}

// CiBuildRun struct for CiBuildRun
type CiBuildRun struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *CiBuildRunAttributes `json:"attributes,omitempty"`
	Relationships *CiBuildRunRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewCiBuildRun instantiates a new CiBuildRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBuildRun(type_ string, id string, links ResourceLinks) *CiBuildRun {
	this := CiBuildRun{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewCiBuildRunWithDefaults instantiates a new CiBuildRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBuildRunWithDefaults() *CiBuildRun {
	this := CiBuildRun{}
	return &this
}

// GetType returns the Type field value
func (o *CiBuildRun) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CiBuildRun) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CiBuildRun) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CiBuildRun) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CiBuildRun) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CiBuildRun) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CiBuildRun) GetAttributes() CiBuildRunAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret CiBuildRunAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRun) GetAttributesOk() (*CiBuildRunAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CiBuildRun) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CiBuildRunAttributes and assigns it to the Attributes field.
func (o *CiBuildRun) SetAttributes(v CiBuildRunAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CiBuildRun) GetRelationships() CiBuildRunRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CiBuildRunRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRun) GetRelationshipsOk() (*CiBuildRunRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CiBuildRun) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CiBuildRunRelationships and assigns it to the Relationships field.
func (o *CiBuildRun) SetRelationships(v CiBuildRunRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *CiBuildRun) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CiBuildRun) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CiBuildRun) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o CiBuildRun) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBuildRun) ToMap() (map[string]interface{}, error) {
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

type NullableCiBuildRun struct {
	value *CiBuildRun
	isSet bool
}

func (v NullableCiBuildRun) Get() *CiBuildRun {
	return v.value
}

func (v *NullableCiBuildRun) Set(val *CiBuildRun) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBuildRun) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBuildRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBuildRun(val *CiBuildRun) *NullableCiBuildRun {
	return &NullableCiBuildRun{value: val, isSet: true}
}

func (v NullableCiBuildRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBuildRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


