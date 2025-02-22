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

// checks if the SubscriptionAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionAttributes{}

// SubscriptionAttributes struct for SubscriptionAttributes
type SubscriptionAttributes struct {
	Name *string `json:"name,omitempty"`
	ProductId *string `json:"productId,omitempty"`
	FamilySharable *bool `json:"familySharable,omitempty"`
	State *string `json:"state,omitempty"`
	SubscriptionPeriod *string `json:"subscriptionPeriod,omitempty"`
	ReviewNote *string `json:"reviewNote,omitempty"`
	GroupLevel *int32 `json:"groupLevel,omitempty"`
	AvailableInAllTerritories *bool `json:"availableInAllTerritories,omitempty"`
}

// NewSubscriptionAttributes instantiates a new SubscriptionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionAttributes() *SubscriptionAttributes {
	this := SubscriptionAttributes{}
	return &this
}

// NewSubscriptionAttributesWithDefaults instantiates a new SubscriptionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionAttributesWithDefaults() *SubscriptionAttributes {
	this := SubscriptionAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubscriptionAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubscriptionAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubscriptionAttributes) SetName(v string) {
	o.Name = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *SubscriptionAttributes) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAttributes) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *SubscriptionAttributes) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *SubscriptionAttributes) SetProductId(v string) {
	o.ProductId = &v
}

// GetFamilySharable returns the FamilySharable field value if set, zero value otherwise.
func (o *SubscriptionAttributes) GetFamilySharable() bool {
	if o == nil || IsNil(o.FamilySharable) {
		var ret bool
		return ret
	}
	return *o.FamilySharable
}

// GetFamilySharableOk returns a tuple with the FamilySharable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAttributes) GetFamilySharableOk() (*bool, bool) {
	if o == nil || IsNil(o.FamilySharable) {
		return nil, false
	}
	return o.FamilySharable, true
}

// HasFamilySharable returns a boolean if a field has been set.
func (o *SubscriptionAttributes) HasFamilySharable() bool {
	if o != nil && !IsNil(o.FamilySharable) {
		return true
	}

	return false
}

// SetFamilySharable gets a reference to the given bool and assigns it to the FamilySharable field.
func (o *SubscriptionAttributes) SetFamilySharable(v bool) {
	o.FamilySharable = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SubscriptionAttributes) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAttributes) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SubscriptionAttributes) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SubscriptionAttributes) SetState(v string) {
	o.State = &v
}

// GetSubscriptionPeriod returns the SubscriptionPeriod field value if set, zero value otherwise.
func (o *SubscriptionAttributes) GetSubscriptionPeriod() string {
	if o == nil || IsNil(o.SubscriptionPeriod) {
		var ret string
		return ret
	}
	return *o.SubscriptionPeriod
}

// GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAttributes) GetSubscriptionPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionPeriod) {
		return nil, false
	}
	return o.SubscriptionPeriod, true
}

// HasSubscriptionPeriod returns a boolean if a field has been set.
func (o *SubscriptionAttributes) HasSubscriptionPeriod() bool {
	if o != nil && !IsNil(o.SubscriptionPeriod) {
		return true
	}

	return false
}

// SetSubscriptionPeriod gets a reference to the given string and assigns it to the SubscriptionPeriod field.
func (o *SubscriptionAttributes) SetSubscriptionPeriod(v string) {
	o.SubscriptionPeriod = &v
}

// GetReviewNote returns the ReviewNote field value if set, zero value otherwise.
func (o *SubscriptionAttributes) GetReviewNote() string {
	if o == nil || IsNil(o.ReviewNote) {
		var ret string
		return ret
	}
	return *o.ReviewNote
}

// GetReviewNoteOk returns a tuple with the ReviewNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAttributes) GetReviewNoteOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewNote) {
		return nil, false
	}
	return o.ReviewNote, true
}

// HasReviewNote returns a boolean if a field has been set.
func (o *SubscriptionAttributes) HasReviewNote() bool {
	if o != nil && !IsNil(o.ReviewNote) {
		return true
	}

	return false
}

// SetReviewNote gets a reference to the given string and assigns it to the ReviewNote field.
func (o *SubscriptionAttributes) SetReviewNote(v string) {
	o.ReviewNote = &v
}

// GetGroupLevel returns the GroupLevel field value if set, zero value otherwise.
func (o *SubscriptionAttributes) GetGroupLevel() int32 {
	if o == nil || IsNil(o.GroupLevel) {
		var ret int32
		return ret
	}
	return *o.GroupLevel
}

// GetGroupLevelOk returns a tuple with the GroupLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAttributes) GetGroupLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.GroupLevel) {
		return nil, false
	}
	return o.GroupLevel, true
}

// HasGroupLevel returns a boolean if a field has been set.
func (o *SubscriptionAttributes) HasGroupLevel() bool {
	if o != nil && !IsNil(o.GroupLevel) {
		return true
	}

	return false
}

// SetGroupLevel gets a reference to the given int32 and assigns it to the GroupLevel field.
func (o *SubscriptionAttributes) SetGroupLevel(v int32) {
	o.GroupLevel = &v
}

// GetAvailableInAllTerritories returns the AvailableInAllTerritories field value if set, zero value otherwise.
func (o *SubscriptionAttributes) GetAvailableInAllTerritories() bool {
	if o == nil || IsNil(o.AvailableInAllTerritories) {
		var ret bool
		return ret
	}
	return *o.AvailableInAllTerritories
}

// GetAvailableInAllTerritoriesOk returns a tuple with the AvailableInAllTerritories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAttributes) GetAvailableInAllTerritoriesOk() (*bool, bool) {
	if o == nil || IsNil(o.AvailableInAllTerritories) {
		return nil, false
	}
	return o.AvailableInAllTerritories, true
}

// HasAvailableInAllTerritories returns a boolean if a field has been set.
func (o *SubscriptionAttributes) HasAvailableInAllTerritories() bool {
	if o != nil && !IsNil(o.AvailableInAllTerritories) {
		return true
	}

	return false
}

// SetAvailableInAllTerritories gets a reference to the given bool and assigns it to the AvailableInAllTerritories field.
func (o *SubscriptionAttributes) SetAvailableInAllTerritories(v bool) {
	o.AvailableInAllTerritories = &v
}

func (o SubscriptionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.FamilySharable) {
		toSerialize["familySharable"] = o.FamilySharable
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.SubscriptionPeriod) {
		toSerialize["subscriptionPeriod"] = o.SubscriptionPeriod
	}
	if !IsNil(o.ReviewNote) {
		toSerialize["reviewNote"] = o.ReviewNote
	}
	if !IsNil(o.GroupLevel) {
		toSerialize["groupLevel"] = o.GroupLevel
	}
	if !IsNil(o.AvailableInAllTerritories) {
		toSerialize["availableInAllTerritories"] = o.AvailableInAllTerritories
	}
	return toSerialize, nil
}

type NullableSubscriptionAttributes struct {
	value *SubscriptionAttributes
	isSet bool
}

func (v NullableSubscriptionAttributes) Get() *SubscriptionAttributes {
	return v.value
}

func (v *NullableSubscriptionAttributes) Set(val *SubscriptionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionAttributes(val *SubscriptionAttributes) *NullableSubscriptionAttributes {
	return &NullableSubscriptionAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


