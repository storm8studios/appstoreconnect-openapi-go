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

// checks if the CiTagStartCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiTagStartCondition{}

// CiTagStartCondition struct for CiTagStartCondition
type CiTagStartCondition struct {
	Source *CiTagPatterns `json:"source,omitempty"`
	FilesAndFoldersRule *CiFilesAndFoldersRule `json:"filesAndFoldersRule,omitempty"`
	AutoCancel *bool `json:"autoCancel,omitempty"`
}

// NewCiTagStartCondition instantiates a new CiTagStartCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiTagStartCondition() *CiTagStartCondition {
	this := CiTagStartCondition{}
	return &this
}

// NewCiTagStartConditionWithDefaults instantiates a new CiTagStartCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiTagStartConditionWithDefaults() *CiTagStartCondition {
	this := CiTagStartCondition{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CiTagStartCondition) GetSource() CiTagPatterns {
	if o == nil || IsNil(o.Source) {
		var ret CiTagPatterns
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiTagStartCondition) GetSourceOk() (*CiTagPatterns, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CiTagStartCondition) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given CiTagPatterns and assigns it to the Source field.
func (o *CiTagStartCondition) SetSource(v CiTagPatterns) {
	o.Source = &v
}

// GetFilesAndFoldersRule returns the FilesAndFoldersRule field value if set, zero value otherwise.
func (o *CiTagStartCondition) GetFilesAndFoldersRule() CiFilesAndFoldersRule {
	if o == nil || IsNil(o.FilesAndFoldersRule) {
		var ret CiFilesAndFoldersRule
		return ret
	}
	return *o.FilesAndFoldersRule
}

// GetFilesAndFoldersRuleOk returns a tuple with the FilesAndFoldersRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiTagStartCondition) GetFilesAndFoldersRuleOk() (*CiFilesAndFoldersRule, bool) {
	if o == nil || IsNil(o.FilesAndFoldersRule) {
		return nil, false
	}
	return o.FilesAndFoldersRule, true
}

// HasFilesAndFoldersRule returns a boolean if a field has been set.
func (o *CiTagStartCondition) HasFilesAndFoldersRule() bool {
	if o != nil && !IsNil(o.FilesAndFoldersRule) {
		return true
	}

	return false
}

// SetFilesAndFoldersRule gets a reference to the given CiFilesAndFoldersRule and assigns it to the FilesAndFoldersRule field.
func (o *CiTagStartCondition) SetFilesAndFoldersRule(v CiFilesAndFoldersRule) {
	o.FilesAndFoldersRule = &v
}

// GetAutoCancel returns the AutoCancel field value if set, zero value otherwise.
func (o *CiTagStartCondition) GetAutoCancel() bool {
	if o == nil || IsNil(o.AutoCancel) {
		var ret bool
		return ret
	}
	return *o.AutoCancel
}

// GetAutoCancelOk returns a tuple with the AutoCancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiTagStartCondition) GetAutoCancelOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoCancel) {
		return nil, false
	}
	return o.AutoCancel, true
}

// HasAutoCancel returns a boolean if a field has been set.
func (o *CiTagStartCondition) HasAutoCancel() bool {
	if o != nil && !IsNil(o.AutoCancel) {
		return true
	}

	return false
}

// SetAutoCancel gets a reference to the given bool and assigns it to the AutoCancel field.
func (o *CiTagStartCondition) SetAutoCancel(v bool) {
	o.AutoCancel = &v
}

func (o CiTagStartCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiTagStartCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.FilesAndFoldersRule) {
		toSerialize["filesAndFoldersRule"] = o.FilesAndFoldersRule
	}
	if !IsNil(o.AutoCancel) {
		toSerialize["autoCancel"] = o.AutoCancel
	}
	return toSerialize, nil
}

type NullableCiTagStartCondition struct {
	value *CiTagStartCondition
	isSet bool
}

func (v NullableCiTagStartCondition) Get() *CiTagStartCondition {
	return v.value
}

func (v *NullableCiTagStartCondition) Set(val *CiTagStartCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableCiTagStartCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableCiTagStartCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiTagStartCondition(val *CiTagStartCondition) *NullableCiTagStartCondition {
	return &NullableCiTagStartCondition{value: val, isSet: true}
}

func (v NullableCiTagStartCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiTagStartCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


