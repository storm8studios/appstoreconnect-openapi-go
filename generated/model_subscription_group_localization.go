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

// checks if the SubscriptionGroupLocalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGroupLocalization{}

// SubscriptionGroupLocalization struct for SubscriptionGroupLocalization
type SubscriptionGroupLocalization struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *SubscriptionGroupLocalizationAttributes `json:"attributes,omitempty"`
	Relationships *SubscriptionGroupLocalizationRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewSubscriptionGroupLocalization instantiates a new SubscriptionGroupLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGroupLocalization(type_ string, id string, links ResourceLinks) *SubscriptionGroupLocalization {
	this := SubscriptionGroupLocalization{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewSubscriptionGroupLocalizationWithDefaults instantiates a new SubscriptionGroupLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGroupLocalizationWithDefaults() *SubscriptionGroupLocalization {
	this := SubscriptionGroupLocalization{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionGroupLocalization) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalization) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionGroupLocalization) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionGroupLocalization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalization) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionGroupLocalization) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubscriptionGroupLocalization) GetAttributes() SubscriptionGroupLocalizationAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SubscriptionGroupLocalizationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalization) GetAttributesOk() (*SubscriptionGroupLocalizationAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubscriptionGroupLocalization) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubscriptionGroupLocalizationAttributes and assigns it to the Attributes field.
func (o *SubscriptionGroupLocalization) SetAttributes(v SubscriptionGroupLocalizationAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SubscriptionGroupLocalization) GetRelationships() SubscriptionGroupLocalizationRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SubscriptionGroupLocalizationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalization) GetRelationshipsOk() (*SubscriptionGroupLocalizationRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SubscriptionGroupLocalization) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SubscriptionGroupLocalizationRelationships and assigns it to the Relationships field.
func (o *SubscriptionGroupLocalization) SetRelationships(v SubscriptionGroupLocalizationRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *SubscriptionGroupLocalization) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalization) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionGroupLocalization) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o SubscriptionGroupLocalization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGroupLocalization) ToMap() (map[string]interface{}, error) {
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

type NullableSubscriptionGroupLocalization struct {
	value *SubscriptionGroupLocalization
	isSet bool
}

func (v NullableSubscriptionGroupLocalization) Get() *SubscriptionGroupLocalization {
	return v.value
}

func (v *NullableSubscriptionGroupLocalization) Set(val *SubscriptionGroupLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGroupLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGroupLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGroupLocalization(val *SubscriptionGroupLocalization) *NullableSubscriptionGroupLocalization {
	return &NullableSubscriptionGroupLocalization{value: val, isSet: true}
}

func (v NullableSubscriptionGroupLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGroupLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


