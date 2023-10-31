package browsers

import "time"

// Timestamp is a container for time.Time
type Timestamp struct {
	time.Time
}

// TimestampFromFloat64 converts float64 timestamps into
// time.Time objects wrapped in a Timestamp object.
func TimestampFromFloat64(ts float64) Timestamp {
	secs := int64(ts)
	nsecs := int64((ts - float64(secs)) * 1e9)
	return Timestamp{time.Unix(secs, nsecs)}
}
