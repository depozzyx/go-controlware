package shared

import (
	"strconv"
	"time"
)

func LastTimestampId(lastTimestamp int64, offset int) string {
	return strconv.FormatInt(lastTimestamp+int64(offset), 20) + Version
}

func LastTimestamp() int64 {
	return time.Now().UnixMilli()
}
