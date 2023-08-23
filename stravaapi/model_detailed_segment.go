/*
Strava API v3

The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stravaapi

import (
	"encoding/json"
	"time"
)

// checks if the DetailedSegment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetailedSegment{}

// DetailedSegment struct for DetailedSegment
type DetailedSegment struct {
	// The unique identifier of this segment
	Id *int64 `json:"id,omitempty"`
	// The name of this segment
	Name *string `json:"name,omitempty"`
	ActivityType *string `json:"activity_type,omitempty"`
	// The segment's distance, in meters
	Distance *float32 `json:"distance,omitempty"`
	// The segment's average grade, in percents
	AverageGrade *float32 `json:"average_grade,omitempty"`
	// The segments's maximum grade, in percents
	MaximumGrade *float32 `json:"maximum_grade,omitempty"`
	// The segments's highest elevation, in meters
	ElevationHigh *float32 `json:"elevation_high,omitempty"`
	// The segments's lowest elevation, in meters
	ElevationLow *float32 `json:"elevation_low,omitempty"`
	// A pair of latitude/longitude coordinates, represented as an array of 2 floating point numbers.
	StartLatlng []float32 `json:"start_latlng,omitempty"`
	// A pair of latitude/longitude coordinates, represented as an array of 2 floating point numbers.
	EndLatlng []float32 `json:"end_latlng,omitempty"`
	// The category of the climb [0, 5]. Higher is harder ie. 5 is Hors catégorie, 0 is uncategorized in climb_category.
	ClimbCategory *int32 `json:"climb_category,omitempty"`
	// The segments's city.
	City *string `json:"city,omitempty"`
	// The segments's state or geographical region.
	State *string `json:"state,omitempty"`
	// The segment's country.
	Country *string `json:"country,omitempty"`
	// Whether this segment is private.
	Private *bool `json:"private,omitempty"`
	AthletePrEffort *SummaryPRSegmentEffort `json:"athlete_pr_effort,omitempty"`
	AthleteSegmentStats *SummarySegmentEffort `json:"athlete_segment_stats,omitempty"`
	// The time at which the segment was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The time at which the segment was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The segment's total elevation gain.
	TotalElevationGain *float32 `json:"total_elevation_gain,omitempty"`
	Map *PolylineMap `json:"map,omitempty"`
	// The total number of efforts for this segment
	EffortCount *int32 `json:"effort_count,omitempty"`
	// The number of unique athletes who have an effort for this segment
	AthleteCount *int32 `json:"athlete_count,omitempty"`
	// Whether this segment is considered hazardous
	Hazardous *bool `json:"hazardous,omitempty"`
	// The number of stars for this segment
	StarCount *int32 `json:"star_count,omitempty"`
}

// NewDetailedSegment instantiates a new DetailedSegment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedSegment() *DetailedSegment {
	this := DetailedSegment{}
	return &this
}

// NewDetailedSegmentWithDefaults instantiates a new DetailedSegment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedSegmentWithDefaults() *DetailedSegment {
	this := DetailedSegment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DetailedSegment) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DetailedSegment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DetailedSegment) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DetailedSegment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DetailedSegment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DetailedSegment) SetName(v string) {
	o.Name = &v
}

// GetActivityType returns the ActivityType field value if set, zero value otherwise.
func (o *DetailedSegment) GetActivityType() string {
	if o == nil || IsNil(o.ActivityType) {
		var ret string
		return ret
	}
	return *o.ActivityType
}

// GetActivityTypeOk returns a tuple with the ActivityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetActivityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityType) {
		return nil, false
	}
	return o.ActivityType, true
}

// HasActivityType returns a boolean if a field has been set.
func (o *DetailedSegment) HasActivityType() bool {
	if o != nil && !IsNil(o.ActivityType) {
		return true
	}

	return false
}

// SetActivityType gets a reference to the given string and assigns it to the ActivityType field.
func (o *DetailedSegment) SetActivityType(v string) {
	o.ActivityType = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *DetailedSegment) GetDistance() float32 {
	if o == nil || IsNil(o.Distance) {
		var ret float32
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetDistanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Distance) {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *DetailedSegment) HasDistance() bool {
	if o != nil && !IsNil(o.Distance) {
		return true
	}

	return false
}

// SetDistance gets a reference to the given float32 and assigns it to the Distance field.
func (o *DetailedSegment) SetDistance(v float32) {
	o.Distance = &v
}

// GetAverageGrade returns the AverageGrade field value if set, zero value otherwise.
func (o *DetailedSegment) GetAverageGrade() float32 {
	if o == nil || IsNil(o.AverageGrade) {
		var ret float32
		return ret
	}
	return *o.AverageGrade
}

// GetAverageGradeOk returns a tuple with the AverageGrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetAverageGradeOk() (*float32, bool) {
	if o == nil || IsNil(o.AverageGrade) {
		return nil, false
	}
	return o.AverageGrade, true
}

// HasAverageGrade returns a boolean if a field has been set.
func (o *DetailedSegment) HasAverageGrade() bool {
	if o != nil && !IsNil(o.AverageGrade) {
		return true
	}

	return false
}

// SetAverageGrade gets a reference to the given float32 and assigns it to the AverageGrade field.
func (o *DetailedSegment) SetAverageGrade(v float32) {
	o.AverageGrade = &v
}

// GetMaximumGrade returns the MaximumGrade field value if set, zero value otherwise.
func (o *DetailedSegment) GetMaximumGrade() float32 {
	if o == nil || IsNil(o.MaximumGrade) {
		var ret float32
		return ret
	}
	return *o.MaximumGrade
}

// GetMaximumGradeOk returns a tuple with the MaximumGrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetMaximumGradeOk() (*float32, bool) {
	if o == nil || IsNil(o.MaximumGrade) {
		return nil, false
	}
	return o.MaximumGrade, true
}

// HasMaximumGrade returns a boolean if a field has been set.
func (o *DetailedSegment) HasMaximumGrade() bool {
	if o != nil && !IsNil(o.MaximumGrade) {
		return true
	}

	return false
}

// SetMaximumGrade gets a reference to the given float32 and assigns it to the MaximumGrade field.
func (o *DetailedSegment) SetMaximumGrade(v float32) {
	o.MaximumGrade = &v
}

// GetElevationHigh returns the ElevationHigh field value if set, zero value otherwise.
func (o *DetailedSegment) GetElevationHigh() float32 {
	if o == nil || IsNil(o.ElevationHigh) {
		var ret float32
		return ret
	}
	return *o.ElevationHigh
}

// GetElevationHighOk returns a tuple with the ElevationHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetElevationHighOk() (*float32, bool) {
	if o == nil || IsNil(o.ElevationHigh) {
		return nil, false
	}
	return o.ElevationHigh, true
}

// HasElevationHigh returns a boolean if a field has been set.
func (o *DetailedSegment) HasElevationHigh() bool {
	if o != nil && !IsNil(o.ElevationHigh) {
		return true
	}

	return false
}

// SetElevationHigh gets a reference to the given float32 and assigns it to the ElevationHigh field.
func (o *DetailedSegment) SetElevationHigh(v float32) {
	o.ElevationHigh = &v
}

// GetElevationLow returns the ElevationLow field value if set, zero value otherwise.
func (o *DetailedSegment) GetElevationLow() float32 {
	if o == nil || IsNil(o.ElevationLow) {
		var ret float32
		return ret
	}
	return *o.ElevationLow
}

// GetElevationLowOk returns a tuple with the ElevationLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetElevationLowOk() (*float32, bool) {
	if o == nil || IsNil(o.ElevationLow) {
		return nil, false
	}
	return o.ElevationLow, true
}

// HasElevationLow returns a boolean if a field has been set.
func (o *DetailedSegment) HasElevationLow() bool {
	if o != nil && !IsNil(o.ElevationLow) {
		return true
	}

	return false
}

// SetElevationLow gets a reference to the given float32 and assigns it to the ElevationLow field.
func (o *DetailedSegment) SetElevationLow(v float32) {
	o.ElevationLow = &v
}

// GetStartLatlng returns the StartLatlng field value if set, zero value otherwise.
func (o *DetailedSegment) GetStartLatlng() []float32 {
	if o == nil || IsNil(o.StartLatlng) {
		var ret []float32
		return ret
	}
	return o.StartLatlng
}

// GetStartLatlngOk returns a tuple with the StartLatlng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetStartLatlngOk() ([]float32, bool) {
	if o == nil || IsNil(o.StartLatlng) {
		return nil, false
	}
	return o.StartLatlng, true
}

// HasStartLatlng returns a boolean if a field has been set.
func (o *DetailedSegment) HasStartLatlng() bool {
	if o != nil && !IsNil(o.StartLatlng) {
		return true
	}

	return false
}

// SetStartLatlng gets a reference to the given []float32 and assigns it to the StartLatlng field.
func (o *DetailedSegment) SetStartLatlng(v []float32) {
	o.StartLatlng = v
}

// GetEndLatlng returns the EndLatlng field value if set, zero value otherwise.
func (o *DetailedSegment) GetEndLatlng() []float32 {
	if o == nil || IsNil(o.EndLatlng) {
		var ret []float32
		return ret
	}
	return o.EndLatlng
}

// GetEndLatlngOk returns a tuple with the EndLatlng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetEndLatlngOk() ([]float32, bool) {
	if o == nil || IsNil(o.EndLatlng) {
		return nil, false
	}
	return o.EndLatlng, true
}

// HasEndLatlng returns a boolean if a field has been set.
func (o *DetailedSegment) HasEndLatlng() bool {
	if o != nil && !IsNil(o.EndLatlng) {
		return true
	}

	return false
}

// SetEndLatlng gets a reference to the given []float32 and assigns it to the EndLatlng field.
func (o *DetailedSegment) SetEndLatlng(v []float32) {
	o.EndLatlng = v
}

// GetClimbCategory returns the ClimbCategory field value if set, zero value otherwise.
func (o *DetailedSegment) GetClimbCategory() int32 {
	if o == nil || IsNil(o.ClimbCategory) {
		var ret int32
		return ret
	}
	return *o.ClimbCategory
}

// GetClimbCategoryOk returns a tuple with the ClimbCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetClimbCategoryOk() (*int32, bool) {
	if o == nil || IsNil(o.ClimbCategory) {
		return nil, false
	}
	return o.ClimbCategory, true
}

// HasClimbCategory returns a boolean if a field has been set.
func (o *DetailedSegment) HasClimbCategory() bool {
	if o != nil && !IsNil(o.ClimbCategory) {
		return true
	}

	return false
}

// SetClimbCategory gets a reference to the given int32 and assigns it to the ClimbCategory field.
func (o *DetailedSegment) SetClimbCategory(v int32) {
	o.ClimbCategory = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *DetailedSegment) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *DetailedSegment) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *DetailedSegment) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DetailedSegment) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DetailedSegment) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DetailedSegment) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *DetailedSegment) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *DetailedSegment) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *DetailedSegment) SetCountry(v string) {
	o.Country = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *DetailedSegment) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *DetailedSegment) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *DetailedSegment) SetPrivate(v bool) {
	o.Private = &v
}

// GetAthletePrEffort returns the AthletePrEffort field value if set, zero value otherwise.
func (o *DetailedSegment) GetAthletePrEffort() SummaryPRSegmentEffort {
	if o == nil || IsNil(o.AthletePrEffort) {
		var ret SummaryPRSegmentEffort
		return ret
	}
	return *o.AthletePrEffort
}

// GetAthletePrEffortOk returns a tuple with the AthletePrEffort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetAthletePrEffortOk() (*SummaryPRSegmentEffort, bool) {
	if o == nil || IsNil(o.AthletePrEffort) {
		return nil, false
	}
	return o.AthletePrEffort, true
}

// HasAthletePrEffort returns a boolean if a field has been set.
func (o *DetailedSegment) HasAthletePrEffort() bool {
	if o != nil && !IsNil(o.AthletePrEffort) {
		return true
	}

	return false
}

// SetAthletePrEffort gets a reference to the given SummaryPRSegmentEffort and assigns it to the AthletePrEffort field.
func (o *DetailedSegment) SetAthletePrEffort(v SummaryPRSegmentEffort) {
	o.AthletePrEffort = &v
}

// GetAthleteSegmentStats returns the AthleteSegmentStats field value if set, zero value otherwise.
func (o *DetailedSegment) GetAthleteSegmentStats() SummarySegmentEffort {
	if o == nil || IsNil(o.AthleteSegmentStats) {
		var ret SummarySegmentEffort
		return ret
	}
	return *o.AthleteSegmentStats
}

// GetAthleteSegmentStatsOk returns a tuple with the AthleteSegmentStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetAthleteSegmentStatsOk() (*SummarySegmentEffort, bool) {
	if o == nil || IsNil(o.AthleteSegmentStats) {
		return nil, false
	}
	return o.AthleteSegmentStats, true
}

// HasAthleteSegmentStats returns a boolean if a field has been set.
func (o *DetailedSegment) HasAthleteSegmentStats() bool {
	if o != nil && !IsNil(o.AthleteSegmentStats) {
		return true
	}

	return false
}

// SetAthleteSegmentStats gets a reference to the given SummarySegmentEffort and assigns it to the AthleteSegmentStats field.
func (o *DetailedSegment) SetAthleteSegmentStats(v SummarySegmentEffort) {
	o.AthleteSegmentStats = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DetailedSegment) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DetailedSegment) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DetailedSegment) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DetailedSegment) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DetailedSegment) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DetailedSegment) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetTotalElevationGain returns the TotalElevationGain field value if set, zero value otherwise.
func (o *DetailedSegment) GetTotalElevationGain() float32 {
	if o == nil || IsNil(o.TotalElevationGain) {
		var ret float32
		return ret
	}
	return *o.TotalElevationGain
}

// GetTotalElevationGainOk returns a tuple with the TotalElevationGain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetTotalElevationGainOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalElevationGain) {
		return nil, false
	}
	return o.TotalElevationGain, true
}

// HasTotalElevationGain returns a boolean if a field has been set.
func (o *DetailedSegment) HasTotalElevationGain() bool {
	if o != nil && !IsNil(o.TotalElevationGain) {
		return true
	}

	return false
}

// SetTotalElevationGain gets a reference to the given float32 and assigns it to the TotalElevationGain field.
func (o *DetailedSegment) SetTotalElevationGain(v float32) {
	o.TotalElevationGain = &v
}

// GetMap returns the Map field value if set, zero value otherwise.
func (o *DetailedSegment) GetMap() PolylineMap {
	if o == nil || IsNil(o.Map) {
		var ret PolylineMap
		return ret
	}
	return *o.Map
}

// GetMapOk returns a tuple with the Map field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetMapOk() (*PolylineMap, bool) {
	if o == nil || IsNil(o.Map) {
		return nil, false
	}
	return o.Map, true
}

// HasMap returns a boolean if a field has been set.
func (o *DetailedSegment) HasMap() bool {
	if o != nil && !IsNil(o.Map) {
		return true
	}

	return false
}

// SetMap gets a reference to the given PolylineMap and assigns it to the Map field.
func (o *DetailedSegment) SetMap(v PolylineMap) {
	o.Map = &v
}

// GetEffortCount returns the EffortCount field value if set, zero value otherwise.
func (o *DetailedSegment) GetEffortCount() int32 {
	if o == nil || IsNil(o.EffortCount) {
		var ret int32
		return ret
	}
	return *o.EffortCount
}

// GetEffortCountOk returns a tuple with the EffortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetEffortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.EffortCount) {
		return nil, false
	}
	return o.EffortCount, true
}

// HasEffortCount returns a boolean if a field has been set.
func (o *DetailedSegment) HasEffortCount() bool {
	if o != nil && !IsNil(o.EffortCount) {
		return true
	}

	return false
}

// SetEffortCount gets a reference to the given int32 and assigns it to the EffortCount field.
func (o *DetailedSegment) SetEffortCount(v int32) {
	o.EffortCount = &v
}

// GetAthleteCount returns the AthleteCount field value if set, zero value otherwise.
func (o *DetailedSegment) GetAthleteCount() int32 {
	if o == nil || IsNil(o.AthleteCount) {
		var ret int32
		return ret
	}
	return *o.AthleteCount
}

// GetAthleteCountOk returns a tuple with the AthleteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetAthleteCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AthleteCount) {
		return nil, false
	}
	return o.AthleteCount, true
}

// HasAthleteCount returns a boolean if a field has been set.
func (o *DetailedSegment) HasAthleteCount() bool {
	if o != nil && !IsNil(o.AthleteCount) {
		return true
	}

	return false
}

// SetAthleteCount gets a reference to the given int32 and assigns it to the AthleteCount field.
func (o *DetailedSegment) SetAthleteCount(v int32) {
	o.AthleteCount = &v
}

// GetHazardous returns the Hazardous field value if set, zero value otherwise.
func (o *DetailedSegment) GetHazardous() bool {
	if o == nil || IsNil(o.Hazardous) {
		var ret bool
		return ret
	}
	return *o.Hazardous
}

// GetHazardousOk returns a tuple with the Hazardous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetHazardousOk() (*bool, bool) {
	if o == nil || IsNil(o.Hazardous) {
		return nil, false
	}
	return o.Hazardous, true
}

// HasHazardous returns a boolean if a field has been set.
func (o *DetailedSegment) HasHazardous() bool {
	if o != nil && !IsNil(o.Hazardous) {
		return true
	}

	return false
}

// SetHazardous gets a reference to the given bool and assigns it to the Hazardous field.
func (o *DetailedSegment) SetHazardous(v bool) {
	o.Hazardous = &v
}

// GetStarCount returns the StarCount field value if set, zero value otherwise.
func (o *DetailedSegment) GetStarCount() int32 {
	if o == nil || IsNil(o.StarCount) {
		var ret int32
		return ret
	}
	return *o.StarCount
}

// GetStarCountOk returns a tuple with the StarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedSegment) GetStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.StarCount) {
		return nil, false
	}
	return o.StarCount, true
}

// HasStarCount returns a boolean if a field has been set.
func (o *DetailedSegment) HasStarCount() bool {
	if o != nil && !IsNil(o.StarCount) {
		return true
	}

	return false
}

// SetStarCount gets a reference to the given int32 and assigns it to the StarCount field.
func (o *DetailedSegment) SetStarCount(v int32) {
	o.StarCount = &v
}

func (o DetailedSegment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetailedSegment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ActivityType) {
		toSerialize["activity_type"] = o.ActivityType
	}
	if !IsNil(o.Distance) {
		toSerialize["distance"] = o.Distance
	}
	if !IsNil(o.AverageGrade) {
		toSerialize["average_grade"] = o.AverageGrade
	}
	if !IsNil(o.MaximumGrade) {
		toSerialize["maximum_grade"] = o.MaximumGrade
	}
	if !IsNil(o.ElevationHigh) {
		toSerialize["elevation_high"] = o.ElevationHigh
	}
	if !IsNil(o.ElevationLow) {
		toSerialize["elevation_low"] = o.ElevationLow
	}
	if !IsNil(o.StartLatlng) {
		toSerialize["start_latlng"] = o.StartLatlng
	}
	if !IsNil(o.EndLatlng) {
		toSerialize["end_latlng"] = o.EndLatlng
	}
	if !IsNil(o.ClimbCategory) {
		toSerialize["climb_category"] = o.ClimbCategory
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if !IsNil(o.AthletePrEffort) {
		toSerialize["athlete_pr_effort"] = o.AthletePrEffort
	}
	if !IsNil(o.AthleteSegmentStats) {
		toSerialize["athlete_segment_stats"] = o.AthleteSegmentStats
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.TotalElevationGain) {
		toSerialize["total_elevation_gain"] = o.TotalElevationGain
	}
	if !IsNil(o.Map) {
		toSerialize["map"] = o.Map
	}
	if !IsNil(o.EffortCount) {
		toSerialize["effort_count"] = o.EffortCount
	}
	if !IsNil(o.AthleteCount) {
		toSerialize["athlete_count"] = o.AthleteCount
	}
	if !IsNil(o.Hazardous) {
		toSerialize["hazardous"] = o.Hazardous
	}
	if !IsNil(o.StarCount) {
		toSerialize["star_count"] = o.StarCount
	}
	return toSerialize, nil
}

type NullableDetailedSegment struct {
	value *DetailedSegment
	isSet bool
}

func (v NullableDetailedSegment) Get() *DetailedSegment {
	return v.value
}

func (v *NullableDetailedSegment) Set(val *DetailedSegment) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedSegment(val *DetailedSegment) *NullableDetailedSegment {
	return &NullableDetailedSegment{value: val, isSet: true}
}

func (v NullableDetailedSegment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


