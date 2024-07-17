package model

import "regexp"

var vatprcAirportRegexp *regexp.Regexp
var vatprcControllerRegexp *regexp.Regexp

func init() {
	vatprcAirportRegexp = regexp.MustCompile("Z[BMSPGJYWLH][A-Z]{2}")
	vatprcControllerRegexp = regexp.MustCompile("^(Z[BSGUHWJPLYM][A-Z0-9]{2}(_[A-Z0-9]*)?_(DEL|GND|TWR|APP|DEP|CTR))|(PRC_FSS)$")
}
