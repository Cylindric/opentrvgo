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
	Humidity          int       `json:"humidity"`
	Light             int       `json:"light"`
	TargetTemperature int       `json:"targettemperature"`
	Valve             int       `json:"valve"`
	Occupancy         int       `json:"occupancy"`
	Battery           float64   `json:"battery"`
	Vacancy           int       `json:"vacancy"`
}

// ParseSensorReport attempts to parse a block of data, for example from a serial port, into a Sample record
func ParseSensorReport(input []byte) (sample Sample, err error) {
	var d map[string]interface{}

	// Timestamp is now
	sample.Timestamp = time.Now()

	// Stringify the byte array.
	str := strings.TrimSpace(string(input))

	// If the line isn't actually a JSON response, just return invalid.
	if strings.HasPrefix(str, `{"@"`) == false {
		return sample, fmt.Errorf("Not a valid sensor report")
	}

	// First try to turn the passed-in data from JSON to an object.
	if err := json.Unmarshal(input, &d); err != nil {
		return sample, fmt.Errorf("could not parse json.")
	}

	// Try to fill the sample with data.

	if _, ok := d["@"]; ok {
		sample.Device = d["@"].(string)
	} else {
		return sample, fmt.Errorf("Record does not contain a device ID field '@'.")
	}

	if _, ok := d["B|cV"]; ok {
		sample.Battery = d["B|cV"].(float64) / 100
	}

	if _, ok := d["H|%"]; ok {
		sample.Humidity = int(d["H|%"].(float64))
	}

	if _, ok := d["L"]; ok {
		sample.Light = int(d["L"].(float64))
	}

	if _, ok := d["tT|C"]; ok {
		sample.TargetTemperature = int(d["tT|C"].(float64))
	}

	if _, ok := d["T|C16"]; ok {
		sample.Temperature = d["T|C16"].(float64) / 16
	}

	if _, ok := d["O"]; ok {
		sample.Occupancy = int(d["O"].(float64))
	}

	if _, ok := d["vac|h"]; ok {
		sample.Vacancy = int(d["vac|h"].(float64))
	}

	if _, ok := d["v|%"]; ok {
		sample.Valve = int(d["v|%"].(float64))
	}

	return sample, nil
}
