# DetailedAthleteAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FollowerCount** | Pointer to **int32** | The athlete&#39;s follower count. | [optional] 
**FriendCount** | Pointer to **int32** | The athlete&#39;s friend count. | [optional] 
**MeasurementPreference** | Pointer to **string** | The athlete&#39;s preferred unit system. | [optional] 
**Ftp** | Pointer to **int32** | The athlete&#39;s FTP (Functional Threshold Power). | [optional] 
**Weight** | Pointer to **float32** | The athlete&#39;s weight. | [optional] 
**Clubs** | Pointer to [**[]SummaryClub**](SummaryClub.md) | The athlete&#39;s clubs. | [optional] 
**Bikes** | Pointer to [**[]SummaryGear**](SummaryGear.md) | The athlete&#39;s bikes. | [optional] 
**Shoes** | Pointer to [**[]SummaryGear**](SummaryGear.md) | The athlete&#39;s shoes. | [optional] 

## Methods

### NewDetailedAthleteAllOf

`func NewDetailedAthleteAllOf() *DetailedAthleteAllOf`

NewDetailedAthleteAllOf instantiates a new DetailedAthleteAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedAthleteAllOfWithDefaults

`func NewDetailedAthleteAllOfWithDefaults() *DetailedAthleteAllOf`

NewDetailedAthleteAllOfWithDefaults instantiates a new DetailedAthleteAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFollowerCount

`func (o *DetailedAthleteAllOf) GetFollowerCount() int32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *DetailedAthleteAllOf) GetFollowerCountOk() (*int32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *DetailedAthleteAllOf) SetFollowerCount(v int32)`

SetFollowerCount sets FollowerCount field to given value.

### HasFollowerCount

`func (o *DetailedAthleteAllOf) HasFollowerCount() bool`

HasFollowerCount returns a boolean if a field has been set.

### GetFriendCount

`func (o *DetailedAthleteAllOf) GetFriendCount() int32`

GetFriendCount returns the FriendCount field if non-nil, zero value otherwise.

### GetFriendCountOk

`func (o *DetailedAthleteAllOf) GetFriendCountOk() (*int32, bool)`

GetFriendCountOk returns a tuple with the FriendCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendCount

`func (o *DetailedAthleteAllOf) SetFriendCount(v int32)`

SetFriendCount sets FriendCount field to given value.

### HasFriendCount

`func (o *DetailedAthleteAllOf) HasFriendCount() bool`

HasFriendCount returns a boolean if a field has been set.

### GetMeasurementPreference

`func (o *DetailedAthleteAllOf) GetMeasurementPreference() string`

GetMeasurementPreference returns the MeasurementPreference field if non-nil, zero value otherwise.

### GetMeasurementPreferenceOk

`func (o *DetailedAthleteAllOf) GetMeasurementPreferenceOk() (*string, bool)`

GetMeasurementPreferenceOk returns a tuple with the MeasurementPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementPreference

`func (o *DetailedAthleteAllOf) SetMeasurementPreference(v string)`

SetMeasurementPreference sets MeasurementPreference field to given value.

### HasMeasurementPreference

`func (o *DetailedAthleteAllOf) HasMeasurementPreference() bool`

HasMeasurementPreference returns a boolean if a field has been set.

### GetFtp

`func (o *DetailedAthleteAllOf) GetFtp() int32`

GetFtp returns the Ftp field if non-nil, zero value otherwise.

### GetFtpOk

`func (o *DetailedAthleteAllOf) GetFtpOk() (*int32, bool)`

GetFtpOk returns a tuple with the Ftp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtp

`func (o *DetailedAthleteAllOf) SetFtp(v int32)`

SetFtp sets Ftp field to given value.

### HasFtp

`func (o *DetailedAthleteAllOf) HasFtp() bool`

HasFtp returns a boolean if a field has been set.

### GetWeight

`func (o *DetailedAthleteAllOf) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DetailedAthleteAllOf) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DetailedAthleteAllOf) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DetailedAthleteAllOf) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetClubs

`func (o *DetailedAthleteAllOf) GetClubs() []SummaryClub`

GetClubs returns the Clubs field if non-nil, zero value otherwise.

### GetClubsOk

`func (o *DetailedAthleteAllOf) GetClubsOk() (*[]SummaryClub, bool)`

GetClubsOk returns a tuple with the Clubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClubs

`func (o *DetailedAthleteAllOf) SetClubs(v []SummaryClub)`

SetClubs sets Clubs field to given value.

### HasClubs

`func (o *DetailedAthleteAllOf) HasClubs() bool`

HasClubs returns a boolean if a field has been set.

### GetBikes

`func (o *DetailedAthleteAllOf) GetBikes() []SummaryGear`

GetBikes returns the Bikes field if non-nil, zero value otherwise.

### GetBikesOk

`func (o *DetailedAthleteAllOf) GetBikesOk() (*[]SummaryGear, bool)`

GetBikesOk returns a tuple with the Bikes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBikes

`func (o *DetailedAthleteAllOf) SetBikes(v []SummaryGear)`

SetBikes sets Bikes field to given value.

### HasBikes

`func (o *DetailedAthleteAllOf) HasBikes() bool`

HasBikes returns a boolean if a field has been set.

### GetShoes

`func (o *DetailedAthleteAllOf) GetShoes() []SummaryGear`

GetShoes returns the Shoes field if non-nil, zero value otherwise.

### GetShoesOk

`func (o *DetailedAthleteAllOf) GetShoesOk() (*[]SummaryGear, bool)`

GetShoesOk returns a tuple with the Shoes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShoes

`func (o *DetailedAthleteAllOf) SetShoes(v []SummaryGear)`

SetShoes sets Shoes field to given value.

### HasShoes

`func (o *DetailedAthleteAllOf) HasShoes() bool`

HasShoes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


