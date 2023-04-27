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

// checks if the ScmGitReferenceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScmGitReferenceAttributes{}

// ScmGitReferenceAttributes struct for ScmGitReferenceAttributes
type ScmGitReferenceAttributes struct {
	Name *string `json:"name,omitempty"`
	CanonicalName *string `json:"canonicalName,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
	Kind *CiGitRefKind `json:"kind,omitempty"`
}

// NewScmGitReferenceAttributes instantiates a new ScmGitReferenceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScmGitReferenceAttributes() *ScmGitReferenceAttributes {
	this := ScmGitReferenceAttributes{}
	return &this
}

// NewScmGitReferenceAttributesWithDefaults instantiates a new ScmGitReferenceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScmGitReferenceAttributesWithDefaults() *ScmGitReferenceAttributes {
	this := ScmGitReferenceAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScmGitReferenceAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmGitReferenceAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScmGitReferenceAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScmGitReferenceAttributes) SetName(v string) {
	o.Name = &v
}

// GetCanonicalName returns the CanonicalName field value if set, zero value otherwise.
func (o *ScmGitReferenceAttributes) GetCanonicalName() string {
	if o == nil || IsNil(o.CanonicalName) {
		var ret string
		return ret
	}
	return *o.CanonicalName
}

// GetCanonicalNameOk returns a tuple with the CanonicalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmGitReferenceAttributes) GetCanonicalNameOk() (*string, bool) {
	if o == nil || IsNil(o.CanonicalName) {
		return nil, false
	}
	return o.CanonicalName, true
}

// HasCanonicalName returns a boolean if a field has been set.
func (o *ScmGitReferenceAttributes) HasCanonicalName() bool {
	if o != nil && !IsNil(o.CanonicalName) {
		return true
	}

	return false
}

// SetCanonicalName gets a reference to the given string and assigns it to the CanonicalName field.
func (o *ScmGitReferenceAttributes) SetCanonicalName(v string) {
	o.CanonicalName = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *ScmGitReferenceAttributes) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmGitReferenceAttributes) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *ScmGitReferenceAttributes) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *ScmGitReferenceAttributes) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ScmGitReferenceAttributes) GetKind() CiGitRefKind {
	if o == nil || IsNil(o.Kind) {
		var ret CiGitRefKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmGitReferenceAttributes) GetKindOk() (*CiGitRefKind, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ScmGitReferenceAttributes) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given CiGitRefKind and assigns it to the Kind field.
func (o *ScmGitReferenceAttributes) SetKind(v CiGitRefKind) {
	o.Kind = &v
}

func (o ScmGitReferenceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScmGitReferenceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CanonicalName) {
		toSerialize["canonicalName"] = o.CanonicalName
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	return toSerialize, nil
}

type NullableScmGitReferenceAttributes struct {
	value *ScmGitReferenceAttributes
	isSet bool
}

func (v NullableScmGitReferenceAttributes) Get() *ScmGitReferenceAttributes {
	return v.value
}

func (v *NullableScmGitReferenceAttributes) Set(val *ScmGitReferenceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableScmGitReferenceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableScmGitReferenceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmGitReferenceAttributes(val *ScmGitReferenceAttributes) *NullableScmGitReferenceAttributes {
	return &NullableScmGitReferenceAttributes{value: val, isSet: true}
}

func (v NullableScmGitReferenceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmGitReferenceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


