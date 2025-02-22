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

// checks if the BetaAppClipInvocationLocalizationInlineCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocationLocalizationInlineCreate{}

// BetaAppClipInvocationLocalizationInlineCreate struct for BetaAppClipInvocationLocalizationInlineCreate
type BetaAppClipInvocationLocalizationInlineCreate struct {
	Type string `json:"type"`
	Id *string `json:"id,omitempty"`
	Attributes BetaAppClipInvocationLocalizationInlineCreateAttributes `json:"attributes"`
	Relationships *BetaAppClipInvocationLocalizationInlineCreateRelationships `json:"relationships,omitempty"`
}

// NewBetaAppClipInvocationLocalizationInlineCreate instantiates a new BetaAppClipInvocationLocalizationInlineCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocationLocalizationInlineCreate(type_ string, attributes BetaAppClipInvocationLocalizationInlineCreateAttributes) *BetaAppClipInvocationLocalizationInlineCreate {
	this := BetaAppClipInvocationLocalizationInlineCreate{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewBetaAppClipInvocationLocalizationInlineCreateWithDefaults instantiates a new BetaAppClipInvocationLocalizationInlineCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationLocalizationInlineCreateWithDefaults() *BetaAppClipInvocationLocalizationInlineCreate {
	this := BetaAppClipInvocationLocalizationInlineCreate{}
	return &this
}

// GetType returns the Type field value
func (o *BetaAppClipInvocationLocalizationInlineCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationLocalizationInlineCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BetaAppClipInvocationLocalizationInlineCreate) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BetaAppClipInvocationLocalizationInlineCreate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationLocalizationInlineCreate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BetaAppClipInvocationLocalizationInlineCreate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BetaAppClipInvocationLocalizationInlineCreate) SetId(v string) {
	o.Id = &v
}

// GetAttributes returns the Attributes field value
func (o *BetaAppClipInvocationLocalizationInlineCreate) GetAttributes() BetaAppClipInvocationLocalizationInlineCreateAttributes {
	if o == nil {
		var ret BetaAppClipInvocationLocalizationInlineCreateAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationLocalizationInlineCreate) GetAttributesOk() (*BetaAppClipInvocationLocalizationInlineCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BetaAppClipInvocationLocalizationInlineCreate) SetAttributes(v BetaAppClipInvocationLocalizationInlineCreateAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BetaAppClipInvocationLocalizationInlineCreate) GetRelationships() BetaAppClipInvocationLocalizationInlineCreateRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BetaAppClipInvocationLocalizationInlineCreateRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationLocalizationInlineCreate) GetRelationshipsOk() (*BetaAppClipInvocationLocalizationInlineCreateRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BetaAppClipInvocationLocalizationInlineCreate) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BetaAppClipInvocationLocalizationInlineCreateRelationships and assigns it to the Relationships field.
func (o *BetaAppClipInvocationLocalizationInlineCreate) SetRelationships(v BetaAppClipInvocationLocalizationInlineCreateRelationships) {
	o.Relationships = &v
}

func (o BetaAppClipInvocationLocalizationInlineCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocationLocalizationInlineCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableBetaAppClipInvocationLocalizationInlineCreate struct {
	value *BetaAppClipInvocationLocalizationInlineCreate
	isSet bool
}

func (v NullableBetaAppClipInvocationLocalizationInlineCreate) Get() *BetaAppClipInvocationLocalizationInlineCreate {
	return v.value
}

func (v *NullableBetaAppClipInvocationLocalizationInlineCreate) Set(val *BetaAppClipInvocationLocalizationInlineCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocationLocalizationInlineCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocationLocalizationInlineCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocationLocalizationInlineCreate(val *BetaAppClipInvocationLocalizationInlineCreate) *NullableBetaAppClipInvocationLocalizationInlineCreate {
	return &NullableBetaAppClipInvocationLocalizationInlineCreate{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocationLocalizationInlineCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocationLocalizationInlineCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


