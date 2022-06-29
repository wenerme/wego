package protox

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UnixToTimestamp converts a unix timestamp to a protobuf timestamp.
func UnixToTimestamp(t int64) *timestamppb.Timestamp {
	if t == 0 {
		return nil
	}
	return &timestamppb.Timestamp{
		Seconds: t,
	}
}

// UnixMilliToTimestamp converts a unix milli timestamp to a protobuf timestamp.
func UnixMilliToTimestamp(t int64) *timestamppb.Timestamp {
	if t == 0 {
		return nil
	}
	return ToTimestamp(time.UnixMilli(t))
}

// ToTimestamp converts a time.Time to a protobuf timestamp.
func ToTimestamp(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}

// PtrToTimestamp converts a pointer to a time.Time to a protobuf timestamp.
func PtrToTimestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return ToTimestamp(*t)
}

// ToDuration converts a time.Duration to a protobuf duration.
var ToDuration = durationpb.New
