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

// checks if the AppStoreVersionLocalizationCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionLocalizationCreateRequestDataAttributes{}

// AppStoreVersionLocalizationCreateRequestDataAttributes struct for AppStoreVersionLocalizationCreateRequestDataAttributes
type AppStoreVersionLocalizationCreateRequestDataAttributes struct {
	Description *string `json:"description,omitempty"`
	Locale string `json:"locale"`
	Keywords *string `json:"keywords,omitempty"`
	MarketingUrl *string `json:"marketingUrl,omitempty"`
	PromotionalText *string `json:"promotionalText,omitempty"`
	SupportUrl *string `json:"supportUrl,omitempty"`
	WhatsNew *string `json:"whatsNew,omitempty"`
}

// NewAppStoreVersionLocalizationCreateRequestDataAttributes instantiates a new AppStoreVersionLocalizationCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionLocalizationCreateRequestDataAttributes(locale string) *AppStoreVersionLocalizationCreateRequestDataAttributes {
	this := AppStoreVersionLocalizationCreateRequestDataAttributes{}
	this.Locale = locale
	return &this
}

// NewAppStoreVersionLocalizationCreateRequestDataAttributesWithDefaults instantiates a new AppStoreVersionLocalizationCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionLocalizationCreateRequestDataAttributesWithDefaults() *AppStoreVersionLocalizationCreateRequestDataAttributes {
	this := AppStoreVersionLocalizationCreateRequestDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetLocale returns the Locale field value
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) SetLocale(v string) {
	o.Locale = v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) GetKeywords() string {
	if o == nil || IsNil(o.Keywords) {
		var ret string
		return ret
	}
	return *o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) GetKeywordsOk() (*string, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given string and assigns it to the Keywords field.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) SetKeywords(v string) {
	o.Keywords = &v
}

// GetMarketingUrl returns the MarketingUrl field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) GetMarketingUrl() string {
	if o == nil || IsNil(o.MarketingUrl) {
		var ret string
		return ret
	}
	return *o.MarketingUrl
}

// GetMarketingUrlOk returns a tuple with the MarketingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) GetMarketingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MarketingUrl) {
		return nil, false
	}
	return o.MarketingUrl, true
}

// HasMarketingUrl returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) HasMarketingUrl() bool {
	if o != nil && !IsNil(o.MarketingUrl) {
		return true
	}

	return false
}

// SetMarketingUrl gets a reference to the given string and assigns it to the MarketingUrl field.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) SetMarketingUrl(v string) {
	o.MarketingUrl = &v
}

// GetPromotionalText returns the PromotionalText field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) GetPromotionalText() string {
	if o == nil || IsNil(o.PromotionalText) {
		var ret string
		return ret
	}
	return *o.PromotionalText
}

// GetPromotionalTextOk returns a tuple with the PromotionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) GetPromotionalTextOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionalText) {
		return nil, false
	}
	return o.PromotionalText, true
}

// HasPromotionalText returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) HasPromotionalText() bool {
	if o != nil && !IsNil(o.PromotionalText) {
		return true
	}

	return false
}

// SetPromotionalText gets a reference to the given string and assigns it to the PromotionalText field.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) SetPromotionalText(v string) {
	o.PromotionalText = &v
}

// GetSupportUrl returns the SupportUrl field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) GetSupportUrl() string {
	if o == nil || IsNil(o.SupportUrl) {
		var ret string
		return ret
	}
	return *o.SupportUrl
}

// GetSupportUrlOk returns a tuple with the SupportUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) GetSupportUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SupportUrl) {
		return nil, false
	}
	return o.SupportUrl, true
}

// HasSupportUrl returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) HasSupportUrl() bool {
	if o != nil && !IsNil(o.SupportUrl) {
		return true
	}

	return false
}

// SetSupportUrl gets a reference to the given string and assigns it to the SupportUrl field.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) SetSupportUrl(v string) {
	o.SupportUrl = &v
}

// GetWhatsNew returns the WhatsNew field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) GetWhatsNew() string {
	if o == nil || IsNil(o.WhatsNew) {
		var ret string
		return ret
	}
	return *o.WhatsNew
}

// GetWhatsNewOk returns a tuple with the WhatsNew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) GetWhatsNewOk() (*string, bool) {
	if o == nil || IsNil(o.WhatsNew) {
		return nil, false
	}
	return o.WhatsNew, true
}

// HasWhatsNew returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) HasWhatsNew() bool {
	if o != nil && !IsNil(o.WhatsNew) {
		return true
	}

	return false
}

// SetWhatsNew gets a reference to the given string and assigns it to the WhatsNew field.
func (o *AppStoreVersionLocalizationCreateRequestDataAttributes) SetWhatsNew(v string) {
	o.WhatsNew = &v
}

func (o AppStoreVersionLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionLocalizationCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["locale"] = o.Locale
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	if !IsNil(o.MarketingUrl) {
		toSerialize["marketingUrl"] = o.MarketingUrl
	}
	if !IsNil(o.PromotionalText) {
		toSerialize["promotionalText"] = o.PromotionalText
	}
	if !IsNil(o.SupportUrl) {
		toSerialize["supportUrl"] = o.SupportUrl
	}
	if !IsNil(o.WhatsNew) {
		toSerialize["whatsNew"] = o.WhatsNew
	}
	return toSerialize, nil
}

type NullableAppStoreVersionLocalizationCreateRequestDataAttributes struct {
	value *AppStoreVersionLocalizationCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppStoreVersionLocalizationCreateRequestDataAttributes) Get() *AppStoreVersionLocalizationCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppStoreVersionLocalizationCreateRequestDataAttributes) Set(val *AppStoreVersionLocalizationCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionLocalizationCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionLocalizationCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionLocalizationCreateRequestDataAttributes(val *AppStoreVersionLocalizationCreateRequestDataAttributes) *NullableAppStoreVersionLocalizationCreateRequestDataAttributes {
	return &NullableAppStoreVersionLocalizationCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppStoreVersionLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionLocalizationCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


