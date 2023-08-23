# SummaryActivityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | The identifier provided at upload time | [optional] 
**UploadId** | Pointer to **int64** | The identifier of the upload that resulted in this activity | [optional] 
**Athlete** | Pointer to [**MetaAthlete**](MetaAthlete.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the activity | [optional] 
**Distance** | Pointer to **float32** | The activity&#39;s distance, in meters | [optional] 
**MovingTime** | Pointer to **int32** | The activity&#39;s moving time, in seconds | [optional] 
**ElapsedTime** | Pointer to **int32** | The activity&#39;s elapsed time, in seconds | [optional] 
**TotalElevationGain** | Pointer to **float32** | The activity&#39;s total elevation gain. | [optional] 
**ElevHigh** | Pointer to **float32** | The activity&#39;s highest elevation, in meters | [optional] 
**ElevLow** | Pointer to **float32** | The activity&#39;s lowest elevation, in meters | [optional] 
**Type** | Pointer to [**ActivityType**](ActivityType.md) |  | [optional] 
**SportType** | Pointer to [**SportType**](SportType.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** | The time at which the activity was started. | [optional] 
**StartDateLocal** | Pointer to **time.Time** | The time at which the activity was started in the local timezone. | [optional] 
**Timezone** | Pointer to **string** | The timezone of the activity | [optional] 
**StartLatlng** | Pointer to **[]float32** | A pair of latitude/longitude coordinates, represented as an array of 2 floating point numbers. | [optional] 
**EndLatlng** | Pointer to **[]float32** | A pair of latitude/longitude coordinates, represented as an array of 2 floating point numbers. | [optional] 
**AchievementCount** | Pointer to **int32** | The number of achievements gained during this activity | [optional] 
**KudosCount** | Pointer to **int32** | The number of kudos given for this activity | [optional] 
**CommentCount** | Pointer to **int32** | The number of comments for this activity | [optional] 
**AthleteCount** | Pointer to **int32** | The number of athletes for taking part in a group activity | [optional] 
**PhotoCount** | Pointer to **int32** | The number of Instagram photos for this activity | [optional] 
**TotalPhotoCount** | Pointer to **int32** | The number of Instagram and Strava photos for this activity | [optional] 
**Map** | Pointer to [**PolylineMap**](PolylineMap.md) |  | [optional] 
**Trainer** | Pointer to **bool** | Whether this activity was recorded on a training machine | [optional] 
**Commute** | Pointer to **bool** | Whether this activity is a commute | [optional] 
**Manual** | Pointer to **bool** | Whether this activity was created manually | [optional] 
**Private** | Pointer to **bool** | Whether this activity is private | [optional] 
**Flagged** | Pointer to **bool** | Whether this activity is flagged | [optional] 
**WorkoutType** | Pointer to **int32** | The activity&#39;s workout type | [optional] 
**UploadIdStr** | Pointer to **string** | The unique identifier of the upload in string format | [optional] 
**AverageSpeed** | Pointer to **float32** | The activity&#39;s average speed, in meters per second | [optional] 
**MaxSpeed** | Pointer to **float32** | The activity&#39;s max speed, in meters per second | [optional] 
**HasKudoed** | Pointer to **bool** | Whether the logged-in athlete has kudoed this activity | [optional] 
**HideFromHome** | Pointer to **bool** | Whether the activity is muted | [optional] 
**GearId** | Pointer to **string** | The id of the gear for the activity | [optional] 
**Kilojoules** | Pointer to **float32** | The total work done in kilojoules during this activity. Rides only | [optional] 
**AverageWatts** | Pointer to **float32** | Average power output in watts during this activity. Rides only | [optional] 
**DeviceWatts** | Pointer to **bool** | Whether the watts are from a power meter, false if estimated | [optional] 
**MaxWatts** | Pointer to **int32** | Rides with power meter data only | [optional] 
**WeightedAverageWatts** | Pointer to **int32** | Similar to Normalized Power. Rides with power meter data only | [optional] 

## Methods

### NewSummaryActivityAllOf

`func NewSummaryActivityAllOf() *SummaryActivityAllOf`

NewSummaryActivityAllOf instantiates a new SummaryActivityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryActivityAllOfWithDefaults

`func NewSummaryActivityAllOfWithDefaults() *SummaryActivityAllOf`

NewSummaryActivityAllOfWithDefaults instantiates a new SummaryActivityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *SummaryActivityAllOf) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SummaryActivityAllOf) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SummaryActivityAllOf) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SummaryActivityAllOf) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUploadId

`func (o *SummaryActivityAllOf) GetUploadId() int64`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *SummaryActivityAllOf) GetUploadIdOk() (*int64, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *SummaryActivityAllOf) SetUploadId(v int64)`

SetUploadId sets UploadId field to given value.

### HasUploadId

`func (o *SummaryActivityAllOf) HasUploadId() bool`

HasUploadId returns a boolean if a field has been set.

### GetAthlete

`func (o *SummaryActivityAllOf) GetAthlete() MetaAthlete`

GetAthlete returns the Athlete field if non-nil, zero value otherwise.

### GetAthleteOk

`func (o *SummaryActivityAllOf) GetAthleteOk() (*MetaAthlete, bool)`

GetAthleteOk returns a tuple with the Athlete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthlete

`func (o *SummaryActivityAllOf) SetAthlete(v MetaAthlete)`

SetAthlete sets Athlete field to given value.

### HasAthlete

`func (o *SummaryActivityAllOf) HasAthlete() bool`

HasAthlete returns a boolean if a field has been set.

### GetName

`func (o *SummaryActivityAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SummaryActivityAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SummaryActivityAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SummaryActivityAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDistance

`func (o *SummaryActivityAllOf) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *SummaryActivityAllOf) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *SummaryActivityAllOf) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *SummaryActivityAllOf) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetMovingTime

`func (o *SummaryActivityAllOf) GetMovingTime() int32`

GetMovingTime returns the MovingTime field if non-nil, zero value otherwise.

### GetMovingTimeOk

`func (o *SummaryActivityAllOf) GetMovingTimeOk() (*int32, bool)`

GetMovingTimeOk returns a tuple with the MovingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovingTime

`func (o *SummaryActivityAllOf) SetMovingTime(v int32)`

SetMovingTime sets MovingTime field to given value.

### HasMovingTime

`func (o *SummaryActivityAllOf) HasMovingTime() bool`

HasMovingTime returns a boolean if a field has been set.

### GetElapsedTime

`func (o *SummaryActivityAllOf) GetElapsedTime() int32`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *SummaryActivityAllOf) GetElapsedTimeOk() (*int32, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *SummaryActivityAllOf) SetElapsedTime(v int32)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *SummaryActivityAllOf) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetTotalElevationGain

`func (o *SummaryActivityAllOf) GetTotalElevationGain() float32`

GetTotalElevationGain returns the TotalElevationGain field if non-nil, zero value otherwise.

### GetTotalElevationGainOk

`func (o *SummaryActivityAllOf) GetTotalElevationGainOk() (*float32, bool)`

GetTotalElevationGainOk returns a tuple with the TotalElevationGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElevationGain

`func (o *SummaryActivityAllOf) SetTotalElevationGain(v float32)`

SetTotalElevationGain sets TotalElevationGain field to given value.

### HasTotalElevationGain

`func (o *SummaryActivityAllOf) HasTotalElevationGain() bool`

HasTotalElevationGain returns a boolean if a field has been set.

### GetElevHigh

`func (o *SummaryActivityAllOf) GetElevHigh() float32`

GetElevHigh returns the ElevHigh field if non-nil, zero value otherwise.

### GetElevHighOk

`func (o *SummaryActivityAllOf) GetElevHighOk() (*float32, bool)`

GetElevHighOk returns a tuple with the ElevHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevHigh

`func (o *SummaryActivityAllOf) SetElevHigh(v float32)`

SetElevHigh sets ElevHigh field to given value.

### HasElevHigh

`func (o *SummaryActivityAllOf) HasElevHigh() bool`

HasElevHigh returns a boolean if a field has been set.

### GetElevLow

`func (o *SummaryActivityAllOf) GetElevLow() float32`

GetElevLow returns the ElevLow field if non-nil, zero value otherwise.

### GetElevLowOk

`func (o *SummaryActivityAllOf) GetElevLowOk() (*float32, bool)`

GetElevLowOk returns a tuple with the ElevLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevLow

`func (o *SummaryActivityAllOf) SetElevLow(v float32)`

SetElevLow sets ElevLow field to given value.

### HasElevLow

`func (o *SummaryActivityAllOf) HasElevLow() bool`

HasElevLow returns a boolean if a field has been set.

### GetType

`func (o *SummaryActivityAllOf) GetType() ActivityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SummaryActivityAllOf) GetTypeOk() (*ActivityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SummaryActivityAllOf) SetType(v ActivityType)`

SetType sets Type field to given value.

### HasType

`func (o *SummaryActivityAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSportType

`func (o *SummaryActivityAllOf) GetSportType() SportType`

GetSportType returns the SportType field if non-nil, zero value otherwise.

### GetSportTypeOk

`func (o *SummaryActivityAllOf) GetSportTypeOk() (*SportType, bool)`

GetSportTypeOk returns a tuple with the SportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSportType

`func (o *SummaryActivityAllOf) SetSportType(v SportType)`

SetSportType sets SportType field to given value.

### HasSportType

`func (o *SummaryActivityAllOf) HasSportType() bool`

HasSportType returns a boolean if a field has been set.

### GetStartDate

`func (o *SummaryActivityAllOf) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SummaryActivityAllOf) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SummaryActivityAllOf) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SummaryActivityAllOf) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartDateLocal

`func (o *SummaryActivityAllOf) GetStartDateLocal() time.Time`

GetStartDateLocal returns the StartDateLocal field if non-nil, zero value otherwise.

### GetStartDateLocalOk

`func (o *SummaryActivityAllOf) GetStartDateLocalOk() (*time.Time, bool)`

GetStartDateLocalOk returns a tuple with the StartDateLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateLocal

`func (o *SummaryActivityAllOf) SetStartDateLocal(v time.Time)`

SetStartDateLocal sets StartDateLocal field to given value.

### HasStartDateLocal

`func (o *SummaryActivityAllOf) HasStartDateLocal() bool`

HasStartDateLocal returns a boolean if a field has been set.

### GetTimezone

`func (o *SummaryActivityAllOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SummaryActivityAllOf) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SummaryActivityAllOf) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SummaryActivityAllOf) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetStartLatlng

`func (o *SummaryActivityAllOf) GetStartLatlng() []float32`

GetStartLatlng returns the StartLatlng field if non-nil, zero value otherwise.

### GetStartLatlngOk

`func (o *SummaryActivityAllOf) GetStartLatlngOk() (*[]float32, bool)`

GetStartLatlngOk returns a tuple with the StartLatlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLatlng

`func (o *SummaryActivityAllOf) SetStartLatlng(v []float32)`

SetStartLatlng sets StartLatlng field to given value.

### HasStartLatlng

`func (o *SummaryActivityAllOf) HasStartLatlng() bool`

HasStartLatlng returns a boolean if a field has been set.

### GetEndLatlng

`func (o *SummaryActivityAllOf) GetEndLatlng() []float32`

GetEndLatlng returns the EndLatlng field if non-nil, zero value otherwise.

### GetEndLatlngOk

`func (o *SummaryActivityAllOf) GetEndLatlngOk() (*[]float32, bool)`

GetEndLatlngOk returns a tuple with the EndLatlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLatlng

`func (o *SummaryActivityAllOf) SetEndLatlng(v []float32)`

SetEndLatlng sets EndLatlng field to given value.

### HasEndLatlng

`func (o *SummaryActivityAllOf) HasEndLatlng() bool`

HasEndLatlng returns a boolean if a field has been set.

### GetAchievementCount

`func (o *SummaryActivityAllOf) GetAchievementCount() int32`

GetAchievementCount returns the AchievementCount field if non-nil, zero value otherwise.

### GetAchievementCountOk

`func (o *SummaryActivityAllOf) GetAchievementCountOk() (*int32, bool)`

GetAchievementCountOk returns a tuple with the AchievementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementCount

`func (o *SummaryActivityAllOf) SetAchievementCount(v int32)`

SetAchievementCount sets AchievementCount field to given value.

### HasAchievementCount

`func (o *SummaryActivityAllOf) HasAchievementCount() bool`

HasAchievementCount returns a boolean if a field has been set.

### GetKudosCount

`func (o *SummaryActivityAllOf) GetKudosCount() int32`

GetKudosCount returns the KudosCount field if non-nil, zero value otherwise.

### GetKudosCountOk

`func (o *SummaryActivityAllOf) GetKudosCountOk() (*int32, bool)`

GetKudosCountOk returns a tuple with the KudosCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKudosCount

`func (o *SummaryActivityAllOf) SetKudosCount(v int32)`

SetKudosCount sets KudosCount field to given value.

### HasKudosCount

`func (o *SummaryActivityAllOf) HasKudosCount() bool`

HasKudosCount returns a boolean if a field has been set.

### GetCommentCount

`func (o *SummaryActivityAllOf) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *SummaryActivityAllOf) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *SummaryActivityAllOf) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *SummaryActivityAllOf) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### GetAthleteCount

`func (o *SummaryActivityAllOf) GetAthleteCount() int32`

GetAthleteCount returns the AthleteCount field if non-nil, zero value otherwise.

### GetAthleteCountOk

`func (o *SummaryActivityAllOf) GetAthleteCountOk() (*int32, bool)`

GetAthleteCountOk returns a tuple with the AthleteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthleteCount

`func (o *SummaryActivityAllOf) SetAthleteCount(v int32)`

SetAthleteCount sets AthleteCount field to given value.

### HasAthleteCount

`func (o *SummaryActivityAllOf) HasAthleteCount() bool`

HasAthleteCount returns a boolean if a field has been set.

### GetPhotoCount

`func (o *SummaryActivityAllOf) GetPhotoCount() int32`

GetPhotoCount returns the PhotoCount field if non-nil, zero value otherwise.

### GetPhotoCountOk

`func (o *SummaryActivityAllOf) GetPhotoCountOk() (*int32, bool)`

GetPhotoCountOk returns a tuple with the PhotoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoCount

`func (o *SummaryActivityAllOf) SetPhotoCount(v int32)`

SetPhotoCount sets PhotoCount field to given value.

### HasPhotoCount

`func (o *SummaryActivityAllOf) HasPhotoCount() bool`

HasPhotoCount returns a boolean if a field has been set.

### GetTotalPhotoCount

`func (o *SummaryActivityAllOf) GetTotalPhotoCount() int32`

GetTotalPhotoCount returns the TotalPhotoCount field if non-nil, zero value otherwise.

### GetTotalPhotoCountOk

`func (o *SummaryActivityAllOf) GetTotalPhotoCountOk() (*int32, bool)`

GetTotalPhotoCountOk returns a tuple with the TotalPhotoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPhotoCount

`func (o *SummaryActivityAllOf) SetTotalPhotoCount(v int32)`

SetTotalPhotoCount sets TotalPhotoCount field to given value.

### HasTotalPhotoCount

`func (o *SummaryActivityAllOf) HasTotalPhotoCount() bool`

HasTotalPhotoCount returns a boolean if a field has been set.

### GetMap

`func (o *SummaryActivityAllOf) GetMap() PolylineMap`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *SummaryActivityAllOf) GetMapOk() (*PolylineMap, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *SummaryActivityAllOf) SetMap(v PolylineMap)`

SetMap sets Map field to given value.

### HasMap

`func (o *SummaryActivityAllOf) HasMap() bool`

HasMap returns a boolean if a field has been set.

### GetTrainer

`func (o *SummaryActivityAllOf) GetTrainer() bool`

GetTrainer returns the Trainer field if non-nil, zero value otherwise.

### GetTrainerOk

`func (o *SummaryActivityAllOf) GetTrainerOk() (*bool, bool)`

GetTrainerOk returns a tuple with the Trainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainer

`func (o *SummaryActivityAllOf) SetTrainer(v bool)`

SetTrainer sets Trainer field to given value.

### HasTrainer

`func (o *SummaryActivityAllOf) HasTrainer() bool`

HasTrainer returns a boolean if a field has been set.

### GetCommute

`func (o *SummaryActivityAllOf) GetCommute() bool`

GetCommute returns the Commute field if non-nil, zero value otherwise.

### GetCommuteOk

`func (o *SummaryActivityAllOf) GetCommuteOk() (*bool, bool)`

GetCommuteOk returns a tuple with the Commute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommute

`func (o *SummaryActivityAllOf) SetCommute(v bool)`

SetCommute sets Commute field to given value.

### HasCommute

`func (o *SummaryActivityAllOf) HasCommute() bool`

HasCommute returns a boolean if a field has been set.

### GetManual

`func (o *SummaryActivityAllOf) GetManual() bool`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *SummaryActivityAllOf) GetManualOk() (*bool, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *SummaryActivityAllOf) SetManual(v bool)`

SetManual sets Manual field to given value.

### HasManual

`func (o *SummaryActivityAllOf) HasManual() bool`

HasManual returns a boolean if a field has been set.

### GetPrivate

`func (o *SummaryActivityAllOf) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *SummaryActivityAllOf) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *SummaryActivityAllOf) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *SummaryActivityAllOf) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetFlagged

`func (o *SummaryActivityAllOf) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *SummaryActivityAllOf) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *SummaryActivityAllOf) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *SummaryActivityAllOf) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### GetWorkoutType

`func (o *SummaryActivityAllOf) GetWorkoutType() int32`

GetWorkoutType returns the WorkoutType field if non-nil, zero value otherwise.

### GetWorkoutTypeOk

`func (o *SummaryActivityAllOf) GetWorkoutTypeOk() (*int32, bool)`

GetWorkoutTypeOk returns a tuple with the WorkoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkoutType

`func (o *SummaryActivityAllOf) SetWorkoutType(v int32)`

SetWorkoutType sets WorkoutType field to given value.

### HasWorkoutType

`func (o *SummaryActivityAllOf) HasWorkoutType() bool`

HasWorkoutType returns a boolean if a field has been set.

### GetUploadIdStr

`func (o *SummaryActivityAllOf) GetUploadIdStr() string`

GetUploadIdStr returns the UploadIdStr field if non-nil, zero value otherwise.

### GetUploadIdStrOk

`func (o *SummaryActivityAllOf) GetUploadIdStrOk() (*string, bool)`

GetUploadIdStrOk returns a tuple with the UploadIdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadIdStr

`func (o *SummaryActivityAllOf) SetUploadIdStr(v string)`

SetUploadIdStr sets UploadIdStr field to given value.

### HasUploadIdStr

`func (o *SummaryActivityAllOf) HasUploadIdStr() bool`

HasUploadIdStr returns a boolean if a field has been set.

### GetAverageSpeed

`func (o *SummaryActivityAllOf) GetAverageSpeed() float32`

GetAverageSpeed returns the AverageSpeed field if non-nil, zero value otherwise.

### GetAverageSpeedOk

`func (o *SummaryActivityAllOf) GetAverageSpeedOk() (*float32, bool)`

GetAverageSpeedOk returns a tuple with the AverageSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageSpeed

`func (o *SummaryActivityAllOf) SetAverageSpeed(v float32)`

SetAverageSpeed sets AverageSpeed field to given value.

### HasAverageSpeed

`func (o *SummaryActivityAllOf) HasAverageSpeed() bool`

HasAverageSpeed returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *SummaryActivityAllOf) GetMaxSpeed() float32`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *SummaryActivityAllOf) GetMaxSpeedOk() (*float32, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *SummaryActivityAllOf) SetMaxSpeed(v float32)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *SummaryActivityAllOf) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetHasKudoed

`func (o *SummaryActivityAllOf) GetHasKudoed() bool`

GetHasKudoed returns the HasKudoed field if non-nil, zero value otherwise.

### GetHasKudoedOk

`func (o *SummaryActivityAllOf) GetHasKudoedOk() (*bool, bool)`

GetHasKudoedOk returns a tuple with the HasKudoed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKudoed

`func (o *SummaryActivityAllOf) SetHasKudoed(v bool)`

SetHasKudoed sets HasKudoed field to given value.

### HasHasKudoed

`func (o *SummaryActivityAllOf) HasHasKudoed() bool`

HasHasKudoed returns a boolean if a field has been set.

### GetHideFromHome

`func (o *SummaryActivityAllOf) GetHideFromHome() bool`

GetHideFromHome returns the HideFromHome field if non-nil, zero value otherwise.

### GetHideFromHomeOk

`func (o *SummaryActivityAllOf) GetHideFromHomeOk() (*bool, bool)`

GetHideFromHomeOk returns a tuple with the HideFromHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromHome

`func (o *SummaryActivityAllOf) SetHideFromHome(v bool)`

SetHideFromHome sets HideFromHome field to given value.

### HasHideFromHome

`func (o *SummaryActivityAllOf) HasHideFromHome() bool`

HasHideFromHome returns a boolean if a field has been set.

### GetGearId

`func (o *SummaryActivityAllOf) GetGearId() string`

GetGearId returns the GearId field if non-nil, zero value otherwise.

### GetGearIdOk

`func (o *SummaryActivityAllOf) GetGearIdOk() (*string, bool)`

GetGearIdOk returns a tuple with the GearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGearId

`func (o *SummaryActivityAllOf) SetGearId(v string)`

SetGearId sets GearId field to given value.

### HasGearId

`func (o *SummaryActivityAllOf) HasGearId() bool`

HasGearId returns a boolean if a field has been set.

### GetKilojoules

`func (o *SummaryActivityAllOf) GetKilojoules() float32`

GetKilojoules returns the Kilojoules field if non-nil, zero value otherwise.

### GetKilojoulesOk

`func (o *SummaryActivityAllOf) GetKilojoulesOk() (*float32, bool)`

GetKilojoulesOk returns a tuple with the Kilojoules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKilojoules

`func (o *SummaryActivityAllOf) SetKilojoules(v float32)`

SetKilojoules sets Kilojoules field to given value.

### HasKilojoules

`func (o *SummaryActivityAllOf) HasKilojoules() bool`

HasKilojoules returns a boolean if a field has been set.

### GetAverageWatts

`func (o *SummaryActivityAllOf) GetAverageWatts() float32`

GetAverageWatts returns the AverageWatts field if non-nil, zero value otherwise.

### GetAverageWattsOk

`func (o *SummaryActivityAllOf) GetAverageWattsOk() (*float32, bool)`

GetAverageWattsOk returns a tuple with the AverageWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageWatts

`func (o *SummaryActivityAllOf) SetAverageWatts(v float32)`

SetAverageWatts sets AverageWatts field to given value.

### HasAverageWatts

`func (o *SummaryActivityAllOf) HasAverageWatts() bool`

HasAverageWatts returns a boolean if a field has been set.

### GetDeviceWatts

`func (o *SummaryActivityAllOf) GetDeviceWatts() bool`

GetDeviceWatts returns the DeviceWatts field if non-nil, zero value otherwise.

### GetDeviceWattsOk

`func (o *SummaryActivityAllOf) GetDeviceWattsOk() (*bool, bool)`

GetDeviceWattsOk returns a tuple with the DeviceWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceWatts

`func (o *SummaryActivityAllOf) SetDeviceWatts(v bool)`

SetDeviceWatts sets DeviceWatts field to given value.

### HasDeviceWatts

`func (o *SummaryActivityAllOf) HasDeviceWatts() bool`

HasDeviceWatts returns a boolean if a field has been set.

### GetMaxWatts

`func (o *SummaryActivityAllOf) GetMaxWatts() int32`

GetMaxWatts returns the MaxWatts field if non-nil, zero value otherwise.

### GetMaxWattsOk

`func (o *SummaryActivityAllOf) GetMaxWattsOk() (*int32, bool)`

GetMaxWattsOk returns a tuple with the MaxWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWatts

`func (o *SummaryActivityAllOf) SetMaxWatts(v int32)`

SetMaxWatts sets MaxWatts field to given value.

### HasMaxWatts

`func (o *SummaryActivityAllOf) HasMaxWatts() bool`

HasMaxWatts returns a boolean if a field has been set.

### GetWeightedAverageWatts

`func (o *SummaryActivityAllOf) GetWeightedAverageWatts() int32`

GetWeightedAverageWatts returns the WeightedAverageWatts field if non-nil, zero value otherwise.

### GetWeightedAverageWattsOk

`func (o *SummaryActivityAllOf) GetWeightedAverageWattsOk() (*int32, bool)`

GetWeightedAverageWattsOk returns a tuple with the WeightedAverageWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedAverageWatts

`func (o *SummaryActivityAllOf) SetWeightedAverageWatts(v int32)`

SetWeightedAverageWatts sets WeightedAverageWatts field to given value.

### HasWeightedAverageWatts

`func (o *SummaryActivityAllOf) HasWeightedAverageWatts() bool`

HasWeightedAverageWatts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


