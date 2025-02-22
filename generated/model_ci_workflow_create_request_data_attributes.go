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

// checks if the CiWorkflowCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiWorkflowCreateRequestDataAttributes{}

// CiWorkflowCreateRequestDataAttributes struct for CiWorkflowCreateRequestDataAttributes
type CiWorkflowCreateRequestDataAttributes struct {
	Name string `json:"name"`
	Description string `json:"description"`
	BranchStartCondition *CiBranchStartCondition `json:"branchStartCondition,omitempty"`
	TagStartCondition *CiTagStartCondition `json:"tagStartCondition,omitempty"`
	PullRequestStartCondition *CiPullRequestStartCondition `json:"pullRequestStartCondition,omitempty"`
	ScheduledStartCondition *CiScheduledStartCondition `json:"scheduledStartCondition,omitempty"`
	Actions []CiAction `json:"actions"`
	IsEnabled bool `json:"isEnabled"`
	IsLockedForEditing *bool `json:"isLockedForEditing,omitempty"`
	Clean bool `json:"clean"`
	ContainerFilePath string `json:"containerFilePath"`
}

// NewCiWorkflowCreateRequestDataAttributes instantiates a new CiWorkflowCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiWorkflowCreateRequestDataAttributes(name string, description string, actions []CiAction, isEnabled bool, clean bool, containerFilePath string) *CiWorkflowCreateRequestDataAttributes {
	this := CiWorkflowCreateRequestDataAttributes{}
	this.Name = name
	this.Description = description
	this.Actions = actions
	this.IsEnabled = isEnabled
	this.Clean = clean
	this.ContainerFilePath = containerFilePath
	return &this
}

// NewCiWorkflowCreateRequestDataAttributesWithDefaults instantiates a new CiWorkflowCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiWorkflowCreateRequestDataAttributesWithDefaults() *CiWorkflowCreateRequestDataAttributes {
	this := CiWorkflowCreateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *CiWorkflowCreateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CiWorkflowCreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *CiWorkflowCreateRequestDataAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CiWorkflowCreateRequestDataAttributes) SetDescription(v string) {
	o.Description = v
}

// GetBranchStartCondition returns the BranchStartCondition field value if set, zero value otherwise.
func (o *CiWorkflowCreateRequestDataAttributes) GetBranchStartCondition() CiBranchStartCondition {
	if o == nil || IsNil(o.BranchStartCondition) {
		var ret CiBranchStartCondition
		return ret
	}
	return *o.BranchStartCondition
}

// GetBranchStartConditionOk returns a tuple with the BranchStartCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataAttributes) GetBranchStartConditionOk() (*CiBranchStartCondition, bool) {
	if o == nil || IsNil(o.BranchStartCondition) {
		return nil, false
	}
	return o.BranchStartCondition, true
}

// HasBranchStartCondition returns a boolean if a field has been set.
func (o *CiWorkflowCreateRequestDataAttributes) HasBranchStartCondition() bool {
	if o != nil && !IsNil(o.BranchStartCondition) {
		return true
	}

	return false
}

// SetBranchStartCondition gets a reference to the given CiBranchStartCondition and assigns it to the BranchStartCondition field.
func (o *CiWorkflowCreateRequestDataAttributes) SetBranchStartCondition(v CiBranchStartCondition) {
	o.BranchStartCondition = &v
}

// GetTagStartCondition returns the TagStartCondition field value if set, zero value otherwise.
func (o *CiWorkflowCreateRequestDataAttributes) GetTagStartCondition() CiTagStartCondition {
	if o == nil || IsNil(o.TagStartCondition) {
		var ret CiTagStartCondition
		return ret
	}
	return *o.TagStartCondition
}

// GetTagStartConditionOk returns a tuple with the TagStartCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataAttributes) GetTagStartConditionOk() (*CiTagStartCondition, bool) {
	if o == nil || IsNil(o.TagStartCondition) {
		return nil, false
	}
	return o.TagStartCondition, true
}

// HasTagStartCondition returns a boolean if a field has been set.
func (o *CiWorkflowCreateRequestDataAttributes) HasTagStartCondition() bool {
	if o != nil && !IsNil(o.TagStartCondition) {
		return true
	}

	return false
}

// SetTagStartCondition gets a reference to the given CiTagStartCondition and assigns it to the TagStartCondition field.
func (o *CiWorkflowCreateRequestDataAttributes) SetTagStartCondition(v CiTagStartCondition) {
	o.TagStartCondition = &v
}

// GetPullRequestStartCondition returns the PullRequestStartCondition field value if set, zero value otherwise.
func (o *CiWorkflowCreateRequestDataAttributes) GetPullRequestStartCondition() CiPullRequestStartCondition {
	if o == nil || IsNil(o.PullRequestStartCondition) {
		var ret CiPullRequestStartCondition
		return ret
	}
	return *o.PullRequestStartCondition
}

// GetPullRequestStartConditionOk returns a tuple with the PullRequestStartCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataAttributes) GetPullRequestStartConditionOk() (*CiPullRequestStartCondition, bool) {
	if o == nil || IsNil(o.PullRequestStartCondition) {
		return nil, false
	}
	return o.PullRequestStartCondition, true
}

// HasPullRequestStartCondition returns a boolean if a field has been set.
func (o *CiWorkflowCreateRequestDataAttributes) HasPullRequestStartCondition() bool {
	if o != nil && !IsNil(o.PullRequestStartCondition) {
		return true
	}

	return false
}

// SetPullRequestStartCondition gets a reference to the given CiPullRequestStartCondition and assigns it to the PullRequestStartCondition field.
func (o *CiWorkflowCreateRequestDataAttributes) SetPullRequestStartCondition(v CiPullRequestStartCondition) {
	o.PullRequestStartCondition = &v
}

// GetScheduledStartCondition returns the ScheduledStartCondition field value if set, zero value otherwise.
func (o *CiWorkflowCreateRequestDataAttributes) GetScheduledStartCondition() CiScheduledStartCondition {
	if o == nil || IsNil(o.ScheduledStartCondition) {
		var ret CiScheduledStartCondition
		return ret
	}
	return *o.ScheduledStartCondition
}

// GetScheduledStartConditionOk returns a tuple with the ScheduledStartCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataAttributes) GetScheduledStartConditionOk() (*CiScheduledStartCondition, bool) {
	if o == nil || IsNil(o.ScheduledStartCondition) {
		return nil, false
	}
	return o.ScheduledStartCondition, true
}

// HasScheduledStartCondition returns a boolean if a field has been set.
func (o *CiWorkflowCreateRequestDataAttributes) HasScheduledStartCondition() bool {
	if o != nil && !IsNil(o.ScheduledStartCondition) {
		return true
	}

	return false
}

// SetScheduledStartCondition gets a reference to the given CiScheduledStartCondition and assigns it to the ScheduledStartCondition field.
func (o *CiWorkflowCreateRequestDataAttributes) SetScheduledStartCondition(v CiScheduledStartCondition) {
	o.ScheduledStartCondition = &v
}

// GetActions returns the Actions field value
func (o *CiWorkflowCreateRequestDataAttributes) GetActions() []CiAction {
	if o == nil {
		var ret []CiAction
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataAttributes) GetActionsOk() ([]CiAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *CiWorkflowCreateRequestDataAttributes) SetActions(v []CiAction) {
	o.Actions = v
}

// GetIsEnabled returns the IsEnabled field value
func (o *CiWorkflowCreateRequestDataAttributes) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataAttributes) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *CiWorkflowCreateRequestDataAttributes) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetIsLockedForEditing returns the IsLockedForEditing field value if set, zero value otherwise.
func (o *CiWorkflowCreateRequestDataAttributes) GetIsLockedForEditing() bool {
	if o == nil || IsNil(o.IsLockedForEditing) {
		var ret bool
		return ret
	}
	return *o.IsLockedForEditing
}

// GetIsLockedForEditingOk returns a tuple with the IsLockedForEditing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataAttributes) GetIsLockedForEditingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLockedForEditing) {
		return nil, false
	}
	return o.IsLockedForEditing, true
}

// HasIsLockedForEditing returns a boolean if a field has been set.
func (o *CiWorkflowCreateRequestDataAttributes) HasIsLockedForEditing() bool {
	if o != nil && !IsNil(o.IsLockedForEditing) {
		return true
	}

	return false
}

// SetIsLockedForEditing gets a reference to the given bool and assigns it to the IsLockedForEditing field.
func (o *CiWorkflowCreateRequestDataAttributes) SetIsLockedForEditing(v bool) {
	o.IsLockedForEditing = &v
}

// GetClean returns the Clean field value
func (o *CiWorkflowCreateRequestDataAttributes) GetClean() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Clean
}

// GetCleanOk returns a tuple with the Clean field value
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataAttributes) GetCleanOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clean, true
}

// SetClean sets field value
func (o *CiWorkflowCreateRequestDataAttributes) SetClean(v bool) {
	o.Clean = v
}

// GetContainerFilePath returns the ContainerFilePath field value
func (o *CiWorkflowCreateRequestDataAttributes) GetContainerFilePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerFilePath
}

// GetContainerFilePathOk returns a tuple with the ContainerFilePath field value
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataAttributes) GetContainerFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerFilePath, true
}

// SetContainerFilePath sets field value
func (o *CiWorkflowCreateRequestDataAttributes) SetContainerFilePath(v string) {
	o.ContainerFilePath = v
}

func (o CiWorkflowCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiWorkflowCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !IsNil(o.BranchStartCondition) {
		toSerialize["branchStartCondition"] = o.BranchStartCondition
	}
	if !IsNil(o.TagStartCondition) {
		toSerialize["tagStartCondition"] = o.TagStartCondition
	}
	if !IsNil(o.PullRequestStartCondition) {
		toSerialize["pullRequestStartCondition"] = o.PullRequestStartCondition
	}
	if !IsNil(o.ScheduledStartCondition) {
		toSerialize["scheduledStartCondition"] = o.ScheduledStartCondition
	}
	toSerialize["actions"] = o.Actions
	toSerialize["isEnabled"] = o.IsEnabled
	if !IsNil(o.IsLockedForEditing) {
		toSerialize["isLockedForEditing"] = o.IsLockedForEditing
	}
	toSerialize["clean"] = o.Clean
	toSerialize["containerFilePath"] = o.ContainerFilePath
	return toSerialize, nil
}

type NullableCiWorkflowCreateRequestDataAttributes struct {
	value *CiWorkflowCreateRequestDataAttributes
	isSet bool
}

func (v NullableCiWorkflowCreateRequestDataAttributes) Get() *CiWorkflowCreateRequestDataAttributes {
	return v.value
}

func (v *NullableCiWorkflowCreateRequestDataAttributes) Set(val *CiWorkflowCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCiWorkflowCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCiWorkflowCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiWorkflowCreateRequestDataAttributes(val *CiWorkflowCreateRequestDataAttributes) *NullableCiWorkflowCreateRequestDataAttributes {
	return &NullableCiWorkflowCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableCiWorkflowCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiWorkflowCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


