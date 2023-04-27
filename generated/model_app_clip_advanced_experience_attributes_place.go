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

// checks if the AppClipAdvancedExperienceAttributesPlace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceAttributesPlace{}

// AppClipAdvancedExperienceAttributesPlace struct for AppClipAdvancedExperienceAttributesPlace
type AppClipAdvancedExperienceAttributesPlace struct {
	PlaceId *string `json:"placeId,omitempty"`
	Names []string `json:"names,omitempty"`
	MainAddress *AppClipAdvancedExperienceAttributesPlaceMainAddress `json:"mainAddress,omitempty"`
	DisplayPoint *AppClipAdvancedExperienceAttributesPlaceDisplayPoint `json:"displayPoint,omitempty"`
	MapAction *string `json:"mapAction,omitempty"`
	Relationship *string `json:"relationship,omitempty"`
	PhoneNumber *AppClipAdvancedExperienceAttributesPlacePhoneNumber `json:"phoneNumber,omitempty"`
	HomePage *string `json:"homePage,omitempty"`
	Categories []string `json:"categories,omitempty"`
}

// NewAppClipAdvancedExperienceAttributesPlace instantiates a new AppClipAdvancedExperienceAttributesPlace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceAttributesPlace() *AppClipAdvancedExperienceAttributesPlace {
	this := AppClipAdvancedExperienceAttributesPlace{}
	return &this
}

// NewAppClipAdvancedExperienceAttributesPlaceWithDefaults instantiates a new AppClipAdvancedExperienceAttributesPlace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceAttributesPlaceWithDefaults() *AppClipAdvancedExperienceAttributesPlace {
	this := AppClipAdvancedExperienceAttributesPlace{}
	return &this
}

// GetPlaceId returns the PlaceId field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlace) GetPlaceId() string {
	if o == nil || IsNil(o.PlaceId) {
		var ret string
		return ret
	}
	return *o.PlaceId
}

// GetPlaceIdOk returns a tuple with the PlaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) GetPlaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlaceId) {
		return nil, false
	}
	return o.PlaceId, true
}

// HasPlaceId returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) HasPlaceId() bool {
	if o != nil && !IsNil(o.PlaceId) {
		return true
	}

	return false
}

// SetPlaceId gets a reference to the given string and assigns it to the PlaceId field.
func (o *AppClipAdvancedExperienceAttributesPlace) SetPlaceId(v string) {
	o.PlaceId = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlace) GetNames() []string {
	if o == nil || IsNil(o.Names) {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) GetNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Names) {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) HasNames() bool {
	if o != nil && !IsNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *AppClipAdvancedExperienceAttributesPlace) SetNames(v []string) {
	o.Names = v
}

// GetMainAddress returns the MainAddress field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlace) GetMainAddress() AppClipAdvancedExperienceAttributesPlaceMainAddress {
	if o == nil || IsNil(o.MainAddress) {
		var ret AppClipAdvancedExperienceAttributesPlaceMainAddress
		return ret
	}
	return *o.MainAddress
}

// GetMainAddressOk returns a tuple with the MainAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) GetMainAddressOk() (*AppClipAdvancedExperienceAttributesPlaceMainAddress, bool) {
	if o == nil || IsNil(o.MainAddress) {
		return nil, false
	}
	return o.MainAddress, true
}

// HasMainAddress returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) HasMainAddress() bool {
	if o != nil && !IsNil(o.MainAddress) {
		return true
	}

	return false
}

// SetMainAddress gets a reference to the given AppClipAdvancedExperienceAttributesPlaceMainAddress and assigns it to the MainAddress field.
func (o *AppClipAdvancedExperienceAttributesPlace) SetMainAddress(v AppClipAdvancedExperienceAttributesPlaceMainAddress) {
	o.MainAddress = &v
}

// GetDisplayPoint returns the DisplayPoint field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlace) GetDisplayPoint() AppClipAdvancedExperienceAttributesPlaceDisplayPoint {
	if o == nil || IsNil(o.DisplayPoint) {
		var ret AppClipAdvancedExperienceAttributesPlaceDisplayPoint
		return ret
	}
	return *o.DisplayPoint
}

// GetDisplayPointOk returns a tuple with the DisplayPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) GetDisplayPointOk() (*AppClipAdvancedExperienceAttributesPlaceDisplayPoint, bool) {
	if o == nil || IsNil(o.DisplayPoint) {
		return nil, false
	}
	return o.DisplayPoint, true
}

// HasDisplayPoint returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) HasDisplayPoint() bool {
	if o != nil && !IsNil(o.DisplayPoint) {
		return true
	}

	return false
}

// SetDisplayPoint gets a reference to the given AppClipAdvancedExperienceAttributesPlaceDisplayPoint and assigns it to the DisplayPoint field.
func (o *AppClipAdvancedExperienceAttributesPlace) SetDisplayPoint(v AppClipAdvancedExperienceAttributesPlaceDisplayPoint) {
	o.DisplayPoint = &v
}

// GetMapAction returns the MapAction field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlace) GetMapAction() string {
	if o == nil || IsNil(o.MapAction) {
		var ret string
		return ret
	}
	return *o.MapAction
}

// GetMapActionOk returns a tuple with the MapAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) GetMapActionOk() (*string, bool) {
	if o == nil || IsNil(o.MapAction) {
		return nil, false
	}
	return o.MapAction, true
}

// HasMapAction returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) HasMapAction() bool {
	if o != nil && !IsNil(o.MapAction) {
		return true
	}

	return false
}

// SetMapAction gets a reference to the given string and assigns it to the MapAction field.
func (o *AppClipAdvancedExperienceAttributesPlace) SetMapAction(v string) {
	o.MapAction = &v
}

// GetRelationship returns the Relationship field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlace) GetRelationship() string {
	if o == nil || IsNil(o.Relationship) {
		var ret string
		return ret
	}
	return *o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) GetRelationshipOk() (*string, bool) {
	if o == nil || IsNil(o.Relationship) {
		return nil, false
	}
	return o.Relationship, true
}

// HasRelationship returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) HasRelationship() bool {
	if o != nil && !IsNil(o.Relationship) {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given string and assigns it to the Relationship field.
func (o *AppClipAdvancedExperienceAttributesPlace) SetRelationship(v string) {
	o.Relationship = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlace) GetPhoneNumber() AppClipAdvancedExperienceAttributesPlacePhoneNumber {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret AppClipAdvancedExperienceAttributesPlacePhoneNumber
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) GetPhoneNumberOk() (*AppClipAdvancedExperienceAttributesPlacePhoneNumber, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given AppClipAdvancedExperienceAttributesPlacePhoneNumber and assigns it to the PhoneNumber field.
func (o *AppClipAdvancedExperienceAttributesPlace) SetPhoneNumber(v AppClipAdvancedExperienceAttributesPlacePhoneNumber) {
	o.PhoneNumber = &v
}

// GetHomePage returns the HomePage field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlace) GetHomePage() string {
	if o == nil || IsNil(o.HomePage) {
		var ret string
		return ret
	}
	return *o.HomePage
}

// GetHomePageOk returns a tuple with the HomePage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) GetHomePageOk() (*string, bool) {
	if o == nil || IsNil(o.HomePage) {
		return nil, false
	}
	return o.HomePage, true
}

// HasHomePage returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) HasHomePage() bool {
	if o != nil && !IsNil(o.HomePage) {
		return true
	}

	return false
}

// SetHomePage gets a reference to the given string and assigns it to the HomePage field.
func (o *AppClipAdvancedExperienceAttributesPlace) SetHomePage(v string) {
	o.HomePage = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceAttributesPlace) GetCategories() []string {
	if o == nil || IsNil(o.Categories) {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) GetCategoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceAttributesPlace) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *AppClipAdvancedExperienceAttributesPlace) SetCategories(v []string) {
	o.Categories = v
}

func (o AppClipAdvancedExperienceAttributesPlace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceAttributesPlace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlaceId) {
		toSerialize["placeId"] = o.PlaceId
	}
	if !IsNil(o.Names) {
		toSerialize["names"] = o.Names
	}
	if !IsNil(o.MainAddress) {
		toSerialize["mainAddress"] = o.MainAddress
	}
	if !IsNil(o.DisplayPoint) {
		toSerialize["displayPoint"] = o.DisplayPoint
	}
	if !IsNil(o.MapAction) {
		toSerialize["mapAction"] = o.MapAction
	}
	if !IsNil(o.Relationship) {
		toSerialize["relationship"] = o.Relationship
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.HomePage) {
		toSerialize["homePage"] = o.HomePage
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceAttributesPlace struct {
	value *AppClipAdvancedExperienceAttributesPlace
	isSet bool
}

func (v NullableAppClipAdvancedExperienceAttributesPlace) Get() *AppClipAdvancedExperienceAttributesPlace {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceAttributesPlace) Set(val *AppClipAdvancedExperienceAttributesPlace) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceAttributesPlace) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceAttributesPlace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceAttributesPlace(val *AppClipAdvancedExperienceAttributesPlace) *NullableAppClipAdvancedExperienceAttributesPlace {
	return &NullableAppClipAdvancedExperienceAttributesPlace{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceAttributesPlace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceAttributesPlace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


