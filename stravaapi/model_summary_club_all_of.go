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

// checks if the SummaryClubAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SummaryClubAllOf{}

// SummaryClubAllOf struct for SummaryClubAllOf
type SummaryClubAllOf struct {
	// URL to a 60x60 pixel profile picture.
	ProfileMedium *string `json:"profile_medium,omitempty"`
	// URL to a ~1185x580 pixel cover photo.
	CoverPhoto *string `json:"cover_photo,omitempty"`
	// URL to a ~360x176  pixel cover photo.
	CoverPhotoSmall *string `json:"cover_photo_small,omitempty"`
	// Deprecated. Prefer to use activity_types.
	SportType *string `json:"sport_type,omitempty"`
	// The activity types that count for a club. This takes precedence over sport_type.
	ActivityTypes []ActivityType `json:"activity_types,omitempty"`
	// The club's city.
	City *string `json:"city,omitempty"`
	// The club's state or geographical region.
	State *string `json:"state,omitempty"`
	// The club's country.
	Country *string `json:"country,omitempty"`
	// Whether the club is private.
	Private *bool `json:"private,omitempty"`
	// The club's member count.
	MemberCount *int32 `json:"member_count,omitempty"`
	// Whether the club is featured or not.
	Featured *bool `json:"featured,omitempty"`
	// Whether the club is verified or not.
	Verified *bool `json:"verified,omitempty"`
	// The club's vanity URL.
	Url *string `json:"url,omitempty"`
}

// NewSummaryClubAllOf instantiates a new SummaryClubAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummaryClubAllOf() *SummaryClubAllOf {
	this := SummaryClubAllOf{}
	return &this
}

// NewSummaryClubAllOfWithDefaults instantiates a new SummaryClubAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummaryClubAllOfWithDefaults() *SummaryClubAllOf {
	this := SummaryClubAllOf{}
	return &this
}

// GetProfileMedium returns the ProfileMedium field value if set, zero value otherwise.
func (o *SummaryClubAllOf) GetProfileMedium() string {
	if o == nil || IsNil(o.ProfileMedium) {
		var ret string
		return ret
	}
	return *o.ProfileMedium
}

// GetProfileMediumOk returns a tuple with the ProfileMedium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryClubAllOf) GetProfileMediumOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileMedium) {
		return nil, false
	}
	return o.ProfileMedium, true
}

// HasProfileMedium returns a boolean if a field has been set.
func (o *SummaryClubAllOf) HasProfileMedium() bool {
	if o != nil && !IsNil(o.ProfileMedium) {
		return true
	}

	return false
}

// SetProfileMedium gets a reference to the given string and assigns it to the ProfileMedium field.
func (o *SummaryClubAllOf) SetProfileMedium(v string) {
	o.ProfileMedium = &v
}

// GetCoverPhoto returns the CoverPhoto field value if set, zero value otherwise.
func (o *SummaryClubAllOf) GetCoverPhoto() string {
	if o == nil || IsNil(o.CoverPhoto) {
		var ret string
		return ret
	}
	return *o.CoverPhoto
}

// GetCoverPhotoOk returns a tuple with the CoverPhoto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryClubAllOf) GetCoverPhotoOk() (*string, bool) {
	if o == nil || IsNil(o.CoverPhoto) {
		return nil, false
	}
	return o.CoverPhoto, true
}

// HasCoverPhoto returns a boolean if a field has been set.
func (o *SummaryClubAllOf) HasCoverPhoto() bool {
	if o != nil && !IsNil(o.CoverPhoto) {
		return true
	}

	return false
}

// SetCoverPhoto gets a reference to the given string and assigns it to the CoverPhoto field.
func (o *SummaryClubAllOf) SetCoverPhoto(v string) {
	o.CoverPhoto = &v
}

// GetCoverPhotoSmall returns the CoverPhotoSmall field value if set, zero value otherwise.
func (o *SummaryClubAllOf) GetCoverPhotoSmall() string {
	if o == nil || IsNil(o.CoverPhotoSmall) {
		var ret string
		return ret
	}
	return *o.CoverPhotoSmall
}

// GetCoverPhotoSmallOk returns a tuple with the CoverPhotoSmall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryClubAllOf) GetCoverPhotoSmallOk() (*string, bool) {
	if o == nil || IsNil(o.CoverPhotoSmall) {
		return nil, false
	}
	return o.CoverPhotoSmall, true
}

// HasCoverPhotoSmall returns a boolean if a field has been set.
func (o *SummaryClubAllOf) HasCoverPhotoSmall() bool {
	if o != nil && !IsNil(o.CoverPhotoSmall) {
		return true
	}

	return false
}

// SetCoverPhotoSmall gets a reference to the given string and assigns it to the CoverPhotoSmall field.
func (o *SummaryClubAllOf) SetCoverPhotoSmall(v string) {
	o.CoverPhotoSmall = &v
}

// GetSportType returns the SportType field value if set, zero value otherwise.
func (o *SummaryClubAllOf) GetSportType() string {
	if o == nil || IsNil(o.SportType) {
		var ret string
		return ret
	}
	return *o.SportType
}

// GetSportTypeOk returns a tuple with the SportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryClubAllOf) GetSportTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SportType) {
		return nil, false
	}
	return o.SportType, true
}

// HasSportType returns a boolean if a field has been set.
func (o *SummaryClubAllOf) HasSportType() bool {
	if o != nil && !IsNil(o.SportType) {
		return true
	}

	return false
}

// SetSportType gets a reference to the given string and assigns it to the SportType field.
func (o *SummaryClubAllOf) SetSportType(v string) {
	o.SportType = &v
}

// GetActivityTypes returns the ActivityTypes field value if set, zero value otherwise.
func (o *SummaryClubAllOf) GetActivityTypes() []ActivityType {
	if o == nil || IsNil(o.ActivityTypes) {
		var ret []ActivityType
		return ret
	}
	return o.ActivityTypes
}

// GetActivityTypesOk returns a tuple with the ActivityTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryClubAllOf) GetActivityTypesOk() ([]ActivityType, bool) {
	if o == nil || IsNil(o.ActivityTypes) {
		return nil, false
	}
	return o.ActivityTypes, true
}

// HasActivityTypes returns a boolean if a field has been set.
func (o *SummaryClubAllOf) HasActivityTypes() bool {
	if o != nil && !IsNil(o.ActivityTypes) {
		return true
	}

	return false
}

// SetActivityTypes gets a reference to the given []ActivityType and assigns it to the ActivityTypes field.
func (o *SummaryClubAllOf) SetActivityTypes(v []ActivityType) {
	o.ActivityTypes = v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *SummaryClubAllOf) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryClubAllOf) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *SummaryClubAllOf) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *SummaryClubAllOf) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SummaryClubAllOf) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryClubAllOf) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SummaryClubAllOf) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SummaryClubAllOf) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *SummaryClubAllOf) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryClubAllOf) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *SummaryClubAllOf) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *SummaryClubAllOf) SetCountry(v string) {
	o.Country = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *SummaryClubAllOf) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryClubAllOf) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *SummaryClubAllOf) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *SummaryClubAllOf) SetPrivate(v bool) {
	o.Private = &v
}

// GetMemberCount returns the MemberCount field value if set, zero value otherwise.
func (o *SummaryClubAllOf) GetMemberCount() int32 {
	if o == nil || IsNil(o.MemberCount) {
		var ret int32
		return ret
	}
	return *o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryClubAllOf) GetMemberCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MemberCount) {
		return nil, false
	}
	return o.MemberCount, true
}

// HasMemberCount returns a boolean if a field has been set.
func (o *SummaryClubAllOf) HasMemberCount() bool {
	if o != nil && !IsNil(o.MemberCount) {
		return true
	}

	return false
}

// SetMemberCount gets a reference to the given int32 and assigns it to the MemberCount field.
func (o *SummaryClubAllOf) SetMemberCount(v int32) {
	o.MemberCount = &v
}

// GetFeatured returns the Featured field value if set, zero value otherwise.
func (o *SummaryClubAllOf) GetFeatured() bool {
	if o == nil || IsNil(o.Featured) {
		var ret bool
		return ret
	}
	return *o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryClubAllOf) GetFeaturedOk() (*bool, bool) {
	if o == nil || IsNil(o.Featured) {
		return nil, false
	}
	return o.Featured, true
}

// HasFeatured returns a boolean if a field has been set.
func (o *SummaryClubAllOf) HasFeatured() bool {
	if o != nil && !IsNil(o.Featured) {
		return true
	}

	return false
}

// SetFeatured gets a reference to the given bool and assigns it to the Featured field.
func (o *SummaryClubAllOf) SetFeatured(v bool) {
	o.Featured = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *SummaryClubAllOf) GetVerified() bool {
	if o == nil || IsNil(o.Verified) {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryClubAllOf) GetVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Verified) {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *SummaryClubAllOf) HasVerified() bool {
	if o != nil && !IsNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *SummaryClubAllOf) SetVerified(v bool) {
	o.Verified = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SummaryClubAllOf) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SummaryClubAllOf) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SummaryClubAllOf) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SummaryClubAllOf) SetUrl(v string) {
	o.Url = &v
}

func (o SummaryClubAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SummaryClubAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileMedium) {
		toSerialize["profile_medium"] = o.ProfileMedium
	}
	if !IsNil(o.CoverPhoto) {
		toSerialize["cover_photo"] = o.CoverPhoto
	}
	if !IsNil(o.CoverPhotoSmall) {
		toSerialize["cover_photo_small"] = o.CoverPhotoSmall
	}
	if !IsNil(o.SportType) {
		toSerialize["sport_type"] = o.SportType
	}
	if !IsNil(o.ActivityTypes) {
		toSerialize["activity_types"] = o.ActivityTypes
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
	if !IsNil(o.MemberCount) {
		toSerialize["member_count"] = o.MemberCount
	}
	if !IsNil(o.Featured) {
		toSerialize["featured"] = o.Featured
	}
	if !IsNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableSummaryClubAllOf struct {
	value *SummaryClubAllOf
	isSet bool
}

func (v NullableSummaryClubAllOf) Get() *SummaryClubAllOf {
	return v.value
}

func (v *NullableSummaryClubAllOf) Set(val *SummaryClubAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSummaryClubAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSummaryClubAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummaryClubAllOf(val *SummaryClubAllOf) *NullableSummaryClubAllOf {
	return &NullableSummaryClubAllOf{value: val, isSet: true}
}

func (v NullableSummaryClubAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummaryClubAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


