# DetailedClubAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Membership** | Pointer to **string** | The membership status of the logged-in athlete. | [optional] 
**Admin** | Pointer to **bool** | Whether the currently logged-in athlete is an administrator of this club. | [optional] 
**Owner** | Pointer to **bool** | Whether the currently logged-in athlete is the owner of this club. | [optional] 
**FollowingCount** | Pointer to **int32** | The number of athletes in the club that the logged-in athlete follows. | [optional] 

## Methods

### NewDetailedClubAllOf

`func NewDetailedClubAllOf() *DetailedClubAllOf`

NewDetailedClubAllOf instantiates a new DetailedClubAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedClubAllOfWithDefaults

`func NewDetailedClubAllOfWithDefaults() *DetailedClubAllOf`

NewDetailedClubAllOfWithDefaults instantiates a new DetailedClubAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembership

`func (o *DetailedClubAllOf) GetMembership() string`

GetMembership returns the Membership field if non-nil, zero value otherwise.

### GetMembershipOk

`func (o *DetailedClubAllOf) GetMembershipOk() (*string, bool)`

GetMembershipOk returns a tuple with the Membership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembership

`func (o *DetailedClubAllOf) SetMembership(v string)`

SetMembership sets Membership field to given value.

### HasMembership

`func (o *DetailedClubAllOf) HasMembership() bool`

HasMembership returns a boolean if a field has been set.

### GetAdmin

`func (o *DetailedClubAllOf) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *DetailedClubAllOf) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *DetailedClubAllOf) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *DetailedClubAllOf) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetOwner

`func (o *DetailedClubAllOf) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DetailedClubAllOf) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DetailedClubAllOf) SetOwner(v bool)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DetailedClubAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetFollowingCount

`func (o *DetailedClubAllOf) GetFollowingCount() int32`

GetFollowingCount returns the FollowingCount field if non-nil, zero value otherwise.

### GetFollowingCountOk

`func (o *DetailedClubAllOf) GetFollowingCountOk() (*int32, bool)`

GetFollowingCountOk returns a tuple with the FollowingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingCount

`func (o *DetailedClubAllOf) SetFollowingCount(v int32)`

SetFollowingCount sets FollowingCount field to given value.

### HasFollowingCount

`func (o *DetailedClubAllOf) HasFollowingCount() bool`

HasFollowingCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


