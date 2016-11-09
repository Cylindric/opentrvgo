package opentrvgo

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Sample represents a single record of data from a sensor
type Sample struct {
	Timestamp         time.Time `json:"timestamp"`
	Device            string    `json:"device"`
	Temperature       float64   `json:"temperature"`
	Humidity          float64   `json:"humidity"`
	Light             float64   `json:"light"`
	TargetTemperature float64   `json:"targettemperature"`
	Valve             float64   `json:"valve"`
	Occupancy         float64   `json:"occupancy"`
	Battery           float64   `json:"battery"`
}

// ParseSensorReport attempts to parse a block of data, for example from a serial port, into a Sample record
func ParseSensorReport(input []byte, smple Sample) (valid bool, err error) {
	var d Sample

	// Stringify the byte array
	s := strings.TrimSpace(string(input))

	// If the line isn't actually a JSON response, just return invalid
	if strings.HasPrefix(s, `{"@"`) == false {
		return false, fmt.Errorf("Not a valid sensor report")
	}

	// First try to turn the passed-in data from JSON to an object
	if err := json.Unmarshal(input, &d); err != nil {
		return false, fmt.Errorf("could not parse json.")
	}

	return false, nil
}
