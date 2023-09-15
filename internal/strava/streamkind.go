package strava

type StreamKind string

const (
	StreamTime           StreamKind = "time"
	StreamDistance       StreamKind = "distance"
	StreamLatlng         StreamKind = "latlng"
	StreamAltitude       StreamKind = "altitude"
	StreamVelicitySmooth StreamKind = "velocity_smooth"
	StreamHeartrate      StreamKind = "heartrate"
	StreamCadence        StreamKind = "cadence"
	StreamWatts          StreamKind = "watts"
	StreamTemp           StreamKind = "temp"
	StreamMoving         StreamKind = "moving"
	StreamGradeSmooth    StreamKind = "grade_smooth"
)

var AllStreamKinds []StreamKind = []StreamKind{
	StreamTime,
	StreamDistance,
	StreamLatlng,
	StreamAltitude,
	StreamVelicitySmooth,
	StreamHeartrate,
	StreamCadence,
	StreamWatts,
	StreamTemp,
	StreamMoving,
	StreamGradeSmooth,
}

var AllStreamKindsString []string = []string{
	string(StreamTime),
	string(StreamDistance),
	string(StreamLatlng),
	string(StreamAltitude),
	string(StreamVelicitySmooth),
	string(StreamHeartrate),
	string(StreamCadence),
	string(StreamWatts),
	string(StreamTemp),
	string(StreamMoving),
	string(StreamGradeSmooth),
}
