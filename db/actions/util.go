package actions

import (
	"encoding/binary"
	"strings"
)

// itob returns an 8-byte big endian representation of v.
func iotb(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// date1 and date2 is define as time.DateTime
// isSameDate determines date1 and date2 whether is same date
// DateTime   = "2006-01-02 15:04:05"
func isSameDate(date1, date2 string) bool {
	return getDate(date1) == getDate(date2)
}

// DateTime   = "2006-01-02 15:04:05"
func getDate(date string) string {
	return strings.Fields(date)[0]
}
