/*
Strava API v3

The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stravaapi

import (
	"encoding/json"
)

// checks if the DistanceStreamAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DistanceStreamAllOf{}

// DistanceStreamAllOf struct for DistanceStreamAllOf
type DistanceStreamAllOf struct {
	// The sequence of distance values for this stream, in meters
	Data []float32 `json:"data,omitempty"`
}

// NewDistanceStreamAllOf instantiates a new DistanceStreamAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistanceStreamAllOf() *DistanceStreamAllOf {
	this := DistanceStreamAllOf{}
	return &this
}

// NewDistanceStreamAllOfWithDefaults instantiates a new DistanceStreamAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistanceStreamAllOfWithDefaults() *DistanceStreamAllOf {
	this := DistanceStreamAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DistanceStreamAllOf) GetData() []float32 {
	if o == nil || IsNil(o.Data) {
		var ret []float32
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistanceStreamAllOf) GetDataOk() ([]float32, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DistanceStreamAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []float32 and assigns it to the Data field.
func (o *DistanceStreamAllOf) SetData(v []float32) {
	o.Data = v
}

func (o DistanceStreamAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DistanceStreamAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDistanceStreamAllOf struct {
	value *DistanceStreamAllOf
	isSet bool
}

func (v NullableDistanceStreamAllOf) Get() *DistanceStreamAllOf {
	return v.value
}

func (v *NullableDistanceStreamAllOf) Set(val *DistanceStreamAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDistanceStreamAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDistanceStreamAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistanceStreamAllOf(val *DistanceStreamAllOf) *NullableDistanceStreamAllOf {
	return &NullableDistanceStreamAllOf{value: val, isSet: true}
}

func (v NullableDistanceStreamAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistanceStreamAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


