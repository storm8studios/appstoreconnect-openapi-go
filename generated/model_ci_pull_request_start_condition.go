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

// checks if the CiPullRequestStartCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiPullRequestStartCondition{}

// CiPullRequestStartCondition struct for CiPullRequestStartCondition
type CiPullRequestStartCondition struct {
	Source *CiBranchPatterns `json:"source,omitempty"`
	Destination *CiBranchPatterns `json:"destination,omitempty"`
	FilesAndFoldersRule *CiFilesAndFoldersRule `json:"filesAndFoldersRule,omitempty"`
	AutoCancel *bool `json:"autoCancel,omitempty"`
}

// NewCiPullRequestStartCondition instantiates a new CiPullRequestStartCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiPullRequestStartCondition() *CiPullRequestStartCondition {
	this := CiPullRequestStartCondition{}
	return &this
}

// NewCiPullRequestStartConditionWithDefaults instantiates a new CiPullRequestStartCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiPullRequestStartConditionWithDefaults() *CiPullRequestStartCondition {
	this := CiPullRequestStartCondition{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CiPullRequestStartCondition) GetSource() CiBranchPatterns {
	if o == nil || IsNil(o.Source) {
		var ret CiBranchPatterns
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiPullRequestStartCondition) GetSourceOk() (*CiBranchPatterns, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CiPullRequestStartCondition) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given CiBranchPatterns and assigns it to the Source field.
func (o *CiPullRequestStartCondition) SetSource(v CiBranchPatterns) {
	o.Source = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *CiPullRequestStartCondition) GetDestination() CiBranchPatterns {
	if o == nil || IsNil(o.Destination) {
		var ret CiBranchPatterns
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiPullRequestStartCondition) GetDestinationOk() (*CiBranchPatterns, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *CiPullRequestStartCondition) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given CiBranchPatterns and assigns it to the Destination field.
func (o *CiPullRequestStartCondition) SetDestination(v CiBranchPatterns) {
	o.Destination = &v
}

// GetFilesAndFoldersRule returns the FilesAndFoldersRule field value if set, zero value otherwise.
func (o *CiPullRequestStartCondition) GetFilesAndFoldersRule() CiFilesAndFoldersRule {
	if o == nil || IsNil(o.FilesAndFoldersRule) {
		var ret CiFilesAndFoldersRule
		return ret
	}
	return *o.FilesAndFoldersRule
}

// GetFilesAndFoldersRuleOk returns a tuple with the FilesAndFoldersRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiPullRequestStartCondition) GetFilesAndFoldersRuleOk() (*CiFilesAndFoldersRule, bool) {
	if o == nil || IsNil(o.FilesAndFoldersRule) {
		return nil, false
	}
	return o.FilesAndFoldersRule, true
}

// HasFilesAndFoldersRule returns a boolean if a field has been set.
func (o *CiPullRequestStartCondition) HasFilesAndFoldersRule() bool {
	if o != nil && !IsNil(o.FilesAndFoldersRule) {
		return true
	}

	return false
}

// SetFilesAndFoldersRule gets a reference to the given CiFilesAndFoldersRule and assigns it to the FilesAndFoldersRule field.
func (o *CiPullRequestStartCondition) SetFilesAndFoldersRule(v CiFilesAndFoldersRule) {
	o.FilesAndFoldersRule = &v
}

// GetAutoCancel returns the AutoCancel field value if set, zero value otherwise.
func (o *CiPullRequestStartCondition) GetAutoCancel() bool {
	if o == nil || IsNil(o.AutoCancel) {
		var ret bool
		return ret
	}
	return *o.AutoCancel
}

// GetAutoCancelOk returns a tuple with the AutoCancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiPullRequestStartCondition) GetAutoCancelOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoCancel) {
		return nil, false
	}
	return o.AutoCancel, true
}

// HasAutoCancel returns a boolean if a field has been set.
func (o *CiPullRequestStartCondition) HasAutoCancel() bool {
	if o != nil && !IsNil(o.AutoCancel) {
		return true
	}

	return false
}

// SetAutoCancel gets a reference to the given bool and assigns it to the AutoCancel field.
func (o *CiPullRequestStartCondition) SetAutoCancel(v bool) {
	o.AutoCancel = &v
}

func (o CiPullRequestStartCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiPullRequestStartCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.FilesAndFoldersRule) {
		toSerialize["filesAndFoldersRule"] = o.FilesAndFoldersRule
	}
	if !IsNil(o.AutoCancel) {
		toSerialize["autoCancel"] = o.AutoCancel
	}
	return toSerialize, nil
}

type NullableCiPullRequestStartCondition struct {
	value *CiPullRequestStartCondition
	isSet bool
}

func (v NullableCiPullRequestStartCondition) Get() *CiPullRequestStartCondition {
	return v.value
}

func (v *NullableCiPullRequestStartCondition) Set(val *CiPullRequestStartCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableCiPullRequestStartCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableCiPullRequestStartCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiPullRequestStartCondition(val *CiPullRequestStartCondition) *NullableCiPullRequestStartCondition {
	return &NullableCiPullRequestStartCondition{value: val, isSet: true}
}

func (v NullableCiPullRequestStartCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiPullRequestStartCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


