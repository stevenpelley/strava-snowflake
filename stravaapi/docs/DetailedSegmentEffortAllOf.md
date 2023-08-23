# DetailedSegmentEffortAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the segment on which this effort was performed | [optional] 
**Activity** | Pointer to [**MetaActivity**](MetaActivity.md) |  | [optional] 
**Athlete** | Pointer to [**MetaAthlete**](MetaAthlete.md) |  | [optional] 
**MovingTime** | Pointer to **int32** | The effort&#39;s moving time | [optional] 
**StartIndex** | Pointer to **int32** | The start index of this effort in its activity&#39;s stream | [optional] 
**EndIndex** | Pointer to **int32** | The end index of this effort in its activity&#39;s stream | [optional] 
**AverageCadence** | Pointer to **float32** | The effort&#39;s average cadence | [optional] 
**AverageWatts** | Pointer to **float32** | The average wattage of this effort | [optional] 
**DeviceWatts** | Pointer to **bool** | For riding efforts, whether the wattage was reported by a dedicated recording device | [optional] 
**AverageHeartrate** | Pointer to **float32** | The heart heart rate of the athlete during this effort | [optional] 
**MaxHeartrate** | Pointer to **float32** | The maximum heart rate of the athlete during this effort | [optional] 
**Segment** | Pointer to [**SummarySegment**](SummarySegment.md) |  | [optional] 
**KomRank** | Pointer to **int32** | The rank of the effort on the global leaderboard if it belongs in the top 10 at the time of upload | [optional] 
**PrRank** | Pointer to **int32** | The rank of the effort on the athlete&#39;s leaderboard if it belongs in the top 3 at the time of upload | [optional] 
**Hidden** | Pointer to **bool** | Whether this effort should be hidden when viewed within an activity | [optional] 

## Methods

### NewDetailedSegmentEffortAllOf

`func NewDetailedSegmentEffortAllOf() *DetailedSegmentEffortAllOf`

NewDetailedSegmentEffortAllOf instantiates a new DetailedSegmentEffortAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedSegmentEffortAllOfWithDefaults

`func NewDetailedSegmentEffortAllOfWithDefaults() *DetailedSegmentEffortAllOf`

NewDetailedSegmentEffortAllOfWithDefaults instantiates a new DetailedSegmentEffortAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DetailedSegmentEffortAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetailedSegmentEffortAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetailedSegmentEffortAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DetailedSegmentEffortAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActivity

`func (o *DetailedSegmentEffortAllOf) GetActivity() MetaActivity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *DetailedSegmentEffortAllOf) GetActivityOk() (*MetaActivity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *DetailedSegmentEffortAllOf) SetActivity(v MetaActivity)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *DetailedSegmentEffortAllOf) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetAthlete

`func (o *DetailedSegmentEffortAllOf) GetAthlete() MetaAthlete`

GetAthlete returns the Athlete field if non-nil, zero value otherwise.

### GetAthleteOk

`func (o *DetailedSegmentEffortAllOf) GetAthleteOk() (*MetaAthlete, bool)`

GetAthleteOk returns a tuple with the Athlete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthlete

`func (o *DetailedSegmentEffortAllOf) SetAthlete(v MetaAthlete)`

SetAthlete sets Athlete field to given value.

### HasAthlete

`func (o *DetailedSegmentEffortAllOf) HasAthlete() bool`

HasAthlete returns a boolean if a field has been set.

### GetMovingTime

`func (o *DetailedSegmentEffortAllOf) GetMovingTime() int32`

GetMovingTime returns the MovingTime field if non-nil, zero value otherwise.

### GetMovingTimeOk

`func (o *DetailedSegmentEffortAllOf) GetMovingTimeOk() (*int32, bool)`

GetMovingTimeOk returns a tuple with the MovingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovingTime

`func (o *DetailedSegmentEffortAllOf) SetMovingTime(v int32)`

SetMovingTime sets MovingTime field to given value.

### HasMovingTime

`func (o *DetailedSegmentEffortAllOf) HasMovingTime() bool`

HasMovingTime returns a boolean if a field has been set.

### GetStartIndex

`func (o *DetailedSegmentEffortAllOf) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *DetailedSegmentEffortAllOf) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *DetailedSegmentEffortAllOf) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *DetailedSegmentEffortAllOf) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetEndIndex

`func (o *DetailedSegmentEffortAllOf) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *DetailedSegmentEffortAllOf) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *DetailedSegmentEffortAllOf) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.

### HasEndIndex

`func (o *DetailedSegmentEffortAllOf) HasEndIndex() bool`

HasEndIndex returns a boolean if a field has been set.

### GetAverageCadence

`func (o *DetailedSegmentEffortAllOf) GetAverageCadence() float32`

GetAverageCadence returns the AverageCadence field if non-nil, zero value otherwise.

### GetAverageCadenceOk

`func (o *DetailedSegmentEffortAllOf) GetAverageCadenceOk() (*float32, bool)`

GetAverageCadenceOk returns a tuple with the AverageCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCadence

`func (o *DetailedSegmentEffortAllOf) SetAverageCadence(v float32)`

SetAverageCadence sets AverageCadence field to given value.

### HasAverageCadence

`func (o *DetailedSegmentEffortAllOf) HasAverageCadence() bool`

HasAverageCadence returns a boolean if a field has been set.

### GetAverageWatts

`func (o *DetailedSegmentEffortAllOf) GetAverageWatts() float32`

GetAverageWatts returns the AverageWatts field if non-nil, zero value otherwise.

### GetAverageWattsOk

`func (o *DetailedSegmentEffortAllOf) GetAverageWattsOk() (*float32, bool)`

GetAverageWattsOk returns a tuple with the AverageWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageWatts

`func (o *DetailedSegmentEffortAllOf) SetAverageWatts(v float32)`

SetAverageWatts sets AverageWatts field to given value.

### HasAverageWatts

`func (o *DetailedSegmentEffortAllOf) HasAverageWatts() bool`

HasAverageWatts returns a boolean if a field has been set.

### GetDeviceWatts

`func (o *DetailedSegmentEffortAllOf) GetDeviceWatts() bool`

GetDeviceWatts returns the DeviceWatts field if non-nil, zero value otherwise.

### GetDeviceWattsOk

`func (o *DetailedSegmentEffortAllOf) GetDeviceWattsOk() (*bool, bool)`

GetDeviceWattsOk returns a tuple with the DeviceWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceWatts

`func (o *DetailedSegmentEffortAllOf) SetDeviceWatts(v bool)`

SetDeviceWatts sets DeviceWatts field to given value.

### HasDeviceWatts

`func (o *DetailedSegmentEffortAllOf) HasDeviceWatts() bool`

HasDeviceWatts returns a boolean if a field has been set.

### GetAverageHeartrate

`func (o *DetailedSegmentEffortAllOf) GetAverageHeartrate() float32`

GetAverageHeartrate returns the AverageHeartrate field if non-nil, zero value otherwise.

### GetAverageHeartrateOk

`func (o *DetailedSegmentEffortAllOf) GetAverageHeartrateOk() (*float32, bool)`

GetAverageHeartrateOk returns a tuple with the AverageHeartrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageHeartrate

`func (o *DetailedSegmentEffortAllOf) SetAverageHeartrate(v float32)`

SetAverageHeartrate sets AverageHeartrate field to given value.

### HasAverageHeartrate

`func (o *DetailedSegmentEffortAllOf) HasAverageHeartrate() bool`

HasAverageHeartrate returns a boolean if a field has been set.

### GetMaxHeartrate

`func (o *DetailedSegmentEffortAllOf) GetMaxHeartrate() float32`

GetMaxHeartrate returns the MaxHeartrate field if non-nil, zero value otherwise.

### GetMaxHeartrateOk

`func (o *DetailedSegmentEffortAllOf) GetMaxHeartrateOk() (*float32, bool)`

GetMaxHeartrateOk returns a tuple with the MaxHeartrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeartrate

`func (o *DetailedSegmentEffortAllOf) SetMaxHeartrate(v float32)`

SetMaxHeartrate sets MaxHeartrate field to given value.

### HasMaxHeartrate

`func (o *DetailedSegmentEffortAllOf) HasMaxHeartrate() bool`

HasMaxHeartrate returns a boolean if a field has been set.

### GetSegment

`func (o *DetailedSegmentEffortAllOf) GetSegment() SummarySegment`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *DetailedSegmentEffortAllOf) GetSegmentOk() (*SummarySegment, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *DetailedSegmentEffortAllOf) SetSegment(v SummarySegment)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *DetailedSegmentEffortAllOf) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetKomRank

`func (o *DetailedSegmentEffortAllOf) GetKomRank() int32`

GetKomRank returns the KomRank field if non-nil, zero value otherwise.

### GetKomRankOk

`func (o *DetailedSegmentEffortAllOf) GetKomRankOk() (*int32, bool)`

GetKomRankOk returns a tuple with the KomRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKomRank

`func (o *DetailedSegmentEffortAllOf) SetKomRank(v int32)`

SetKomRank sets KomRank field to given value.

### HasKomRank

`func (o *DetailedSegmentEffortAllOf) HasKomRank() bool`

HasKomRank returns a boolean if a field has been set.

### GetPrRank

`func (o *DetailedSegmentEffortAllOf) GetPrRank() int32`

GetPrRank returns the PrRank field if non-nil, zero value otherwise.

### GetPrRankOk

`func (o *DetailedSegmentEffortAllOf) GetPrRankOk() (*int32, bool)`

GetPrRankOk returns a tuple with the PrRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrRank

`func (o *DetailedSegmentEffortAllOf) SetPrRank(v int32)`

SetPrRank sets PrRank field to given value.

### HasPrRank

`func (o *DetailedSegmentEffortAllOf) HasPrRank() bool`

HasPrRank returns a boolean if a field has been set.

### GetHidden

`func (o *DetailedSegmentEffortAllOf) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *DetailedSegmentEffortAllOf) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *DetailedSegmentEffortAllOf) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *DetailedSegmentEffortAllOf) HasHidden() bool`

HasHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


