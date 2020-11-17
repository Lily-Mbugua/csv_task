package sylabs

import (
	"log"
	"strconv"
	"time"
)

// parseRFC3339Time is a helpful utility function that will assist in converting the string containing
// time in RFC3339 into a friendly Time.Time value
func parseRFC3339Time(s string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Fatal(err)
	}
	return parsedTime
}

// parseBool is a helpful utility function that will assist in converting a string into a boolean
func parseBool(s string) bool {
	// lets parse a boolean from string
	val, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

// parseUInt64 is a helpful utility function that will assist in converting a string into a int64
func parseUInt64(s string) uint64 {
	// lets parse a bolean from string
	val, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
