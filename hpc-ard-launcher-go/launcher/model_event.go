/*
Dispatch Centre API

The Dispatch Centre API is the execution layer for the Capsules framework.  It handles all the details of launching and monitoring runtime environments.

API version: 2.7.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package launcher

import (
	"encoding/json"
)

// Event struct for Event
type Event struct {
	Level *string `json:"level,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	Reporter *string `json:"reporter,omitempty"`
	Message *string `json:"message,omitempty"`
	ExtraInfo *map[string]string `json:"extraInfo,omitempty"`
	AdditionalPropertiesField *map[string]interface{} `json:"additionalProperties,omitempty"`
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent() *Event {
	this := Event{}
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *Event) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *Event) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *Event) SetLevel(v string) {
	o.Level = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Event) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Event) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *Event) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetReporter returns the Reporter field value if set, zero value otherwise.
func (o *Event) GetReporter() string {
	if o == nil || o.Reporter == nil {
		var ret string
		return ret
	}
	return *o.Reporter
}

// GetReporterOk returns a tuple with the Reporter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetReporterOk() (*string, bool) {
	if o == nil || o.Reporter == nil {
		return nil, false
	}
	return o.Reporter, true
}

// HasReporter returns a boolean if a field has been set.
func (o *Event) HasReporter() bool {
	if o != nil && o.Reporter != nil {
		return true
	}

	return false
}

// SetReporter gets a reference to the given string and assigns it to the Reporter field.
func (o *Event) SetReporter(v string) {
	o.Reporter = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Event) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Event) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Event) SetMessage(v string) {
	o.Message = &v
}

// GetExtraInfo returns the ExtraInfo field value if set, zero value otherwise.
func (o *Event) GetExtraInfo() map[string]string {
	if o == nil || o.ExtraInfo == nil {
		var ret map[string]string
		return ret
	}
	return *o.ExtraInfo
}

// GetExtraInfoOk returns a tuple with the ExtraInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetExtraInfoOk() (*map[string]string, bool) {
	if o == nil || o.ExtraInfo == nil {
		return nil, false
	}
	return o.ExtraInfo, true
}

// HasExtraInfo returns a boolean if a field has been set.
func (o *Event) HasExtraInfo() bool {
	if o != nil && o.ExtraInfo != nil {
		return true
	}

	return false
}

// SetExtraInfo gets a reference to the given map[string]string and assigns it to the ExtraInfo field.
func (o *Event) SetExtraInfo(v map[string]string) {
	o.ExtraInfo = &v
}

// GetAdditionalPropertiesField returns the AdditionalPropertiesField field value if set, zero value otherwise.
func (o *Event) GetAdditionalPropertiesField() map[string]interface{} {
	if o == nil || o.AdditionalPropertiesField == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AdditionalPropertiesField
}

// GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetAdditionalPropertiesFieldOk() (*map[string]interface{}, bool) {
	if o == nil || o.AdditionalPropertiesField == nil {
		return nil, false
	}
	return o.AdditionalPropertiesField, true
}

// HasAdditionalPropertiesField returns a boolean if a field has been set.
func (o *Event) HasAdditionalPropertiesField() bool {
	if o != nil && o.AdditionalPropertiesField != nil {
		return true
	}

	return false
}

// SetAdditionalPropertiesField gets a reference to the given map[string]interface{} and assigns it to the AdditionalPropertiesField field.
func (o *Event) SetAdditionalPropertiesField(v map[string]interface{}) {
	o.AdditionalPropertiesField = &v
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Reporter != nil {
		toSerialize["reporter"] = o.Reporter
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ExtraInfo != nil {
		toSerialize["extraInfo"] = o.ExtraInfo
	}
	if o.AdditionalPropertiesField != nil {
		toSerialize["additionalProperties"] = o.AdditionalPropertiesField
	}
	return json.Marshal(toSerialize)
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

