# DetailedActivityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the activity | [optional] 
**Photos** | Pointer to [**PhotosSummary**](PhotosSummary.md) |  | [optional] 
**Gear** | Pointer to [**SummaryGear**](SummaryGear.md) |  | [optional] 
**Calories** | Pointer to **float32** | The number of kilocalories consumed during this activity | [optional] 
**SegmentEfforts** | Pointer to [**[]DetailedSegmentEffort**](DetailedSegmentEffort.md) |  | [optional] 
**DeviceName** | Pointer to **string** | The name of the device used to record the activity | [optional] 
**EmbedToken** | Pointer to **string** | The token used to embed a Strava activity | [optional] 
**SplitsMetric** | Pointer to [**[]Split**](Split.md) | The splits of this activity in metric units (for runs) | [optional] 
**SplitsStandard** | Pointer to [**[]Split**](Split.md) | The splits of this activity in imperial units (for runs) | [optional] 
**Laps** | Pointer to [**[]Lap**](Lap.md) |  | [optional] 
**BestEfforts** | Pointer to [**[]DetailedSegmentEffort**](DetailedSegmentEffort.md) |  | [optional] 

## Methods

### NewDetailedActivityAllOf

`func NewDetailedActivityAllOf() *DetailedActivityAllOf`

NewDetailedActivityAllOf instantiates a new DetailedActivityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedActivityAllOfWithDefaults

`func NewDetailedActivityAllOfWithDefaults() *DetailedActivityAllOf`

NewDetailedActivityAllOfWithDefaults instantiates a new DetailedActivityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DetailedActivityAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DetailedActivityAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DetailedActivityAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DetailedActivityAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPhotos

`func (o *DetailedActivityAllOf) GetPhotos() PhotosSummary`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *DetailedActivityAllOf) GetPhotosOk() (*PhotosSummary, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *DetailedActivityAllOf) SetPhotos(v PhotosSummary)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *DetailedActivityAllOf) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### GetGear

`func (o *DetailedActivityAllOf) GetGear() SummaryGear`

GetGear returns the Gear field if non-nil, zero value otherwise.

### GetGearOk

`func (o *DetailedActivityAllOf) GetGearOk() (*SummaryGear, bool)`

GetGearOk returns a tuple with the Gear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGear

`func (o *DetailedActivityAllOf) SetGear(v SummaryGear)`

SetGear sets Gear field to given value.

### HasGear

`func (o *DetailedActivityAllOf) HasGear() bool`

HasGear returns a boolean if a field has been set.

### GetCalories

`func (o *DetailedActivityAllOf) GetCalories() float32`

GetCalories returns the Calories field if non-nil, zero value otherwise.

### GetCaloriesOk

`func (o *DetailedActivityAllOf) GetCaloriesOk() (*float32, bool)`

GetCaloriesOk returns a tuple with the Calories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalories

`func (o *DetailedActivityAllOf) SetCalories(v float32)`

SetCalories sets Calories field to given value.

### HasCalories

`func (o *DetailedActivityAllOf) HasCalories() bool`

HasCalories returns a boolean if a field has been set.

### GetSegmentEfforts

`func (o *DetailedActivityAllOf) GetSegmentEfforts() []DetailedSegmentEffort`

GetSegmentEfforts returns the SegmentEfforts field if non-nil, zero value otherwise.

### GetSegmentEffortsOk

`func (o *DetailedActivityAllOf) GetSegmentEffortsOk() (*[]DetailedSegmentEffort, bool)`

GetSegmentEffortsOk returns a tuple with the SegmentEfforts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentEfforts

`func (o *DetailedActivityAllOf) SetSegmentEfforts(v []DetailedSegmentEffort)`

SetSegmentEfforts sets SegmentEfforts field to given value.

### HasSegmentEfforts

`func (o *DetailedActivityAllOf) HasSegmentEfforts() bool`

HasSegmentEfforts returns a boolean if a field has been set.

### GetDeviceName

`func (o *DetailedActivityAllOf) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DetailedActivityAllOf) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DetailedActivityAllOf) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DetailedActivityAllOf) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetEmbedToken

`func (o *DetailedActivityAllOf) GetEmbedToken() string`

GetEmbedToken returns the EmbedToken field if non-nil, zero value otherwise.

### GetEmbedTokenOk

`func (o *DetailedActivityAllOf) GetEmbedTokenOk() (*string, bool)`

GetEmbedTokenOk returns a tuple with the EmbedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedToken

`func (o *DetailedActivityAllOf) SetEmbedToken(v string)`

SetEmbedToken sets EmbedToken field to given value.

### HasEmbedToken

`func (o *DetailedActivityAllOf) HasEmbedToken() bool`

HasEmbedToken returns a boolean if a field has been set.

### GetSplitsMetric

`func (o *DetailedActivityAllOf) GetSplitsMetric() []Split`

GetSplitsMetric returns the SplitsMetric field if non-nil, zero value otherwise.

### GetSplitsMetricOk

`func (o *DetailedActivityAllOf) GetSplitsMetricOk() (*[]Split, bool)`

GetSplitsMetricOk returns a tuple with the SplitsMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitsMetric

`func (o *DetailedActivityAllOf) SetSplitsMetric(v []Split)`

SetSplitsMetric sets SplitsMetric field to given value.

### HasSplitsMetric

`func (o *DetailedActivityAllOf) HasSplitsMetric() bool`

HasSplitsMetric returns a boolean if a field has been set.

### GetSplitsStandard

`func (o *DetailedActivityAllOf) GetSplitsStandard() []Split`

GetSplitsStandard returns the SplitsStandard field if non-nil, zero value otherwise.

### GetSplitsStandardOk

`func (o *DetailedActivityAllOf) GetSplitsStandardOk() (*[]Split, bool)`

GetSplitsStandardOk returns a tuple with the SplitsStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitsStandard

`func (o *DetailedActivityAllOf) SetSplitsStandard(v []Split)`

SetSplitsStandard sets SplitsStandard field to given value.

### HasSplitsStandard

`func (o *DetailedActivityAllOf) HasSplitsStandard() bool`

HasSplitsStandard returns a boolean if a field has been set.

### GetLaps

`func (o *DetailedActivityAllOf) GetLaps() []Lap`

GetLaps returns the Laps field if non-nil, zero value otherwise.

### GetLapsOk

`func (o *DetailedActivityAllOf) GetLapsOk() (*[]Lap, bool)`

GetLapsOk returns a tuple with the Laps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaps

`func (o *DetailedActivityAllOf) SetLaps(v []Lap)`

SetLaps sets Laps field to given value.

### HasLaps

`func (o *DetailedActivityAllOf) HasLaps() bool`

HasLaps returns a boolean if a field has been set.

### GetBestEfforts

`func (o *DetailedActivityAllOf) GetBestEfforts() []DetailedSegmentEffort`

GetBestEfforts returns the BestEfforts field if non-nil, zero value otherwise.

### GetBestEffortsOk

`func (o *DetailedActivityAllOf) GetBestEffortsOk() (*[]DetailedSegmentEffort, bool)`

GetBestEffortsOk returns a tuple with the BestEfforts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestEfforts

`func (o *DetailedActivityAllOf) SetBestEfforts(v []DetailedSegmentEffort)`

SetBestEfforts sets BestEfforts field to given value.

### HasBestEfforts

`func (o *DetailedActivityAllOf) HasBestEfforts() bool`

HasBestEfforts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


