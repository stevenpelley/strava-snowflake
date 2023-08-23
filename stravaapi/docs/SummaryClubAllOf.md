# SummaryClubAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileMedium** | Pointer to **string** | URL to a 60x60 pixel profile picture. | [optional] 
**CoverPhoto** | Pointer to **string** | URL to a ~1185x580 pixel cover photo. | [optional] 
**CoverPhotoSmall** | Pointer to **string** | URL to a ~360x176  pixel cover photo. | [optional] 
**SportType** | Pointer to **string** | Deprecated. Prefer to use activity_types. | [optional] 
**ActivityTypes** | Pointer to [**[]ActivityType**](ActivityType.md) | The activity types that count for a club. This takes precedence over sport_type. | [optional] 
**City** | Pointer to **string** | The club&#39;s city. | [optional] 
**State** | Pointer to **string** | The club&#39;s state or geographical region. | [optional] 
**Country** | Pointer to **string** | The club&#39;s country. | [optional] 
**Private** | Pointer to **bool** | Whether the club is private. | [optional] 
**MemberCount** | Pointer to **int32** | The club&#39;s member count. | [optional] 
**Featured** | Pointer to **bool** | Whether the club is featured or not. | [optional] 
**Verified** | Pointer to **bool** | Whether the club is verified or not. | [optional] 
**Url** | Pointer to **string** | The club&#39;s vanity URL. | [optional] 

## Methods

### NewSummaryClubAllOf

`func NewSummaryClubAllOf() *SummaryClubAllOf`

NewSummaryClubAllOf instantiates a new SummaryClubAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryClubAllOfWithDefaults

`func NewSummaryClubAllOfWithDefaults() *SummaryClubAllOf`

NewSummaryClubAllOfWithDefaults instantiates a new SummaryClubAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileMedium

`func (o *SummaryClubAllOf) GetProfileMedium() string`

GetProfileMedium returns the ProfileMedium field if non-nil, zero value otherwise.

### GetProfileMediumOk

`func (o *SummaryClubAllOf) GetProfileMediumOk() (*string, bool)`

GetProfileMediumOk returns a tuple with the ProfileMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMedium

`func (o *SummaryClubAllOf) SetProfileMedium(v string)`

SetProfileMedium sets ProfileMedium field to given value.

### HasProfileMedium

`func (o *SummaryClubAllOf) HasProfileMedium() bool`

HasProfileMedium returns a boolean if a field has been set.

### GetCoverPhoto

`func (o *SummaryClubAllOf) GetCoverPhoto() string`

GetCoverPhoto returns the CoverPhoto field if non-nil, zero value otherwise.

### GetCoverPhotoOk

`func (o *SummaryClubAllOf) GetCoverPhotoOk() (*string, bool)`

GetCoverPhotoOk returns a tuple with the CoverPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverPhoto

`func (o *SummaryClubAllOf) SetCoverPhoto(v string)`

SetCoverPhoto sets CoverPhoto field to given value.

### HasCoverPhoto

`func (o *SummaryClubAllOf) HasCoverPhoto() bool`

HasCoverPhoto returns a boolean if a field has been set.

### GetCoverPhotoSmall

`func (o *SummaryClubAllOf) GetCoverPhotoSmall() string`

GetCoverPhotoSmall returns the CoverPhotoSmall field if non-nil, zero value otherwise.

### GetCoverPhotoSmallOk

`func (o *SummaryClubAllOf) GetCoverPhotoSmallOk() (*string, bool)`

GetCoverPhotoSmallOk returns a tuple with the CoverPhotoSmall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverPhotoSmall

`func (o *SummaryClubAllOf) SetCoverPhotoSmall(v string)`

SetCoverPhotoSmall sets CoverPhotoSmall field to given value.

### HasCoverPhotoSmall

`func (o *SummaryClubAllOf) HasCoverPhotoSmall() bool`

HasCoverPhotoSmall returns a boolean if a field has been set.

### GetSportType

`func (o *SummaryClubAllOf) GetSportType() string`

GetSportType returns the SportType field if non-nil, zero value otherwise.

### GetSportTypeOk

`func (o *SummaryClubAllOf) GetSportTypeOk() (*string, bool)`

GetSportTypeOk returns a tuple with the SportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSportType

`func (o *SummaryClubAllOf) SetSportType(v string)`

SetSportType sets SportType field to given value.

### HasSportType

`func (o *SummaryClubAllOf) HasSportType() bool`

HasSportType returns a boolean if a field has been set.

### GetActivityTypes

`func (o *SummaryClubAllOf) GetActivityTypes() []ActivityType`

GetActivityTypes returns the ActivityTypes field if non-nil, zero value otherwise.

### GetActivityTypesOk

`func (o *SummaryClubAllOf) GetActivityTypesOk() (*[]ActivityType, bool)`

GetActivityTypesOk returns a tuple with the ActivityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTypes

`func (o *SummaryClubAllOf) SetActivityTypes(v []ActivityType)`

SetActivityTypes sets ActivityTypes field to given value.

### HasActivityTypes

`func (o *SummaryClubAllOf) HasActivityTypes() bool`

HasActivityTypes returns a boolean if a field has been set.

### GetCity

`func (o *SummaryClubAllOf) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SummaryClubAllOf) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SummaryClubAllOf) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SummaryClubAllOf) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *SummaryClubAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SummaryClubAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SummaryClubAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SummaryClubAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *SummaryClubAllOf) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SummaryClubAllOf) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SummaryClubAllOf) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SummaryClubAllOf) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPrivate

`func (o *SummaryClubAllOf) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *SummaryClubAllOf) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *SummaryClubAllOf) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *SummaryClubAllOf) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetMemberCount

`func (o *SummaryClubAllOf) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *SummaryClubAllOf) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *SummaryClubAllOf) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *SummaryClubAllOf) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetFeatured

`func (o *SummaryClubAllOf) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *SummaryClubAllOf) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *SummaryClubAllOf) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *SummaryClubAllOf) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetVerified

`func (o *SummaryClubAllOf) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *SummaryClubAllOf) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *SummaryClubAllOf) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *SummaryClubAllOf) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetUrl

`func (o *SummaryClubAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SummaryClubAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SummaryClubAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SummaryClubAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


