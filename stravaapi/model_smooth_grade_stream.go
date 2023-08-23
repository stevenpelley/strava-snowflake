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

// checks if the SmoothGradeStream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmoothGradeStream{}

// SmoothGradeStream struct for SmoothGradeStream
type SmoothGradeStream struct {
	// The number of data points in this stream
	OriginalSize *int32 `json:"original_size,omitempty"`
	// The level of detail (sampling) in which this stream was returned
	Resolution *string `json:"resolution,omitempty"`
	// The base series used in the case the stream was downsampled
	SeriesType *string `json:"series_type,omitempty"`
	// The sequence of grade values for this stream, as percents of a grade
	Data []float32 `json:"data,omitempty"`
}

// NewSmoothGradeStream instantiates a new SmoothGradeStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmoothGradeStream() *SmoothGradeStream {
	this := SmoothGradeStream{}
	return &this
}

// NewSmoothGradeStreamWithDefaults instantiates a new SmoothGradeStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmoothGradeStreamWithDefaults() *SmoothGradeStream {
	this := SmoothGradeStream{}
	return &this
}

// GetOriginalSize returns the OriginalSize field value if set, zero value otherwise.
func (o *SmoothGradeStream) GetOriginalSize() int32 {
	if o == nil || IsNil(o.OriginalSize) {
		var ret int32
		return ret
	}
	return *o.OriginalSize
}

// GetOriginalSizeOk returns a tuple with the OriginalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmoothGradeStream) GetOriginalSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginalSize) {
		return nil, false
	}
	return o.OriginalSize, true
}

// HasOriginalSize returns a boolean if a field has been set.
func (o *SmoothGradeStream) HasOriginalSize() bool {
	if o != nil && !IsNil(o.OriginalSize) {
		return true
	}

	return false
}

// SetOriginalSize gets a reference to the given int32 and assigns it to the OriginalSize field.
func (o *SmoothGradeStream) SetOriginalSize(v int32) {
	o.OriginalSize = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *SmoothGradeStream) GetResolution() string {
	if o == nil || IsNil(o.Resolution) {
		var ret string
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmoothGradeStream) GetResolutionOk() (*string, bool) {
	if o == nil || IsNil(o.Resolution) {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *SmoothGradeStream) HasResolution() bool {
	if o != nil && !IsNil(o.Resolution) {
		return true
	}

	return false
}

// SetResolution gets a reference to the given string and assigns it to the Resolution field.
func (o *SmoothGradeStream) SetResolution(v string) {
	o.Resolution = &v
}

// GetSeriesType returns the SeriesType field value if set, zero value otherwise.
func (o *SmoothGradeStream) GetSeriesType() string {
	if o == nil || IsNil(o.SeriesType) {
		var ret string
		return ret
	}
	return *o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmoothGradeStream) GetSeriesTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SeriesType) {
		return nil, false
	}
	return o.SeriesType, true
}

// HasSeriesType returns a boolean if a field has been set.
func (o *SmoothGradeStream) HasSeriesType() bool {
	if o != nil && !IsNil(o.SeriesType) {
		return true
	}

	return false
}

// SetSeriesType gets a reference to the given string and assigns it to the SeriesType field.
func (o *SmoothGradeStream) SetSeriesType(v string) {
	o.SeriesType = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SmoothGradeStream) GetData() []float32 {
	if o == nil || IsNil(o.Data) {
		var ret []float32
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmoothGradeStream) GetDataOk() ([]float32, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SmoothGradeStream) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []float32 and assigns it to the Data field.
func (o *SmoothGradeStream) SetData(v []float32) {
	o.Data = v
}

func (o SmoothGradeStream) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmoothGradeStream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OriginalSize) {
		toSerialize["original_size"] = o.OriginalSize
	}
	if !IsNil(o.Resolution) {
		toSerialize["resolution"] = o.Resolution
	}
	if !IsNil(o.SeriesType) {
		toSerialize["series_type"] = o.SeriesType
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSmoothGradeStream struct {
	value *SmoothGradeStream
	isSet bool
}

func (v NullableSmoothGradeStream) Get() *SmoothGradeStream {
	return v.value
}

func (v *NullableSmoothGradeStream) Set(val *SmoothGradeStream) {
	v.value = val
	v.isSet = true
}

func (v NullableSmoothGradeStream) IsSet() bool {
	return v.isSet
}

func (v *NullableSmoothGradeStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmoothGradeStream(val *SmoothGradeStream) *NullableSmoothGradeStream {
	return &NullableSmoothGradeStream{value: val, isSet: true}
}

func (v NullableSmoothGradeStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmoothGradeStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


