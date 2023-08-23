# DetailedSegmentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The time at which the segment was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time at which the segment was last updated. | [optional] 
**TotalElevationGain** | Pointer to **float32** | The segment&#39;s total elevation gain. | [optional] 
**Map** | Pointer to [**PolylineMap**](PolylineMap.md) |  | [optional] 
**EffortCount** | Pointer to **int32** | The total number of efforts for this segment | [optional] 
**AthleteCount** | Pointer to **int32** | The number of unique athletes who have an effort for this segment | [optional] 
**Hazardous** | Pointer to **bool** | Whether this segment is considered hazardous | [optional] 
**StarCount** | Pointer to **int32** | The number of stars for this segment | [optional] 

## Methods

### NewDetailedSegmentAllOf

`func NewDetailedSegmentAllOf() *DetailedSegmentAllOf`

NewDetailedSegmentAllOf instantiates a new DetailedSegmentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedSegmentAllOfWithDefaults

`func NewDetailedSegmentAllOfWithDefaults() *DetailedSegmentAllOf`

NewDetailedSegmentAllOfWithDefaults instantiates a new DetailedSegmentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DetailedSegmentAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DetailedSegmentAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DetailedSegmentAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DetailedSegmentAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DetailedSegmentAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DetailedSegmentAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DetailedSegmentAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DetailedSegmentAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTotalElevationGain

`func (o *DetailedSegmentAllOf) GetTotalElevationGain() float32`

GetTotalElevationGain returns the TotalElevationGain field if non-nil, zero value otherwise.

### GetTotalElevationGainOk

`func (o *DetailedSegmentAllOf) GetTotalElevationGainOk() (*float32, bool)`

GetTotalElevationGainOk returns a tuple with the TotalElevationGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElevationGain

`func (o *DetailedSegmentAllOf) SetTotalElevationGain(v float32)`

SetTotalElevationGain sets TotalElevationGain field to given value.

### HasTotalElevationGain

`func (o *DetailedSegmentAllOf) HasTotalElevationGain() bool`

HasTotalElevationGain returns a boolean if a field has been set.

### GetMap

`func (o *DetailedSegmentAllOf) GetMap() PolylineMap`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *DetailedSegmentAllOf) GetMapOk() (*PolylineMap, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *DetailedSegmentAllOf) SetMap(v PolylineMap)`

SetMap sets Map field to given value.

### HasMap

`func (o *DetailedSegmentAllOf) HasMap() bool`

HasMap returns a boolean if a field has been set.

### GetEffortCount

`func (o *DetailedSegmentAllOf) GetEffortCount() int32`

GetEffortCount returns the EffortCount field if non-nil, zero value otherwise.

### GetEffortCountOk

`func (o *DetailedSegmentAllOf) GetEffortCountOk() (*int32, bool)`

GetEffortCountOk returns a tuple with the EffortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffortCount

`func (o *DetailedSegmentAllOf) SetEffortCount(v int32)`

SetEffortCount sets EffortCount field to given value.

### HasEffortCount

`func (o *DetailedSegmentAllOf) HasEffortCount() bool`

HasEffortCount returns a boolean if a field has been set.

### GetAthleteCount

`func (o *DetailedSegmentAllOf) GetAthleteCount() int32`

GetAthleteCount returns the AthleteCount field if non-nil, zero value otherwise.

### GetAthleteCountOk

`func (o *DetailedSegmentAllOf) GetAthleteCountOk() (*int32, bool)`

GetAthleteCountOk returns a tuple with the AthleteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthleteCount

`func (o *DetailedSegmentAllOf) SetAthleteCount(v int32)`

SetAthleteCount sets AthleteCount field to given value.

### HasAthleteCount

`func (o *DetailedSegmentAllOf) HasAthleteCount() bool`

HasAthleteCount returns a boolean if a field has been set.

### GetHazardous

`func (o *DetailedSegmentAllOf) GetHazardous() bool`

GetHazardous returns the Hazardous field if non-nil, zero value otherwise.

### GetHazardousOk

`func (o *DetailedSegmentAllOf) GetHazardousOk() (*bool, bool)`

GetHazardousOk returns a tuple with the Hazardous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHazardous

`func (o *DetailedSegmentAllOf) SetHazardous(v bool)`

SetHazardous sets Hazardous field to given value.

### HasHazardous

`func (o *DetailedSegmentAllOf) HasHazardous() bool`

HasHazardous returns a boolean if a field has been set.

### GetStarCount

`func (o *DetailedSegmentAllOf) GetStarCount() int32`

GetStarCount returns the StarCount field if non-nil, zero value otherwise.

### GetStarCountOk

`func (o *DetailedSegmentAllOf) GetStarCountOk() (*int32, bool)`

GetStarCountOk returns a tuple with the StarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarCount

`func (o *DetailedSegmentAllOf) SetStarCount(v int32)`

SetStarCount sets StarCount field to given value.

### HasStarCount

`func (o *DetailedSegmentAllOf) HasStarCount() bool`

HasStarCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


