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

// checks if the UserAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAttributes{}

// UserAttributes struct for UserAttributes
type UserAttributes struct {
	Username *string `json:"username,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	Roles []UserRole `json:"roles,omitempty"`
	AllAppsVisible *bool `json:"allAppsVisible,omitempty"`
	ProvisioningAllowed *bool `json:"provisioningAllowed,omitempty"`
}

// NewUserAttributes instantiates a new UserAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAttributes() *UserAttributes {
	this := UserAttributes{}
	return &this
}

// NewUserAttributesWithDefaults instantiates a new UserAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAttributesWithDefaults() *UserAttributes {
	this := UserAttributes{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserAttributes) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAttributes) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserAttributes) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserAttributes) SetUsername(v string) {
	o.Username = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserAttributes) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAttributes) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserAttributes) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserAttributes) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserAttributes) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAttributes) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserAttributes) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UserAttributes) SetLastName(v string) {
	o.LastName = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserAttributes) GetRoles() []UserRole {
	if o == nil || IsNil(o.Roles) {
		var ret []UserRole
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAttributes) GetRolesOk() ([]UserRole, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserAttributes) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []UserRole and assigns it to the Roles field.
func (o *UserAttributes) SetRoles(v []UserRole) {
	o.Roles = v
}

// GetAllAppsVisible returns the AllAppsVisible field value if set, zero value otherwise.
func (o *UserAttributes) GetAllAppsVisible() bool {
	if o == nil || IsNil(o.AllAppsVisible) {
		var ret bool
		return ret
	}
	return *o.AllAppsVisible
}

// GetAllAppsVisibleOk returns a tuple with the AllAppsVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAttributes) GetAllAppsVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.AllAppsVisible) {
		return nil, false
	}
	return o.AllAppsVisible, true
}

// HasAllAppsVisible returns a boolean if a field has been set.
func (o *UserAttributes) HasAllAppsVisible() bool {
	if o != nil && !IsNil(o.AllAppsVisible) {
		return true
	}

	return false
}

// SetAllAppsVisible gets a reference to the given bool and assigns it to the AllAppsVisible field.
func (o *UserAttributes) SetAllAppsVisible(v bool) {
	o.AllAppsVisible = &v
}

// GetProvisioningAllowed returns the ProvisioningAllowed field value if set, zero value otherwise.
func (o *UserAttributes) GetProvisioningAllowed() bool {
	if o == nil || IsNil(o.ProvisioningAllowed) {
		var ret bool
		return ret
	}
	return *o.ProvisioningAllowed
}

// GetProvisioningAllowedOk returns a tuple with the ProvisioningAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAttributes) GetProvisioningAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.ProvisioningAllowed) {
		return nil, false
	}
	return o.ProvisioningAllowed, true
}

// HasProvisioningAllowed returns a boolean if a field has been set.
func (o *UserAttributes) HasProvisioningAllowed() bool {
	if o != nil && !IsNil(o.ProvisioningAllowed) {
		return true
	}

	return false
}

// SetProvisioningAllowed gets a reference to the given bool and assigns it to the ProvisioningAllowed field.
func (o *UserAttributes) SetProvisioningAllowed(v bool) {
	o.ProvisioningAllowed = &v
}

func (o UserAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.AllAppsVisible) {
		toSerialize["allAppsVisible"] = o.AllAppsVisible
	}
	if !IsNil(o.ProvisioningAllowed) {
		toSerialize["provisioningAllowed"] = o.ProvisioningAllowed
	}
	return toSerialize, nil
}

type NullableUserAttributes struct {
	value *UserAttributes
	isSet bool
}

func (v NullableUserAttributes) Get() *UserAttributes {
	return v.value
}

func (v *NullableUserAttributes) Set(val *UserAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAttributes(val *UserAttributes) *NullableUserAttributes {
	return &NullableUserAttributes{value: val, isSet: true}
}

func (v NullableUserAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


