# SummaryAthleteAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceState** | Pointer to **int32** | Resource state, indicates level of detail. Possible values: 1 -&gt; \&quot;meta\&quot;, 2 -&gt; \&quot;summary\&quot;, 3 -&gt; \&quot;detail\&quot; | [optional] 
**Firstname** | Pointer to **string** | The athlete&#39;s first name. | [optional] 
**Lastname** | Pointer to **string** | The athlete&#39;s last name. | [optional] 
**ProfileMedium** | Pointer to **string** | URL to a 62x62 pixel profile picture. | [optional] 
**Profile** | Pointer to **string** | URL to a 124x124 pixel profile picture. | [optional] 
**City** | Pointer to **string** | The athlete&#39;s city. | [optional] 
**State** | Pointer to **string** | The athlete&#39;s state or geographical region. | [optional] 
**Country** | Pointer to **string** | The athlete&#39;s country. | [optional] 
**Sex** | Pointer to **string** | The athlete&#39;s sex. | [optional] 
**Premium** | Pointer to **bool** | Deprecated.  Use summit field instead. Whether the athlete has any Summit subscription. | [optional] 
**Summit** | Pointer to **bool** | Whether the athlete has any Summit subscription. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time at which the athlete was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time at which the athlete was last updated. | [optional] 

## Methods

### NewSummaryAthleteAllOf

`func NewSummaryAthleteAllOf() *SummaryAthleteAllOf`

NewSummaryAthleteAllOf instantiates a new SummaryAthleteAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryAthleteAllOfWithDefaults

`func NewSummaryAthleteAllOfWithDefaults() *SummaryAthleteAllOf`

NewSummaryAthleteAllOfWithDefaults instantiates a new SummaryAthleteAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceState

`func (o *SummaryAthleteAllOf) GetResourceState() int32`

GetResourceState returns the ResourceState field if non-nil, zero value otherwise.

### GetResourceStateOk

`func (o *SummaryAthleteAllOf) GetResourceStateOk() (*int32, bool)`

GetResourceStateOk returns a tuple with the ResourceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceState

`func (o *SummaryAthleteAllOf) SetResourceState(v int32)`

SetResourceState sets ResourceState field to given value.

### HasResourceState

`func (o *SummaryAthleteAllOf) HasResourceState() bool`

HasResourceState returns a boolean if a field has been set.

### GetFirstname

`func (o *SummaryAthleteAllOf) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *SummaryAthleteAllOf) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *SummaryAthleteAllOf) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *SummaryAthleteAllOf) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *SummaryAthleteAllOf) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *SummaryAthleteAllOf) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *SummaryAthleteAllOf) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *SummaryAthleteAllOf) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetProfileMedium

`func (o *SummaryAthleteAllOf) GetProfileMedium() string`

GetProfileMedium returns the ProfileMedium field if non-nil, zero value otherwise.

### GetProfileMediumOk

`func (o *SummaryAthleteAllOf) GetProfileMediumOk() (*string, bool)`

GetProfileMediumOk returns a tuple with the ProfileMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMedium

`func (o *SummaryAthleteAllOf) SetProfileMedium(v string)`

SetProfileMedium sets ProfileMedium field to given value.

### HasProfileMedium

`func (o *SummaryAthleteAllOf) HasProfileMedium() bool`

HasProfileMedium returns a boolean if a field has been set.

### GetProfile

`func (o *SummaryAthleteAllOf) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SummaryAthleteAllOf) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SummaryAthleteAllOf) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SummaryAthleteAllOf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetCity

`func (o *SummaryAthleteAllOf) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SummaryAthleteAllOf) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SummaryAthleteAllOf) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SummaryAthleteAllOf) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *SummaryAthleteAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SummaryAthleteAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SummaryAthleteAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SummaryAthleteAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *SummaryAthleteAllOf) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SummaryAthleteAllOf) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SummaryAthleteAllOf) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SummaryAthleteAllOf) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetSex

`func (o *SummaryAthleteAllOf) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *SummaryAthleteAllOf) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *SummaryAthleteAllOf) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *SummaryAthleteAllOf) HasSex() bool`

HasSex returns a boolean if a field has been set.

### GetPremium

`func (o *SummaryAthleteAllOf) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *SummaryAthleteAllOf) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *SummaryAthleteAllOf) SetPremium(v bool)`

SetPremium sets Premium field to given value.

### HasPremium

`func (o *SummaryAthleteAllOf) HasPremium() bool`

HasPremium returns a boolean if a field has been set.

### GetSummit

`func (o *SummaryAthleteAllOf) GetSummit() bool`

GetSummit returns the Summit field if non-nil, zero value otherwise.

### GetSummitOk

`func (o *SummaryAthleteAllOf) GetSummitOk() (*bool, bool)`

GetSummitOk returns a tuple with the Summit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummit

`func (o *SummaryAthleteAllOf) SetSummit(v bool)`

SetSummit sets Summit field to given value.

### HasSummit

`func (o *SummaryAthleteAllOf) HasSummit() bool`

HasSummit returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SummaryAthleteAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SummaryAthleteAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SummaryAthleteAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SummaryAthleteAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SummaryAthleteAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SummaryAthleteAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SummaryAthleteAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SummaryAthleteAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


