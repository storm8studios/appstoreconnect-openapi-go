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

// checks if the EndUserLicenseAgreementUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndUserLicenseAgreementUpdateRequestData{}

// EndUserLicenseAgreementUpdateRequestData struct for EndUserLicenseAgreementUpdateRequestData
type EndUserLicenseAgreementUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *BetaLicenseAgreementAttributes `json:"attributes,omitempty"`
	Relationships *EndUserLicenseAgreementUpdateRequestDataRelationships `json:"relationships,omitempty"`
}

// NewEndUserLicenseAgreementUpdateRequestData instantiates a new EndUserLicenseAgreementUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndUserLicenseAgreementUpdateRequestData(type_ string, id string) *EndUserLicenseAgreementUpdateRequestData {
	this := EndUserLicenseAgreementUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewEndUserLicenseAgreementUpdateRequestDataWithDefaults instantiates a new EndUserLicenseAgreementUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndUserLicenseAgreementUpdateRequestDataWithDefaults() *EndUserLicenseAgreementUpdateRequestData {
	this := EndUserLicenseAgreementUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *EndUserLicenseAgreementUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EndUserLicenseAgreementUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *EndUserLicenseAgreementUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EndUserLicenseAgreementUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EndUserLicenseAgreementUpdateRequestData) GetAttributes() BetaLicenseAgreementAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret BetaLicenseAgreementAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementUpdateRequestData) GetAttributesOk() (*BetaLicenseAgreementAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EndUserLicenseAgreementUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BetaLicenseAgreementAttributes and assigns it to the Attributes field.
func (o *EndUserLicenseAgreementUpdateRequestData) SetAttributes(v BetaLicenseAgreementAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *EndUserLicenseAgreementUpdateRequestData) GetRelationships() EndUserLicenseAgreementUpdateRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret EndUserLicenseAgreementUpdateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementUpdateRequestData) GetRelationshipsOk() (*EndUserLicenseAgreementUpdateRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *EndUserLicenseAgreementUpdateRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given EndUserLicenseAgreementUpdateRequestDataRelationships and assigns it to the Relationships field.
func (o *EndUserLicenseAgreementUpdateRequestData) SetRelationships(v EndUserLicenseAgreementUpdateRequestDataRelationships) {
	o.Relationships = &v
}

func (o EndUserLicenseAgreementUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndUserLicenseAgreementUpdateRequestData) ToMap() (map[string]interface{}, error) {
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

type NullableEndUserLicenseAgreementUpdateRequestData struct {
	value *EndUserLicenseAgreementUpdateRequestData
	isSet bool
}

func (v NullableEndUserLicenseAgreementUpdateRequestData) Get() *EndUserLicenseAgreementUpdateRequestData {
	return v.value
}

func (v *NullableEndUserLicenseAgreementUpdateRequestData) Set(val *EndUserLicenseAgreementUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableEndUserLicenseAgreementUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableEndUserLicenseAgreementUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndUserLicenseAgreementUpdateRequestData(val *EndUserLicenseAgreementUpdateRequestData) *NullableEndUserLicenseAgreementUpdateRequestData {
	return &NullableEndUserLicenseAgreementUpdateRequestData{value: val, isSet: true}
}

func (v NullableEndUserLicenseAgreementUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndUserLicenseAgreementUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


