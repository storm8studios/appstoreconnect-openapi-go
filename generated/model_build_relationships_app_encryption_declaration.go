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

// checks if the BuildRelationshipsAppEncryptionDeclaration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildRelationshipsAppEncryptionDeclaration{}

// BuildRelationshipsAppEncryptionDeclaration struct for BuildRelationshipsAppEncryptionDeclaration
type BuildRelationshipsAppEncryptionDeclaration struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Data *AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData `json:"data,omitempty"`
}

// NewBuildRelationshipsAppEncryptionDeclaration instantiates a new BuildRelationshipsAppEncryptionDeclaration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildRelationshipsAppEncryptionDeclaration() *BuildRelationshipsAppEncryptionDeclaration {
	this := BuildRelationshipsAppEncryptionDeclaration{}
	return &this
}

// NewBuildRelationshipsAppEncryptionDeclarationWithDefaults instantiates a new BuildRelationshipsAppEncryptionDeclaration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildRelationshipsAppEncryptionDeclarationWithDefaults() *BuildRelationshipsAppEncryptionDeclaration {
	this := BuildRelationshipsAppEncryptionDeclaration{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BuildRelationshipsAppEncryptionDeclaration) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationshipsAppEncryptionDeclaration) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BuildRelationshipsAppEncryptionDeclaration) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *BuildRelationshipsAppEncryptionDeclaration) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BuildRelationshipsAppEncryptionDeclaration) GetData() AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData {
	if o == nil || IsNil(o.Data) {
		var ret AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationshipsAppEncryptionDeclaration) GetDataOk() (*AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BuildRelationshipsAppEncryptionDeclaration) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData and assigns it to the Data field.
func (o *BuildRelationshipsAppEncryptionDeclaration) SetData(v AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData) {
	o.Data = &v
}

func (o BuildRelationshipsAppEncryptionDeclaration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildRelationshipsAppEncryptionDeclaration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBuildRelationshipsAppEncryptionDeclaration struct {
	value *BuildRelationshipsAppEncryptionDeclaration
	isSet bool
}

func (v NullableBuildRelationshipsAppEncryptionDeclaration) Get() *BuildRelationshipsAppEncryptionDeclaration {
	return v.value
}

func (v *NullableBuildRelationshipsAppEncryptionDeclaration) Set(val *BuildRelationshipsAppEncryptionDeclaration) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildRelationshipsAppEncryptionDeclaration) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildRelationshipsAppEncryptionDeclaration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildRelationshipsAppEncryptionDeclaration(val *BuildRelationshipsAppEncryptionDeclaration) *NullableBuildRelationshipsAppEncryptionDeclaration {
	return &NullableBuildRelationshipsAppEncryptionDeclaration{value: val, isSet: true}
}

func (v NullableBuildRelationshipsAppEncryptionDeclaration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildRelationshipsAppEncryptionDeclaration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


