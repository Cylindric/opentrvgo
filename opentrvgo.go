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
	Message           string    `json:"message"`
	Sequence          int       `json:"sequence"`
}

// ParseSensorReport attempts to parse a block of data, for example from a serial port, into a Sample record
func ParseSensorReport(input []byte) (output map[string]interface{}, err error) {
	var d map[string]interface{}
	// var output map[string]interface{}

	output = make(map[string]interface{})

	// Timestamp is now
	output["timestamp"] = time.Now()

	// Stringify the byte array.
	str := strings.TrimSpace(string(input))
	output["message"] = str

	// If the line isn't actually a JSON response, just return invalid.
	if strings.HasPrefix(str, `{"@"`) == false {
		return output, fmt.Errorf("Not a valid sensor report")
	}

	// First try to turn the passed-in data from JSON to an object.
	if err := json.Unmarshal(input, &d); err != nil {
		return output, fmt.Errorf("could not parse json")
	}

	// Try to fill the sample with data.
	if _, ok := d["@"]; ok {
		output["device"] = d["@"].(string)
	} else {
		return output, fmt.Errorf("record does not contain a device ID field '@'")
	}

	if _, ok := d["+"]; ok {
		output["sequence"] = int(d["+"].(float64))
	}

	if _, ok := d["B|cV"]; ok {
		output["battery"] = d["B|cV"].(float64) / 100
	}

	if _, ok := d["H|%"]; ok {
		output["humidity"] = int(d["H|%"].(float64))
	}

	if _, ok := d["L"]; ok {
		output["light"] = int(d["L"].(float64))
	}

	if _, ok := d["tT|C"]; ok {
		output["target_temperature"] = int(d["tT|C"].(float64))
	}

	if _, ok := d["T|C16"]; ok {
		output["temperature"] = d["T|C16"].(float64) / 16
	}

	if _, ok := d["O"]; ok {
		output["occupancy"] = int(d["O"].(float64))
	}

	if _, ok := d["vac|h"]; ok {
		output["vacancy"] = int(d["vac|h"].(float64))
	}

	if _, ok := d["v|%"]; ok {
		output["valve"] = int(d["v|%"].(float64))
	}

	return output, nil
}
