package strava

type StreamKind string

const (
	time           StreamKind = "time"
	distance       StreamKind = "distance"
	latlng         StreamKind = "latlng"
	altitude       StreamKind = "altitude"
	velicitySmooth StreamKind = "velocity_smooth"
	heartrate      StreamKind = "heartrate"
	cadence        StreamKind = "cadence"
	watts          StreamKind = "watts"
	temp           StreamKind = "temp"
	moving         StreamKind = "moving"
	grade_smooth   StreamKind = "grade_smooth"
)

var AllStreamKinds []StreamKind = []StreamKind{
	time,
	distance,
	latlng,
	altitude,
	velicitySmooth,
	heartrate,
	cadence,
	watts,
	temp,
	moving,
	grade_smooth,
}

var AllStreamKindsString []string = []string{
	string(time),
	string(distance),
	string(latlng),
	string(altitude),
	string(velicitySmooth),
	string(heartrate),
	string(cadence),
	string(watts),
	string(temp),
	string(moving),
	string(grade_smooth),
}
