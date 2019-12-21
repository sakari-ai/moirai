package ptypes

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func Timestamp(ts *timestamp.Timestamp) time.Time {
	value, _ := ptypes.Timestamp(ts)
	return value
}

func TimestampProto(ts time.Time) *timestamp.Timestamp {
	value, _ := ptypes.TimestampProto(ts)
	return value
}

func ShiftTimestamp(tt *timestamp.Timestamp, sendAt *timestamp.Timestamp, now time.Time) (*timestamp.Timestamp, bool) {
	shift := doubleCheckTimestamp(Timestamp(sendAt), Timestamp(tt), now)
	if shift != 0 {
		shiftTime := Timestamp(tt)
		shiftTime = shiftTime.Add(time.Duration(shift))
		return TimestampProto(shiftTime), true
	}
	return tt, false
}

func doubleCheckTimestamp(checkedTime time.Time, ttTime time.Time, now time.Time) int64 {
	if checkedTime.UnixNano() > now.UnixNano() || checkedTime.UnixNano() < time.Now().Add(-30*24*time.Hour).UnixNano() {
		shift := now.UnixNano() - checkedTime.UnixNano()
		if ttTime.UnixNano()+shift <= now.UnixNano() {
			return shift
		}
	}
	if ttTime.UnixNano() > now.UnixNano() || ttTime.UnixNano() < time.Now().Add(-30*24*time.Hour).UnixNano() {
		shift := now.UnixNano() - ttTime.UnixNano()
		return shift
	}
	return 0
}
