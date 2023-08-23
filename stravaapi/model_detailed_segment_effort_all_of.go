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

// checks if the DetailedSegmentEffortAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetailedSegmentEffortAllOf{}

// DetailedSegmentEffortAllOf struct for DetailedSegmentEffortAllOf
type DetailedSegmentEffortAllOf struct {
	// The name of the segment on which this effort was performed
	Name *string `json:"name,omitempty"`
	Activity *MetaActivity `json:"activity,omitempty"`
	Athlete *MetaAthlete `json:"athlete,omitempty"`
	// The effort's moving time
	MovingTime *int32 `json:"moving_time,omitempty"`
	// The start index of this effort in its activity's stream
	StartIndex *int32 `json:"start_index,omitempty"`
	// The end index of this effort in its activity's stream
	EndIndex *int32 `json:"end_index,omitempty"`
	// The effort's average cadence
	AverageCadence *float32 `json:"average_cadence,omitempty"`
	// The average wattage of this effort
	AverageWatts *float32 `json:"average_watts,omitempty"`
	// For riding efforts, whether the wattage was reported by a dedicated recording device
	DeviceWatts *bool `json:"device_watts,omitempty"`
	// The heart heart rate of the athlete during this effort
	AverageHeartrate *float32 `json:"average_heartrate,omitempty"`
	// The maximum heart rate of the athlete during this effort
	MaxHeartrate *float32 `json:"max_heartrate,omitempty"`
	Segment *SummarySegment `json:"segment,omitempty"`
	// The rank of the effort on the global leaderboard if it belongs in the top 10 at the time of upload
	KomRank *int32 `json:"kom_rank,omitempty"`
	// The rank of the effort on the athlete's leaderboard if it belongs in the top 3 at the time of upload
	PrRank *int32 `json:"pr_rank,omitempty"`
	// Whether this effort should be hidden when viewed within an activity
	Hidden *bool `json:"hidden,omitempty"`
}

// NewDetailedSegmentEffortAllOf instantiates a new DetailedSegmentEffortAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedSegmentEffortAllOf() *DetailedSegmentEffortAllOf {
	this := DetailedSegmentEffortAllOf{}
	return &this
}

// NewDetailedSegmentEffortAllOfWithDefaults instantiates a new DetailedSegmentEffortAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedSegmentEffortAllOfWithDefaults() *DetailedSegmentEffortAllOf {
	this := DetailedSegmentEffortAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DetailedSegmentEffortAllOf) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegmentEffortAllOf) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DetailedSegmentEffortAllOf) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DetailedSegmentEffortAllOf) SetName(v string) {
	o.Name = &v
}

// GetActivity returns the Activity field value if set, zero value otherwise.
func (o *DetailedSegmentEffortAllOf) GetActivity() MetaActivity {
	if o == nil || IsNil(o.Activity) {
		var ret MetaActivity
		return ret
	}
	return *o.Activity
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegmentEffortAllOf) GetActivityOk() (*MetaActivity, bool) {
	if o == nil || IsNil(o.Activity) {
		return nil, false
	}
	return o.Activity, true
}

// HasActivity returns a boolean if a field has been set.
func (o *DetailedSegmentEffortAllOf) HasActivity() bool {
	if o != nil && !IsNil(o.Activity) {
		return true
	}

	return false
}

// SetActivity gets a reference to the given MetaActivity and assigns it to the Activity field.
func (o *DetailedSegmentEffortAllOf) SetActivity(v MetaActivity) {
	o.Activity = &v
}

// GetAthlete returns the Athlete field value if set, zero value otherwise.
func (o *DetailedSegmentEffortAllOf) GetAthlete() MetaAthlete {
	if o == nil || IsNil(o.Athlete) {
		var ret MetaAthlete
		return ret
	}
	return *o.Athlete
}

// GetAthleteOk returns a tuple with the Athlete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegmentEffortAllOf) GetAthleteOk() (*MetaAthlete, bool) {
	if o == nil || IsNil(o.Athlete) {
		return nil, false
	}
	return o.Athlete, true
}

// HasAthlete returns a boolean if a field has been set.
func (o *DetailedSegmentEffortAllOf) HasAthlete() bool {
	if o != nil && !IsNil(o.Athlete) {
		return true
	}

	return false
}

// SetAthlete gets a reference to the given MetaAthlete and assigns it to the Athlete field.
func (o *DetailedSegmentEffortAllOf) SetAthlete(v MetaAthlete) {
	o.Athlete = &v
}

// GetMovingTime returns the MovingTime field value if set, zero value otherwise.
func (o *DetailedSegmentEffortAllOf) GetMovingTime() int32 {
	if o == nil || IsNil(o.MovingTime) {
		var ret int32
		return ret
	}
	return *o.MovingTime
}

// GetMovingTimeOk returns a tuple with the MovingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegmentEffortAllOf) GetMovingTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.MovingTime) {
		return nil, false
	}
	return o.MovingTime, true
}

// HasMovingTime returns a boolean if a field has been set.
func (o *DetailedSegmentEffortAllOf) HasMovingTime() bool {
	if o != nil && !IsNil(o.MovingTime) {
		return true
	}

	return false
}

// SetMovingTime gets a reference to the given int32 and assigns it to the MovingTime field.
func (o *DetailedSegmentEffortAllOf) SetMovingTime(v int32) {
	o.MovingTime = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *DetailedSegmentEffortAllOf) GetStartIndex() int32 {
	if o == nil || IsNil(o.StartIndex) {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegmentEffortAllOf) GetStartIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.StartIndex) {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *DetailedSegmentEffortAllOf) HasStartIndex() bool {
	if o != nil && !IsNil(o.StartIndex) {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *DetailedSegmentEffortAllOf) SetStartIndex(v int32) {
	o.StartIndex = &v
}

// GetEndIndex returns the EndIndex field value if set, zero value otherwise.
func (o *DetailedSegmentEffortAllOf) GetEndIndex() int32 {
	if o == nil || IsNil(o.EndIndex) {
		var ret int32
		return ret
	}
	return *o.EndIndex
}

// GetEndIndexOk returns a tuple with the EndIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegmentEffortAllOf) GetEndIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.EndIndex) {
		return nil, false
	}
	return o.EndIndex, true
}

// HasEndIndex returns a boolean if a field has been set.
func (o *DetailedSegmentEffortAllOf) HasEndIndex() bool {
	if o != nil && !IsNil(o.EndIndex) {
		return true
	}

	return false
}

// SetEndIndex gets a reference to the given int32 and assigns it to the EndIndex field.
func (o *DetailedSegmentEffortAllOf) SetEndIndex(v int32) {
	o.EndIndex = &v
}

// GetAverageCadence returns the AverageCadence field value if set, zero value otherwise.
func (o *DetailedSegmentEffortAllOf) GetAverageCadence() float32 {
	if o == nil || IsNil(o.AverageCadence) {
		var ret float32
		return ret
	}
	return *o.AverageCadence
}

// GetAverageCadenceOk returns a tuple with the AverageCadence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegmentEffortAllOf) GetAverageCadenceOk() (*float32, bool) {
	if o == nil || IsNil(o.AverageCadence) {
		return nil, false
	}
	return o.AverageCadence, true
}

// HasAverageCadence returns a boolean if a field has been set.
func (o *DetailedSegmentEffortAllOf) HasAverageCadence() bool {
	if o != nil && !IsNil(o.AverageCadence) {
		return true
	}

	return false
}

// SetAverageCadence gets a reference to the given float32 and assigns it to the AverageCadence field.
func (o *DetailedSegmentEffortAllOf) SetAverageCadence(v float32) {
	o.AverageCadence = &v
}

// GetAverageWatts returns the AverageWatts field value if set, zero value otherwise.
func (o *DetailedSegmentEffortAllOf) GetAverageWatts() float32 {
	if o == nil || IsNil(o.AverageWatts) {
		var ret float32
		return ret
	}
	return *o.AverageWatts
}

// GetAverageWattsOk returns a tuple with the AverageWatts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegmentEffortAllOf) GetAverageWattsOk() (*float32, bool) {
	if o == nil || IsNil(o.AverageWatts) {
		return nil, false
	}
	return o.AverageWatts, true
}

// HasAverageWatts returns a boolean if a field has been set.
func (o *DetailedSegmentEffortAllOf) HasAverageWatts() bool {
	if o != nil && !IsNil(o.AverageWatts) {
		return true
	}

	return false
}

// SetAverageWatts gets a reference to the given float32 and assigns it to the AverageWatts field.
func (o *DetailedSegmentEffortAllOf) SetAverageWatts(v float32) {
	o.AverageWatts = &v
}

// GetDeviceWatts returns the DeviceWatts field value if set, zero value otherwise.
func (o *DetailedSegmentEffortAllOf) GetDeviceWatts() bool {
	if o == nil || IsNil(o.DeviceWatts) {
		var ret bool
		return ret
	}
	return *o.DeviceWatts
}

// GetDeviceWattsOk returns a tuple with the DeviceWatts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegmentEffortAllOf) GetDeviceWattsOk() (*bool, bool) {
	if o == nil || IsNil(o.DeviceWatts) {
		return nil, false
	}
	return o.DeviceWatts, true
}

// HasDeviceWatts returns a boolean if a field has been set.
func (o *DetailedSegmentEffortAllOf) HasDeviceWatts() bool {
	if o != nil && !IsNil(o.DeviceWatts) {
		return true
	}

	return false
}

// SetDeviceWatts gets a reference to the given bool and assigns it to the DeviceWatts field.
func (o *DetailedSegmentEffortAllOf) SetDeviceWatts(v bool) {
	o.DeviceWatts = &v
}

// GetAverageHeartrate returns the AverageHeartrate field value if set, zero value otherwise.
func (o *DetailedSegmentEffortAllOf) GetAverageHeartrate() float32 {
	if o == nil || IsNil(o.AverageHeartrate) {
		var ret float32
		return ret
	}
	return *o.AverageHeartrate
}

// GetAverageHeartrateOk returns a tuple with the AverageHeartrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegmentEffortAllOf) GetAverageHeartrateOk() (*float32, bool) {
	if o == nil || IsNil(o.AverageHeartrate) {
		return nil, false
	}
	return o.AverageHeartrate, true
}

// HasAverageHeartrate returns a boolean if a field has been set.
func (o *DetailedSegmentEffortAllOf) HasAverageHeartrate() bool {
	if o != nil && !IsNil(o.AverageHeartrate) {
		return true
	}

	return false
}

// SetAverageHeartrate gets a reference to the given float32 and assigns it to the AverageHeartrate field.
func (o *DetailedSegmentEffortAllOf) SetAverageHeartrate(v float32) {
	o.AverageHeartrate = &v
}

// GetMaxHeartrate returns the MaxHeartrate field value if set, zero value otherwise.
func (o *DetailedSegmentEffortAllOf) GetMaxHeartrate() float32 {
	if o == nil || IsNil(o.MaxHeartrate) {
		var ret float32
		return ret
	}
	return *o.MaxHeartrate
}

// GetMaxHeartrateOk returns a tuple with the MaxHeartrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegmentEffortAllOf) GetMaxHeartrateOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxHeartrate) {
		return nil, false
	}
	return o.MaxHeartrate, true
}

// HasMaxHeartrate returns a boolean if a field has been set.
func (o *DetailedSegmentEffortAllOf) HasMaxHeartrate() bool {
	if o != nil && !IsNil(o.MaxHeartrate) {
		return true
	}

	return false
}

// SetMaxHeartrate gets a reference to the given float32 and assigns it to the MaxHeartrate field.
func (o *DetailedSegmentEffortAllOf) SetMaxHeartrate(v float32) {
	o.MaxHeartrate = &v
}

// GetSegment returns the Segment field value if set, zero value otherwise.
func (o *DetailedSegmentEffortAllOf) GetSegment() SummarySegment {
	if o == nil || IsNil(o.Segment) {
		var ret SummarySegment
		return ret
	}
	return *o.Segment
}

// GetSegmentOk returns a tuple with the Segment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegmentEffortAllOf) GetSegmentOk() (*SummarySegment, bool) {
	if o == nil || IsNil(o.Segment) {
		return nil, false
	}
	return o.Segment, true
}

// HasSegment returns a boolean if a field has been set.
func (o *DetailedSegmentEffortAllOf) HasSegment() bool {
	if o != nil && !IsNil(o.Segment) {
		return true
	}

	return false
}

// SetSegment gets a reference to the given SummarySegment and assigns it to the Segment field.
func (o *DetailedSegmentEffortAllOf) SetSegment(v SummarySegment) {
	o.Segment = &v
}

// GetKomRank returns the KomRank field value if set, zero value otherwise.
func (o *DetailedSegmentEffortAllOf) GetKomRank() int32 {
	if o == nil || IsNil(o.KomRank) {
		var ret int32
		return ret
	}
	return *o.KomRank
}

// GetKomRankOk returns a tuple with the KomRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegmentEffortAllOf) GetKomRankOk() (*int32, bool) {
	if o == nil || IsNil(o.KomRank) {
		return nil, false
	}
	return o.KomRank, true
}

// HasKomRank returns a boolean if a field has been set.
func (o *DetailedSegmentEffortAllOf) HasKomRank() bool {
	if o != nil && !IsNil(o.KomRank) {
		return true
	}

	return false
}

// SetKomRank gets a reference to the given int32 and assigns it to the KomRank field.
func (o *DetailedSegmentEffortAllOf) SetKomRank(v int32) {
	o.KomRank = &v
}

// GetPrRank returns the PrRank field value if set, zero value otherwise.
func (o *DetailedSegmentEffortAllOf) GetPrRank() int32 {
	if o == nil || IsNil(o.PrRank) {
		var ret int32
		return ret
	}
	return *o.PrRank
}

// GetPrRankOk returns a tuple with the PrRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegmentEffortAllOf) GetPrRankOk() (*int32, bool) {
	if o == nil || IsNil(o.PrRank) {
		return nil, false
	}
	return o.PrRank, true
}

// HasPrRank returns a boolean if a field has been set.
func (o *DetailedSegmentEffortAllOf) HasPrRank() bool {
	if o != nil && !IsNil(o.PrRank) {
		return true
	}

	return false
}

// SetPrRank gets a reference to the given int32 and assigns it to the PrRank field.
func (o *DetailedSegmentEffortAllOf) SetPrRank(v int32) {
	o.PrRank = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *DetailedSegmentEffortAllOf) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegmentEffortAllOf) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *DetailedSegmentEffortAllOf) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *DetailedSegmentEffortAllOf) SetHidden(v bool) {
	o.Hidden = &v
}

func (o DetailedSegmentEffortAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetailedSegmentEffortAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Activity) {
		toSerialize["activity"] = o.Activity
	}
	if !IsNil(o.Athlete) {
		toSerialize["athlete"] = o.Athlete
	}
	if !IsNil(o.MovingTime) {
		toSerialize["moving_time"] = o.MovingTime
	}
	if !IsNil(o.StartIndex) {
		toSerialize["start_index"] = o.StartIndex
	}
	if !IsNil(o.EndIndex) {
		toSerialize["end_index"] = o.EndIndex
	}
	if !IsNil(o.AverageCadence) {
		toSerialize["average_cadence"] = o.AverageCadence
	}
	if !IsNil(o.AverageWatts) {
		toSerialize["average_watts"] = o.AverageWatts
	}
	if !IsNil(o.DeviceWatts) {
		toSerialize["device_watts"] = o.DeviceWatts
	}
	if !IsNil(o.AverageHeartrate) {
		toSerialize["average_heartrate"] = o.AverageHeartrate
	}
	if !IsNil(o.MaxHeartrate) {
		toSerialize["max_heartrate"] = o.MaxHeartrate
	}
	if !IsNil(o.Segment) {
		toSerialize["segment"] = o.Segment
	}
	if !IsNil(o.KomRank) {
		toSerialize["kom_rank"] = o.KomRank
	}
	if !IsNil(o.PrRank) {
		toSerialize["pr_rank"] = o.PrRank
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	return toSerialize, nil
}

type NullableDetailedSegmentEffortAllOf struct {
	value *DetailedSegmentEffortAllOf
	isSet bool
}

func (v NullableDetailedSegmentEffortAllOf) Get() *DetailedSegmentEffortAllOf {
	return v.value
}

func (v *NullableDetailedSegmentEffortAllOf) Set(val *DetailedSegmentEffortAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedSegmentEffortAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedSegmentEffortAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedSegmentEffortAllOf(val *DetailedSegmentEffortAllOf) *NullableDetailedSegmentEffortAllOf {
	return &NullableDetailedSegmentEffortAllOf{value: val, isSet: true}
}

func (v NullableDetailedSegmentEffortAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedSegmentEffortAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


