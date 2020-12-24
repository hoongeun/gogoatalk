package util

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/nleeper/goment"
)

func TsToTime(ts *timestamp.Timestamp) time.Time {
	t, err := ptypes.Timestamp(ts)
	if err != nil {
		return time.Now()
	}
	return t.In(time.Local)
}

func ToLiteralTime(ut int64) string {
	now, err := goment.New()
	if err != nil {
		panic(err)
	}
	tl, err := goment.Unix(ut)
	if err != nil {
		panic(err)
	}
	tl.Local()
	if tl.IsBefore(now.Subtract(1, "days")) {
		// 2014/10/2
		return tl.Format("LLL")
	}
	return tl.FromNow()
}
