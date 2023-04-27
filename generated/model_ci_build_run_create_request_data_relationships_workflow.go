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

// checks if the CiBuildRunCreateRequestDataRelationshipsWorkflow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBuildRunCreateRequestDataRelationshipsWorkflow{}

// CiBuildRunCreateRequestDataRelationshipsWorkflow struct for CiBuildRunCreateRequestDataRelationshipsWorkflow
type CiBuildRunCreateRequestDataRelationshipsWorkflow struct {
	Data *CiBuildRunRelationshipsWorkflowData `json:"data,omitempty"`
}

// NewCiBuildRunCreateRequestDataRelationshipsWorkflow instantiates a new CiBuildRunCreateRequestDataRelationshipsWorkflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBuildRunCreateRequestDataRelationshipsWorkflow() *CiBuildRunCreateRequestDataRelationshipsWorkflow {
	this := CiBuildRunCreateRequestDataRelationshipsWorkflow{}
	return &this
}

// NewCiBuildRunCreateRequestDataRelationshipsWorkflowWithDefaults instantiates a new CiBuildRunCreateRequestDataRelationshipsWorkflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBuildRunCreateRequestDataRelationshipsWorkflowWithDefaults() *CiBuildRunCreateRequestDataRelationshipsWorkflow {
	this := CiBuildRunCreateRequestDataRelationshipsWorkflow{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CiBuildRunCreateRequestDataRelationshipsWorkflow) GetData() CiBuildRunRelationshipsWorkflowData {
	if o == nil || IsNil(o.Data) {
		var ret CiBuildRunRelationshipsWorkflowData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunCreateRequestDataRelationshipsWorkflow) GetDataOk() (*CiBuildRunRelationshipsWorkflowData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CiBuildRunCreateRequestDataRelationshipsWorkflow) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CiBuildRunRelationshipsWorkflowData and assigns it to the Data field.
func (o *CiBuildRunCreateRequestDataRelationshipsWorkflow) SetData(v CiBuildRunRelationshipsWorkflowData) {
	o.Data = &v
}

func (o CiBuildRunCreateRequestDataRelationshipsWorkflow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBuildRunCreateRequestDataRelationshipsWorkflow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCiBuildRunCreateRequestDataRelationshipsWorkflow struct {
	value *CiBuildRunCreateRequestDataRelationshipsWorkflow
	isSet bool
}

func (v NullableCiBuildRunCreateRequestDataRelationshipsWorkflow) Get() *CiBuildRunCreateRequestDataRelationshipsWorkflow {
	return v.value
}

func (v *NullableCiBuildRunCreateRequestDataRelationshipsWorkflow) Set(val *CiBuildRunCreateRequestDataRelationshipsWorkflow) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBuildRunCreateRequestDataRelationshipsWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBuildRunCreateRequestDataRelationshipsWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBuildRunCreateRequestDataRelationshipsWorkflow(val *CiBuildRunCreateRequestDataRelationshipsWorkflow) *NullableCiBuildRunCreateRequestDataRelationshipsWorkflow {
	return &NullableCiBuildRunCreateRequestDataRelationshipsWorkflow{value: val, isSet: true}
}

func (v NullableCiBuildRunCreateRequestDataRelationshipsWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBuildRunCreateRequestDataRelationshipsWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


