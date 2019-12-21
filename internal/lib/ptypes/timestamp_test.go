package ptypes

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func TestShiftTimestamp(t *testing.T) {
	now := time.Now()
	type args struct {
		tt     *timestamp.Timestamp
		sendAt *timestamp.Timestamp
		now    time.Time
	}
	tests := []struct {
		name  string
		args  args
		want  *timestamp.Timestamp
		want1 bool
	}{
		{
			name: "#1: Nonshift time",
			args: args{
				tt:     TimestampProto(now.Add(-2 * time.Hour)),
				sendAt: TimestampProto(now.Add(-2 * time.Hour)),
				now:    now,
			},
			want:  TimestampProto(now.Add(-2 * time.Hour)),
			want1: false,
		},
		{
			name: "#2: Must shift 1hour from now",
			args: args{
				tt:     TimestampProto(now.Add(2 * time.Hour)),
				sendAt: TimestampProto(now.Add(3 * time.Hour)),
				now:    now,
			},
			want:  TimestampProto(now.Add(-1 * time.Hour)),
			want1: true,
		},
		{
			name: "#3: Must shift 2hour from now",
			args: args{
				tt:     TimestampProto(now.Add(2 * time.Hour)),
				sendAt: TimestampProto(now.Add(4 * time.Hour)),
				now:    now,
			},
			want:  TimestampProto(now.Add(-2 * time.Hour)),
			want1: true,
		},
		{
			name: "#4: Must shift a month ago to now",
			args: args{
				tt:     TimestampProto(now.Add(-30 * 24 * time.Hour)),
				sendAt: TimestampProto(now),
				now:    now,
			},
			want:  TimestampProto(now),
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ShiftTimestamp(tt.args.tt, tt.args.sendAt, tt.args.now)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShiftTimestamp() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ShiftTimestamp() got1 = %v, want %v", got1, tt.want1)
			}
			assert.True(t, time.Unix(got.Seconds, int64(got.Nanos)).UnixNano() <= now.UnixNano(), "Shifted time must lower than now")
		})
	}
}
